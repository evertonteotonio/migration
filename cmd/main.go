package main

import (
	"github.com/crgimenes/goconfig"
	"github.com/gosidekick/migration"
)

type config struct {
	Database string `cfg:"db" cfgRequired:"true"`
	Source   string `cfg:"source" cfgRequired:"true"`
	Migrate  string `cfg:"migrate" cfgRequired:"true"`
}

func main() {
	cfg := config{}

	err := goconfig.Parse(&cfg)
	if err != nil {
		println(err.Error())
		return
	}

	err = migration.Run(cfg.Source, cfg.Database, cfg.Migrate)
	if err != nil {
		println(err.Error())
	}
}
