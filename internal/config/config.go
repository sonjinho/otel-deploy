package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath        string `yaml:"db_path"`
	ServerAddress string `yaml:"server_address"`
	Port          string `yaml:"port"`
	LogLevel      string `yaml:"log_level"`
}

func LoadConfig() *Config {

	_ = godotenv.Load()

	cfg := &Config{
		DBPath:        "./data",
		ServerAddress: "0.0.0.0",
		Port:          "8080",
		LogLevel:      "info",
	}

	yamlFile, err := os.ReadFile("config/default.yaml")
	if err == nil {
		if err := yaml.Unmarshal(yamlFile, cfg); err != nil {
			log.Fatalf("Error parsing Yaml file: %v", err)
		}
	} else {
		log.Println("No default.yaml found, using default config")
	}

	cfg.DBPath = getEnv("DB_PATH", cfg.DBPath)
	cfg.Port = getEnv("PORT", cfg.Port)
	cfg.LogLevel = getEnv("LOG_LEVEL", cfg.LogLevel)

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
