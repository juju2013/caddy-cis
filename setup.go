package caseinsensitive

import (

  "github.com/mholt/caddy"
  "github.com/mholt/caddy/caddyhttp/httpserver"
)

func setup(c *caddy.Controller) error {

  c.Next() // Skip "caseinsensitive" literal

    if len(c.RemainingArgs()) > 1 {
      return c.Err("Unexpected value " + c.Val())
  }

	cfg := httpserver.GetConfig(c)
  mid := func(next httpserver.Handler) httpserver.Handler {
    return CisHandler{Next: next}
  }
  cfg.AddMiddleware(mid)
  return nil
}
