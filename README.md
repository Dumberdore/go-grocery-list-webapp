# Grocery List Application

A full-stack web application for managing grocery lists, built with Go and React.

## Tech Stack

### Backend
- Go 1.23
- PostgreSQL 16
- Docker & Docker Compose

### Frontend
- React 19
- TypeScript
- Vite 6

## Prerequisites

- Go 1.23 or higher
- Node.js 20 or higher
- Docker and Docker Compose
- Make (optional, but recommended)

## Getting Started

1. Clone the repository:
```bash
gh repo clone Dumberdore/go-grocery-list-webapp
cd go-grocery-list-webapp
```

2. Create a `.env` file in the root directory:
```bash
APP_ENV=development
PORT=8080
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=grocerylist
```

3. Start the application:

Using Make:
```bash
make docker-run    # Starts all services using Docker
```

Or manually:
```bash
docker compose up --build
```

The application will be available at:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- Database: localhost:5432

## Development

### Available Make Commands

```bash
make build        # Build the Go application
make run          # Run both frontend and backend
make test         # Run Go tests
make watch        # Run Go application with hot reload
make docker-run   # Start all services with Docker
make docker-down  # Stop all Docker services
make clean        # Clean build artifacts
```

### Hot Reload

The project supports hot reload for both frontend and backend:
- Frontend uses Vite's built-in HMR
- Backend uses Air for live reload (install with `go install github.com/air-verse/air@latest`)

### Project Structure

```
├── cmd/
│   └── api/          # Application entrypoint
├── frontend/         # React frontend application
├── internal/         # Go package code
│   ├── database/     # Database connection and migrations
│   ├── handlers/     # HTTP handlers
│   ├── models/       # Data models
│   ├── repository/   # Database operations
│   └── server/       # HTTP server setup
├── docker-compose.yml
├── Dockerfile
└── Makefile
```

## API Endpoints

- `GET /api/items` - Get all grocery items
- `POST /api/items` - Create a new grocery item
- `PUT /api/items/{id}` - Update a grocery item
- `DELETE /api/items/{id}` - Delete a grocery item

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
