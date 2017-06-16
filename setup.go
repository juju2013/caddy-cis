package caseinsensitive

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

// plugin setup, parse configuration file and register Handler
func setup(c *caddy.Controller) error {

	for c.Next() {
		if c.NextArg() { // expect noarguement
			return c.ArgErr()
		}
	}
	cfg := httpserver.GetConfig(c)
	mid := func(next httpserver.Handler) httpserver.Handler {
		return CisHandler{Root: cfg.Root, Next: next}
	}
	cfg.AddMiddleware(mid)
	return nil
}
