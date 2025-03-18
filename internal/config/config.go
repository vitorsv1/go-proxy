package config

// Config holds the proxy server configuration
type Config struct {
	Port      string
	LogToFile bool
	Blocklist []string
}
