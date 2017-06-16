package caseinsensitive

import (
	"testing"

	"github.com/mholt/caddy"
  "github.com/mholt/caddy/caddyhttp/httpserver"
)

func TestSetup(t *testing.T) {
	c := caddy.NewTestController("http", `caseinsensitive`)
	err := setup(c)
	if err != nil {
		t.Errorf("Expected no errors, but got: %v", err)
	}
	mids := httpserver.GetConfig(c).Middleware()
	if len(mids) == 0 {
		t.Fatal("Expected middleware, got 0 instead")
	}

	handler := mids[0](httpserver.EmptyNext)
	myHandler, ok := handler.(CisHandler)
	if !ok {
		t.Fatalf("Expected handler to be type CisHandler, got: %#v", handler)
	}

	if !httpserver.SameNext(myHandler.Next, httpserver.EmptyNext) {
		t.Error("'Next' field of handler was not set properly")
	}
}
