package ssh

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"testing"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/lexisync/jawad-os/internal/app"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui"
	gossh "golang.org/x/crypto/ssh"
)

func getTestConfig() *config.PortfolioConfig {
	cfg, _ := config.Load("")
	return cfg
}

func TestNewServerConfig(t *testing.T) {
	cfg := getTestConfig()
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "test_key")

	srv, err := NewServer(
		cfg,
		WithHost("127.0.0.1"),
		WithPort("22222"),
		WithHostKeyPath(keyPath),
	)
	if err != nil {
		t.Fatalf("expected nil error creating server, got: %v", err)
	}

	if srv.Addr() != "127.0.0.1:22222" {
		t.Errorf("expected addr '127.0.0.1:22222', got '%s'", srv.Addr())
	}
}

func TestEnvironmentOverrides(t *testing.T) {
	cfg := getTestConfig()
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "env_key")

	t.Setenv("SSH_HOST", "127.0.0.2")
	t.Setenv("PORT", "23456")
	t.Setenv("SSH_KEY_PATH", keyPath)

	srv, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("expected nil error with env vars, got: %v", err)
	}

	if srv.Addr() != "127.0.0.2:23456" {
		t.Errorf("expected addr '127.0.0.2:23456', got '%s'", srv.Addr())
	}
}

// mockSSHSession implements ssh.Session for unit testing session isolation.
type mockSSHSession struct {
	user string
	ssh.Session
}

func (m mockSSHSession) User() string {
	return m.user
}

func (m mockSSHSession) Pty() (ssh.Pty, <-chan ssh.Window, bool) {
	return ssh.Pty{
		Window: ssh.Window{Width: 100, Height: 30},
	}, nil, true
}

func TestSessionHandlerIsolation(t *testing.T) {
	cfg := getTestConfig()
	handler := SessionHandler(cfg)

	// Visitor 1: Alice
	sess1 := mockSSHSession{user: "alice"}
	model1, _ := handler(sess1)
	m1 := model1.(app.Model)

	if m1.User != "alice" {
		t.Errorf("expected user 'alice', got '%s'", m1.User)
	}

	// Visitor 2: Bob
	sess2 := mockSSHSession{user: "bob"}
	model2, _ := handler(sess2)
	m2 := model2.(app.Model)

	if m2.User != "bob" {
		t.Errorf("expected user 'bob', got '%s'", m2.User)
	}

	// Mutate Visitor 1 state (navigate to Skills, enter command mode)
	m1.NavigateTo(ui.ScreenSkills)
	m1.CommandMode = true

	// Verify Visitor 2 state is completely unaffected
	if m2.CurrentScreen != ui.ScreenBoot {
		t.Errorf("expected visitor 2 to remain on ScreenBoot, got %v", m2.CurrentScreen)
	}
	if m2.CommandMode {
		t.Errorf("expected visitor 2 CommandMode to be false")
	}
}

func getFreePort(t *testing.T) int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen on free port: %v", err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func TestSSHServerLifecycleAndConnection(t *testing.T) {
	cfg := getTestConfig()
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "server_ed25519")

	port := getFreePort(t)
	portStr := fmt.Sprintf("%d", port)

	srv, err := NewServer(
		cfg,
		WithHost("127.0.0.1"),
		WithPort(portStr),
		WithHostKeyPath(keyPath),
	)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	// Start server in background
	serverErrChan := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
			serverErrChan <- err
		}
	}()

	// Wait for server to bind
	time.Sleep(200 * time.Millisecond)

	// Connect as an SSH client
	sshConfig := &gossh.ClientConfig{
		User: "recruiter",
		Auth: []gossh.AuthMethod{
			gossh.Password("anonymous"),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         2 * time.Second,
	}

	client, err := gossh.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port), sshConfig)
	if err != nil {
		t.Fatalf("failed to dial SSH server: %v", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		t.Fatalf("failed to open session: %v", err)
	}
	defer session.Close()

	// Request PTY
	modes := gossh.TerminalModes{
		gossh.ECHO:          0,
		gossh.TTY_OP_ISPEED: 14400,
		gossh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm-256color", 24, 80, modes); err != nil {
		t.Fatalf("failed to request pty: %v", err)
	}

	// Close client session to simulate visitor disconnect
	session.Close()
	client.Close()

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Errorf("error during server shutdown: %v", err)
	}
}

func TestConcurrentConnections(t *testing.T) {
	cfg := getTestConfig()
	tmpDir := t.TempDir()
	keyPath := filepath.Join(tmpDir, "server_ed25519_concurrent")

	port := getFreePort(t)
	portStr := fmt.Sprintf("%d", port)

	srv, err := NewServer(
		cfg,
		WithHost("127.0.0.1"),
		WithPort(portStr),
		WithHostKeyPath(keyPath),
	)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	go func() {
		_ = srv.ListenAndServe()
	}()

	time.Sleep(200 * time.Millisecond)

	// Dial 3 concurrent clients
	const numClients = 3
	clients := make([]*gossh.Client, numClients)
	sessions := make([]*gossh.Session, numClients)

	for i := 0; i < numClients; i++ {
		sshConfig := &gossh.ClientConfig{
			User: fmt.Sprintf("user-%d", i),
			Auth: []gossh.AuthMethod{
				gossh.Password("test"),
			},
			HostKeyCallback: gossh.InsecureIgnoreHostKey(),
			Timeout:         2 * time.Second,
		}

		c, err := gossh.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port), sshConfig)
		if err != nil {
			t.Fatalf("client %d failed to dial: %v", i, err)
		}
		clients[i] = c

		s, err := c.NewSession()
		if err != nil {
			t.Fatalf("client %d failed to open session: %v", i, err)
		}
		sessions[i] = s

		modes := gossh.TerminalModes{
			gossh.ECHO: 0,
		}
		if err := s.RequestPty("xterm-256color", 24, 80, modes); err != nil {
			t.Fatalf("client %d failed to request pty: %v", i, err)
		}
	}

	// Close all clients
	for i := 0; i < numClients; i++ {
		sessions[i].Close()
		clients[i].Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Errorf("error during server shutdown: %v", err)
	}
}
