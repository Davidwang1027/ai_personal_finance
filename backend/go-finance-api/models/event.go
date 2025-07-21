package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// PlaidAPIEvent represents a logged Plaid API request and response
type PlaidAPIEvent struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	UserID        uuid.UUID       `json:"user_id" db:"user_id"`
	ItemID        *uuid.UUID      `json:"item_id" db:"item_id"` // Can be null for some operations
	Endpoint      string          `json:"endpoint" db:"endpoint"`
	RequestBody   json.RawMessage `json:"request_body" db:"request_body"`
	ResponseBody  json.RawMessage `json:"response_body" db:"response_body"`
	StatusCode    int             `json:"status_code" db:"status_code"`
	ErrorCode     *string         `json:"error_code" db:"error_code"`
	ErrorMessage  *string         `json:"error_message" db:"error_message"`
	RequestID     string          `json:"request_id" db:"request_id"`
	RequestTime   time.Time       `json:"request_time" db:"request_time"`
	ResponseTime  time.Time       `json:"response_time" db:"response_time"`
	ExecutionTime int             `json:"execution_time" db:"execution_time"` // in ms
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}

// NewPlaidAPIEvent creates a new Plaid API event log
func NewPlaidAPIEvent(
	userID uuid.UUID,
	itemID *uuid.UUID,
	endpoint string,
	requestBody json.RawMessage,
	requestTime time.Time,
) *PlaidAPIEvent {
	now := time.Now().UTC()
	return &PlaidAPIEvent{
		ID:          uuid.New(),
		UserID:      userID,
		ItemID:      itemID,
		Endpoint:    endpoint,
		RequestBody: requestBody,
		RequestTime: requestTime,
		CreatedAt:   now,
	}
}

// SetResponse sets the response data for an API event
func (e *PlaidAPIEvent) SetResponse(
	responseBody json.RawMessage,
	statusCode int,
	requestID string,
	errorCode *string,
	errorMessage *string,
) {
	responseTime := time.Now().UTC()
	e.ResponseBody = responseBody
	e.StatusCode = statusCode
	e.RequestID = requestID
	e.ErrorCode = errorCode
	e.ErrorMessage = errorMessage
	e.ResponseTime = responseTime
	e.ExecutionTime = int(responseTime.Sub(e.RequestTime).Milliseconds())
}

// LinkEvent represents a Plaid Link event triggered by user interaction
type LinkEvent struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	UserID          uuid.UUID       `json:"user_id" db:"user_id"`
	ItemID          *uuid.UUID      `json:"item_id" db:"item_id"` // Can be null for initial link events
	EventName       string          `json:"event_name" db:"event_name"`
	EventMetadata   json.RawMessage `json:"event_metadata" db:"event_metadata"`
	LinkSessionID   string          `json:"link_session_id" db:"link_session_id"`
	RequestID       string          `json:"request_id" db:"request_id"`
	ErrorCode       *string         `json:"error_code" db:"error_code"`
	ErrorMessage    *string         `json:"error_message" db:"error_message"`
	Status          string          `json:"status" db:"status"`
	InstitutionID   *string         `json:"institution_id" db:"institution_id"`
	InstitutionName *string         `json:"institution_name" db:"institution_name"`
	Timestamp       time.Time       `json:"timestamp" db:"timestamp"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
}

// NewLinkEvent creates a new Link event log
func NewLinkEvent(
	userID uuid.UUID,
	itemID *uuid.UUID,
	eventName string,
	eventMetadata json.RawMessage,
	linkSessionID string,
	requestID string,
	status string,
	timestamp time.Time,
) *LinkEvent {
	now := time.Now().UTC()
	return &LinkEvent{
		ID:            uuid.New(),
		UserID:        userID,
		ItemID:        itemID,
		EventName:     eventName,
		EventMetadata: eventMetadata,
		LinkSessionID: linkSessionID,
		RequestID:     requestID,
		Status:        status,
		Timestamp:     timestamp,
		CreatedAt:     now,
	}
}

// SetError adds error information to a Link event
func (e *LinkEvent) SetError(errorCode, errorMessage string) {
	e.ErrorCode = &errorCode
	e.ErrorMessage = &errorMessage
}

// SetInstitution adds institution information to a Link event
func (e *LinkEvent) SetInstitution(institutionID, institutionName string) {
	e.InstitutionID = &institutionID
	e.InstitutionName = &institutionName
}
