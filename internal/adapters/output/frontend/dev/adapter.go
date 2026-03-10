package dev

import (
	"context"
	"os"
	"os/exec"
)

// Server implements ports/output/frontend.FrontendServer by running pnpm dev.
type Server struct {
	workingDir string
}

// NewServer returns a frontend server that runs the dev server in workingDir.
func NewServer(workingDir string) *Server {
	return &Server{
		workingDir: workingDir,
	}
}

// Start runs pnpm dev until ctx is cancelled.
func (s *Server) Start(ctx context.Context) error {
	if s.workingDir == "" {
		return nil
	}

	command := exec.CommandContext(ctx, "pnpm", "run", "dev", "--log-order=stream")
	command.Dir = s.workingDir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	return command.Process.Kill()
}
