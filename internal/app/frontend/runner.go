package frontend

import (
	"context"

	outputfrontend "github.com/pranavskurup/aidoris-go/internal/ports/output/frontend"
	inputfrontend "github.com/pranavskurup/aidoris-go/internal/ports/input/frontend"
)

// Runner implements the FrontendRunner input port by delegating to the provided FrontendServer.
type Runner struct{}

// NewRunner returns a Runner that satisfies inputfrontend.FrontendRunner.
func NewRunner() inputfrontend.FrontendRunner {
	return &Runner{}
}

// Run starts the given frontend server until ctx is cancelled.
func (r *Runner) Run(ctx context.Context, server outputfrontend.FrontendServer) error {
	return server.Start(ctx)
}
