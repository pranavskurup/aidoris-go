package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/pranavskurup/aidoris-go/internal/ports/frontend"
	"github.com/pranavskurup/aidoris-go/internal/ports/frontend/devadapter"
	"github.com/pranavskurup/aidoris-go/internal/ports/frontend/embeddedadapter"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var server frontend.FrontendServer

	if isDevMode() {
		server = devadapter.NewDevFrontendServer("./web")
	} else {
		server = embeddedadapter.NewEmbeddedFrontendServer(":3000")
	}

	if err := server.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("frontend server error: %v", err)
	}
}

func isDevMode() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	// When run via "go run ./cmd/aidoris", the binary lives in a temp path like .../go-build<hash>/b001/exe/aidoris
	return strings.Contains(exe, "go-build")
}
