# Go Todo API

A RESTful API for managing todos built with Go, GORM, and PostgreSQL.

## Features

- RESTful API endpoints for todo management
- User authentication with JWT
- PostgreSQL database with GORM
- Swagger documentation
- Dockerized application

## Prerequisites

- Go 1.24+
- PostgreSQL
- Docker and Docker Compose (for containerized deployment)

## Running Locally

1. Clone the repository
2. Configure the database in `configs/config.yaml`
3. Run the application:

```bash
go run cmd/main.go
```

Or use the Makefile:

```bash
make run
```

## API Documentation

Swagger documentation is available at `/swagger/` when the application is running.

## Docker Deployment

### Using Docker Compose (Recommended)

The easiest way to run the application is using Docker Compose:

```bash
# Build and start the containers
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the containers
docker-compose down
```

### Using Docker

You can also build and run the Docker container manually:

```bash
# Build the Docker image
docker build -t go-todo-app .

# Run the container
docker run -p 8080:8080 --name go-todo-app go-todo-app
```

## Environment Configuration

The application uses the configuration in `configs/config.yaml`. When running with Docker, you can override database settings using environment variables in the docker-compose.yml file.

## API Endpoints

- Authentication
  - POST `/api/auth/signup` - Register a new user
  - POST `/api/auth/signin` - Login and get JWT tokens
  - POST `/api/auth/signout` - Logout

- Todos
  - GET `/api/todos` - Get all todos
  - GET `/api/todos/{id}` - Get a specific todo
  - POST `/api/todos/create` - Create a new todo
  - PUT `/api/todos/update/{id}` - Update a todo
  - DELETE `/api/todos/delete/{id}` - Delete a todo

- Users
  - GET `/api/users` - Get all users
  - GET `/api/users/{id}` - Get a specific user
  - PUT `/api/users/update/{id}` - Update a user
  - DELETE `/api/users/delete/{id}` - Delete a user