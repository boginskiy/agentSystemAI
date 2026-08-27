package container

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type DockerFile struct {
	Image    string
	Name     string
	Path     string
	Commands map[string][]string
}

func NewDockerFile(image string, path string) *DockerFile {
	pathArr := strings.Split(path, "/")
	nameFile := pathArr[len(pathArr)-1]

	return &DockerFile{
		Image: image,
		Name:  nameFile,
		Path:  path,
		Commands: map[string][]string{
			"up":     {"run", image},
			"id":     {},
			"status": {},
		},
	}
}

func (c *DockerFile) Down(ctx context.Context) error {
	return "", fmt.Errorf("method 'Down' is not implemented in 'DockerFile' struct")
}

func (d *DockerFile) Up(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "docker", d.Commands["up"]...)
	cmd.Dir = filepath.Dir(d.Path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s up failed: %w\nOutput: %s", d.Name, err, output)
	}
	return nil
}

func (d *DockerFile) GetID(ctx context.Context) (string, error) {
	// TODO
	return "", fmt.Errorf("method 'GetID' is not implemented in 'DockerFile' struct")
}

func (d *DockerFile) GetStatus(ctx context.Context) (string, error) {
	// TODO
	return "", fmt.Errorf("method 'GerStatus' is not implemented in 'DockerFile' struct")
}

func (d *DockerFile) GetName() string {
	return d.Name
}
