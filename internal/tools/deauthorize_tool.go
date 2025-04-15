package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/idushes/hh-api-tool/internal/service"
	"github.com/mark3labs/mcp-go/mcp"
)

// DeauthorizeTool представляет инструмент для деактивации токена HeadHunter API
type DeauthorizeTool struct {
	hhAuthClient *service.HeadHunterAuthClient
}

// NewDeauthorizeTool создает новый инструмент деактивации токена
func NewDeauthorizeTool() *DeauthorizeTool {
	return &DeauthorizeTool{
		hhAuthClient: service.NewHeadHunterAuthClient(),
	}
}

// GetToolDefinition возвращает определение инструмента MCP
func (t *DeauthorizeTool) GetToolDefinition() mcp.Tool {
	return mcp.NewTool("deauthorize",
		mcp.WithDescription(`Деактивация токена доступа API HeadHunter.

OAuth2 Token Invalidation Information (EN):
- This tool is designed to invalidate user tokens, not application tokens
- According to HeadHunter API documentation: "Application tokens obtained with client_credentials cannot be invalidated through this method"
- This tool will only work with tokens obtained through user-based authentication flows (not implemented in this tool set)
- If you need to stop using an application token (obtained via the authorize tool), simply stop using it
- For security best practices, application tokens should be stored securely and rotated regularly
`),
		mcp.WithString("access_token",
			mcp.Required(),
			mcp.Description("Токен доступа, который нужно деактивировать (неприменимо к токенам приложения)"),
		),
	)
}

// HandleRequest обрабатывает запрос на деактивацию токена
func (t *DeauthorizeTool) HandleRequest(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	accessToken := request.Params.Arguments["access_token"].(string)

	// Проверка, что токен не пустой
	if accessToken == "" {
		return mcp.NewToolResultError("Токен доступа обязателен"), nil
	}

	// Попытка деактивировать токен
	err := t.hhAuthClient.RevokeToken(accessToken)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка деактивации токена: %v", err)), nil
	}

	// Формирование успешного ответа
	result := map[string]interface{}{
		"status":  "success",
		"message": "Токен успешно деактивирован",
	}

	// Преобразование результата в строку JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка кодирования результата: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonResult)), nil
}
