package container

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boginskiy/agentSystemAI/two_agent/pkg/docker"
)

var StorePath = "two_agent/store"
var FileName = "docker-compose.yml"

type Container struct {
	Command string
	Docker  docker.Docker
}

func NewContainer(command string) *Container {
	return &Container{
		Command: command,
		Docker:  docker.NewDockerManager(),
	}
}

func (c *Container) CallCommand() string {
	return c.Command
}

func (c *Container) Do(ctx context.Context, conditions []string) error {
	// Проверка, что есть дополнительные условия в введенной команде
	// Пример. "/create container: <описание параметров>"
	if len(conditions) > 1 {
		_ = conditions[1:] // extraCond
		fmt.Println("Not Start Container - 2")

	} else {
		// Поступил запрос на испольнение команды без дополнительных условий.
		// Пример. "/create container
		return c.createContainer(ctx)
	}

	return nil
}

func (c *Container) createContainer(ctx context.Context) error {
	// Проверка наличия Dockerfile файла в папке "store".
	if isfile := c.checkFile(FileName); isfile {
		// Запускаем Dockerfile.

		fmt.Println("Start Docker container...")
		path := filepath.Join(StorePath, FileName)

		err := c.Docker.Up(ctx, path)
		return err

		// TODO Делать развертывание контейнера. Обращение к LLM и т.п.
	}

	// TODO. Пока заглушка и возвращаем ошибку.
	// Dockerfile отсутствует.
	// Нужно генерировать через GigaChat API,делать запрос с соответствующим контекстом.

	fmt.Println("Not Start Container")

	return fmt.Errorf("error: not start container %s", FileName)
}

func (c *Container) checkFile(name string) bool {
	fullPath := filepath.Join(StorePath, name)

	_, err := os.Stat(fullPath)
	if err != nil {
		return false
	}
	return true
}

// TODO
// Генерация Dockerfile, при ошибках запуска контейнера, при отсутствии файла, далее логика запуска новоСгенерированного файла
// Проверка, что есть файл для запуска контейнера.
// Запуск контейнера. Проверка что запуск успешен. Проверка БД
// Как понять, если не жестко кодировать, когда модель должна обратиться в tools, а когда использовать генератор текста
