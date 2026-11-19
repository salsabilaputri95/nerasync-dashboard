package config

import (
	"os"
	"strings"
)

// Config menyimpan konfigurasi runtime backend API SINERGI.
type Config struct {
	Port               string
	AppEnv             string
	LogLevel           string
	DatabaseURLPublic  string
	CORSAllowedOrigins []string
}

// Load membaca konfigurasi dari variabel lingkungan dengan nilai default yang aman.
func Load() Config {
	originsRaw := getEnvOrDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	origins := make([]string, 0)
	for _, item := range strings.Split(originsRaw, ",") {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}
	if len(origins) == 0 {
		origins = []string{"http://localhost:3000"}
	}

	return Config{
		Port:               getEnvOrDefault("PORT", "8080"),
		AppEnv:             strings.ToLower(getEnvOrDefault("APP_ENV", "development")),
		LogLevel:           strings.ToLower(getEnvOrDefault("LOG_LEVEL", "info")),
		DatabaseURLPublic:  strings.TrimSpace(os.Getenv("DATABASE_URL_PUBLIC")),
		CORSAllowedOrigins: origins,
	}
}

func getEnvOrDefault(key string, fallback string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return fallback
	}
	return val
}
