package container

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boginskiy/agentSystemAI/two_agent/internal/cli"
)

//                                   Commands                                      //
// "/create container: <описание параметров>" || "/create ctr: <описание параметров>"
// "/status container: <описание параметров>" || "/status ctr: <описание параметров>"
// "/delete container: <описание параметров>" || "/delete ctr: <описание параметров>"

type Container struct {
	ID            string
	Command       string
	DockerManager DockerManager
	Formater      cli.Formater
}

func NewContainer(command string, formater cli.Formater) (*Container, error) {
	cmd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dockerManager, err := NewDManager(filepath.Join(cmd, StorePath, FileName))
	if err != nil {
		// DockerManager is nill without return error
		fmt.Println(formater.LineMessWithErr("error creating Docker Manager", err))
	}

	return &Container{
		Command:       command,
		DockerManager: dockerManager,
		Formater:      formater,
	}, nil
}

func (c *Container) GetCommand() string {
	return c.Command
}

func (c *Container) Do(ctx context.Context, conditions []string) error {
	if c.DockerManager == nil {
		return fmt.Errorf("DManager is not created")
	}

	if len(conditions) == 0 {
		return fmt.Errorf("there is not command for a start of container")
	}

	// TODO
	// Проверка, что есть дополнительные условия в введенной команде
	// Пример. "/create container: <описание параметров>"
	if len(conditions) > 1 {
		_ = conditions[1:] // extraCond
		c.Formater.LineMess("There is extra conditions with a start of container")
		return nil
	}

	// Поступил запрос на исполнение команды без дополнительных условий.
	// Пример. "/create container" || "/create ctr"

	switch conditions[0] {
	case "/create container", "/create ctr":
		return c.createContainer(ctx)

	case "/status container", "/status ctr":
		return c.statusContainer(ctx)

	case "/delete container", "/delete ctr":
		return c.deleteContainer(ctx)
	}

	return nil
}

// TODO!
func (c *Container) deleteContainer(ctx context.Context) error {
	err := c.DockerManager.Down(ctx)
	if err == nil {
		c.Formater.LineMess("")
	}
}

func (c *Container) statusContainer(ctx context.Context) error {
	status, err := c.DockerManager.Status(ctx)
	c.Formater.LineMess(status)
	return err
}

func (c *Container) createContainer(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	c.Formater.LineMess("start the docker container...")

	errChan := make(chan error, 1)

	go func(errCh chan error) {
		defer close(errChan)
		id, err := c.DockerManager.Up(ctx)
		c.ID = id // присваиваем внешней переменной
		errChan <- err
	}(errChan)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		// Контейнер успешно создан
		if err == nil {
			fmt.Printf("docker container %s has been created successfully\n", c.ID)
			return nil

		} else {
			// Ошибка при сооздании контейнера
			fmt.Printf("docker container has not been created\n")
			return err
		}

	}
	// TODO. Пока заглушка и возвращаем ошибку.
	// Dockerfile отсутствует.
	// Нужно генерировать через GigaChat API,делать запрос с соответствующим контекстом.
	// return fmt.Errorf("error: not start container %s", FileName)
}

// TODO
// Генерация Dockerfile, при ошибках запуска контейнера, при отсутствии файла, далее логика запуска новоСгенерированного файла
// Проверка, что есть файл для запуска контейнера.
// Запуск контейнера. Проверка что запуск успешен. Проверка БД
// Как понять, если не жестко кодировать, когда модель должна обратиться в tools, а когда использовать генератор текста
