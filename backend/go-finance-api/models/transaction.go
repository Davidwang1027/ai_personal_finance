package models

import (
	"time"

	"github.com/google/uuid"
)

// Transaction represents a financial transaction from Plaid
type Transaction struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	AccountID          uuid.UUID `json:"account_id" db:"account_id"`
	UserID             uuid.UUID `json:"user_id" db:"user_id"`
	PlaidTransactionID string    `json:"plaid_transaction_id" db:"plaid_transaction_id"`
	CategoryID         string    `json:"category_id" db:"category_id"`
	Category           []string  `json:"category" db:"category"`
	Name               string    `json:"name" db:"name"`
	MerchantName       string    `json:"merchant_name" db:"merchant_name"`
	Amount             float64   `json:"amount" db:"amount"`
	IsoCurrencyCode    string    `json:"iso_currency_code" db:"iso_currency_code"`
	Date               time.Time `json:"date" db:"date"`
	Pending            bool      `json:"pending" db:"pending"`
	PaymentChannel     string    `json:"payment_channel" db:"payment_channel"`
	Address            string    `json:"address" db:"address"`
	City               string    `json:"city" db:"city"`
	Region             string    `json:"region" db:"region"`
	PostalCode         string    `json:"postal_code" db:"postal_code"`
	Country            string    `json:"country" db:"country"`
	Latitude           float64   `json:"latitude" db:"latitude"`
	Longitude          float64   `json:"longitude" db:"longitude"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// NewTransaction creates a new Transaction record
func NewTransaction(
	accountID uuid.UUID,
	userID uuid.UUID,
	plaidTransactionID string,
	categoryID string,
	category []string,
	name string,
	merchantName string,
	amount float64,
	isoCurrencyCode string,
	date time.Time,
	pending bool,
	paymentChannel string,
) *Transaction {
	now := time.Now().UTC()
	return &Transaction{
		ID:                 uuid.New(),
		AccountID:          accountID,
		UserID:             userID,
		PlaidTransactionID: plaidTransactionID,
		CategoryID:         categoryID,
		Category:           category,
		Name:               name,
		MerchantName:       merchantName,
		Amount:             amount,
		IsoCurrencyCode:    isoCurrencyCode,
		Date:               date,
		Pending:            pending,
		PaymentChannel:     paymentChannel,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

// SetLocation sets the transaction location details
func (t *Transaction) SetLocation(address, city, region, postalCode, country string, latitude, longitude float64) {
	t.Address = address
	t.City = city
	t.Region = region
	t.PostalCode = postalCode
	t.Country = country
	t.Latitude = latitude
	t.Longitude = longitude
	t.UpdatedAt = time.Now().UTC()
}

// UpdatePendingStatus updates the pending status of a transaction
func (t *Transaction) UpdatePendingStatus(pending bool) {
	t.Pending = pending
	t.UpdatedAt = time.Now().UTC()
}
