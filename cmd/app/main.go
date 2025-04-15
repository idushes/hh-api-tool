package main

import (
	"fmt"
	"log"

	"github.com/idushes/hh-api-tool/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Создаем новый MCP сервер
	s := server.NewMCPServer(
		"HeadHunter API Tools",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Регистрируем все инструменты
	tools.RegisterTools(s)

	// Запускаем сервер
	log.Println("Запуск HeadHunter API Tools MCP сервера...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Ошибка сервера: %v\n", err)
	}
}
