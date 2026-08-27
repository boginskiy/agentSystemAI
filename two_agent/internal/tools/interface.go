package tools

import "context"

type Toolmaker interface {
	GetCommand() string
	Do(ctx context.Context, conditions []string) error
}

type Creater interface {
}
