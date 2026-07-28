package tools

import "context"

type Toolmaker interface {
	CallCommand() string
	Do(ctx context.Context, conditions []string) error
}

type Creater interface {
}
