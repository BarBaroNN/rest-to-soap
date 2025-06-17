package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	generators "rest-to-soap/core/build/generators"
	config "rest-to-soap/core/config"

	"go.uber.org/zap"
)

var (
	configPath = flag.String("config", "config/config.json", "path to config file")
)

func main() {
	flag.Parse()

	// Ensure config directory exists
	configDir := filepath.Dir(*configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatalf("Failed to create config directory: %v", err)
	}

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger, err := initLogger(cfg.Logging)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize template generator
	templateGen := generators.NewTemplateGenerator()
	if err := templateGen.GenerateTemplates(cfg); err != nil {
		logger.Fatal("Failed to generate templates", zap.Error(err))
	}

	// Initialize registry generator
	registryGen := generators.NewRegistryGenerator()
	if err := registryGen.GenerateRegistry(cfg); err != nil {
		logger.Fatal("Failed to generate registry", zap.Error(err))
	}
}

func initLogger(cfg config.LogConfig) (*zap.Logger, error) {
	var config zap.Config

	if cfg.Format == "json" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	// Set log level
	switch cfg.Level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	return config.Build()
}
