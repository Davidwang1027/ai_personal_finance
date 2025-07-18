package models

import (
	"time"
)

// Transaction represents a financial transaction from Plaid
type Transaction struct {
	ID                      int64     `json:"id"`
	AccountID               int64     `json:"account_id"`            // References our Account model
	PlaidTransactionID      string    `json:"plaid_transaction_id"`  // Plaid's transaction_id
	CategoryID              int64     `json:"category_id,omitempty"` // References our Category model
	Name                    string    `json:"name"`
	MerchantName            string    `json:"merchant_name,omitempty"`
	Amount                  float64   `json:"amount"`
	IsoCurrencyCode         string    `json:"iso_currency_code"`
	Date                    time.Time `json:"date"` // Date when transaction occurred
	Pending                 bool      `json:"pending"`
	PaymentChannel          string    `json:"payment_channel"`
	Description             string    `json:"description,omitempty"`
	Categories              []string  `json:"categories,omitempty"` // Array of category strings
	Category                string    `json:"category,omitempty"`   // Primary category
	Location                string    `json:"location,omitempty"`   // JSON string of location data
	PersonalFinanceCategory string    `json:"personal_finance_category,omitempty"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

// TransactionCategory represents a category for transactions
type TransactionCategory struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ParentID  int64     `json:"parent_id,omitempty"` // For hierarchical categories
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TransactionSync represents a response for a transaction sync operation
type TransactionSync struct {
	Added      []Transaction `json:"added"`
	Modified   []Transaction `json:"modified"`
	Removed    []string      `json:"removed"` // Only transaction IDs
	NextCursor string        `json:"next_cursor"`
}

// PaymentChannel represents the payment channel of a transaction
const (
	PaymentChannelOnline  = "online"
	PaymentChannelInStore = "in store"
	PaymentChannelOther   = "other"
)
