package models

import (
	"time"

	"github.com/google/uuid"
)

// Item represents a Plaid Item (a connection to a financial institution)
type Item struct {
	ID              uuid.UUID `json:"id" db:"id"`
	UserID          uuid.UUID `json:"user_id" db:"user_id"`
	PlaidItemID     string    `json:"plaid_item_id" db:"plaid_item_id"`
	AccessToken     string    `json:"-" db:"access_token"` // Sensitive - don't expose in JSON
	InstitutionID   string    `json:"institution_id" db:"institution_id"`
	InstitutionName string    `json:"institution_name" db:"institution_name"`
	Status          string    `json:"status" db:"status"`
	WebhookURL      string    `json:"webhook_url" db:"webhook_url"`
	Consent         string    `json:"consent" db:"consent"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// NewItem creates a new Plaid Item record
func NewItem(
	userID uuid.UUID,
	plaidItemID string,
	accessToken string,
	institutionID string,
	institutionName string,
) *Item {
	now := time.Now().UTC()
	return &Item{
		ID:              uuid.New(),
		UserID:          userID,
		PlaidItemID:     plaidItemID,
		AccessToken:     accessToken,
		InstitutionID:   institutionID,
		InstitutionName: institutionName,
		Status:          "active",
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// UpdateStatus updates the status of the Item
func (i *Item) UpdateStatus(status string) {
	i.Status = status
	i.UpdatedAt = time.Now().UTC()
}

// SetWebhook updates the webhook URL for the Item
func (i *Item) SetWebhook(webhookURL string) {
	i.WebhookURL = webhookURL
	i.UpdatedAt = time.Now().UTC()
}

// ItemStatus represents the possible statuses of an Item
const (
	ItemStatusActive            = "active"
	ItemStatusLoginRequired     = "login_required"
	ItemStatusErrored           = "errored"
	ItemStatusPendingExpiration = "pending_expiration"
	ItemStatusRevoked           = "revoked"
)
