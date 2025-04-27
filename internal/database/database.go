package database

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

type Database struct {
    DB *sql.DB
}

func NewDatabase() (*Database, error) {
    db, err := sql.Open("postgres", getConnectionString())
    if err != nil {
        return nil, fmt.Errorf("error opening database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to the database: %w", err)
    }

    if err := createTables(db); err != nil {
        return nil, fmt.Errorf("error creating tables: %w", err)
    }

    return &Database{
        DB: db,
    }, nil
}

func getConnectionString() string {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        getEnvOrDefault("DB_HOST", "localhost"),
        getEnvOrDefault("DB_PORT", "5432"),
        getEnvOrDefault("DB_USER", "postgres"),
        getEnvOrDefault("DB_PASSWORD", "postgres"),
        getEnvOrDefault("DB_NAME", "grocerylist"),
    )
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func createTables(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS grocery_items (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        completed BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := db.Exec(query)
    return err
}