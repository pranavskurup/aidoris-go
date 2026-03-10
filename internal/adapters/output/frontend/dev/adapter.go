package dev

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
)

// Server implements ports/output/frontend.FrontendServer by running pnpm dev.
type Server struct {
	workingDir string
	address    string
}

// NewServer returns a frontend server that runs the dev server in workingDir at address (e.g. "localhost:55001").
func NewServer(workingDir, address string) *Server {
	return &Server{
		workingDir: workingDir,
		address:    address,
	}
}

// Start runs pnpm dev until ctx is cancelled.
func (s *Server) Start(ctx context.Context) error {
	if s.workingDir == "" {
		return nil
	}

	host, port, err := net.SplitHostPort(s.address)
	if err != nil {
		return fmt.Errorf("frontend address: %w", err)
	}
	if host == "" {
		host = "localhost"
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return fmt.Errorf("port not free: %w", err)
	}
	_ = listener.Close()

	command := exec.CommandContext(ctx, "pnpm", "run", "dev", "--log-order=stream", "--", "--port="+port, "--host="+host)
	command.Dir = s.workingDir
	command.Env = append(os.Environ(), "HOST="+host, "PORT="+port, "AIDORIS_WEB_DEV_PORT="+port)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	return command.Process.Kill()
}
