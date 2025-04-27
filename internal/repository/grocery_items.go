package repository

import (
	"database/sql"
	"sample_project/internal/models"
)

type GroceryItemRepository struct {
	db *sql.DB
}

func NewGroceryItemRepository(db *sql.DB) *GroceryItemRepository {
	return &GroceryItemRepository{
		db: db,
	}
}

func (r *GroceryItemRepository) GetAll() ([]models.GroceryItem, error) {
	rows, err := r.db.Query("SELECT id, name, completed, created_at FROM grocery_items ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]models.GroceryItem, 0)
	for rows.Next() {
		var item models.GroceryItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Completed, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *GroceryItemRepository) Create(item *models.GroceryItem) error {
	return r.db.QueryRow(
		"INSERT INTO grocery_items (name, completed) VALUES ($1, $2) RETURNING id, created_at",
		item.Name, item.Completed,
	).Scan(&item.ID, &item.CreatedAt)
}

func (r *GroceryItemRepository) Update(id int, item models.GroceryItem) (bool, error) {
	result, err := r.db.Exec(
		"UPDATE grocery_items SET name = $1, completed = $2 WHERE id = $3",
		item.Name, item.Completed, id,
	)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (r *GroceryItemRepository) Delete(id int) (bool, error) {
	result, err := r.db.Exec("DELETE FROM grocery_items WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}
