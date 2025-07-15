package server

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewTCPServer,
	NewWebSocketServer,
	NewKCPServer,
	NewGRPCServer,
	NewHTTPServer,
	NewRegistrar,
)
