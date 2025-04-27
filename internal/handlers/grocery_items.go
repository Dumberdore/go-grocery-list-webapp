package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sample_project/internal/models"
	"sample_project/internal/repository"
	"strconv"
)

type GroceryItemHandler struct {
	repo *repository.GroceryItemRepository
}

func NewGroceryItemHandler(db *sql.DB) *GroceryItemHandler {
	return &GroceryItemHandler{
		repo: repository.NewGroceryItemRepository(db),
	}
}

// HandleAllItems handles requests for the collection of items (/api/items)
func (h *GroceryItemHandler) HandleAllItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllItems(w, r)
	case http.MethodPost:
		h.CreateItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleItem handles requests for individual items (/api/items/{id})
func (h *GroceryItemHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		h.UpdateItem(w, r)
	case http.MethodDelete:
		h.DeleteItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetAllItems retrieves all grocery items
func (h *GroceryItemHandler) GetAllItems(w http.ResponseWriter, _ *http.Request) {
	items, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateItem creates a new grocery item
func (h *GroceryItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.GroceryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.Create(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateItem updates an existing grocery item
func (h *GroceryItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/api/items/"):])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.GroceryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.repo.Update(id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteItem deletes a grocery item
func (h *GroceryItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/api/items/"):])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	deleted, err := h.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
