package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the global configuration for VaultMesh.
type Config struct {
	Redundancy RedundancyConfig `mapstructure:"redundancy"`
}

// RedundancyConfig holds parameters for erasure coding.
type RedundancyConfig struct {
	DataShards   int `mapstructure:"data_shards"`
	ParityShards int `mapstructure:"parity_shards"`
}

// Load loads the configuration from a file or environment variables.
func Load(path string) (*Config, error) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.vaultmesh")
	}

	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("redundancy.data_shards", 3)
	viper.SetDefault("redundancy.parity_shards", 2)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
