package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port              string `env:"PORT" envDefault:"8080"`
	GeminiApiKey      string `env:"GEMINI_API_KEY,required"`
	LangfusePublicKey string `env:"LANGFUSE_PUBLIC_KEY"`
	LangfuseSecretKey string `env:"LANGFUSE_SECRET_KEY"`
	LangfuseApiUrl    string `env:"LANGFUSE_API_URL"`
}

func Load() (*Config, error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading Env file")
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
