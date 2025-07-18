package models

import (
	"time"
)

// Account represents a financial account from Plaid
type Account struct {
	ID                     int64     `json:"id"`
	ItemID                 int64     `json:"item_id"`          // References our Item model
	PlaidAccountID         string    `json:"plaid_account_id"` // Plaid's account_id
	Name                   string    `json:"name"`
	OfficialName           string    `json:"official_name,omitempty"`
	Mask                   string    `json:"mask"`              // Last 4 digits
	Type                   string    `json:"type"`              // depository, credit, loan, investment, etc.
	Subtype                string    `json:"subtype,omitempty"` // checking, savings, etc.
	CurrentBalance         float64   `json:"current_balance"`
	AvailableBalance       float64   `json:"available_balance"`
	CreditLimit            float64   `json:"credit_limit,omitempty"`
	IsoCurrencyCode        string    `json:"iso_currency_code"`
	UnofficialCurrencyCode string    `json:"unofficial_currency_code,omitempty"`
	LastUpdated            time.Time `json:"last_updated"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

// AccountType represents the type of account
const (
	AccountTypeDepository = "depository"
	AccountTypeCredit     = "credit"
	AccountTypeLoan       = "loan"
	AccountTypeInvestment = "investment"
	AccountTypeOther      = "other"
)

// AccountSubtype represents the subtype of account
const (
	AccountSubtypeChecking      = "checking"
	AccountSubtypeSavings       = "savings"
	AccountSubtypeCreditCard    = "credit card"
	AccountSubtypeMortgage      = "mortgage"
	AccountSubtypeStudentLoan   = "student loan"
	AccountSubtypenvestment401k = "401k"
)
