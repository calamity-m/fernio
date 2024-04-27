package server

import (
	"github.com/calamity-m/fernio/lib/logging"
)

type Server struct {
	Config Config
	Log    *logging.Logger
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
