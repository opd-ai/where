// Command server launches the Where dedicated game server.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/opd-ai/where/config"
	"github.com/opd-ai/where/pkg/network"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv := network.NewServer(cfg.Server.Address, cfg.Server.Port, cfg.Server.TickRate)
	if err := srv.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "server start: %v\n", err)
		os.Exit(1)
	}

	log.Printf("server listening on %s:%d (tick rate: %d)", cfg.Server.Address, cfg.Server.Port, cfg.Server.TickRate)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("shutting down...")
	if err := srv.Stop(); err != nil {
		fmt.Fprintf(os.Stderr, "server stop: %v\n", err)
	}
}
