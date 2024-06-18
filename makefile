# Variabel
CMD_DIR = ./src/cmd
DOCS_DIR = ./docs

# Target Help
.PHONY: help
help:
	@echo "Use make for:"
	@echo "  - init-swagger : Generate Swagger docs"
	@echo " - run : Running service"

# For generate swagger docs
.PHONY: init-swagger
init-swagger:
	@echo "Generating Swagger docs..."
	swag init -g ./src/cmd/main.go


# Run Service
.PHONY: run
run: 
	@echo "Run Service"
	go run ./src/cmd/main.go

