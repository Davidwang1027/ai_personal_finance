package plaid

import (
	"context"
	"log"

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
