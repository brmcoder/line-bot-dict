package main

import (
	"log"

	"github.com/brmcoder/line-bot-dict/controller"
	"github.com/brmcoder/line-bot-dict/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := controller.NewServer(config)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
