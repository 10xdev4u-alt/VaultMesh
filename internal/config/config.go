// Package config defines the configuration structure and loading logic for VaultMesh.
package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config represents the complete system configuration.
type Config struct {
	Network   NetworkConfig   `mapstructure:"network"`
	Storage   StorageConfig   `mapstructure:"storage"`
	Crypto    CryptoConfig    `mapstructure:"crypto"`
	Incentive IncentiveConfig `mapstructure:"incentive"`
	API       APIConfig       `mapstructure:"api"`
	TUI       TUIConfig       `mapstructure:"tui"`
}

// NetworkConfig defines LibP2P and network-related settings.
type NetworkConfig struct {
	ListenAddrs []string `mapstructure:"listen_addrs"`
	Bootstrap   []string `mapstructure:"bootstrap_peers"`
	PeerStore   string   `mapstructure:"peer_store"`
	PublicAddr  string   `mapstructure:"public_addr"`
}

// StorageConfig defines local storage and chunking settings.
type StorageConfig struct {
	DataDir   string `mapstructure:"data_dir"`
	ChunkSize int64  `mapstructure:"chunk_size"`
	DBPath    string `mapstructure:"db_path"`
}

// CryptoConfig defines encryption settings.
type CryptoConfig struct {
	KeyPath string `mapstructure:"key_path"`
}

// IncentiveConfig defines reputation and proof settings.
type IncentiveConfig struct {
	ReputationEnabled bool `mapstructure:"reputation_enabled"`
}

// APIConfig defines REST/gRPC server settings.
type APIConfig struct {
	HTTPAddr string `mapstructure:"http_addr"`
	GRPCAddr string `mapstructure:"grpc_addr"`
}

// TUIConfig defines Terminal UI settings.
type TUIConfig struct {
	Theme string `mapstructure:"theme"`
}

// Load loads the configuration from a file or environment variables.
func Load(configPath string, onReload func(newConfig *Config)) (*Config, error) {
	v := viper.New()

	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("/etc/vaultmesh/")
		v.AddConfigPath("$HOME/.vaultmesh/")
	}

	v.SetEnvPrefix("VAULTMESH")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Set default values
	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	if onReload != nil {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			var newCfg Config
			if err := v.Unmarshal(&newCfg); err == nil {
				if err := newCfg.Validate(); err == nil {
					onReload(&newCfg)
				}
			}
		})
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("network.listen_addrs", []string{"/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/udp/4001/quic-v1"})
	v.SetDefault("storage.data_dir", "./data")
	v.SetDefault("storage.chunk_size", 1024*1024) // 1MB default
	v.SetDefault("storage.db_path", "./data/badger")
	v.SetDefault("api.http_addr", ":8080")
	v.SetDefault("api.grpc_addr", ":9090")
	v.SetDefault("tui.theme", "dark")
	v.SetDefault("crypto.key_path", "./keys/master.key")
}

// Validate ensures the configuration is valid.
func (c *Config) Validate() error {
	if c.Storage.ChunkSize <= 0 {
		return fmt.Errorf("chunk_size must be greater than 0")
	}
	if c.Storage.DataDir == "" {
		return fmt.Errorf("data_dir cannot be empty")
	}
	return nil
}
