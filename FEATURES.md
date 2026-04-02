# Features — PumpFun Sniper Bot

Full API reference with request/response examples for all endpoints.

> All endpoints require `Authorization: Bearer <your-license-key>` header.
> Base URL: `http://your-server:8080`

---

## 1. Multi-Wallet Sniping

Coordinate simultaneous buys across multiple wallets the moment a new pump.fun token launches. Each wallet sends an independent RPC transaction in parallel, maximizing fill probability.

**Endpoint:** `POST /api/snipe/fire`

```bash
curl -X POST http://localhost:8080/api/snipe/fire \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenId": "TokenMintAddress...",
    "wallets": ["wallet1_pubkey", "wallet2_pubkey", "wallet3_pubkey"],
    "buyPercent": 25.0,
    "slippageBps": 500,
    "devBuySOL": 0.5
  }'
```

**Response:**
```json
{
  "success": true,
  "signatures": [
    "5KtP...tx1",
    "7mNq...tx2",
    "9pRz...tx3"
  ],
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenId` | string | Pump.fun token mint address |
| `wallets` | []string | Wallet public keys to snipe with |
| `buyPercent` | float64 | % of each wallet's SOL balance to spend |
| `slippageBps` | uint16 | Slippage tolerance in basis points (500 = 5%) |
| `devBuySOL` | float64 | Expected dev buy amount (for bonding curve calculation) |

---

## 2. Jito MEV Bundle Protection

Submit buy transactions as a Jito bundle for MEV protection. Bundles are executed atomically — all wallets buy in the same block, preventing sandwich attacks and front-running.

**Endpoint:** `POST /api/snipe/bundle-fire`

```bash
curl -X POST http://localhost:8080/api/snipe/bundle-fire \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenId": "TokenMintAddress...",
    "wallets": ["wallet1_pubkey", "wallet2_pubkey"],
    "jitoTipSOL": 0.01,
    "rpcFallbackDelayMs": 500
  }'
```

**Response:**
```json
{
  "success": true,
  "bundleId": "bundle_abc123...",
  "signatures": ["5KtP...tx1", "7mNq...tx2"],
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenId` | string | Pump.fun token mint address |
| `wallets` | []string | Wallet public keys for bundle |
| `jitoTipSOL` | float64 | Jito validator tip in SOL (higher = faster inclusion) |
| `rpcFallbackDelayMs` | int64 | Fallback to RPC if bundle not confirmed within N ms |

---

## 3. Volume Bot

Generate automated buy/sell cycles across wallets to create organic-looking trading volume on pump.fun tokens. Useful for maintaining bonding curve momentum.

**Endpoint:** `POST /api/volume/start`

```bash
curl -X POST http://localhost:8080/api/volume/start \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenId": "TokenMintAddress...",
    "wallets": ["wallet1_pubkey", "wallet2_pubkey", "wallet3_pubkey"],
    "cycles": 10,
    "delayMs": 2000,
    "amountSOL": 0.1
  }'
```

**Response:**
```json
{
  "success": true,
  "jobId": "vol_job_xyz789",
  "estimatedDurationMs": 20000,
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenId` | string | Pump.fun token mint address |
| `wallets` | []string | Wallets to rotate through |
| `cycles` | int | Number of buy/sell cycles to execute |
| `delayMs` | int64 | Delay between cycles in milliseconds |
| `amountSOL` | float64 | SOL amount per buy transaction |

---

## 4. WebSocket Monitoring

Subscribe to real-time transaction events for any pump.fun token. Receive live price updates, buy/sell events, take-profit triggers, and stop-loss alerts.

**Endpoint:** `GET /api/monitor/ws` (WebSocket upgrade)

```javascript
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/api/monitor/ws', {
  headers: { 'Authorization': 'Bearer YOUR_LICENSE_KEY' }
});

ws.on('open', () => {
  // Subscribe to a token
  ws.send(JSON.stringify({
    action: 'subscribe',
    tokenId: 'TokenMintAddress...'
  }));
});

ws.on('message', (data) => {
  const event = JSON.parse(data);
  // event.type: 'trade' | 'take_profit' | 'stop_loss' | 'price_update'
  console.log(`[${event.type}] ${event.tokenId} — ${event.action} ${event.amountSOL} SOL`);
});
```

**Event payload:**
```json
{
  "type": "trade",
  "tokenId": "TokenMintAddress...",
  "action": "buy",
  "amountSOL": 1.5,
  "priceUSD": 0.00042,
  "marketCap": 42000,
  "timestamp": "2026-04-01T12:00:00Z",
  "signature": "5KtP...txhash"
}
```

---

## 5. Token Analysis

Fetch current bonding curve data for any pump.fun token — market cap, liquidity, progress to graduation, and creator info.

**Endpoint:** `GET /api/token/:id`

```bash
curl http://localhost:8080/api/token/TokenMintAddress... \
  -H "Authorization: Bearer YOUR_LICENSE_KEY"
```

**Response:**
```json
{
  "tokenId": "TokenMintAddress...",
  "marketCapSOL": 42.5,
  "liquiditySOL": 12.3,
  "progress": 0.68,
  "complete": false,
  "creator": "CreatorWalletAddress...",
  "isCashbackCoin": false
}
```

---

## 6. Bundle Sell

Execute coordinated sells across multiple wallets via Jito bundle. Ensures all wallets exit in the same block to minimize price impact.

**Endpoint:** `POST /api/sell/bundle`

```bash
curl -X POST http://localhost:8080/api/sell/bundle \
  -H "Authorization: Bearer YOUR_LICENSE_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "tokenId": "TokenMintAddress...",
    "wallets": ["wallet1_pubkey", "wallet2_pubkey"],
    "sellPercent": 100.0,
    "jitoTipSOL": 0.005
  }'
```

**Response:**
```json
{
  "success": true,
  "signatures": ["5KtP...tx1", "7mNq...tx2"],
  "solReceived": 3.42,
  "error": ""
}
```

**Parameters:**
| Field | Type | Description |
|-------|------|-------------|
| `tokenId` | string | Pump.fun token mint address |
| `wallets` | []string | Wallets holding tokens to sell |
| `sellPercent` | float64 | % of each wallet's token balance to sell |
| `jitoTipSOL` | float64 | Jito validator tip in SOL |

---

## Authentication

All endpoints require a valid license key:

```
Authorization: Bearer YOUR_LICENSE_KEY
```

License keys are delivered via email after purchase at [memesnipe.com](https://memesnipe.com).

---

## Error Responses

```json
{
  "success": false,
  "error": "insufficient SOL balance in wallet wallet1_pubkey"
}
```

Common error codes:
- `401 Unauthorized` — invalid or expired license key
- `400 Bad Request` — missing required fields or invalid parameters
- `429 Too Many Requests` — rate limit exceeded
- `500 Internal Server Error` — RPC or blockchain error (retryable)
