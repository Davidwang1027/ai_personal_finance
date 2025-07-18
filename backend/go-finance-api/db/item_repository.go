package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/models"
)

// ItemRepository handles database operations for Plaid items
type ItemRepository struct {
	db *Database
}

// NewItemRepository creates a new item repository
func NewItemRepository(db *Database) *ItemRepository {
	return &ItemRepository{db}
}

// Create inserts a new item into the database
func (r *ItemRepository) Create(item *models.Item) error {
	query := `
	INSERT INTO items (
		user_id, item_id, access_token, institution_id, institution_name, 
		webhook_url, status, error, last_success_sync, transaction_cursor, 
		created_at, updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		item.UserID,
		item.ItemID,
		item.AccessToken,
		item.InstitutionID,
		item.InstitutionName,
		item.WebhookURL,
		item.Status,
		item.Error,
		item.LastSuccessSync,
		item.TransactionCursor,
		now,
		now,
	).Scan(&item.ID)

	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	item.CreatedAt = now
	item.UpdatedAt = now
	return nil
}

// GetByID retrieves an item by ID
func (r *ItemRepository) GetByID(id int64) (*models.Item, error) {
	query := `
	SELECT id, user_id, item_id, access_token, institution_id, institution_name,
		webhook_url, status, error, last_success_sync, transaction_cursor,
		created_at, updated_at
	FROM items
	WHERE id = $1`

	item := &models.Item{}
	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.UserID,
		&item.ItemID,
		&item.AccessToken,
		&item.InstitutionID,
		&item.InstitutionName,
		&item.WebhookURL,
		&item.Status,
		&item.Error,
		&item.LastSuccessSync,
		&item.TransactionCursor,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

// GetByPlaidItemID retrieves an item by Plaid's item_id
func (r *ItemRepository) GetByPlaidItemID(itemID string) (*models.Item, error) {
	query := `
	SELECT id, user_id, item_id, access_token, institution_id, institution_name,
		webhook_url, status, error, last_success_sync, transaction_cursor,
		created_at, updated_at
	FROM items
	WHERE item_id = $1`

	item := &models.Item{}
	err := r.db.QueryRow(query, itemID).Scan(
		&item.ID,
		&item.UserID,
		&item.ItemID,
		&item.AccessToken,
		&item.InstitutionID,
		&item.InstitutionName,
		&item.WebhookURL,
		&item.Status,
		&item.Error,
		&item.LastSuccessSync,
		&item.TransactionCursor,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return item, nil
}

// GetByUserID retrieves all items for a user
func (r *ItemRepository) GetByUserID(userID int64) ([]*models.Item, error) {
	query := `
	SELECT id, user_id, item_id, access_token, institution_id, institution_name,
		webhook_url, status, error, last_success_sync, transaction_cursor,
		created_at, updated_at
	FROM items
	WHERE user_id = $1
	ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	defer rows.Close()

	items := []*models.Item{}
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.ItemID,
			&item.AccessToken,
			&item.InstitutionID,
			&item.InstitutionName,
			&item.WebhookURL,
			&item.Status,
			&item.Error,
			&item.LastSuccessSync,
			&item.TransactionCursor,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item row: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating items rows: %w", err)
	}

	return items, nil
}

// Update updates an item
func (r *ItemRepository) Update(item *models.Item) error {
	query := `
	UPDATE items
	SET institution_id = $1, institution_name = $2, webhook_url = $3, 
		status = $4, error = $5, last_success_sync = $6, 
		transaction_cursor = $7, updated_at = $8
	WHERE id = $9`

	now := time.Now()
	_, err := r.db.Exec(
		query,
		item.InstitutionID,
		item.InstitutionName,
		item.WebhookURL,
		item.Status,
		item.Error,
		item.LastSuccessSync,
		item.TransactionCursor,
		now,
		item.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	item.UpdatedAt = now
	return nil
}

// UpdateStatus updates an item's status
func (r *ItemRepository) UpdateStatus(id int64, status string, errorMessage string) error {
	query := `
	UPDATE items
	SET status = $1, error = $2, updated_at = $3
	WHERE id = $4`

	now := time.Now()
	_, err := r.db.Exec(
		query,
		status,
		errorMessage,
		now,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to update item status: %w", err)
	}

	return nil
}

// UpdateTransactionCursor updates an item's transaction cursor
func (r *ItemRepository) UpdateTransactionCursor(id int64, cursor string) error {
	query := `
	UPDATE items
	SET transaction_cursor = $1, last_success_sync = $2, updated_at = $3
	WHERE id = $4`

	now := time.Now()
	_, err := r.db.Exec(
		query,
		cursor,
		now,
		now,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to update transaction cursor: %w", err)
	}

	return nil
}

// Delete deletes an item
func (r *ItemRepository) Delete(id int64) error {
	query := `DELETE FROM items WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}
