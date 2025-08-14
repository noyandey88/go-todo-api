# Makefile

# Variables
APP_CMD = cmd/main.go

# Default target
run: swag-run

build: swag-build

# Generate Swagger docs and run the app
swag-run:
	@echo "Running the application..."
	swag init -g $(APP_CMD)
	go run $(APP_CMD)

# Generate Swagger docs and run the app
swag-build:
	@echo "Building the application..."
	swag init -g $(APP_CMD)
	go build $(APP_CMD)

# Optional: only generate Swagger docs
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g $(APP_CMD)
	@echo "Swagger documentation generated successfully."
