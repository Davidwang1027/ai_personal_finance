package db

import (
	"database/sql"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// TransactionRepository handles database operations for transactions
type TransactionRepository struct {
	db *Database
}

// NewTransactionRepository creates a new TransactionRepository
func NewTransactionRepository(db *Database) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create inserts a new transaction into the database
func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, 
			$14, $15, $16, $17, $18, $19, $20, $21, $22
		)
	`
	_, err := r.db.Exec(
		query,
		transaction.ID,
		transaction.AccountID,
		transaction.UserID,
		transaction.PlaidTransactionID,
		transaction.CategoryID,
		pq.Array(transaction.Category),
		transaction.Name,
		transaction.MerchantName,
		transaction.Amount,
		transaction.IsoCurrencyCode,
		transaction.Date,
		transaction.Pending,
		transaction.PaymentChannel,
		transaction.Address,
		transaction.City,
		transaction.Region,
		transaction.PostalCode,
		transaction.Country,
		transaction.Latitude,
		transaction.Longitude,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
	return err
}

// GetByID retrieves a transaction by ID
func (r *TransactionRepository) GetByID(id uuid.UUID) (*models.Transaction, error) {
	query := `
		SELECT 
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		FROM transactions
		WHERE id = $1
	`
	var transaction models.Transaction
	err := r.db.QueryRow(query, id).Scan(
		&transaction.ID,
		&transaction.AccountID,
		&transaction.UserID,
		&transaction.PlaidTransactionID,
		&transaction.CategoryID,
		pq.Array(&transaction.Category),
		&transaction.Name,
		&transaction.MerchantName,
		&transaction.Amount,
		&transaction.IsoCurrencyCode,
		&transaction.Date,
		&transaction.Pending,
		&transaction.PaymentChannel,
		&transaction.Address,
		&transaction.City,
		&transaction.Region,
		&transaction.PostalCode,
		&transaction.Country,
		&transaction.Latitude,
		&transaction.Longitude,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Transaction not found
		}
		return nil, err
	}
	return &transaction, nil
}

// GetByPlaidTransactionID retrieves a transaction by its Plaid transaction ID
func (r *TransactionRepository) GetByPlaidTransactionID(plaidTransactionID string) (*models.Transaction, error) {
	query := `
		SELECT 
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		FROM transactions
		WHERE plaid_transaction_id = $1
	`
	var transaction models.Transaction
	err := r.db.QueryRow(query, plaidTransactionID).Scan(
		&transaction.ID,
		&transaction.AccountID,
		&transaction.UserID,
		&transaction.PlaidTransactionID,
		&transaction.CategoryID,
		pq.Array(&transaction.Category),
		&transaction.Name,
		&transaction.MerchantName,
		&transaction.Amount,
		&transaction.IsoCurrencyCode,
		&transaction.Date,
		&transaction.Pending,
		&transaction.PaymentChannel,
		&transaction.Address,
		&transaction.City,
		&transaction.Region,
		&transaction.PostalCode,
		&transaction.Country,
		&transaction.Latitude,
		&transaction.Longitude,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Transaction not found
		}
		return nil, err
	}
	return &transaction, nil
}

// GetByAccountID retrieves transactions for a specific account
func (r *TransactionRepository) GetByAccountID(accountID uuid.UUID, limit, offset int) ([]*models.Transaction, error) {
	query := `
		SELECT 
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		FROM transactions
		WHERE account_id = $1
		ORDER BY date DESC, created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, accountID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.UserID,
			&transaction.PlaidTransactionID,
			&transaction.CategoryID,
			pq.Array(&transaction.Category),
			&transaction.Name,
			&transaction.MerchantName,
			&transaction.Amount,
			&transaction.IsoCurrencyCode,
			&transaction.Date,
			&transaction.Pending,
			&transaction.PaymentChannel,
			&transaction.Address,
			&transaction.City,
			&transaction.Region,
			&transaction.PostalCode,
			&transaction.Country,
			&transaction.Latitude,
			&transaction.Longitude,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByUserID retrieves transactions for a specific user
func (r *TransactionRepository) GetByUserID(userID uuid.UUID, limit, offset int) ([]*models.Transaction, error) {
	query := `
		SELECT 
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY date DESC, created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.UserID,
			&transaction.PlaidTransactionID,
			&transaction.CategoryID,
			pq.Array(&transaction.Category),
			&transaction.Name,
			&transaction.MerchantName,
			&transaction.Amount,
			&transaction.IsoCurrencyCode,
			&transaction.Date,
			&transaction.Pending,
			&transaction.PaymentChannel,
			&transaction.Address,
			&transaction.City,
			&transaction.Region,
			&transaction.PostalCode,
			&transaction.Country,
			&transaction.Latitude,
			&transaction.Longitude,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByDateRange retrieves transactions for a specific user within a date range
func (r *TransactionRepository) GetByDateRange(userID uuid.UUID, startDate, endDate time.Time, limit, offset int) ([]*models.Transaction, error) {
	query := `
		SELECT 
			id, account_id, user_id, plaid_transaction_id, category_id, category,
			name, merchant_name, amount, iso_currency_code, date, pending, 
			payment_channel, address, city, region, postal_code, country,
			latitude, longitude, created_at, updated_at
		FROM transactions
		WHERE user_id = $1 AND date BETWEEN $2 AND $3
		ORDER BY date DESC, created_at DESC
		LIMIT $4 OFFSET $5
	`
	rows, err := r.db.Query(query, userID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.UserID,
			&transaction.PlaidTransactionID,
			&transaction.CategoryID,
			pq.Array(&transaction.Category),
			&transaction.Name,
			&transaction.MerchantName,
			&transaction.Amount,
			&transaction.IsoCurrencyCode,
			&transaction.Date,
			&transaction.Pending,
			&transaction.PaymentChannel,
			&transaction.Address,
			&transaction.City,
			&transaction.Region,
			&transaction.PostalCode,
			&transaction.Country,
			&transaction.Latitude,
			&transaction.Longitude,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

// UpdatePendingStatus updates the pending status of a transaction
func (r *TransactionRepository) UpdatePendingStatus(id uuid.UUID, pending bool) error {
	query := `
		UPDATE transactions
		SET pending = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(query, pending, time.Now().UTC(), id)
	return err
}

// Delete removes a transaction from the database
func (r *TransactionRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM transactions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// DeleteByPlaidTransactionID removes a transaction by its Plaid transaction ID
func (r *TransactionRepository) DeleteByPlaidTransactionID(plaidTransactionID string) error {
	query := `DELETE FROM transactions WHERE plaid_transaction_id = $1`
	_, err := r.db.Exec(query, plaidTransactionID)
	return err
}
