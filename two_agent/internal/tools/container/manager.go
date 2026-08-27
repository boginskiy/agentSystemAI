package container

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type DManager struct {
	Path       string
	DCommander DockerCommander
}

// NewDManager.
func NewDManager(path string) (*DManager, error) {
	// Проверка наличия файла для создания контейнера
	_, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("file not found for up docker container %w", err)
	}

	// Определяем типы команд, которые будем вызывать в методах.
	dCommander, err := choiseCommander(path)
	if err != nil {
		return nil, err
	}

	return &DManager{
		Path:       path,
		DCommander: dCommander,
	}, nil
}

func choiseCommander(path string) (DockerCommander, error) {
	pathArr := strings.Split(path, "/")
	nameFile := pathArr[len(pathArr)-1]

	switch nameFile {
	case "docker-compose.yml":
		// Name of Image
		nameImage := "postgres:16-alpine"
		return NewComposeFile(nameImage, path), nil

	case "Dockerfile":
		// Name of Image
		nameImage := "d_postgres"
		return NewDockerFile(nameImage, path), nil

	default:
		return nil, fmt.Errorf("docker file is %s. It is not defined ", nameFile)
	}
}

func (d *DManager) Up(ctx context.Context) (string, error) {
	err := d.DCommander.Up(ctx)
	if err != nil {
		return "", err
	}
	return d.DCommander.GetID(ctx)
}

func (d *DManager) Status(ctx context.Context) (string, error) {
	return d.DCommander.GetStatus(ctx)
}

func (d *DManager) Down(ctx context.Context) (string, error) {
	d.DCommander.Down(ctx)
}

// TODO
// Проверка запущен ли контейнер уже, если да, его надо остановить и перезапустить
// или настроить логику свою...

// Если контейнер уже создан и мы его пересоздаем, ошибки нет, но контейнер остается со старой версией
