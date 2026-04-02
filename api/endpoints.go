package api

// API endpoint constants
const (
	RouteSnipeFire       = "/api/snipe/fire"
	RouteBundleSnipeFire = "/api/snipe/bundle-fire"
	RouteVolumeStart     = "/api/volume/start"
	RouteSellBundle      = "/api/sell/bundle"
	RouteMonitorWS       = "/api/monitor/ws"
	RouteTokenInfo       = "/api/token/:id"
	RouteTokenMeta       = "/api/token/:id/meta"
	RouteHealth          = "/api/health"
)

// RegisterRoutes sets up all API routes on the given mux.
// Handler implementations are not included in this preview.
//
// Authentication: All endpoints require Bearer token via
// Authorization header. See FEATURES.md for details.
