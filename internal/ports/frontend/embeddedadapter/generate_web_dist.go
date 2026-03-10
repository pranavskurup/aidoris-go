//go:build ignore

package embeddedadapter

// This file documents the expected directory structure for the embedded web assets.
//
// Before building the Go binary for distribution, run the web build so that
// the contents of `web/apps/web/dist` are available. Then copy or link that
// directory into `internal/ports/frontend/embeddedadapter/dist` so that
// the embed directive in `embedded_adapter.go` can include the correct files.
