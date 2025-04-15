package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools регистрирует все инструменты в сервере MCP
func RegisterTools(s *server.MCPServer) {
	// Регистрация инструмента авторизации
	authTool := NewAuthTool()
	s.AddTool(authTool.GetToolDefinition(), authTool.HandleRequest)

	// Регистрация инструмента деактивации токена
	deauthTool := NewDeauthorizeTool()
	s.AddTool(deauthTool.GetToolDefinition(), deauthTool.HandleRequest)

	// Регистрация инструмента получения информации о пользователе
	userInfoTool := NewUserInfoTool()
	s.AddTool(userInfoTool.GetToolDefinition(), userInfoTool.HandleRequest)

	// Здесь будут регистрироваться другие инструменты
	// ...
}
