package config

var (
	DefaultGRPSyncListeningAddr  = ":7580"
	DefaultGRPCNodeListeningAddr = ":7559"
	DefaultGRPCCTLListeningAddr  = ":7560"
	Tls                          bool
	CaFile                       string
	ServerHostOverride           string
)
