package api

import (
	"sync"
	"time"
)

// SnipeFireRequest — multi-wallet snipe via RPC
type SnipeFireRequest struct {
	TokenID     string   `json:"tokenId"`
	Wallets     []string `json:"wallets"`
	BuyPercent  float64  `json:"buyPercent"`
	SlippageBps uint16   `json:"slippageBps"`
	DevBuySOL   float64  `json:"devBuySOL"`
}

// BundleSnipeFireRequest — Jito MEV-protected bundle snipe
type BundleSnipeFireRequest struct {
	TokenID            string   `json:"tokenId"`
	Wallets            []string `json:"wallets"`
	JitoTipSOL         float64  `json:"jitoTipSOL"`
	RpcFallbackDelayMs int64    `json:"rpcFallbackDelayMs"`
}

// MonitorState — per-token monitoring state
type MonitorState struct {
	TokenID    string    `json:"tokenId"`
	CreatedAt  time.Time `json:"createdAt"`
	LastTxTime time.Time `json:"lastTxTime"`
	TxCount    int64     `json:"txCount"`
	Profit     float64   `json:"profit"`
}

// VolumeRequest — automated volume generation
type VolumeRequest struct {
	TokenID   string   `json:"tokenId"`
	Wallets   []string `json:"wallets"`
	Cycles    int      `json:"cycles"`
	DelayMs   int64    `json:"delayMs"`
	AmountSOL float64  `json:"amountSOL"`
}

// SellBundleRequest — coordinated multi-wallet sell
type SellBundleRequest struct {
	TokenID     string   `json:"tokenId"`
	Wallets     []string `json:"wallets"`
	SellPercent float64  `json:"sellPercent"`
	JitoTipSOL  float64  `json:"jitoTipSOL"`
}

// SnipeResponse — returned after snipe execution
type SnipeResponse struct {
	Success    bool     `json:"success"`
	Signatures []string `json:"signatures"`
	Error      string   `json:"error,omitempty"`
}

// wsHub — WebSocket hub for real-time monitoring streams
type wsHub struct {
	mu        sync.RWMutex
	clients   map[string]map[*wsClient]bool
	broadcast chan wsMessage
}
