package models

import "time"

type GroceryItem struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}