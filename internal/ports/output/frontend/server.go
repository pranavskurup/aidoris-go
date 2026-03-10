package frontend

import "context"

// FrontendServer is the output port for serving the frontend.
// Implementations (output adapters) provide embedded binary or dev server.
type FrontendServer interface {
	Start(ctx context.Context) error
}
