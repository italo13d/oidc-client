package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/italo13d/oidc-client/internal/oidc"
)

type Handlers struct {
	oidc *oidc.Client
}

func New(c *oidc.Client) *Handlers { return &Handlers{oidc: c} }

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	url := h.oidc.AuthCodeURL("state-token")
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handlers) Callback(w http.ResponseWriter, r *http.Request) {
	claims, err := h.oidc.ExchangeAndVerify(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		switch {
		case errors.Is(err, oidc.ErrEmailNotVerified):
			http.Error(w,
				"Falha de autenticação: seu e-mail não foi verificado no provedor.",
				http.StatusForbidden)
		case strings.Contains(err.Error(), "não está na lista"):
			http.Error(w,
				"Falha de autorização: você não tem permissão para acessar este sistema.",
				http.StatusForbidden)
		default:
			http.Error(w,
				"Falha de autenticação: "+err.Error(),
				http.StatusUnauthorized)
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
      <h1>Login bem-sucedido!</h1>
      <p>Olá <strong>%s</strong> (<em>%s</em>), seu acesso está autorizado.</p>
    `, claims.Name, claims.Email)
}
