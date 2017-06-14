package caseinsensitive

import (
  "net/http"
  "github.com/mholt/caddy"
  "github.com/mholt/caddy/caddyhttp/httpserver"
)

type CisHandler struct {
	Next httpserver.Handler
}

func init() {
	caddy.RegisterPlugin("caseinsensitive", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func (h CisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	return h.Next.ServeHTTP(w, r)
}
