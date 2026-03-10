package frontend

import "context"

type FrontendServer interface {
	Start(ctx context.Context) error
}

