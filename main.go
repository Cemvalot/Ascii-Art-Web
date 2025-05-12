package main

import (
	"ascii-art-web-stylize/internal/config"
	"ascii-art-web-stylize/internal/server"
	"log"
)

func main() {
	cfg := config.Conf()
	srv := server.Create(cfg)

	log.Printf("Server starting on %s...", cfg.Port)
	log.Fatal(srv.Start())
}
