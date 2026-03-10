package embeddedadapter

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type EmbeddedFrontendServer struct {
	address string
}

func NewEmbeddedFrontendServer(address string) *EmbeddedFrontendServer {
	return &EmbeddedFrontendServer{
		address: address,
	}
}

func (s *EmbeddedFrontendServer) Start(ctx context.Context) error {
	if len(embeddedWebBinary) == 0 {
		return nil
	}

	pattern := "aidoris-web-*"
	binaryPath := ""

	if runtime.GOOS == "windows" {
		localAppData := os.Getenv("LOCALAPPDATA")

		if localAppData == "" {
			// Fallback to temporary file if LOCALAPPDATA is not set.
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
	command := exec.CommandContext(ctx, binaryPath)
	command.Env = append(os.Environ(), "PORT=55001")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		return err
	}

	<-ctx.Done()

	_ = command.Process.Kill()

	return nil
}

