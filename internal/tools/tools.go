package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools регистрирует все инструменты в сервере MCP
func RegisterTools(s *server.MCPServer) {
	// Регистрация инструмента авторизации
	authTool := NewAuthTool()
	s.AddTool(authTool.GetToolDefinition(), authTool.HandleRequest)

	// Здесь будут регистрироваться другие инструменты
	// ...
}
