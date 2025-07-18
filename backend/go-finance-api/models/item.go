package models

import (
	"time"
)

// Item represents a Plaid item (a set of credentials at a financial institution)
type Item struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	ItemID            string    `json:"item_id"`            // Plaid's item_id
	AccessToken       string    `json:"-"`                  // Plaid access token (sensitive)
	InstitutionID     string    `json:"institution_id"`     // ID of the financial institution
	InstitutionName   string    `json:"institution_name"`   // Name of the financial institution
	WebhookURL        string    `json:"webhook_url"`        // URL for receiving webhooks
	Status            string    `json:"status"`             // "good", "error", etc.
	Error             string    `json:"error,omitempty"`    // Error message if status is "error"
	LastSuccessSync   time.Time `json:"last_success_sync"`  // Last successful sync time
	TransactionCursor string    `json:"transaction_cursor"` // Cursor for transactions/sync
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// ItemResponse is the response for item information, excluding sensitive data
type ItemResponse struct {
	ID              int64     `json:"id"`
	ItemID          string    `json:"item_id"`
	InstitutionID   string    `json:"institution_id"`
	InstitutionName string    `json:"institution_name"`
	Status          string    `json:"status"`
	Error           string    `json:"error,omitempty"`
	LastSuccessSync time.Time `json:"last_success_sync"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ItemStatus represents the status of a Plaid item
const (
	ItemStatusGood            = "good"
	ItemStatusError           = "error"
	ItemStatusUserPermRevoked = "user_permission_revoked"
	ItemStatusLoginRequired   = "login_required"
)
