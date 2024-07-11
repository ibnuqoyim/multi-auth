# Multi Auth

## Project Overview
This project is a login service built with Golang, Echo framework, and MongoDB. It supports multiple authentication methods including Email/Password, OAuth, SSO, and Two-Factor Authentication (2FA).

## Requirements
- Golang
- MongoDB
- Redis

## Setup
1. Clone the repository.
2. Create a `.env` file based on the example provided.
3. Run `go mod tidy` to install dependencies.
4. Run the application using `go run cmd/main.go`.

## API Documentation
API documentation is available at `/swagger/index.html`.

## Folder Structure
- `cmd/` - Main application entry point.
- `configs/` - Configuration files.
- `controllers/` - API controllers.
- `middlewares/` - Middleware functions.
- `models/` - Data models.
- `repositories/` - Database repositories.
- `services/` - Business logic services.
- `utils/` - Utility functions.
- `test/` - Unit tests.
- `docs/` - Swagger documentation.