package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPlaidClient is a mock implementation of the Plaid client
type MockPlaidClient struct {
	mock.Mock
}

// CreateLinkToken mocks the CreateLinkToken method
func (m *MockPlaidClient) CreateLinkToken(userID string, clientName string, products []interface{}) (string, error) {
	args := m.Called(userID, clientName, products)
	return args.String(0), args.Error(1)
}

// ExchangePublicToken mocks the ExchangePublicToken method
func (m *MockPlaidClient) ExchangePublicToken(publicToken string) (string, string, error) {
	args := m.Called(publicToken)
	return args.String(0), args.String(1), args.Error(2)
}

// GetAccounts mocks the GetAccounts method
func (m *MockPlaidClient) GetAccounts(accessToken string) (interface{}, error) {
	args := m.Called(accessToken)
	return args.Get(0), args.Error(1)
}

// GetTransactions mocks the GetTransactions method
func (m *MockPlaidClient) GetTransactions(accessToken string, startDate, endDate time.Time, options interface{}) (interface{}, error) {
	args := m.Called(accessToken, startDate, endDate, options)
	return args.Get(0), args.Error(1)
}

// SyncTransactions mocks the SyncTransactions method
func (m *MockPlaidClient) SyncTransactions(accessToken string, cursor string) (interface{}, error) {
	args := m.Called(accessToken, cursor)
	return args.Get(0), args.Error(1)
}

// GetItem mocks the GetItem method
func (m *MockPlaidClient) GetItem(accessToken string) (interface{}, error) {
	args := m.Called(accessToken)
	return args.Get(0), args.Error(1)
}

// UpdateItemWebhook mocks the UpdateItemWebhook method
func (m *MockPlaidClient) UpdateItemWebhook(accessToken, webhookURL string) (interface{}, error) {
	args := m.Called(accessToken, webhookURL)
	return args.Get(0), args.Error(1)
}

// VerifyWebhook mocks the VerifyWebhook method
func (m *MockPlaidClient) VerifyWebhook(body []byte) bool {
	args := m.Called(body)
	return args.Bool(0)
}

// SetupRouter creates a test router for testing handlers
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

// TestCreateLinkToken tests the CreateLinkToken handler
func TestCreateLinkToken(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.POST("/api/plaid/create_link_token", handler.CreateLinkToken)

	// Mock client behavior
	mockClient.On("CreateLinkToken", mock.Anything, mock.Anything, mock.Anything).Return("test_link_token", nil)

	// Create request
	reqBody := CreateLinkTokenRequest{
		ClientUserID: "test_user",
		ClientName:   "Test App",
		Products:     []string{"transactions"},
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/plaid/create_link_token", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test_link_token", response["link_token"])
}

// TestGetTransactions tests the GetTransactions handler
func TestGetTransactions(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.POST("/api/plaid/transactions", handler.GetTransactions)

	// Mock client behavior - returning a simple struct to simulate response
	mockResponse := map[string]interface{}{
		"transactions": []map[string]interface{}{
			{
				"id":             "tx1",
				"amount":         100.0,
				"date":           "2025-07-01",
				"name":           "Test Transaction",
				"payment_method": "credit",
			},
		},
	}
	mockClient.On("GetTransactions", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil)

	// Create request
	reqBody := GetTransactionsRequest{
		AccessToken: "test_access_token",
		StartDate:   "2025-06-01",
		EndDate:     "2025-07-01",
		Count:       10,
		Offset:      0,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/plaid/transactions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestSyncTransactions tests the SyncTransactions handler
func TestSyncTransactions(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.POST("/api/plaid/transactions/sync", handler.SyncTransactions)

	// Mock client behavior - returning a simple struct to simulate response
	mockResponse := map[string]interface{}{
		"added": []map[string]interface{}{
			{
				"id":     "tx1",
				"amount": 100.0,
				"date":   "2025-07-01",
				"name":   "Test Transaction",
			},
		},
		"modified":    []map[string]interface{}{},
		"removed":     []map[string]interface{}{},
		"next_cursor": "next-cursor-value",
	}
	mockClient.On("SyncTransactions", mock.Anything, mock.Anything).Return(mockResponse, nil)

	// Create request
	reqBody := SyncTransactionsRequest{
		AccessToken: "test_access_token",
		Cursor:      "test_cursor",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/plaid/transactions/sync", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestGetItem tests the GetItem handler
func TestGetItem(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.GET("/api/plaid/item", handler.GetItem)

	// Mock client behavior
	mockResponse := map[string]interface{}{
		"item": map[string]interface{}{
			"item_id":            "item-id-123",
			"institution_id":     "ins_123",
			"available_products": []string{"transactions", "auth"},
			"billed_products":    []string{"identity"},
		},
	}
	mockClient.On("GetItem", mock.Anything).Return(mockResponse, nil)

	// Create request
	req := httptest.NewRequest("GET", "/api/plaid/item?access_token=test_access_token", nil)

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestUpdateItemWebhook tests the UpdateItemWebhook handler
func TestUpdateItemWebhook(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.POST("/api/plaid/item/webhook", handler.UpdateItemWebhook)

	// Mock client behavior
	mockResponse := map[string]interface{}{
		"item": map[string]interface{}{
			"webhook": "https://example.com/webhook",
		},
	}
	mockClient.On("UpdateItemWebhook", mock.Anything, mock.Anything).Return(mockResponse, nil)

	// Create request
	reqBody := UpdateItemWebhookRequest{
		AccessToken: "test_access_token",
		WebhookURL:  "https://example.com/webhook",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/plaid/item/webhook", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestHandleWebhook tests the HandleWebhook handler
func TestHandleWebhook(t *testing.T) {
	// Set up router and handler with mock client
	router := SetupRouter()
	mockClient := new(MockPlaidClient)
	handler := &PlaidHandler{plaidClient: mockClient}

	// Setup route
	router.POST("/api/plaid/webhook", handler.HandleWebhook)

	// Mock client behavior
	mockClient.On("VerifyWebhook", mock.Anything).Return(true)

	// Create webhook payload
	webhookBody := map[string]interface{}{
		"webhook_type": "TRANSACTIONS",
		"webhook_code": "DEFAULT_UPDATE",
		"item_id":      "item-id-123",
	}
	jsonBody, _ := json.Marshal(webhookBody)
	req := httptest.NewRequest("POST", "/api/plaid/webhook", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request and record response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert response
	assert.Equal(t, http.StatusOK, resp.Code)
}
