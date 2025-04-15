package main

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// Демонстрационное использование библиотеки mcp-go
	fmt.Println("Версия протокола MCP:", mcp.LATEST_PROTOCOL_VERSION)
	log.Println("Структура проекта с зависимостью mcp-go создана успешно")
}
