# Architecture

## System Overview

PumpFun Sniper Bot is a Go REST API that interfaces with the Solana blockchain
for automated pump.fun token trading.

## Components

### API Layer
- HTTP server (net/http) with middleware for auth, rate limiting, audit logging
- WebSocket hub for real-time monitoring streams
- JSON request/response with structured error handling

### Core Engine
- **Snipe Manager** — Coordinates multi-wallet buy execution (RPC + Jito bundle)
- **Volume Manager** — Automated buy/sell cycles for volume generation
- **Monitor Loop** — Yellowstone gRPC subscription for real-time transaction monitoring

### Solana Layer
- Direct RPC calls for single-wallet operations
- Jito block engine for MEV-protected bundle submissions
- Yellowstone gRPC for low-latency transaction streaming
- Bonding curve calculations for price estimation

### Data Flow

1. Token launch detected via Yellowstone gRPC stream
2. Bonding curve account parsed, market cap evaluated
3. Snipe request constructed with wallet allocation
4. Transaction sent via Jito bundle (priority) or RPC fallback
5. Monitor loop tracks position and triggers take-profit/stop-loss

### Security
- Bearer token authentication on all API endpoints
- Rate limiting per IP and per token
- Audit logging for all trading operations
- No private keys stored in config (loaded from encrypted keystore)
