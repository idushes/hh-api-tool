package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/idushes/hh-api-tool/internal/service"
	"github.com/mark3labs/mcp-go/mcp"
)

// UserInfoTool представляет инструмент для получения информации о пользователе
type UserInfoTool struct {
	hhUserClient *service.HeadHunterUserClient
}

// NewUserInfoTool создает новый инструмент для получения информации о пользователе
func NewUserInfoTool() *UserInfoTool {
	return &UserInfoTool{
		hhUserClient: service.NewHeadHunterUserClient(),
	}
}

// GetToolDefinition возвращает определение инструмента MCP
func (t *UserInfoTool) GetToolDefinition() mcp.Tool {
	return mcp.NewTool("get_user_info",
		mcp.WithDescription(`Получение информации о текущем пользователе API HeadHunter.

Endpoint: GET /me

Данный инструмент возвращает информацию о текущем пользователе, ассоциированном с предоставленным access token.
Информация включает в себя данные профиля, статусы (соискатель/работодатель/админ) и URLs к связанным ресурсам.
`),
		mcp.WithString("access_token",
			mcp.Required(),
			mcp.Description("Access Token полученный в результате авторизации"),
		),
	)
}

// HandleRequest обрабатывает запрос на получение информации о пользователе
func (t *UserInfoTool) HandleRequest(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	accessToken := request.Params.Arguments["access_token"].(string)

	// Проверка, что токен не пустой
	if accessToken == "" {
		return mcp.NewToolResultError("Access Token обязателен"), nil
	}

	// Получение информации о пользователе из HeadHunter API
	userResp, err := t.hhUserClient.GetCurrentUser(accessToken)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка получения информации о пользователе: %v", err)), nil
	}

	// Преобразование ответа в структуру для вывода
	result := map[string]interface{}{
		"status":  "success",
		"message": "Успешное получение информации о пользователе",
		"user":    userResp,
	}

	// Преобразование результата в строку JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Ошибка кодирования результата: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonResult)), nil
}
