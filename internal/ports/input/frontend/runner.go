package frontend

import (
	"context"

	outputfrontend "github.com/pranavskurup/aidoris-go/internal/ports/output/frontend"
)

// FrontendRunner is the input port for running the frontend application.
// Input adapters (e.g. CLI) call Run with the chosen FrontendServer (output adapter).
type FrontendRunner interface {
	Run(ctx context.Context, server outputfrontend.FrontendServer) error
}
