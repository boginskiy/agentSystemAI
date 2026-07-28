package docker

import (
	"context"
	"fmt"
	"os/exec"
)

type DockerFile struct {
	Name string
}

func NewDockerFile(name string) *DockerFile {
	return &DockerFile{
		Name: name, // docker-compose
	}
}

// TODO!!! Разбираемся с тем как запустить ! Docker-compose
func (d *DockerFile) Up(ctx context.Context, path string) error {
	cmd := exec.CommandContext(ctx, d.Name, "up", "-d")

	fmt.Println(path)
	cmd.Dir = path

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
	cmd := exec.CommandContext(ctx, d.Name, "ps")
	cmd.Dir = path

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s ps failed: %w", d.Name, err)
	}

	return string(output), nil
}
