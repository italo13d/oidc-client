package oidc

import (
	"context"
	"errors"
	"fmt"

	gooidc "github.com/coreos/go-oidc"
	"github.com/italo13d/oidc-client/internal/config"
	"golang.org/x/oauth2"
)

var (
	ErrEmailNotVerified = errors.New("e-mail não verificado junto ao provedor")
)

type Client struct {
	oauth2        *oauth2.Config
	verifier      *gooidc.IDTokenVerifier
	allowedEmails map[string]struct{}
}

func New(cfg *config.Config) (*Client, error) {
	ctx := context.Background()

	provider, err := gooidc.NewProvider(ctx, cfg.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("falha no discovery do provedor: %w", err)
	}

	emailMap := make(map[string]struct{}, len(cfg.TestUsers))
	for _, e := range cfg.TestUsers {
		emailMap[e] = struct{}{}
	}

	return &Client{
		oauth2: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{gooidc.ScopeOpenID, "profile", "email"},
		},
		verifier:      provider.Verifier(&gooidc.Config{ClientID: cfg.ClientID}),
		allowedEmails: emailMap,
	}, nil
}

func (c *Client) AuthCodeURL(state string) string {
	return c.oauth2.AuthCodeURL(state)
}

type Claims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
}

func (c *Client) ExchangeAndVerify(ctx context.Context, code string) (*Claims, error) {
	tok, err := c.oauth2.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("falha ao trocar código por token: %w", err)
	}
	rawID, ok := tok.Extra("id_token").(string)
	if !ok || rawID == "" {
		return nil, errors.New("id_token ausente na resposta do provedor")
	}
	idToken, err := c.verifier.Verify(ctx, rawID)
	if err != nil {
		return nil, fmt.Errorf("token inválido: %w", err)
	}
	var claims Claims
	if err := idToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("falha ao ler claims do token: %w", err)
	}
	if !claims.EmailVerified {
		return nil, ErrEmailNotVerified
	}
	if _, ok := c.allowedEmails[claims.Email]; !ok {
		return nil, fmt.Errorf("usuário %q não está na lista de acesso autorizado", claims.Email)
	}
	return &claims, nil
}
