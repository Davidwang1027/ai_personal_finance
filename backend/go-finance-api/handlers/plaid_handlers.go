package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/davidwang/go-finance-api/go-finance-api/plaid"
	"github.com/gin-gonic/gin"
	plaidlib "github.com/plaid/plaid-go/v20/plaid"
)

// PlaidHandler handles Plaid API related requests
type PlaidHandler struct {
	plaidClient *plaid.Client
}

// NewPlaidHandler creates a new PlaidHandler
func NewPlaidHandler(plaidClient *plaid.Client) *PlaidHandler {
	return &PlaidHandler{
		plaidClient: plaidClient,
	}
}

// CreateLinkTokenRequest is the request body for creating a link token
type CreateLinkTokenRequest struct {
	ClientUserID string   `json:"client_user_id" binding:"required"`
	ClientName   string   `json:"client_name" binding:"required"`
	Products     []string `json:"products"`
}

// CreateLinkToken creates a link token for Plaid Link
func (h *PlaidHandler) CreateLinkToken(c *gin.Context) {
	var req CreateLinkTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string products to Plaid products enum
	products := make([]plaidlib.Products, 0, len(req.Products))
	for _, p := range req.Products {
		switch p {
		case "auth":
			products = append(products, plaidlib.PRODUCTS_AUTH)
		case "transactions":
			products = append(products, plaidlib.PRODUCTS_TRANSACTIONS)
		case "identity":
			products = append(products, plaidlib.PRODUCTS_IDENTITY)
		case "investments":
			products = append(products, plaidlib.PRODUCTS_INVESTMENTS)
		case "liabilities":
			products = append(products, plaidlib.PRODUCTS_LIABILITIES)
		case "assets":
			products = append(products, plaidlib.PRODUCTS_ASSETS)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product: " + p})
			return
		}
	}

	linkToken, err := h.plaidClient.CreateLinkToken(req.ClientUserID, req.ClientName, products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
}

// ExchangePublicTokenRequest is the request body for exchanging a public token
type ExchangePublicTokenRequest struct {
	PublicToken string `json:"public_token" binding:"required"`
}

// ExchangePublicToken exchanges a public token for an access token
func (h *PlaidHandler) ExchangePublicToken(c *gin.Context) {
	var req ExchangePublicTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, itemID, err := h.plaidClient.ExchangePublicToken(req.PublicToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// In a real app, you'd save these values to your database
	// associated with the current user
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})
}

// GetAccounts retrieves accounts for an access token
func (h *PlaidHandler) GetAccounts(c *gin.Context) {
	// In a real app, you'd get the access token from your database
	// based on the authenticated user
	accessToken := c.Query("access_token")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token is required"})
		return
	}

	accounts, err := h.plaidClient.GetAccounts(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// GetTransactionsRequest is the request body for getting transactions
type GetTransactionsRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"` // YYYY-MM-DD
	EndDate     string `json:"end_date" binding:"required"`   // YYYY-MM-DD
	Count       int32  `json:"count"`
	Offset      int32  `json:"offset"`
}

// GetTransactions retrieves transactions for the specified date range
func (h *PlaidHandler) GetTransactions(c *gin.Context) {
	var req GetTransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
		return
	}

	// Create options object if count or offset is specified
	var options *plaidlib.TransactionsGetRequestOptions
	if req.Count > 0 || req.Offset > 0 {
		options = &plaidlib.TransactionsGetRequestOptions{}
		if req.Count > 0 {
			options.SetCount(req.Count)
		}
		if req.Offset > 0 {
			options.SetOffset(req.Offset)
		}
	}

	transactions, err := h.plaidClient.GetTransactions(req.AccessToken, startDate, endDate, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// SyncTransactionsRequest is the request body for syncing transactions
type SyncTransactionsRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
	Cursor      string `json:"cursor"`
}

// SyncTransactions uses the transactions/sync endpoint to get transaction updates
func (h *PlaidHandler) SyncTransactions(c *gin.Context) {
	var req SyncTransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	syncResponse, err := h.plaidClient.SyncTransactions(req.AccessToken, req.Cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, syncResponse)
}

// GetItem retrieves item information
func (h *PlaidHandler) GetItem(c *gin.Context) {
	accessToken := c.Query("access_token")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token is required"})
		return
	}

	item, err := h.plaidClient.GetItem(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// UpdateItemWebhookRequest is the request body for updating an item's webhook
type UpdateItemWebhookRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
	WebhookURL  string `json:"webhook_url" binding:"required"`
}

// UpdateItemWebhook updates the webhook URL for an item
func (h *PlaidHandler) UpdateItemWebhook(c *gin.Context) {
	var req UpdateItemWebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.plaidClient.UpdateItemWebhook(req.AccessToken, req.WebhookURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// HandleWebhook processes webhooks from Plaid
func (h *PlaidHandler) HandleWebhook(c *gin.Context) {
	// Read the body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// In production, you would verify the webhook using the Plaid-Verification header
	if !h.plaidClient.VerifyWebhook(body) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid webhook"})
		return
	}

	// Parse the webhook
	var webhookData map[string]interface{}
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid webhook data"})
		return
	}

	// Process different webhook types
	webhookType, _ := webhookData["webhook_type"].(string)
	webhookCode, _ := webhookData["webhook_code"].(string)

	// Here you would process the webhook based on its type and code
	// For example:
	switch webhookType {
	case "TRANSACTIONS":
		// Process transaction webhooks
		switch webhookCode {
		case "DEFAULT_UPDATE":
			// New transactions available
			// In a real implementation, you'd use the transactions/sync endpoint to get the updates
		case "HISTORICAL_UPDATE":
			// Historical transactions available
		case "TRANSACTIONS_REMOVED":
			// Transactions have been removed
		}
	case "ITEM":
		// Process item webhooks
		switch webhookCode {
		case "ERROR":
			// Item error occurred
		case "PENDING_EXPIRATION":
			// Access token is expiring soon
		case "USER_PERMISSION_REVOKED":
			// User revoked access
		}
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"status": "webhook received"})
}
