package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"mstoney/httphandler"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}

func main() {
	fx.New(
		fx.Provide(
			ProvideConfig,
			ProvideLogger,
			http.NewServeMux,
			httphandler.New,
		),
		fx.Invoke(httphandler.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, cfg *Config, h *httphandler.Handler) {
	fmt.Println(*h.Mux)
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go http.ListenAndServe(cfg.ApplicationConfig.Address, h.Mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		fmt.Println("Error reading config file")
	}
	fmt.Println(data)
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Println("Error unmarshaling data")
	}
	fmt.Println(conf)
	return &conf
}

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	return slogger
}
