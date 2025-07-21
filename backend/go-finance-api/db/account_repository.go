package db

import (
	"database/sql"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/models"
	"github.com/google/uuid"
)

// AccountRepository handles database operations for accounts
type AccountRepository struct {
	db *Database
}

// NewAccountRepository creates a new AccountRepository
func NewAccountRepository(db *Database) *AccountRepository {
	return &AccountRepository{db: db}
}

// Create inserts a new account into the database
func (r *AccountRepository) Create(account *models.Account) error {
	query := `
		INSERT INTO accounts (
			id, item_id, user_id, plaid_account_id, name, official_name, 
			type, subtype, mask, available_balance, current_balance, 
			currency_code, last_updated, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err := r.db.Exec(
		query,
		account.ID,
		account.ItemID,
		account.UserID,
		account.PlaidAccountID,
		account.Name,
		account.OfficialName,
		account.Type,
		account.Subtype,
		account.Mask,
		account.AvailableBalance,
		account.CurrentBalance,
		account.CurrencyCode,
		account.LastUpdated,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return err
}

// GetByID retrieves an account by ID
func (r *AccountRepository) GetByID(id uuid.UUID) (*models.Account, error) {
	query := `
		SELECT 
			id, item_id, user_id, plaid_account_id, name, official_name, 
			type, subtype, mask, available_balance, current_balance, 
			currency_code, last_updated, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`
	var account models.Account
	err := r.db.QueryRow(query, id).Scan(
		&account.ID,
		&account.ItemID,
		&account.UserID,
		&account.PlaidAccountID,
		&account.Name,
		&account.OfficialName,
		&account.Type,
		&account.Subtype,
		&account.Mask,
		&account.AvailableBalance,
		&account.CurrentBalance,
		&account.CurrencyCode,
		&account.LastUpdated,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Account not found
		}
		return nil, err
	}
	return &account, nil
}

// GetByPlaidAccountID retrieves an account by its Plaid account ID
func (r *AccountRepository) GetByPlaidAccountID(plaidAccountID string) (*models.Account, error) {
	query := `
		SELECT 
			id, item_id, user_id, plaid_account_id, name, official_name, 
			type, subtype, mask, available_balance, current_balance, 
			currency_code, last_updated, created_at, updated_at
		FROM accounts
		WHERE plaid_account_id = $1
	`
	var account models.Account
	err := r.db.QueryRow(query, plaidAccountID).Scan(
		&account.ID,
		&account.ItemID,
		&account.UserID,
		&account.PlaidAccountID,
		&account.Name,
		&account.OfficialName,
		&account.Type,
		&account.Subtype,
		&account.Mask,
		&account.AvailableBalance,
		&account.CurrentBalance,
		&account.CurrencyCode,
		&account.LastUpdated,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Account not found
		}
		return nil, err
	}
	return &account, nil
}

// GetByItemID retrieves all accounts for a specific item
func (r *AccountRepository) GetByItemID(itemID uuid.UUID) ([]*models.Account, error) {
	query := `
		SELECT 
			id, item_id, user_id, plaid_account_id, name, official_name, 
			type, subtype, mask, available_balance, current_balance, 
			currency_code, last_updated, created_at, updated_at
		FROM accounts
		WHERE item_id = $1
		ORDER BY name
	`
	rows, err := r.db.Query(query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.ID,
			&account.ItemID,
			&account.UserID,
			&account.PlaidAccountID,
			&account.Name,
			&account.OfficialName,
			&account.Type,
			&account.Subtype,
			&account.Mask,
			&account.AvailableBalance,
			&account.CurrentBalance,
			&account.CurrencyCode,
			&account.LastUpdated,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetByUserID retrieves all accounts for a specific user
func (r *AccountRepository) GetByUserID(userID uuid.UUID) ([]*models.Account, error) {
	query := `
		SELECT 
			id, item_id, user_id, plaid_account_id, name, official_name, 
			type, subtype, mask, available_balance, current_balance, 
			currency_code, last_updated, created_at, updated_at
		FROM accounts
		WHERE user_id = $1
		ORDER BY name
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.ID,
			&account.ItemID,
			&account.UserID,
			&account.PlaidAccountID,
			&account.Name,
			&account.OfficialName,
			&account.Type,
			&account.Subtype,
			&account.Mask,
			&account.AvailableBalance,
			&account.CurrentBalance,
			&account.CurrencyCode,
			&account.LastUpdated,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

// UpdateBalances updates the account balances
func (r *AccountRepository) UpdateBalances(id uuid.UUID, availableBalance, currentBalance float64) error {
	query := `
		UPDATE accounts
		SET available_balance = $1, current_balance = $2, last_updated = $3, updated_at = $4
		WHERE id = $5
	`
	now := time.Now().UTC()
	_, err := r.db.Exec(query, availableBalance, currentBalance, now, now, id)
	return err
}

// Delete removes an account from the database
func (r *AccountRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
