package container

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

type Command string
type Args []string

type DockerFile struct {
	docker   string
	Commands map[Command]Args
	Path     string
}

// NewDockerFile. command: "docker", args: []string{"compose", "up", "-d"}
func NewDockerFile(path string) *DockerFile {
	return &DockerFile{
		docker: "docker",
		Path:   path,
		Commands: map[Command]Args{
			"build":         []string{"build", "-t", "my-image", "."}, // Dockerfile
			"run":           []string{"run", "my-image"},              // Dockerfile
			"compose up":    []string{"compose", "up", "-d"},          // docker compose
			"compose build": []string{"compose", "build"},             // docker compose
		},
	}
}

func (d *DockerFile) Execute(ctx context.Context, path string, arg ...string) error {
	cmd := exec.CommandContext(ctx, d.docker, arg...)
	cmd.Dir = filepath.Dir(path)
}

func (d *DockerFile) Up(ctx context.Context, path string) error {
	arr := []string{"docker", "compose", "up", "-d"}

	cmd := exec.CommandContext(ctx, arr...)

	cmd.Dir = filepath.Dir(path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s up failed: %w\nOutput: %s", d.Name, err, output)
	}

	status, err := d.checkStatus(ctx, path)
	if err != nil {
		return err
	}

	fmt.Printf("Compose up output: %s, status: %s\n", output, status)
	return nil
}

func (d *DockerFile) checkStatus(ctx context.Context, path string) (string, error) {
	cmd := exec.CommandContext(ctx, "docker", "compose", "ps")
	cmd.Dir = filepath.Dir(path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s ps failed: %w", d.Name, err)
	}

	return string(output), nil
}
