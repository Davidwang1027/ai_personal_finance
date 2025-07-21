package plaid

import (
	"testing"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/config"
	plaidlib "github.com/plaid/plaid-go/v20/plaid"
)

// TestNewClient tests the client initialization
func TestNewClient(t *testing.T) {
	// Load config
	cfg := config.Load()

	// If environment variables are not set, set them for testing
	if cfg.PlaidClientID == "" {
		cfg.PlaidClientID = "test_client_id"
	}
	if cfg.PlaidSecret == "" {
		cfg.PlaidSecret = "test_secret"
	}

	// Create client
	client := NewClient(cfg)
	if client == nil {
		t.Fatal("Failed to create Plaid client")
	}

	// Verify client is initialized
	if client.GetClient() == nil {
		t.Error("Plaid API client is nil")
	}

	// Check environment setting
	switch cfg.PlaidEnv {
	case "sandbox", "development", "production":
		// valid environments
	default:
		t.Logf("Warning: Using non-standard Plaid environment: %s", cfg.PlaidEnv)
	}

	t.Log("Plaid client initialized successfully")
}

// setupTestConfig creates a test config with mock values for testing
func setupTestConfig() *config.Config {
	cfg := config.Load()

	// Set test values if environment variables are not set
	if cfg.PlaidClientID == "" {
		cfg.PlaidClientID = "test_client_id"
	}
	if cfg.PlaidSecret == "" {
		cfg.PlaidSecret = "test_secret"
	}
	if cfg.PlaidEnv == "" {
		cfg.PlaidEnv = "sandbox"
	}

	return cfg
}

// TestGetTransactions tests the GetTransactions method
func TestGetTransactions(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test - in a real test we'd use a mock or test token
	// For now, we'll just check that the function doesn't panic when called with invalid token
	accessToken := "invalid_access_token"
	startDate := time.Now().AddDate(0, -3, 0) // 3 months ago
	endDate := time.Now()
	options := &plaidlib.TransactionsGetRequestOptions{
		Count:  plaidlib.PtrInt32(10),
		Offset: plaidlib.PtrInt32(0),
	}

	// This should fail with an auth error, but shouldn't panic
	_, err := client.GetTransactions(accessToken, startDate, endDate, options)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestSyncTransactions tests the SyncTransactions method
func TestSyncTransactions(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.SyncTransactions(accessToken, "")
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestGetItem tests the GetItem method
func TestGetItem(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.GetItem(accessToken)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestUpdateItemWebhook tests the UpdateItemWebhook method
func TestUpdateItemWebhook(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"
	webhookURL := "https://example.com/webhook"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.UpdateItemWebhook(accessToken, webhookURL)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestVerifyWebhook tests the VerifyWebhook method
func TestVerifyWebhook(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// Simple test to check that function runs without error
	body := []byte(`{"webhook_type": "TRANSACTIONS", "webhook_code": "DEFAULT_UPDATE"}`)
	result := client.VerifyWebhook(body)

	// In our simple implementation, this should always be true
	if !result {
		t.Error("Expected VerifyWebhook to return true, but got false")
	}
}

}

// setupTestConfig creates a test config with mock values for testing
func setupTestConfig() *config.Config {
	cfg := config.Load()

	// Set test values if environment variables are not set
	if cfg.PlaidClientID == "" {
		cfg.PlaidClientID = "test_client_id"
	}
	if cfg.PlaidSecret == "" {
		cfg.PlaidSecret = "test_secret"
	}
	if cfg.PlaidEnv == "" {
		cfg.PlaidEnv = "sandbox"
	}

	return cfg
}

// TestGetTransactions tests the GetTransactions method
func TestGetTransactions(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test - in a real test we'd use a mock or test token
	// For now, we'll just check that the function doesn't panic when called with invalid token
	accessToken := "invalid_access_token"
	startDate := time.Now().AddDate(0, -3, 0) // 3 months ago
	endDate := time.Now()
	options := &plaidlib.TransactionsGetRequestOptions{
		Count:  plaidlib.PtrInt32(10),
		Offset: plaidlib.PtrInt32(0),
	}

	// This should fail with an auth error, but shouldn't panic
	_, err := client.GetTransactions(accessToken, startDate, endDate, options)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestSyncTransactions tests the SyncTransactions method
func TestSyncTransactions(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.SyncTransactions(accessToken, "")
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestGetItem tests the GetItem method
func TestGetItem(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.GetItem(accessToken)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestUpdateItemWebhook tests the UpdateItemWebhook method
func TestUpdateItemWebhook(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// We need an access token for this test
	accessToken := "invalid_access_token"
	webhookURL := "https://example.com/webhook"

	// This should fail with an auth error, but shouldn't panic
	_, err := client.UpdateItemWebhook(accessToken, webhookURL)
	if err == nil {
		t.Error("Expected error with invalid access token, but got none")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// TestVerifyWebhook tests the VerifyWebhook method
func TestVerifyWebhook(t *testing.T) {
	// Setup test config
	cfg := setupTestConfig()

	// Create client
	client := NewClient(cfg)

	// Simple test to check that function runs without error
	body := []byte(`{"webhook_type": "TRANSACTIONS", "webhook_code": "DEFAULT_UPDATE"}`)
	result := client.VerifyWebhook(body)

	// In our simple implementation, this should always be true
	if !result {
		t.Error("Expected VerifyWebhook to return true, but got false")
	}
}
