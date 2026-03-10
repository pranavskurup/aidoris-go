package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	srcDistDir = "web/apps/web/dist"
	dstDistDir = "internal/ports/frontend/embeddedadapter/dist"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	srcPath := filepath.Join(workingDir, srcDistDir)
	dstPath := filepath.Join(workingDir, dstDistDir)

	if err := ensureDirExists(srcPath); err != nil {
		log.Fatalf("source dist directory not available at %q: %v", srcPath, err)
	}

	if err := os.RemoveAll(dstPath); err != nil {
		log.Fatalf("failed to remove existing destination directory %q: %v", dstPath, err)
	}

	if err := copyDir(srcPath, dstPath); err != nil {
		log.Fatalf("failed to copy dist directory from %q to %q: %v", srcPath, dstPath, err)
	}

	fmt.Printf("Copied web dist from %q to %q\n", srcPath, dstPath)
}

func ensureDirExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", path)
	}

	return nil
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
				return err
			}

			return nil
		}

		if err := copyFile(path, targetPath, info.Mode()); err != nil {
			return err
		}

		return nil
	})
}

func copyFile(src, dst string, mode os.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

