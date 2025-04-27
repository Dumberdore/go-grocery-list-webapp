package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sample_project/internal/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Initialize handlers
	groceryHandler := handlers.NewGroceryItemHandler(s.db.DB)

	// Register routes
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/api/items", groceryHandler.HandleAllItems)
	mux.HandleFunc("/api/items/", groceryHandler.HandleItem)

	// Chain middlewares
	handler := s.contentTypeMiddleware(mux)
	handler = s.corsMiddleware(handler)

	return handler
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
