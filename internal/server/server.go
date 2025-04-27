package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"sample_project/internal/database"
)

type Server struct {
	port int
	db   *database.Database
}

func NewServer() (*http.Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	server := &Server{
		port: port,
		db:   db,
	}

	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer, nil
}
