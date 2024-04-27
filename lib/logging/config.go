// Configuration for servers based on a dysfunctional options pattern - https://rednafi.com/go/dysfunctional_options_pattern
package logging

import (
	"log/slog"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Level       slog.Level `mapstructure:"level" json:"level"`
	Structured  bool       `mapstructure:"structured" json:"structured"`
	AddSource   bool       `mapstructure:"add-source" json:"certkeypath"`
	System      string     `mapstructure:"system" json:"system"`
	Environment string     `mapstructure:"environment" json:"environment"`
}

// Defaults
var (
	defaultLevel      = slog.LevelInfo
	defaultStructured = true
	defaultAddSource  = false
)

func ParseLevel(level string) (slog.Level, error) {
	var lvl slog.Level

	err := lvl.UnmarshalText([]byte(level))

	if err != nil {
		return slog.LevelDebug, err
	}
	return lvl, nil
}

func (cfg *LogConfig) WithLevel(level slog.Level) *LogConfig {
	cfg.Level = level
	return cfg
}

func (cfg *LogConfig) WithStructured(structured bool) *LogConfig {
	cfg.Structured = structured
	return cfg
}

// Create a config instance with blank values, intended to be used with With... functions to create and modify
// logging config
func NewConfg() *LogConfig {
	return &LogConfig{}
}

// Create a config instance with defaults set
func NewDefault() *LogConfig {
	return NewConfg().WithLevel(defaultLevel).WithStructured(defaultStructured)
}

// Create a config instance through viper, which will read configuration from disk,
// environment variables and defaults in that order.
//
// This method assumes that the provided viper has a configured filename/path set, such as:
// using viper.SetConfigFile(...)
func NewFromViper(v *viper.Viper, overwrite bool) (*LogConfig, error) {
	cfg := NewConfg()

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
		return NewDefault(), err
	}

	return cfg, nil
}
