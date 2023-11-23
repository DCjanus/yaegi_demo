package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/dcjanus/yaegi_demo/internal/symbol"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type Engine struct {
	handle http.HandlerFunc
	mutex  sync.RWMutex

	rulePath   string
	ruleDigest string
}

func NewEngine(ctx context.Context, rulePath string) (*Engine, error) {
	ret := &Engine{
		rulePath: rulePath,
	}
	if err := ret.reload(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to reload")
	}
	go ret.proc(ctx)
	return ret, nil
}

func (e *Engine) proc(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := e.reload(ctx); err != nil {
				logrus.WithError(err).Error("failed to reload")
				continue
			}
		}
	}

}

func (e *Engine) reload(ctx context.Context) error {
	e.mutex.RLock()
	rulePath := e.rulePath
	e.mutex.RUnlock()

	ruleContent, err := os.ReadFile(rulePath)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}
	ruleSumBin := sha256.Sum256(ruleContent)
	ruleDigest := hex.EncodeToString(ruleSumBin[:])

	e.mutex.RLock()
	ruleDigestOld := e.ruleDigest
	e.mutex.RUnlock()

	if ruleDigest == ruleDigestOld {
		return nil
	}

	i := interp.New(interp.Options{
		GoPath: "/dev/null",
	})
	if err := i.Use(stdlib.Symbols); err != nil {
		return errors.Wrap(err, "failed to use stdlib symbols")
	}
	if err := i.Use(symbol.Symbols); err != nil {
		return errors.Wrap(err, "failed to use symbol symbols")
	}

	if _, err := i.EvalWithContext(ctx, string(ruleContent)); err != nil {
		return errors.Wrap(err, "failed to eval rule")
	}
	ret, err := i.EvalWithContext(ctx, "rule.Handle")
	if err != nil {
		return errors.Wrap(err, "failed to eval Handle")
	}
	handle, ok := ret.Interface().(func(http.ResponseWriter, *http.Request))
	if !ok {
		return errors.New("failed to convert Handle to func(http.ResponseWriter, *http.Request)")
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()
	if ruleDigest == e.ruleDigest {
		return nil
	}
	e.handle = handle
	e.ruleDigest = ruleDigest
	logrus.WithField("digest", ruleDigest).Info("rule reloaded")
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.mutex.RLock()
	handle := e.handle
	e.mutex.RUnlock()

	handle.ServeHTTP(w, r)
}
