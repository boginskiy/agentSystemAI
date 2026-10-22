package main

import (
	"agentTz/models"
	"agentTz/ollama"
	"agentTz/request"
	"agentTz/response"
	"agentTz/tools"
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const NameModel = "qwen3-8b-8k"

var StopWorking = func() { fmt.Fprintln(os.Stdout, "\nЗавершение работы...") }

// Dialog
var Dialog []models.ChatMessage
var ChatReq *request.ChatRequest

func main() {
	client := ollama.NewClient("")
	ctx := context.Background()

	inputCh := make(chan string)                  // Канал для ввода пользователя
	outputCh := make(chan *response.ChatResponse) // Канал для вывода агента
	errCh := make(chan error, 1)                  // Канал с ошибками

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	ctxExit, cancel := context.WithCancel(ctx)
	defer cancel()

	// Горутина чтения: не блокирует основной цикл
	go func() {
		defer close(inputCh)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			select {
			case <-ctxExit.Done():
				return
			case inputCh <- scanner.Text():
			}
		}
	}()

	fmt.Fprintln(os.Stdout, "Агент запущен. Для выхода нажмите Ctrl+C.")

	// Горутина запроса в модель
	go func() {
		defer cancel()        // выполнится последним
		defer close(errCh)    // выполнится вторым
		defer close(outputCh) // выполнится первым

		for {
			select {
			case <-ctxExit.Done():
				return

			case userInput, ok := <-inputCh:
				if !ok || userInput == "exit" {
					return
				}

				userMsg := models.ChatMessage{Role: "user", Content: userInput}
				output, err := DoRequest(ctxExit, client, userMsg)

				if err != nil {
					select {
					case errCh <- err:
					case <-ctxExit.Done():
						return
					}
					continue
				}

				select {
				case outputCh <- output:
				case <-ctxExit.Done():
					return
				}
			}
		}
	}()

	// Основной поток. Чтение ответа модели
loop:
	for {
		select {
		case err, ok := <-errCh:
			if !ok {
				break loop
			}
			fmt.Fprintln(os.Stderr, "Ошибка агента:", err) // Error

		case out, ok := <-outputCh:
			if !ok {
				break loop
			}
			fmt.Fprintln(os.Stdout, out.Choices[0].Message.Content) // Content

		case <-ctxExit.Done():
			break loop
		}
	}

	StopWorking()
}

func DoRequest(ctx context.Context, client *ollama.Client, msg models.ChatMessage) (*response.ChatResponse, error) {
	// New Dialog
	if len(Dialog) == 0 {
		Dialog = NewDialog()
	}
	Dialog = append(Dialog, msg)

	// Tools
	toolList := []models.Tool{tools.TimeTool()} //

	const maxSteps = 5

	for step := 1; step <= maxSteps; step++ {
		req := request.ChatRequest{
			Model:    NameModel,
			Messages: Dialog,
			Tools:    toolList,
			Stream:   false,
			Think:    false, // отключаем размышления
		}

		ctxChat, cancelChat := context.WithTimeout(ctx, 120*time.Second)
		output, err := client.Chat(ctxChat, req)
		cancelChat()

		if err != nil {
			return nil, err
		}

		msg := output.Choices[0].Message

		// Добавляем ответ ассистента в историю.
		Dialog = append(Dialog, msg)

		// Если модель не просит инструментов — это финальный ответ.
		if len(msg.ToolCalls) == 0 {
			return output, nil
		}

		// Модель просит вызвать инструменты. Выполняем каждый.
		for _, call := range msg.ToolCalls {
			result := tools.ExecuteTool(call.Function.Name, call.Function.Arguments)
			Dialog = append(Dialog, models.ChatMessage{
				Role:       "tool",
				Content:    result,
				ToolCallID: call.ID, // важно на /v1: модель ожидает связь tool-ответа с вызовом
			})
		}
	}

	return nil, nil
}

func NewDialog() []models.ChatMessage {
	chatMsg := make([]models.ChatMessage, 0, 5)

	systemPrompt := fmt.Sprintln(
		// "You are a helpful assistant. If the question concerns time or date.",
		// "Be sure to use the get_current_time tool. Don’t make up the time.",
		"",
	)
	// systemPrompt := fmt.Sprintln(
	// 	"Ты — полезный ассистент. Если вопрос касается времени или даты, ",
	// 	"Обязательно используй инструмент get_current_time. Не выдумывай время.")

	systemMsg := models.ChatMessage{Role: "system", Content: systemPrompt}
	chatMsg = append(chatMsg, systemMsg)
	return chatMsg
}

// TODO
// Модель с каждым разом получает все больше и больше сообщений? Это и есть контекст, который может держать модель?
// Переходи на llama.cpp с Go-обёрткой
