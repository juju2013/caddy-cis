package caseinsensitive

import (
  "net/http"
  "strings"
  "os"
  "path"
  "path/filepath"
  
  "github.com/mholt/caddy"
  "github.com/mholt/caddy/caddyhttp/httpserver"
)

type CisHandler struct {
  Root string
	Next httpserver.Handler
}

func init() {
	caddy.RegisterPlugin("caseinsensitive", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func (h CisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
  p := httpserver.SafePath(h.Root, r.URL.Path)
  if _, err := os.Stat(p); err == nil {
    // file exists, nothing to do
    return h.Next.ServeHTTP(w, r)
  }
  elements := strings.Split(p[len(h.Root):], "/")
  progress:=h.Root
  for i, d := range(elements) {
    if len(d) < 1 {
      continue
    }
    if _, err := os.Stat(filepath.Join(progress, d)); err == nil {
      progress = filepath.Join(progress, d)
      continue // found, change nothing
    }
    // not found
    if i < len(elements) - 1 {
      d=strings.ToUpper(d)
      if _, err := os.Stat(filepath.Join(progress, d)); err == nil {
        progress = filepath.Join(progress, d)
        elements[i]=d
        continue
      }
      d=strings.ToLower(d)
      if _, err := os.Stat(filepath.Join(progress, d)); err == nil {
        progress = filepath.Join(progress, d)
        elements[i]=d
        continue
      }
      // lower+UPPER case not found, stop processing (will 404)
      break
    } else {
      // procession file, try (lower+UPPER name)+(lower+UPPER extension)
      for _, lowUPname := range(mixCaseFile(d)) {
        if _, err := os.Stat(filepath.Join(progress, lowUPname)); err == nil {
          elements[i]=lowUPname
          break
        }
      }
    }
  }
  // reassemble elements to path
  p=""
  for _, d := range(elements) {
    p=path.Join(p, d)
  }
  r.URL.Path=p
	return h.Next.ServeHTTP(w, r)
}

func mixCaseFile(f string) []string {
  l := strings.LastIndex(f, ".")
  return []string {
    f,
    strings.ToLower(f),
    strings.ToUpper(f),
    strings.ToLower(f[:l])+strings.ToUpper(f[l:]),
    strings.ToUpper(f[:l])+strings.ToLower(f[l:]),
  }
}
