package server

import (
    "net/http"
    "strings"
)

// contentTypeMiddleware sets JSON Content-Type for all API routes
func (s *Server) contentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasPrefix(r.URL.Path, "/api/") {
            w.Header().Set("Content-Type", "application/json")
        }
        next.ServeHTTP(w, r)
    })
}

// corsMiddleware handles CORS headers and preflight requests
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins in production
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
        w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

        // Handle preflight OPTIONS requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}