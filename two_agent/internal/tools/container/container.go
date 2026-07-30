package container

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli"
)

var StorePath = "two_agent/store"
var FileName = "docker-compose.yml"
var TimeOut = time.Duration(time.Second * 5)

type Container struct {
	Command  string
	Docker   Docker
	Formater cli.Formater
}

func NewContainer(command string, formater cli.Formater) *Container {
	return &Container{
		Command:  command,
		Docker:   NewDockerManager(),
		Formater: formater,
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
		c.Formater.LineMess("There is extra conditions with a start of container")

	} else {
		// Поступил запрос на испольнение команды без дополнительных условий.
		// Пример. "/create container
		return c.createContainer(ctx)
	}

	return nil
}

// Поженить контейнер! Сделать максимальную абстракцию.
// Проработать архитектуру.

func (c *Container) createContainer(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	// Проверка наличия Dockerfile файла в папке "store".
	if isfile := c.checkFile(FileName); isfile {
		cmd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(cmd, StorePath, FileName)

		c.Formater.LineMess("Start the docker container...")

		errChan := make(chan error, 1)
		go func(errCh chan error) {
			errChan <- c.Docker.Up(ctx, path)
			close(errChan)
		}(errChan)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			if err != nil {
				// При развертывании контейнера произошла ошибка.
				// Проверяем был ли при этом создан контейнер.

				return err
			}
		}

		return
		// TODO Делать развертывание контейнера. Обращение к LLM и т.п.
	}

	// TODO. Пока заглушка и возвращаем ошибку.
	// Dockerfile отсутствует.
	// Нужно генерировать через GigaChat API,делать запрос с соответствующим контекстом.

	c.Formater.LineMess("LLM is not connected. Not start container")
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
