// Package rule contains the server handler logic.
// This file is not built into the binary; Yaegi loads it from disk at runtime.
// You can edit it in-place and still benefit from IDE tooling.
package rule

import (
	"net/http"

	"github.com/dcjanus/yaegi_demo/internal/helper"

	"github.com/sirupsen/logrus"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	// Demonstrate using a third-party package.
	logrus.
		WithField("url", r.URL.String()).
		WithField("method", r.Method).
		Info("request received")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	// ref: https://github.com/DCjanus/yaegi_demo/issues/1
	if r.Header.Get("Content-Type") == "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Content-Type should not be application/json\n"))
		return
	}

	// Demonstrate using an internal package.
	helper.UselessHelper(w, helper.UselessHeader)

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello!\n"))
	w.Write([]byte("Your Content-Type is " + r.Header.Get("Content-Type") + "\n"))
	w.Write([]byte("Your User-Agent is " + r.Header.Get("User-Agent") + "\n"))
	w.Write([]byte("Your Host is " + r.Host + "\n"))
	w.Write([]byte("Your RemoteAddr is " + r.RemoteAddr + "\n"))
	w.Write([]byte("Your RequestURI is " + r.RequestURI + "\n"))
	w.Write([]byte("Your Method is " + r.Method + "\n"))
	w.Write([]byte("Your Proto is " + r.Proto + "\n"))
	w.Write([]byte("Your URL is " + r.URL.String() + "\n"))
	w.Write([]byte("Code above is wrote by GitHub Copilot.\n"))
	w.Write([]byte("\n"))
	w.Write([]byte(":)"))
}
