// Configuration for servers based on a dysfunctional options pattern - https://rednafi.com/go/dysfunctional_options_pattern
package logging

import (
	"log/slog"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Level      slog.Level `mapstructure:"level" json:"level"`
	Structured bool       `mapstructure:"structured" json:"structured"`
	AddSource  bool       `mapstructure:"add-source" json:"certkeypath"`
}

// Defaults
var (
	defaultLevel      = slog.LevelInfo
	defaultStructured = true
	defaultAddSource  = false
)

func NewLogConfig(opts ...func(*LogConfig)) *LogConfig {
	cfg := &LogConfig{
		Level:      defaultLevel,
		Structured: defaultStructured,
		AddSource:  defaultAddSource,
	}

	for _, fn := range opts {
		fn(cfg)
	}

	return cfg
}

// Create a config instance through viper, which will read configuration from disk,
// environment variables and defaults in that order.
//
// This method assumes that the provided viper has a configured filename/path set, such as:
// using viper.SetConfigFile(...)
func NewFromViper(v *viper.Viper, overwrite bool) (*LogConfig, error) {
	cfg := NewLogConfig()

	v.SetDefault("logging.level", defaultLevel)
	v.SetDefault("logging.structured", defaultStructured)
	v.SetDefault("logging.add-source", defaultAddSource)

	if err := v.ReadInConfig(); err != nil {
		// We failed to find config file, try to write it instead
		if err := v.SafeWriteConfig(); err != nil {
			return nil, err
		}
	}

	if overwrite {
		if err := v.WriteConfig(); err != nil {
			return nil, err
		}
	}

	if err := v.UnmarshalKey("logging", &cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc())); err != nil {
		return cfg, err
	}

	return cfg, nil
}
