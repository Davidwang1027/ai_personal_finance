package plaid

import (
	"context"
	"log"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/config"
	"github.com/plaid/plaid-go/v20/plaid"
)

// Client wraps the Plaid client and configuration
type Client struct {
	client *plaid.APIClient
	config *config.Config
}

// NewClient initializes and returns a new Plaid client
func NewClient(cfg *config.Config) *Client {
	// Configure Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", cfg.PlaidClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", cfg.PlaidSecret)

	// Set the environment based on configuration
	var env plaid.Environment
	switch cfg.PlaidEnv {
	case "sandbox":
		env = plaid.Sandbox
	case "development":
		env = plaid.Development
	case "production":
		env = plaid.Production
	default:
		log.Printf("Warning: Unknown Plaid environment '%s', defaulting to sandbox", cfg.PlaidEnv)
		env = plaid.Sandbox
	}

	// Use the environment
	configuration.UseEnvironment(env)

	// Create the Plaid API client
	client := plaid.NewAPIClient(configuration)

	return &Client{
		client: client,
		config: cfg,
	}
}

// GetClient returns the Plaid API client
func (c *Client) GetClient() *plaid.APIClient {
	return c.client
}

// CreateLinkToken generates a link token for initializing Plaid Link
func (c *Client) CreateLinkToken(userID string, clientName string, products []plaid.Products) (string, error) {
	ctx := context.Background()

	// Create user object
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: userID,
	}

	// Configure the link token create request
	request := plaid.NewLinkTokenCreateRequest(
		clientName,
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)

	// Add requested products
	request.SetProducts(products)

	// Add webhooks if needed (optional)
	// request.SetWebhook("https://your-domain.com/plaid/webhook")

	// Execute the request
	resp, _, err := c.client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		return "", err
	}

	return resp.GetLinkToken(), nil
}

// ExchangePublicToken exchanges a public token from Plaid Link for an access token
func (c *Client) ExchangePublicToken(publicToken string) (string, string, error) {
	ctx := context.Background()

	request := plaid.NewItemPublicTokenExchangeRequest(publicToken)
	resp, _, err := c.client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*request).Execute()
	if err != nil {
		return "", "", err
	}

	return resp.GetAccessToken(), resp.GetItemId(), nil
}

// GetAccounts retrieves account data for an Item
func (c *Client) GetAccounts(accessToken string) (*plaid.AccountsGetResponse, error) {
	ctx := context.Background()

	request := plaid.NewAccountsGetRequest(accessToken)
	resp, _, err := c.client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetTransactions retrieves transactions for a specific date range
func (c *Client) GetTransactions(accessToken string, startDate, endDate time.Time, options *plaid.TransactionsGetRequestOptions) (*plaid.TransactionsGetResponse, error) {
	ctx := context.Background()

	// Format dates in ISO 8601 format (YYYY-MM-DD)
	startDateStr := startDate.Format("2006-01-02")
	endDateStr := endDate.Format("2006-01-02")

	request := plaid.NewTransactionsGetRequest(accessToken, startDateStr, endDateStr)

	if options != nil {
		request.SetOptions(*options)
	}

	resp, _, err := c.client.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// SyncTransactions uses the newer transactions/sync endpoint
func (c *Client) SyncTransactions(accessToken string, cursor string) (*plaid.TransactionsSyncResponse, error) {
	ctx := context.Background()

	request := plaid.NewTransactionsSyncRequest(accessToken)

	if cursor != "" {
		request.SetCursor(cursor)
	}

	resp, _, err := c.client.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetItem retrieves item information
func (c *Client) GetItem(accessToken string) (*plaid.ItemGetResponse, error) {
	ctx := context.Background()

	request := plaid.NewItemGetRequest(accessToken)
	resp, _, err := c.client.PlaidApi.ItemGet(ctx).ItemGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateItemWebhook updates the webhook URL for an item
func (c *Client) UpdateItemWebhook(accessToken, webhookURL string) (*plaid.ItemWebhookUpdateResponse, error) {
	ctx := context.Background()

	request := plaid.NewItemWebhookUpdateRequest(accessToken)
	request.SetWebhook(webhookURL)
	resp, _, err := c.client.PlaidApi.ItemWebhookUpdate(ctx).ItemWebhookUpdateRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// VerifyWebhook verifies that a webhook is from Plaid
// This is a simplified implementation - in production you would use the Plaid-Verification header
func (c *Client) VerifyWebhook(body []byte) bool {
	// In a real implementation, you would:
	// 1. Extract the Plaid-Verification header from the request
	// 2. Use the verification key to verify the signature
	// 3. Check that the request is not expired

	// For now we'll just return true as this is for development
	return true
}
