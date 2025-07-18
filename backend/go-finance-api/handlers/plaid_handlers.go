package handlers

import (
	"net/http"

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
