package store

import "time"

// Wallet represents a stored wallet/address entry (for tracking and auditing)
type Wallet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"index" json:"type"` // e.g. eth, btc
	Address   string    `gorm:"uniqueIndex" json:"address"`
	Meta      string    `json:"meta"` // JSON or notes
	CreatedAt time.Time `json:"created_at"`
}

// PriceCache is a very small cache model used by price adapter
type PriceCache struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Symbol    string    `gorm:"index" json:"symbol"`
	PriceUSD  float64   `json:"price_usd"`
	UpdatedAt time.Time `json:"updated_at"`
}
