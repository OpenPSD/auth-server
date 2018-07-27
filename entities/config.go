package entities

import (
	"fmt"
	"os"
)

type Config struct {
	HydraURL string
	Port     string
}

func NewConfig() Config {
	config := Config{
		HydraURL: "http://localhost:4444",
		Port:     ":8000",
	}
	paramHydraURL := os.Getenv("HYDRA_URL")
	paramPort := os.Getenv("PORT")

	if paramHydraURL != "" {
		config.HydraURL = paramHydraURL
	}

	if paramPort != "" {
		config.Port = fmt.Sprintf(":%s", paramPort)
	}

	return config
}
