package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"Jayce/core/conf"
	"Jayce/tcp"
)

var (
	path       = filepath.Join("conf", "server.json")
	configFile = flag.String("file", path, "the config file")
)

func main() {
	flag.Parse()

	var (
		config tcp.Config
	)

	if err := conf.LoadConfig(*configFile, &config); err != nil {
		fmt.Printf("load config from file: %s occurs err: %v\n", *configFile, err)
		return
	}

	server, err := tcp.NewServer(config)
	if err != nil {
		fmt.Printf("new tcp server occurs err: %v\n", err)
	}

	server.Logger.Info("tcp new server start...")

	server.Server()
}
