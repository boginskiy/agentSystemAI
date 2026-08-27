package container

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type ComposeFile struct {
	Image    string
	Name     string
	Path     string
	Commands map[string][]string
}

func NewComposeFile(image string, path string) *ComposeFile {
	pathArr := strings.Split(path, "/")
	nameFile := pathArr[len(pathArr)-1]
	return &ComposeFile{
		Image: image,
		Path:  path,
		Name:  nameFile,
		Commands: map[string][]string{
			"up":     {"compose", "up", "-d"},
			"id":     {"compose", "ps", "-q"},
			"status": {"compose", "ps"},
			"down":   {"stop"},
		},
	}
}

func (c *ComposeFile) Up(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "docker", c.Commands["up"]...)
	cmd.Dir = filepath.Dir(c.Path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s up failed: %w\nOutput: %s", c.Name, err, output)
	}
	return nil
}

func (c *ComposeFile) Down(ctx context.Context) error {
	id, err := c.GetID(ctx)
	if err != nil {
		return err
	}

	c.Commands["down"] = append(c.Commands["down"], id)

	cmd := exec.CommandContext(ctx, "docker", c.Commands["down"]...)
	cmd.Dir = filepath.Dir(c.Path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s down failed: %w\nOutput: %s", c.Name, err, output)
	}
	return nil
}

func (c *ComposeFile) GetID(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "docker", c.Commands["id"]...)
	cmd.Dir = filepath.Dir(c.Path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s container ID failed: %w\nOutput: %s", c.Name, err, output)
	}
	// Сокращаем полный ID 68dcbd8bf311602fb46bc619c8c57cd3cf29185458c1782492c61e951e938be3
	//      на короткий ID 68dcbd8bf311
	return string(output)[:12], nil
}

func (c *ComposeFile) GetStatus(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "docker", c.Commands["status"]...)
	cmd.Dir = filepath.Dir(c.Path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s status failed: %w\nOutput: %s", c.Name, err, output)
	}
	return string(output), nil
}

func (c *ComposeFile) GetName() string {
	return c.Name
}
