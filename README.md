# Auth Service - Golang Learning Project

Golang authentication service dengan JWT dan OAuth2 - project pembelajaran untuk Java engineers yang belajar Go.

## Features

- ✅ User registration & login
- ✅ JWT-based authentication (access token + refresh token)
- ✅ Password hashing dengan bcrypt
- ⏳ OAuth2 integration (Google, GitHub)
- ✅ PostgreSQL database dengan GORM
- ✅ Standard library HTTP (no framework - untuk belajar fundamental)
- ✅ Clean architecture (handler → service → repository)

## Tech Stack

- **Go 1.25**
- **Database:** PostgreSQL + GORM
- **Authentication:** JWT (golang-jwt/jwt)
- **Password:** bcrypt
- **Config:** godotenv
- **HTTP:** Standard library (`net/http`)

## Project Structure

```
auth-service/
├── cmd/
│   └── api/              # Entry point aplikasi
│       └── main.go
├── internal/             # Private application code
│   ├── config/          # Configuration management
│   ├── domain/          # Domain models/entities
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic
│   ├── handler/         # HTTP handlers (controllers)
│   ├── middleware/      # HTTP middleware (auth, logging)
│   └── util/           # Utilities (JWT, password)
├── migrations/          # Database migrations
├── .env                 # Environment variables (gitignored)
├── .env.example         # Environment template
└── go.mod              # Dependencies
```

## Setup

### 1. Prerequisites
- Go 1.25+
- PostgreSQL 14+

### 2. Install Dependencies
```bash
go mod download
```

### 3. Configuration
Copy `.env.example` ke `.env` dan sesuaikan:
```bash
cp .env.example .env
```

Edit `.env`:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_service_db
JWT_SECRET=your_secret_key
```

### 4. Database Setup
```bash
# Create database
createdb auth_service_db

# Run migrations (auto-migrate saat app start)
```

### 5. Run Application
```bash
go run cmd/api/main.go
```

## API Endpoints

### Authentication

#### Register
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

#### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "access_token": "eyJhbGc...",
  "refresh_token": "eyJhbGc...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Refresh Token
```http
POST /api/auth/refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGc..."
}
```

#### Logout
```http
POST /api/auth/logout
Authorization: Bearer <access_token>
Content-Type: application/json

{
  "refresh_token": "eyJhbGc..."
}
```

#### Get Current User
```http
GET /api/auth/me
Authorization: Bearer <access_token>
```

## Development

### Run Tests
```bash
go test ./...
```

### Format Code
```bash
goimports -w .
```

### Build
```bash
go build -o bin/api cmd/api/main.go
```

## Progress

- [x] Project setup & dependencies
- [x] Configuration management
- [x] Domain models (User, RefreshToken)
- [ ] Database connection & migrations
- [ ] JWT utilities
- [ ] Password utilities
- [ ] Repository layer
- [ ] Service layer
- [ ] HTTP handlers
- [ ] Middleware
- [ ] Main application
- [ ] OAuth2 integration
- [ ] Docker setup