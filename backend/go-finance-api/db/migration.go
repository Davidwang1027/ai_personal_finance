package db

import (
	"log"
)

// Migrations contains SQL statements for database migrations
var Migrations = []string{
	// Migration 1: Create users table
	`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 2: Create items table
	`CREATE TABLE IF NOT EXISTS items (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		plaid_item_id VARCHAR(255) UNIQUE NOT NULL,
		access_token VARCHAR(255) NOT NULL,
		institution_id VARCHAR(255),
		institution_name VARCHAR(255),
		status VARCHAR(50) NOT NULL,
		webhook_url VARCHAR(255),
		consent VARCHAR(50),
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 3: Create accounts table
	`CREATE TABLE IF NOT EXISTS accounts (
		id UUID PRIMARY KEY,
		item_id UUID NOT NULL REFERENCES items(id) ON DELETE CASCADE,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		plaid_account_id VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		official_name VARCHAR(255),
		type VARCHAR(50) NOT NULL,
		subtype VARCHAR(50),
		mask VARCHAR(50),
		available_balance DECIMAL(19, 4),
		current_balance DECIMAL(19, 4),
		currency_code VARCHAR(3),
		last_updated TIMESTAMP WITH TIME ZONE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 4: Create transactions table
	`CREATE TABLE IF NOT EXISTS transactions (
		id UUID PRIMARY KEY,
		account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		plaid_transaction_id VARCHAR(255) UNIQUE NOT NULL,
		category_id VARCHAR(255),
		category TEXT[],
		name VARCHAR(255) NOT NULL,
		merchant_name VARCHAR(255),
		amount DECIMAL(19, 4) NOT NULL,
		iso_currency_code VARCHAR(3),
		date DATE NOT NULL,
		pending BOOLEAN NOT NULL,
		payment_channel VARCHAR(50),
		address VARCHAR(255),
		city VARCHAR(255),
		region VARCHAR(255),
		postal_code VARCHAR(20),
		country VARCHAR(2),
		latitude DECIMAL(9, 6),
		longitude DECIMAL(9, 6),
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 5: Create plaid_api_events table
	`CREATE TABLE IF NOT EXISTS plaid_api_events (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		item_id UUID REFERENCES items(id) ON DELETE SET NULL,
		endpoint VARCHAR(255) NOT NULL,
		request_body JSONB NOT NULL,
		response_body JSONB,
		status_code INTEGER,
		error_code VARCHAR(255),
		error_message TEXT,
		request_id VARCHAR(255),
		request_time TIMESTAMP WITH TIME ZONE NOT NULL,
		response_time TIMESTAMP WITH TIME ZONE,
		execution_time INTEGER,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 6: Create link_events table
	`CREATE TABLE IF NOT EXISTS link_events (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		item_id UUID REFERENCES items(id) ON DELETE SET NULL,
		event_name VARCHAR(255) NOT NULL,
		event_metadata JSONB NOT NULL,
		link_session_id VARCHAR(255) NOT NULL,
		request_id VARCHAR(255),
		error_code VARCHAR(255),
		error_message TEXT,
		status VARCHAR(50) NOT NULL,
		institution_id VARCHAR(255),
		institution_name VARCHAR(255),
		timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL
	)`,

	// Migration 7: Create indices
	`CREATE INDEX idx_items_user_id ON items(user_id);
	 CREATE INDEX idx_accounts_item_id ON accounts(item_id);
	 CREATE INDEX idx_accounts_user_id ON accounts(user_id);
	 CREATE INDEX idx_transactions_account_id ON transactions(account_id);
	 CREATE INDEX idx_transactions_user_id ON transactions(user_id);
	 CREATE INDEX idx_transactions_date ON transactions(date);
	 CREATE INDEX idx_plaid_api_events_user_id ON plaid_api_events(user_id);
	 CREATE INDEX idx_plaid_api_events_item_id ON plaid_api_events(item_id);
	 CREATE INDEX idx_link_events_user_id ON link_events(user_id);
	 CREATE INDEX idx_link_events_item_id ON link_events(item_id);`,
}

// MigrateDB executes all migrations on the database
func (db *Database) MigrateDB() error {
	log.Println("Running database migrations...")

	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			migration_number INTEGER UNIQUE NOT NULL,
			applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	// Get the latest migration number
	var latestMigration int
	err = db.QueryRow("SELECT COALESCE(MAX(migration_number), 0) FROM migrations").Scan(&latestMigration)
	if err != nil {
		return err
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Apply new migrations
	for i := latestMigration; i < len(Migrations); i++ {
		migrationNum := i + 1
		log.Printf("Applying migration %d...", migrationNum)

		// Execute migration
		_, err = tx.Exec(Migrations[i])
		if err != nil {
			return err
		}

		// Record migration
		_, err = tx.Exec("INSERT INTO migrations (migration_number) VALUES ($1)", migrationNum)
		if err != nil {
			return err
		}

		log.Printf("Migration %d applied successfully", migrationNum)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
