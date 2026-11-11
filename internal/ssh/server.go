package ssh

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/lexisync/jawad-os/internal/config"
)

type Config struct {
	Host        string
	Port        string
	HostKeyPath string
}

type Server struct {
	sshServer *ssh.Server
	cfg       Config
}

// Option represents a functional option for configuring the SSH server.
type Option func(*Config)

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port string) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithHostKeyPath(path string) Option {
	return func(c *Config) {
		c.HostKeyPath = path
	}
}

// NewServer creates a new Wish SSH server instance.
func NewServer(portfolioCfg *config.PortfolioConfig, opts ...Option) (*Server, error) {
	sshCfg := Config{
		Host:        "0.0.0.0",
		Port:        "2222",
		HostKeyPath: ".ssh/id_ed25519",
	}

	// Read environment overrides if available
	if h := os.Getenv("SSH_HOST"); h != "" {
		sshCfg.Host = h
	}
	if p := os.Getenv("PORT"); p != "" {
		sshCfg.Port = p
	} else if p := os.Getenv("SSH_PORT"); p != "" {
		sshCfg.Port = p
	}
	if k := os.Getenv("SSH_KEY_PATH"); k != "" {
		sshCfg.HostKeyPath = k
	}

	for _, opt := range opts {
		opt(&sshCfg)
	}

	// Ensure the directory for the host key exists
	if dir := filepath.Dir(sshCfg.HostKeyPath); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("creating host key directory %s: %w", dir, err)
		}
	}

	addr := net.JoinHostPort(sshCfg.Host, sshCfg.Port)

	s, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(sshCfg.HostKeyPath),
		// Public portfolio: accept any public key or passwordless connection
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithPasswordAuth(func(ctx ssh.Context, password string) bool {
			return true
		}),
		wish.WithMiddleware(
			bubbletea.Middleware(SessionHandler(portfolioCfg)),
			logging.Middleware(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing wish server: %w", err)
	}

	return &Server{
		sshServer: s,
		cfg:       sshCfg,
	}, nil
}

// ListenAndServe starts the SSH server.
func (s *Server) ListenAndServe() error {
	return s.sshServer.ListenAndServe()
}

// Shutdown gracefully stops the SSH server within the provided context.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.sshServer.Shutdown(ctx)
}

// Addr returns the configured address.
func (s *Server) Addr() string {
	return net.JoinHostPort(s.cfg.Host, s.cfg.Port)
}
