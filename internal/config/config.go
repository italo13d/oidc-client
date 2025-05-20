package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IssuerURL    string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	ServerAddr   string
	TestUsers    []string
}

// Load lê o .env e carrega todas as configs
func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		IssuerURL:    os.Getenv("OIDC_ISSUER_URL"),
		ClientID:     os.Getenv("OIDC_CLIENT_ID"),
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OIDC_REDIRECT"),
		ServerAddr:   ":8080",
		TestUsers:    []string{},
	}

	if err := loadAuthorizedUsers(cfg); err != nil {
		log.Printf("⚠️  não foi possível carregar usuários autorizados: %v", err)
	}

	if cfg.IssuerURL == "" ||
		cfg.ClientID == "" ||
		cfg.ClientSecret == "" ||
		cfg.RedirectURL == "" {
		return nil, errors.New("variáveis de ambiente OIDC incompletas")
	}

	return cfg, nil
}
