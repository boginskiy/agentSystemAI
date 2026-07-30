package container

import "context"

type Docker interface {
	Up(ctx context.Context, path string) error
	// CheckStatus(ctx context.Context, path string) (string, error)
}
