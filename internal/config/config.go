package config

import (
	"bufio"
	"os"
	"strings"
)

// Config holds every environment-specific value the server reads.
type Config struct {
	Port      string
	StaticDir string
	BaseURL   string
	AIAPIKey  string
	AIBaseURL string
	AIModel   string
}

// Load reads .env (if present) then resolves the runtime configuration.
// The file never overrides variables already set in the process environment.
func Load() Config {
	loadEnvFile(".env")
	return Config{
		Port:      env("PORT", "3000"),
		StaticDir: env("STATIC_DIR", "public"),
		BaseURL:   env("NEXT_PUBLIC_BASE_URL", "http://localhost:3000"),
		AIAPIKey:  os.Getenv("AI_API_KEY"),
		AIBaseURL: env("AI_BASE_URL", "https://integrate.api.nvidia.com/v1"),
		AIModel:   env("AI_MODEL", "meta/llama-3.1-8b-instruct"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func loadEnvFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}
}
