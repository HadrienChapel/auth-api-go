# Authentication API in Go

## Overview
This is a RESTful API built with Go for user authentication. It supports user registration, login, and JWT-based authentication for protected routes.

## Features
- User registration with bcrypt password hashing.
- User login with JWT generation.
- Middleware to protect routes with token validation.
- SQLite database for user storage.

## Prerequisites
- Go 1.19+

## Setup and Run
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd auth-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. The API will be available at `http://localhost:8080`.

## API Endpoints
### `POST /register`
Register a new user. Example body:
```json
{
  "username": "user1",
  "password": "password123"
}
```

### `POST /login`
Log in and receive a JWT. Example body:
```json
{
  "username": "user1",
  "password": "password123"
}
```

### `GET /protected/dashboard`
Access a protected route (requires token). Add the token to the `Authorization` header:
```
Authorization: <JWT_TOKEN>
```

## Notes
- The secret key for JWT is defined in `utils/jwt.go`. Change it for production.
- Database migrations are handled automatically using GORM.
