package devadapter

import (
	"context"
	"os"
	"os/exec"
)

type DevFrontendServer struct {
	workingDir string
}

func NewDevFrontendServer(workingDir string) *DevFrontendServer {
	return &DevFrontendServer{
		workingDir: workingDir,
	}
}

func (s *DevFrontendServer) Start(ctx context.Context) error {
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
