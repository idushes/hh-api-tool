package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/idushes/hh-api-tool/internal/service"
	"github.com/mark3labs/mcp-go/mcp"
)

// AuthTool представляет инструмент для авторизации через HeadHunter API
type AuthTool struct {
	hhAuthClient *service.HeadHunterAuthClient
}

// NewAuthTool создает новый инструмент авторизации
func NewAuthTool() *AuthTool {
	return &AuthTool{
		hhAuthClient: service.NewHeadHunterAuthClient(),
	}
}

// GetToolDefinition возвращает определение инструмента MCP
func (t *AuthTool) GetToolDefinition() mcp.Tool {
	return mcp.NewTool("authorize",
		mcp.WithDescription("Авторизация в API HeadHunter используя учетные данные клиента"),
		mcp.WithString("client_id",
			mcp.Required(),
			mcp.Description("Client ID полученный из HeadHunter API"),
		),
		mcp.WithString("client_secret",
			mcp.Required(),
			mcp.Description("Client Secret полученный из HeadHunter API"),
		),
	)
}

// HandleRequest обрабатывает запрос на авторизацию
func (t *AuthTool) HandleRequest(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientID := request.Params.Arguments["client_id"].(string)
	clientSecret := request.Params.Arguments["client_secret"].(string)

	// Для тестирования можно использовать переменные окружения
	if envClientID := os.Getenv("CLIENT_ID"); envClientID != "" && clientID == "use_env" {
		clientID = envClientID
	}

	if envClientSecret := os.Getenv("CLIENT_SECRET"); envClientSecret != "" && clientSecret == "use_env" {
		clientSecret = envClientSecret
	}

	// Проверка, что учетные данные не пустые
	if clientID == "" || clientSecret == "" {
		return mcp.NewToolResultError("Client ID и Client Secret обязательны"), nil
	}

	// Получение токена из HeadHunter API
	tokenResp, err := t.hhAuthClient.GetClientCredentialsToken(clientID, clientSecret)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка авторизации в HeadHunter API: %v", err)), nil
	}

	// Преобразование ответа с токеном в строку JSON
	result := map[string]interface{}{
		"status":       "success",
		"message":      "Успешная авторизация в HeadHunter API",
		"access_token": tokenResp.AccessToken,
		"token_type":   tokenResp.TokenType,
		"expires_in":   tokenResp.ExpiresIn,
	}

	// Преобразование результата в строку JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка кодирования результата: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonResult)), nil
}
