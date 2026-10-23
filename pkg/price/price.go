package price

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Simple CoinGecko adapter (no third-party libs)
type Client struct{
	http *http.Client
	cache map[string]cached
}

type cached struct{
	price float64
	at time.Time
}

func NewClient() *Client{
	return &Client{http: &http.Client{Timeout: 5*time.Second}, cache: make(map[string]cached)}
}

var symbolMap = map[string]string{
	"btc":"bitcoin",
	"eth":"ethereum",
}

func (c *Client) GetPriceUSD(symbol string) (float64, error) {
	sym := strings.ToLower(symbol)
	if s, ok := c.cache[sym]; ok && time.Since(s.at) < 60*time.Second {
		return s.price, nil
	}
	id, ok := symbolMap[sym]
	if !ok {
		return 0, errors.New("unsupported symbol")
	}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", id)
	resp, err := c.http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("coingecko returned %d", resp.StatusCode)
	}
	var body map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, err
	}
	p, ok := body[id]["usd"]
	if !ok {
		return 0, errors.New("unexpected response")
	}
	// cache
	c.cache[sym] = cached{price: p, at: time.Now()}
	return p, nil
}
