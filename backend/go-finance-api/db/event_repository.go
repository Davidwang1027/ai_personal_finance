package db

import (
	"github.com/davidwang/go-finance-api/go-finance-api/models"
	"github.com/google/uuid"
)

// PlaidAPIEventRepository handles database operations for Plaid API events
type PlaidAPIEventRepository struct {
	db *Database
}

// NewPlaidAPIEventRepository creates a new PlaidAPIEventRepository
func NewPlaidAPIEventRepository(db *Database) *PlaidAPIEventRepository {
	return &PlaidAPIEventRepository{db: db}
}

// Create logs a new Plaid API event
func (r *PlaidAPIEventRepository) Create(event *models.PlaidAPIEvent) error {
	query := `
		INSERT INTO plaid_api_events (
			id, user_id, item_id, endpoint, request_body, response_body, status_code,
			error_code, error_message, request_id, request_time, response_time, 
			execution_time, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := r.db.Exec(
		query,
		event.ID,
		event.UserID,
		event.ItemID,
		event.Endpoint,
		event.RequestBody,
		event.ResponseBody,
		event.StatusCode,
		event.ErrorCode,
		event.ErrorMessage,
		event.RequestID,
		event.RequestTime,
		event.ResponseTime,
		event.ExecutionTime,
		event.CreatedAt,
	)
	return err
}

// GetByUserID retrieves API events for a specific user
func (r *PlaidAPIEventRepository) GetByUserID(userID uuid.UUID, limit, offset int) ([]*models.PlaidAPIEvent, error) {
	query := `
		SELECT 
			id, user_id, item_id, endpoint, request_body, response_body, status_code,
			error_code, error_message, request_id, request_time, response_time, 
			execution_time, created_at
		FROM plaid_api_events
		WHERE user_id = $1
		ORDER BY request_time DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.PlaidAPIEvent
	for rows.Next() {
		var event models.PlaidAPIEvent
		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.ItemID,
			&event.Endpoint,
			&event.RequestBody,
			&event.ResponseBody,
			&event.StatusCode,
			&event.ErrorCode,
			&event.ErrorMessage,
			&event.RequestID,
			&event.RequestTime,
			&event.ResponseTime,
			&event.ExecutionTime,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// GetByItemID retrieves API events for a specific item
func (r *PlaidAPIEventRepository) GetByItemID(itemID uuid.UUID, limit, offset int) ([]*models.PlaidAPIEvent, error) {
	query := `
		SELECT 
			id, user_id, item_id, endpoint, request_body, response_body, status_code,
			error_code, error_message, request_id, request_time, response_time, 
			execution_time, created_at
		FROM plaid_api_events
		WHERE item_id = $1
		ORDER BY request_time DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, itemID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.PlaidAPIEvent
	for rows.Next() {
		var event models.PlaidAPIEvent
		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.ItemID,
			&event.Endpoint,
			&event.RequestBody,
			&event.ResponseBody,
			&event.StatusCode,
			&event.ErrorCode,
			&event.ErrorMessage,
			&event.RequestID,
			&event.RequestTime,
			&event.ResponseTime,
			&event.ExecutionTime,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// LinkEventRepository handles database operations for Plaid Link events
type LinkEventRepository struct {
	db *Database
}

// NewLinkEventRepository creates a new LinkEventRepository
func NewLinkEventRepository(db *Database) *LinkEventRepository {
	return &LinkEventRepository{db: db}
}

// Create logs a new Link event
func (r *LinkEventRepository) Create(event *models.LinkEvent) error {
	query := `
		INSERT INTO link_events (
			id, user_id, item_id, event_name, event_metadata, link_session_id,
			request_id, error_code, error_message, status, institution_id,
			institution_name, timestamp, created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := r.db.Exec(
		query,
		event.ID,
		event.UserID,
		event.ItemID,
		event.EventName,
		event.EventMetadata,
		event.LinkSessionID,
		event.RequestID,
		event.ErrorCode,
		event.ErrorMessage,
		event.Status,
		event.InstitutionID,
		event.InstitutionName,
		event.Timestamp,
		event.CreatedAt,
	)
	return err
}

// GetByUserID retrieves Link events for a specific user
func (r *LinkEventRepository) GetByUserID(userID uuid.UUID, limit, offset int) ([]*models.LinkEvent, error) {
	query := `
		SELECT 
			id, user_id, item_id, event_name, event_metadata, link_session_id,
			request_id, error_code, error_message, status, institution_id,
			institution_name, timestamp, created_at
		FROM link_events
		WHERE user_id = $1
		ORDER BY timestamp DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.LinkEvent
	for rows.Next() {
		var event models.LinkEvent
		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.ItemID,
			&event.EventName,
			&event.EventMetadata,
			&event.LinkSessionID,
			&event.RequestID,
			&event.ErrorCode,
			&event.ErrorMessage,
			&event.Status,
			&event.InstitutionID,
			&event.InstitutionName,
			&event.Timestamp,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// GetByItemID retrieves Link events for a specific item
func (r *LinkEventRepository) GetByItemID(itemID uuid.UUID, limit, offset int) ([]*models.LinkEvent, error) {
	query := `
		SELECT 
			id, user_id, item_id, event_name, event_metadata, link_session_id,
			request_id, error_code, error_message, status, institution_id,
			institution_name, timestamp, created_at
		FROM link_events
		WHERE item_id = $1
		ORDER BY timestamp DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.Query(query, itemID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.LinkEvent
	for rows.Next() {
		var event models.LinkEvent
		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.ItemID,
			&event.EventName,
			&event.EventMetadata,
			&event.LinkSessionID,
			&event.RequestID,
			&event.ErrorCode,
			&event.ErrorMessage,
			&event.Status,
			&event.InstitutionID,
			&event.InstitutionName,
			&event.Timestamp,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// GetByLinkSessionID retrieves Link events for a specific Link session
func (r *LinkEventRepository) GetByLinkSessionID(linkSessionID string) ([]*models.LinkEvent, error) {
	query := `
		SELECT 
			id, user_id, item_id, event_name, event_metadata, link_session_id,
			request_id, error_code, error_message, status, institution_id,
			institution_name, timestamp, created_at
		FROM link_events
		WHERE link_session_id = $1
		ORDER BY timestamp DESC
	`
	rows, err := r.db.Query(query, linkSessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.LinkEvent
	for rows.Next() {
		var event models.LinkEvent
		err := rows.Scan(
			&event.ID,
			&event.UserID,
			&event.ItemID,
			&event.EventName,
			&event.EventMetadata,
			&event.LinkSessionID,
			&event.RequestID,
			&event.ErrorCode,
			&event.ErrorMessage,
			&event.Status,
			&event.InstitutionID,
			&event.InstitutionName,
			&event.Timestamp,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
