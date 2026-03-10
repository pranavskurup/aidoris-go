package main

import (
	"context"
	"errors"
	"log"
	"os/signal"
	"syscall"

	"github.com/pranavskurup/aidoris-go/internal/adapters/input/cli"
	"github.com/pranavskurup/aidoris-go/internal/app/frontend"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	runner := frontend.NewRunner()
	frontendCLI := cli.NewFrontendCLI(runner)

	if err := frontendCLI.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("frontend error: %v", err)
	}
}
