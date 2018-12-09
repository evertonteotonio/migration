package main

import (
	"fmt"

	"github.com/crgimenes/goconfig"
)

func main() {
	config := struct {
		Database string `cfg:"db" crgRequired:"true" cfgHelper:""`
		Source   string `cfg:"source" crgRequired:"true" cfgHelper:""`
		Migrate  string `cfg:"migrate" crgRequired:"true" cfgHelper:""`
	}{}

	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	fmt.Println(config)

}
