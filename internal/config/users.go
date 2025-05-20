package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func loadAuthorizedUsers(cfg *Config) error {
	candidates := []string{
		"internal/config/authorized_users.json",
		filepath.Join(".", "internal/config/authorized_users.json"),
		filepath.Join("..", "internal/config/authorized_users.json"),
	}

	var data []byte
	var err error
	var path string
	for _, p := range candidates {
		data, err = os.ReadFile(p)
		if err == nil {
			path = p
			break
		}
	}
	if err != nil {
		return fmt.Errorf("arquivo não encontrado: %w", err)
	}

	var file struct {
		Emails []string `json:"emails"`
	}
	if err := json.Unmarshal(data, &file); err != nil {
		return fmt.Errorf("erro ao decodificar %s: %w", path, err)
	}

	cfg.TestUsers = file.Emails
	log.Printf("✅ carregados %d usuários de %s", len(file.Emails), path)
	return nil
}
