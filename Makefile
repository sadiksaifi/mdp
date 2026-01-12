.PHONY: help build run test coverage release clean install tidy

BINARY := mdp
DIST_DIR := dist
CMD_DIR := ./cmd/mdp

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-12s %s\n", $$1, $$2}'

build: ## Build the binary
	go build -o $(BINARY) $(CMD_DIR)

run: build ## Build and run with CLAUDE.md
	./$(BINARY) CLAUDE.md

test: ## Run all tests
	go test ./...

coverage: ## Run tests with coverage
	go test -cover ./...

release: ## Build release binaries for multiple platforms
	@mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY)-darwin-amd64 $(CMD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY)-darwin-arm64 $(CMD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY)-linux-amd64 $(CMD_DIR)
	@echo "Release binaries built in $(DIST_DIR)/"

clean: ## Remove built artifacts
	rm -f $(BINARY)
	rm -rf $(DIST_DIR)

install: build ## Install binary to /usr/local/bin
	cp $(BINARY) /usr/local/bin/$(BINARY)
	@echo "Installed $(BINARY) to /usr/local/bin/"

tidy: ## Run go mod tidy
	go mod tidy
