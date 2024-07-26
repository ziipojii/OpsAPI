# Variabel
CMD_DIR = ./src/cmd
DOCS_DIR = ./docs

# Target Help
.PHONY: help
help:
	@echo "Use make for:"
	@echo " - init-swagger : Generate Swagger docs"
	@echo " - run-local         : Running service locally"

# For generate swagger docs
.PHONY: init-swagger
init-swagger:
	@echo "Generating Swagger docs..."
	swag init -g ./src/cmd/main.go

# Run Service
.PHONY: run-local
run-local:
	@echo "Run Service"
	cp .env.local .env
	go run $(CMD_DIR)/main.go
