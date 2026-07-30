package container

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type DockerManager struct {
	commands map[string]Docker
}

func NewDockerManager() *DockerManager {
	dockerCompose := NewDockerFile("docker compose")
	docker := NewDockerFile("docker")

	return &DockerManager{
		commands: map[string]Docker{
			"docker compose": dockerCompose,
			"docker":         docker}, // TODO Команда не валидна.
	}
}

func (d *DockerManager) Up(ctx context.Context, path string) error {
	// Проверка наличия файла для создания контейнера
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("file not found for up container %w", err)
	}

	pathArr := strings.Split(path, "/")
	if len(pathArr) == 0 {
		return fmt.Errorf("path is not defined for up container")
	}

	nameFile := pathArr[len(pathArr)-1]

	switch nameFile {
	case "docker-compose.yml":
		if d, ok := d.commands["docker compose"]; ok {
			err := d.Up(ctx, path)
			if err != nil {
				return err
			}
			return nil
		}

	case "Dockerfile":
		// TODO. Нет логики.
		return nil
	default:
		return fmt.Errorf("name of file is not valid %s", nameFile)
	}
	return nil
}
