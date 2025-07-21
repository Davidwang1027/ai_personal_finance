package db

import (
	"database/sql"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/models"
	"github.com/google/uuid"
)

// ItemRepository handles database operations for Plaid items
type ItemRepository struct {
	db *Database
}

// NewItemRepository creates a new ItemRepository
func NewItemRepository(db *Database) *ItemRepository {
	return &ItemRepository{db: db}
}

// Create inserts a new item into the database
func (r *ItemRepository) Create(item *models.Item) error {
	query := `
		INSERT INTO items (id, user_id, plaid_item_id, access_token, institution_id, institution_name, status, webhook_url, consent, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Exec(
		query,
		item.ID,
		item.UserID,
		item.PlaidItemID,
		item.AccessToken,
		item.InstitutionID,
		item.InstitutionName,
		item.Status,
		item.WebhookURL,
		item.Consent,
		item.CreatedAt,
		item.UpdatedAt,
	)
	return err
}

// GetByID retrieves an item by ID
func (r *ItemRepository) GetByID(id uuid.UUID) (*models.Item, error) {
	query := `
		SELECT id, user_id, plaid_item_id, access_token, institution_id, institution_name, status, webhook_url, consent, created_at, updated_at
		FROM items
		WHERE id = $1
	`
	var item models.Item
	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.UserID,
		&item.PlaidItemID,
		&item.AccessToken,
		&item.InstitutionID,
		&item.InstitutionName,
		&item.Status,
		&item.WebhookURL,
		&item.Consent,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Item not found
		}
		return nil, err
	}
	return &item, nil
}

// GetByPlaidItemID retrieves an item by its Plaid item ID
func (r *ItemRepository) GetByPlaidItemID(plaidItemID string) (*models.Item, error) {
	query := `
		SELECT id, user_id, plaid_item_id, access_token, institution_id, institution_name, status, webhook_url, consent, created_at, updated_at
		FROM items
		WHERE plaid_item_id = $1
	`
	var item models.Item
	err := r.db.QueryRow(query, plaidItemID).Scan(
		&item.ID,
		&item.UserID,
		&item.PlaidItemID,
		&item.AccessToken,
		&item.InstitutionID,
		&item.InstitutionName,
		&item.Status,
		&item.WebhookURL,
		&item.Consent,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Item not found
		}
		return nil, err
	}
	return &item, nil
}

// GetByUserID retrieves all items for a specific user
func (r *ItemRepository) GetByUserID(userID uuid.UUID) ([]*models.Item, error) {
	query := `
		SELECT id, user_id, plaid_item_id, access_token, institution_id, institution_name, status, webhook_url, consent, created_at, updated_at
		FROM items
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.PlaidItemID,
			&item.AccessToken,
			&item.InstitutionID,
			&item.InstitutionName,
			&item.Status,
			&item.WebhookURL,
			&item.Consent,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Update updates an existing item
func (r *ItemRepository) Update(item *models.Item) error {
	query := `
		UPDATE items
		SET status = $1, webhook_url = $2, consent = $3, institution_id = $4, institution_name = $5, updated_at = $6
		WHERE id = $7
	`
	item.UpdatedAt = time.Now().UTC()
	_, err := r.db.Exec(
		query,
		item.Status,
		item.WebhookURL,
		item.Consent,
		item.InstitutionID,
		item.InstitutionName,
		item.UpdatedAt,
		item.ID,
	)
	return err
}

// UpdateAccessToken updates just the access token for an item
func (r *ItemRepository) UpdateAccessToken(id uuid.UUID, accessToken string) error {
	query := `
		UPDATE items
		SET access_token = $1, updated_at = $2
		WHERE id = $3
	`
	updatedAt := time.Now().UTC()
	_, err := r.db.Exec(query, accessToken, updatedAt, id)
	return err
}

// Delete removes an item from the database
func (r *ItemRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
