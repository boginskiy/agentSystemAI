package container

import (
	"fmt"
	"os"
	"path/filepath"
)

var StorePath = "two_agent/store"
var FileName = "Dockerfile"

type Container struct {
	Command string
}

func NewContainer(command string) *Container {
	return &Container{
		Command: command,
	}
}

func (c *Container) CallCommand() string {
	return c.Command
}

func (c *Container) Do(conditions []string) error {
	// Проверка, что есть дополнительные условия в введенной команде
	// Пример. "/create container: <описание параметров>"
	if len(conditions) > 1 {
		_ = conditions[1:] // extraCond
		fmt.Println("Not Start Container - 2")

	} else {
		// Поступил запрос на испольнение команды без дополнительных условий.
		// Пример. "/create container
		c.createContainer()
	}

	return nil
}

func (c *Container) createContainer() error {
	// Проверка наличия Dockerfile файла в папке "store".
	if isfile := c.checkFile(FileName); isfile {
		// Запускаем Dockerfile.

		fmt.Println("Start Container")
		return nil

		// TODO Делать развертывание контейнера. Обращение к LLM и т.п.
	}

	fmt.Println("Not Start Container")

	return nil

	// Dockerfile отсутствует.
	// TODO. Нужно генерировать через GigaChat API,делать запрос с соответствующим контекстом.

}

func (c *Container) checkFile(name string) bool {
	fullPath := filepath.Join(StorePath, name)

	_, err := os.Stat(fullPath)
	if err != nil {
		return false
	}
	return true
}

// Генерация Dockerfile, при ошибках запуска контейнера, при отсутствии файла, далее логика запуска новоСгенерированного файла
// Проверка, что есть файл для запуска контейнера.
// Запуск контейнера. Проверка что запуск успешен. Проверка БД

// Как понять, если не жестко кодировать, когда модель должна обратиться в tools, а когда использовать генератор текста
