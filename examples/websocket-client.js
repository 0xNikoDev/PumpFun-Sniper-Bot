// Minimal WebSocket client for PumpFun Sniper Bot monitoring
const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:8080/api/monitor/ws', {
  headers: { 'Authorization': 'Bearer YOUR_LICENSE_KEY' }
});

ws.on('open', () => {
  console.log('Connected to monitor');
  // Subscribe to a specific token
  ws.send(JSON.stringify({
    action: 'subscribe',
    tokenId: 'TokenMintAddress...'
  }));
});

ws.on('message', (data) => {
  const event = JSON.parse(data);
  console.log(`[${event.type}] ${event.tokenId} — ${event.action} ${event.amountSOL} SOL`);
});

ws.on('close', () => console.log('Disconnected'));
