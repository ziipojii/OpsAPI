# Variables
CMD_DIR = ./src/cmd
DOCS_DIR = ./docs

# Target Help
.PHONY: help
help:
	@echo "Use make for:"
	@echo " - init-swagger  : Generate Swagger docs"
	@echo " - run-local     : Running service locally"
	@echo " - docker-build  : Build Docker image"
	@echo " - docker-up     : Run Docker Compose"

# For generate swagger docs
.PHONY: init-swagger
init-swagger:
	@echo "Generating Swagger docs..."
	swag init -g $(CMD_DIR)/main.go

# Run Service Locally
.PHONY: run-local
run-local:
	@echo "Run Service"
	cp .env.local .env
	go run $(CMD_DIR)/main.go

# Build Docker Image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t devops_api .

# Run Docker Compose
.PHONY: docker-up
docker-up:
	@echo "Running Docker Compose..."
	docker-compose up --build
