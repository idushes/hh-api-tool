package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/idushes/hh-api-tool/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Настраиваем обработку сигналов прерывания
	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Создаем новый MCP сервер
	s := server.NewMCPServer(
		"HeadHunter API Tools",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Регистрируем все инструменты
	tools.RegisterTools(s)

	// Запускаем горутину для обработки сигналов
	go func() {
		sig := <-signalChan
		log.Printf("Получен сигнал: %v. Завершаю работу...", sig)

		// Даем время на корректное завершение операций
		time.Sleep(500 * time.Millisecond)

		done <- true
	}()

	// Проверяем наличие переменной окружения MCP_SERVER_PORT
	serverPort := os.Getenv("MCP_SERVER_PORT")

	if serverPort != "" {
		// Запускаем SSE сервер, если указан порт
		sseServer := server.NewSSEServer(s)

		go func() {
			addr := fmt.Sprintf(":%s", serverPort)
			log.Printf("Запуск HeadHunter API Tools в режиме SSE на порту %s...", serverPort)
			if err := sseServer.Start(addr); err != nil && err != http.ErrServerClosed {
				log.Printf("Ошибка SSE сервера: %v\n", err)
				done <- true
			}
		}()

		// Ожидаем сигнал завершения
		<-done

		// Корректно останавливаем SSE сервер
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := sseServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("Ошибка при остановке SSE сервера: %v\n", err)
		}

		log.Println("Сервер остановлен")
	} else {
		// Запускаем сервер в режиме stdio, если порт не указан
		go func() {
			log.Println("Запуск HeadHunter API Tools в режиме stdio...")
			if err := server.ServeStdio(s); err != nil {
				fmt.Printf("Ошибка сервера: %v\n", err)
				done <- true
			}
		}()

		// Ожидаем сигнала завершения
		<-done
		log.Println("Сервер остановлен")
	}
}
