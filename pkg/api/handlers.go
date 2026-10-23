package api

import (
	"net/http"
	"time"

	"github.com/example/cipher-arsenal/internal/store"
	"github.com/example/cipher-arsenal/pkg/price"
	walletpkg "github.com/example/cipher-arsenal/pkg/wallet"
	cryptopkg "github.com/example/cipher-arsenal/pkg/crypto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db      *gorm.DB
	cfg     *Config
	priceCl *price.Client
}

type Config struct {
	Port string
}

func NewHandler(db *gorm.DB, conf *Config) *Handler {
	pc := price.NewClient()
	return &Handler{db: db, cfg: conf, priceCl: pc}
}

func NewHandlerFromDB(db *gorm.DB, cfg *Config) *Handler {
	return NewHandler(db, cfg)
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", h.Health)
	rg.POST("/wallets", h.CreateWallet)
	rg.GET("/validate", h.ValidateAddress)
	rg.GET("/price/:symbol", h.GetPrice)
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().UTC()})
}

type createWalletReq struct {
	Network string `json:"network" binding:"required"
}

func (h *Handler) CreateWallet(c *gin.Context) {
	var req createWalletReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	priv, addr, err := walletpkg.Generate(req.Network)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// store address for audit
	h.db.Create(&store.Wallet{Type: req.Network, Address: addr})
	// WARNING: returning priv key is only for development/test scenarios
	c.JSON(http.StatusOK, gin.H{"network": req.Network, "address": addr, "private_key": priv})
}

func (h *Handler) ValidateAddress(c *gin.Context) {
	network := c.Query("network")
	address := c.Query("address")
	ok, err := cryptopkg.Validate(network, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"network": network, "address": address, "valid": ok})
}

func (h *Handler) GetPrice(c *gin.Context) {
	symbol := c.Param("symbol")
	price, err := h.priceCl.GetPriceUSD(symbol)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"symbol": symbol, "price_usd": price})
}
