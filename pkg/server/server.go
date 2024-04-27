package server

import (
	"github.com/calamity-m/fernio/pkg/logging"
	"github.com/spf13/viper"
)

var (
	defaultHost            string = "localhost"
	defaultPort            int    = 8000
	defaultTlsEnabled      bool   = false
	defaultCertPath        string = "/run/secret/cert.crt"
	defaultCertKeyPath     string = "/run/secret/cert.key"
	defaultEnvironment     string = "dev"
	defaultRequestIdHeader string = "X-Request-Id"
	DefaultSystem          string = "fernio"
)

type TLS struct {
	Enabled     bool   `json:"enabled,omitempty"`
	CertPath    string `mapstructure:"cert-path" json:"certpath,omitempty"`
	CertKeyPath string `mapstructure:"cert-key-path" json:"certkeypath,omitempty"`
}

type Config struct {
	System          string `mapstructure:"system" json:"system,omitempty"`
	Environment     string `mapstructure:"environment" json:"environment,omitempty"`
	RequestIdHeader string `mapstructure:"request-id-header" json:"request-id-header,omitempty"`
	Host            string `mapstructure:"host" json:"host,omitempty"`
	Port            int    `mapstructure:"Port" json:"port,omitempty"`
	Tls             TLS
}

type Server struct {
	Config Config
	Log    *logging.Logger
}

func NewConfg(opts ...func(*Config)) *Config {
	cfg := &Config{Host: defaultHost, Port: defaultPort, Tls: TLS{Enabled: defaultTlsEnabled}}

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
func NewViperConfig(v *viper.Viper, overwrite bool) (*Config, error) {
	cfg := NewConfg()

	v.SetDefault("server.system", DefaultSystem)
	v.SetDefault("server.environment", defaultEnvironment)
	v.SetDefault("server.request-id-header", defaultRequestIdHeader)
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
		return cfg, err
	}

	return cfg, nil
}

func New(opts ...func(*Server)) *Server {
	s := &Server{}

	for _, fn := range opts {
		fn(s)
	}
	return s
}

func WithConfig(c Config) func(*Server) {
	return func(s *Server) {
		s.Config = c
	}
}

func WithLogger(l *logging.Logger) func(*Server) {
	return func(s *Server) {
		s.Log = l
	}
}
