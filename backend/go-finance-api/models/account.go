package models

import (
	"time"

	"github.com/google/uuid"
)

// Account represents a financial account from Plaid
type Account struct {
	ID               uuid.UUID `json:"id" db:"id"`
	ItemID           uuid.UUID `json:"item_id" db:"item_id"`
	UserID           uuid.UUID `json:"user_id" db:"user_id"`
	PlaidAccountID   string    `json:"plaid_account_id" db:"plaid_account_id"`
	Name             string    `json:"name" db:"name"`
	OfficialName     string    `json:"official_name" db:"official_name"`
	Type             string    `json:"type" db:"type"`
	Subtype          string    `json:"subtype" db:"subtype"`
	Mask             string    `json:"mask" db:"mask"`
	AvailableBalance float64   `json:"available_balance" db:"available_balance"`
	CurrentBalance   float64   `json:"current_balance" db:"current_balance"`
	CurrencyCode     string    `json:"currency_code" db:"currency_code"`
	LastUpdated      time.Time `json:"last_updated" db:"last_updated"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// NewAccount creates a new Account record
func NewAccount(
	itemID uuid.UUID,
	userID uuid.UUID,
	plaidAccountID string,
	name string,
	officialName string,
	accountType string,
	accountSubtype string,
	mask string,
	availableBalance float64,
	currentBalance float64,
	currencyCode string,
) *Account {
	now := time.Now().UTC()
	return &Account{
		ID:               uuid.New(),
		ItemID:           itemID,
		UserID:           userID,
		PlaidAccountID:   plaidAccountID,
		Name:             name,
		OfficialName:     officialName,
		Type:             accountType,
		Subtype:          accountSubtype,
		Mask:             mask,
		AvailableBalance: availableBalance,
		CurrentBalance:   currentBalance,
		CurrencyCode:     currencyCode,
		LastUpdated:      now,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}

// UpdateBalances updates the account balances
func (a *Account) UpdateBalances(availableBalance, currentBalance float64) {
	now := time.Now().UTC()
	a.AvailableBalance = availableBalance
	a.CurrentBalance = currentBalance
	a.LastUpdated = now
	a.UpdatedAt = now
}
