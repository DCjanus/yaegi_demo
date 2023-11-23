package main

import (
	"context"
	"github.com/dcjanus/yaegi_demo/internal/engine"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	pflag.String("rule", "./rule/rule.go", "rule file path")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eng, err := engine.NewEngine(ctx, viper.GetString("rule"))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: eng,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
