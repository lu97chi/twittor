package config

import "os"

// MongoConfig config
type MongoConfig struct {
	User     string
	Password string
	URL      string
	Database string
	Options  string
}

// JSONTokenCofing config
type JSONTokenCofing struct {
	Seed string
}

// NetworkConfig config
type NetworkConfig struct {
	Port string
}

// Config main config file
type Config struct {
	MongoDB MongoConfig
	JWT     JSONTokenCofing
	Network NetworkConfig
}

// New returns a new config
func New() *Config {
	return &Config{
		MongoDB: MongoConfig{
			User:     getEnv("MONGO_USER", ""),
			Password: getEnv("MONGO_PASSWORD", ""),
			URL:      getEnv("MONGO_URL", ""),
			Database: getEnv("MONGO_DATABASE", ""),
			Options:  getEnv("MONGO_CONNECTION_OPTIONS", ""),
		},
		JWT: JSONTokenCofing{
			Seed: getEnv("JSON_WEB_TOKEN_SEED", ""),
		},
		Network: NetworkConfig{
			Port: getEnv("PORT", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
