package cli

import (
	"context"
	"os"
	"strings"

	"github.com/pranavskurup/aidoris-go/internal/adapters/output/frontend/dev"
	"github.com/pranavskurup/aidoris-go/internal/adapters/output/frontend/embedded"
	inputfrontend "github.com/pranavskurup/aidoris-go/internal/ports/input/frontend"
	outputfrontend "github.com/pranavskurup/aidoris-go/internal/ports/output/frontend"
)

// FrontendCLI is the input adapter that runs the frontend (dev or embedded) via the FrontendRunner use case.
type FrontendCLI struct {
	runner inputfrontend.FrontendRunner
}

// NewFrontendCLI returns a CLI adapter that uses the given runner.
func NewFrontendCLI(runner inputfrontend.FrontendRunner) *FrontendCLI {
	return &FrontendCLI{
		runner: runner,
	}
}

// Run selects the frontend server (dev or embedded) and runs it until ctx is cancelled.
func (c *FrontendCLI) Run(ctx context.Context) error {
	var server outputfrontend.FrontendServer
	if isDevMode() {
		server = dev.NewServer("./web", ":55001")
	} else {
		server = embedded.NewServer(":55001")
	}
	return c.runner.Run(ctx, server)
}

func isDevMode() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	return strings.Contains(exe, "go-build")
}
