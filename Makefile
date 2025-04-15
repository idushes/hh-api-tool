# Makefile для проекта hh-api-tool

# Переменные проекта
APP_NAME := hh-api-tool
BUILD_DIR := build
VERSION := 1.0.0
MAIN_PATH := cmd/app/main.go

# Go переменные
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOGET := $(GOCMD) get

# Поддерживаемые платформы
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

# Цель по умолчанию
.PHONY: all
all: clean deps build

# Очистка директории сборки
.PHONY: clean
clean:
	@echo "Очистка директории сборки..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)

# Установка зависимостей
.PHONY: deps
deps:
	@echo "Установка зависимостей..."
	@$(GOMOD) tidy

# Запуск тестов
.PHONY: test
test:
	@echo "Запуск тестов..."
	@$(GOTEST) -v ./...

# Сборка для текущей платформы
.PHONY: build
build:
	@echo "Сборка для текущей платформы..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Сборка для всех платформ
.PHONY: build-all
build-all: clean deps
	@echo "Сборка для всех платформ..."
	$(foreach platform,$(PLATFORMS),\
		$(eval GOOS := $(word 1,$(subst /, ,$(platform))))\
		$(eval GOARCH := $(word 2,$(subst /, ,$(platform))))\
		$(eval EXTENSION := $(if $(filter windows,$(GOOS)),.exe,))\
		$(eval PLATFORM_NAME := $(APP_NAME)-$(VERSION)-$(GOOS)-$(GOARCH)$(EXTENSION))\
		@echo "Сборка для $(GOOS)/$(GOARCH)..." \
		&& GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILD_DIR)/$(PLATFORM_NAME) $(MAIN_PATH) || exit 1 \
	;)

# Сборка только для Linux
.PHONY: build-linux
build-linux: clean deps
	@echo "Сборка для Linux..."
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(MAIN_PATH)
	@GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 $(MAIN_PATH)

# Сборка только для macOS
.PHONY: build-darwin
build-darwin: clean deps
	@echo "Сборка для macOS..."
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 $(MAIN_PATH)

# Сборка только для Windows
.PHONY: build-windows
build-windows: clean deps
	@echo "Сборка для Windows..."
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(MAIN_PATH)

# Очистка следов сборки
.PHONY: distclean
distclean: clean
	@echo "Полная очистка..."
	@rm -rf vendor

# Помощь
.PHONY: help
help:
	@echo "HH-API-Tool Makefile"
	@echo ""
	@echo "Доступные команды:"
	@echo "  make             - Сборка для текущей платформы (выполняет clean, deps, build)"
	@echo "  make build       - Сборка для текущей платформы"
	@echo "  make build-all   - Сборка для всех поддерживаемых платформ"
	@echo "  make build-linux - Сборка только для Linux (amd64, arm64)"
	@echo "  make build-darwin - Сборка только для macOS (amd64, arm64)"
	@echo "  make build-windows - Сборка только для Windows (amd64)"
	@echo "  make clean       - Очистка директории сборки"
	@echo "  make deps        - Установка зависимостей"
	@echo "  make test        - Запуск тестов"
	@echo "  make distclean   - Полная очистка, включая vendor"
	@echo "  make help        - Показ этой справки"
	@echo ""
	@echo "Пример использования:"
	@echo "  make             - Собирает приложение для текущей платформы"
	@echo "  make build-all   - Собирает приложение для всех платформ"
	@echo ""
	@echo "Результаты сборки будут помещены в директорию $(BUILD_DIR)" 