package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfgPath := flag.String("config", "config.toml", "Path to config")
	flag.Parse()

	cfg, err := initConfig(*cfgPath)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()

	s, err := NewServer(cfg)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

	s.ListenAndServe(ctx)
}
