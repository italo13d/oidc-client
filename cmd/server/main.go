package main

import (
	"log"
	"net/http"

	"github.com/italo13d/oidc-client/internal/config"
	"github.com/italo13d/oidc-client/internal/handlers"
	"github.com/italo13d/oidc-client/internal/oidc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	oidcClient, err := oidc.New(cfg)
	if err != nil {
		log.Fatalf("oidc error: %v", err)
	}

	h := handlers.New(oidcClient)

	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/callback", h.Callback)

	log.Printf("Servidor ouvindo em %s …", cfg.ServerAddr)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, nil))
}
