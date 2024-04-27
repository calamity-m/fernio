// Configuration for servers based on a dysfunctional options pattern - https://rednafi.com/go/dysfunctional_options_pattern
package server

import (
	"github.com/spf13/viper"
)

type TLS struct {
	Enabled     bool   `json:"enabled,omitempty"`
	CertPath    string `mapstructure:"cert-path" json:"certpath,omitempty"`
	CertKeyPath string `mapstructure:"cert-key-path" json:"certkeypath,omitempty"`
}

type Config struct {
	Host string
	Port int
	Tls  TLS
}

// Defaults
var (
	defaultHost        string = "localhost"
	defaultPort        int    = 8000
	defaultTlsEnabled  bool   = false
	defaultCertPath    string = "/run/secret/cert.crt"
	defaultCertKeyPath string = "/run/secret/cert.key"
)

func (cfg *Config) WithHost(host string) *Config {
	cfg.Host = host
	return cfg
}

func (cfg *Config) WithPort(port int) *Config {
	cfg.Port = port
	return cfg
}

func (cfg *Config) WithTLS(tls *TLS) *Config {
	cfg.Tls = *tls
	return cfg
}

func (cfg *Config) WithTLSCertKeyPath(certKeyPath string) *Config {
	cfg.Tls.CertKeyPath = certKeyPath
	return cfg
}

func (cfg *Config) WithTLSCertPath(certPath string) *Config {
	cfg.Tls.CertPath = certPath
	return cfg
}

func (cfg *Config) WithTLSEnabled(enabled bool) *Config {
	cfg.Tls.Enabled = enabled
	return cfg
}

// Create a config instance with blank values, intended to be used with With... functions to create and modify
// server config
func NewConfg() *Config {
	return &Config{}
}

// Create a config instance with defaults set
func NewDefault() *Config {
	return NewConfg().WithHost(defaultHost).WithPort(defaultPort).WithTLSEnabled(defaultTlsEnabled)
}

// Create a config instance through viper, which will read configuration from disk,
// environment variables and defaults in that order.
//
// This method assumes that the provided viper has a configured filename/path set, such as:
// using viper.SetConfigFile(...)
func NewFromViper(v *viper.Viper, overwrite bool) (*Config, error) {
	cfg := NewConfg()

	v.SetDefault("server.host", defaultHost)
	v.SetDefault("server.port", defaultPort)
	v.SetDefault("server.tls.enabled", defaultTlsEnabled)
	v.SetDefault("server.tls.cert-path", defaultCertPath)
	v.SetDefault("server.tls.cert-key-path", defaultCertKeyPath)

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

	if err := v.UnmarshalKey("server", &cfg); err != nil {
		return NewDefault(), err
	}

	return cfg, nil
}
