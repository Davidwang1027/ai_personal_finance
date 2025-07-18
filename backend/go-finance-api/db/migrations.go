package db

import (
	"fmt"
	"log"
)

// MigrateDB creates all necessary tables in the database
func (db *Database) MigrateDB() error {
	log.Println("Running database migrations...")

	// Create users table
	if err := db.createUsersTable(); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create items table
	if err := db.createItemsTable(); err != nil {
		return fmt.Errorf("failed to create items table: %w", err)
	}

	// Create accounts table
	if err := db.createAccountsTable(); err != nil {
		return fmt.Errorf("failed to create accounts table: %w", err)
	}

	// Create transaction_categories table
	if err := db.createTransactionCategoriesTable(); err != nil {
		return fmt.Errorf("failed to create transaction_categories table: %w", err)
	}

	// Create transactions table
	if err := db.createTransactionsTable(); err != nil {
		return fmt.Errorf("failed to create transactions table: %w", err)
	}

	// Create events table
	if err := db.createEventsTable(); err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// Create users table
func (db *Database) createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	return err
}

// Create items table
func (db *Database) createItemsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		item_id VARCHAR(100) UNIQUE NOT NULL,
		access_token VARCHAR(255) NOT NULL,
		institution_id VARCHAR(100),
		institution_name VARCHAR(255),
		webhook_url VARCHAR(255),
		status VARCHAR(50) NOT NULL DEFAULT 'good',
		error TEXT,
		last_success_sync TIMESTAMP,
		transaction_cursor TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	return err
}

// Create accounts table
func (db *Database) createAccountsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		item_id INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
		plaid_account_id VARCHAR(100) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		official_name VARCHAR(255),
		mask VARCHAR(4),
		type VARCHAR(50) NOT NULL,
		subtype VARCHAR(50),
		current_balance NUMERIC(19, 4),
		available_balance NUMERIC(19, 4),
		credit_limit NUMERIC(19, 4),
		iso_currency_code VARCHAR(3),
		unofficial_currency_code VARCHAR(3),
		last_updated TIMESTAMP NOT NULL DEFAULT NOW(),
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	return err
}

// Create transaction_categories table
func (db *Database) createTransactionCategoriesTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS transaction_categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		parent_id INTEGER REFERENCES transaction_categories(id),
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	return err
}

// Create transactions table
func (db *Database) createTransactionsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
		plaid_transaction_id VARCHAR(100) UNIQUE NOT NULL,
		category_id INTEGER REFERENCES transaction_categories(id),
		name VARCHAR(255) NOT NULL,
		merchant_name VARCHAR(255),
		amount NUMERIC(19, 4) NOT NULL,
		iso_currency_code VARCHAR(3),
		date DATE NOT NULL,
		pending BOOLEAN NOT NULL DEFAULT FALSE,
		payment_channel VARCHAR(50),
		description TEXT,
		categories JSONB,
		category VARCHAR(255),
		location JSONB,
		personal_finance_category VARCHAR(255),
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	return err
}

// Create events table
func (db *Database) createEventsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
		item_id INTEGER REFERENCES items(id) ON DELETE SET NULL,
		type VARCHAR(50) NOT NULL,
		event_name VARCHAR(100) NOT NULL,
		request_id VARCHAR(100),
		status VARCHAR(50) NOT NULL,
		description TEXT,
		metadata JSONB,
		endpoint VARCHAR(255),
		request_body TEXT,
		response TEXT,
		error TEXT,
		link_session_id VARCHAR(100),
		link_event_name VARCHAR(100),
		link_event_metadata JSONB,
		webhook_type VARCHAR(50),
		webhook_code VARCHAR(50),
		payload JSONB,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	
	-- Create index for events table
	CREATE INDEX IF NOT EXISTS idx_events_user_id ON events(user_id);
	CREATE INDEX IF NOT EXISTS idx_events_item_id ON events(item_id);
	CREATE INDEX IF NOT EXISTS idx_events_type ON events(type);
	CREATE INDEX IF NOT EXISTS idx_events_created_at ON events(created_at);`

	_, err := db.Exec(query)
	return err
}
