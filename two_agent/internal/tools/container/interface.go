package container

import "context"

type DockerManager interface {
	Status(ctx context.Context) (string, error)
	Up(ctx context.Context) (string, error)
	Down(ctx context.Context) error
}

type DockerCommander interface {
	GetStatus(ctx context.Context) (string, error)
	GetID(ctx context.Context) (string, error)
	GetName() string
	Down(ctx context.Context) error
	Up(ctx context.Context) error
}
