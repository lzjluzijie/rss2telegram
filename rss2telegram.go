package main

import (
	"log"

	"github.com/lzjluzijie/rss2telegram/core"
)

func main() {
	app, err := core.GetApp()
	if err != nil {
		log.Fatalf("can not load config: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
