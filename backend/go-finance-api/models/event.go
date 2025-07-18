package models

import (
	"time"
)

// EventType represents the type of event
type EventType string

const (
	// Event types
	EventTypePlaidAPI   EventType = "plaid_api"   // Plaid API requests/responses
	EventTypePlaidLink  EventType = "plaid_link"  // Plaid Link interactions
	EventTypeWebhook    EventType = "webhook"     // Plaid webhook events
	EventTypeUserAction EventType = "user_action" // User interactions
	EventTypeSystem     EventType = "system"      // System events
)

// Event represents a logged event
type Event struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id,omitempty"`
	ItemID      int64     `json:"item_id,omitempty"`
	Type        EventType `json:"type"`
	EventName   string    `json:"event_name"`
	RequestID   string    `json:"request_id,omitempty"` // Plaid request ID
	Status      string    `json:"status"`               // success, error
	Description string    `json:"description,omitempty"`
	Metadata    string    `json:"metadata,omitempty"` // JSON string of additional data
	CreatedAt   time.Time `json:"created_at"`
}

// PlaidAPIEvent represents a Plaid API event
type PlaidAPIEvent struct {
	Event
	Endpoint    string `json:"endpoint"` // Plaid endpoint called
	RequestBody string `json:"request_body,omitempty"`
	Response    string `json:"response,omitempty"`
	Error       string `json:"error,omitempty"`
}

// PlaidLinkEvent represents a Plaid Link event
type PlaidLinkEvent struct {
	Event
	LinkSessionID     string `json:"link_session_id"`
	LinkEventName     string `json:"link_event_name"`     // onSuccess, onExit, etc.
	LinkEventMetadata string `json:"link_event_metadata"` // JSON string of event metadata
	Error             string `json:"error,omitempty"`
}

// WebhookEvent represents a Plaid webhook event
type WebhookEvent struct {
	Event
	WebhookType string `json:"webhook_type"` // TRANSACTIONS, ITEM, etc.
	WebhookCode string `json:"webhook_code"` // DEFAULT_UPDATE, ERROR, etc.
	Payload     string `json:"payload"`      // JSON string of webhook payload
}

// EventStatus represents the status of an event
const (
	EventStatusSuccess = "success"
	EventStatusError   = "error"
	EventStatusWarning = "warning"
	EventStatusInfo    = "info"
)
