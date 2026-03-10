package embedded

import (
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Server implements ports/output/frontend.FrontendServer by running the embedded web binary.
type Server struct {
	address string
}

// NewServer returns a frontend server that serves via the embedded binary.
func NewServer(address string) *Server {
	return &Server{
		address: address,
	}
}

// Start runs the embedded web binary until ctx is cancelled.
func (s *Server) Start(ctx context.Context) error {
	if len(embeddedWebBinary) == 0 {
		return nil
	}

	pattern := "aidoris-web-*"
	binaryPath := ""

	if runtime.GOOS == "windows" {
		localAppData := os.Getenv("LOCALAPPDATA")

		if localAppData == "" {
			tempFile, err := os.CreateTemp("", pattern+".exe")
			if err != nil {
				return err
			}

			if _, err := tempFile.Write(embeddedWebBinary); err != nil {
				tempFile.Close()
				return err
			}

			if err := tempFile.Chmod(0o700); err != nil {
				tempFile.Close()
				return err
			}

			if err := tempFile.Close(); err != nil {
				return err
			}

			binaryPath = tempFile.Name()
		} else {
			appDir := filepath.Join(localAppData, "AidoriS")
			if err := os.MkdirAll(appDir, 0o755); err != nil {
				return err
			}

			binaryPath = filepath.Join(appDir, "aidoris-web.exe")
			if err := os.WriteFile(binaryPath, embeddedWebBinary, 0o700); err != nil {
				return err
			}
		}
	} else {
		tempFile, err := os.CreateTemp("", pattern)
		if err != nil {
			return err
		}

		if _, err := tempFile.Write(embeddedWebBinary); err != nil {
			tempFile.Close()
			return err
		}

		if err := tempFile.Chmod(0o700); err != nil {
			tempFile.Close()
			return err
		}

		if err := tempFile.Close(); err != nil {
			return err
		}

		binaryPath = tempFile.Name()
	}

	host, port, err := net.SplitHostPort(s.address)
	if err != nil {
		return err
	}
	if host == "" {
		host = "localhost"
	}

	command := exec.CommandContext(ctx, binaryPath)
	command.Env = append(os.Environ(), "HOST="+host, "PORT="+port)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	_ = command.Process.Kill()

	return nil
}
