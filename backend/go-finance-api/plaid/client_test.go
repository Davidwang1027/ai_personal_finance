package plaid

import (
	"testing"

	"github.com/davidwang/go-finance-api/go-finance-api/config"
)

// TestNewClient tests the client initialization
func TestNewClient(t *testing.T) {
	// Load config
	cfg := config.Load()

	// Validate required config fields
	if cfg.PlaidClientID == "" {
		t.Skip("PLAID_CLIENT_ID not set in environment, skipping test")
	}
	if cfg.PlaidSecret == "" {
		t.Skip("PLAID_SECRET not set in environment, skipping test")
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
