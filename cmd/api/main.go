package main

import (
	"context"
	"fmt"
	"log"
	"rest/config"
	"rest/internal/server"
)

func main() {
	log.Println("Starting server")
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
		//return
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
		//return
	}
	log.Println("Config loaded")

	s := server.NewServer(cfg)
	if err = s.Run(context.TODO()); err != nil {
		fmt.Println("Cannot start server: %v", err)
	}
}
