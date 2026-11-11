package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lexisync/jawad-os/internal/app"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ssh"
	"github.com/lexisync/jawad-os/internal/ui"
)

func main() {
	var (
		configPath = flag.String("config", "configs/portfolio.yaml", "path to portfolio configuration yaml")
		skipBoot   = flag.Bool("skip-boot", false, "skip initial boot animation in local mode")
		sshMode    = flag.Bool("ssh", false, "run as Wish SSH server")
		sshHost    = flag.String("host", "0.0.0.0", "SSH listen host")
		sshPort    = flag.String("port", "2222", "SSH listen port")
		sshKey     = flag.String("key", ".ssh/id_ed25519", "SSH host key path")
	)

	if len(os.Args) > 1 && (os.Args[1] == "ssh" || os.Args[1] == "serve") {
		*sshMode = true
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	if *sshMode {
		runSSHServer(cfg, *sshHost, *sshPort, *sshKey)
		return
	}

	runLocal(cfg, *skipBoot)
}

func runLocal(cfg *config.PortfolioConfig, skipBoot bool) {
	var modelOpts []app.Option
	if skipBoot {
		modelOpts = append(modelOpts, app.WithInitialScreen(ui.ScreenDashboard))
	}

	p := app.NewProgram(cfg, modelOpts)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running JAWAD.OS: %v\n", err)
		os.Exit(1)
	}
}

func runSSHServer(cfg *config.PortfolioConfig, host, port, keyPath string) {
	srv, err := ssh.NewServer(
		cfg,
		ssh.WithHost(host),
		ssh.WithPort(port),
		ssh.WithHostKeyPath(keyPath),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SSH server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("┌────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│ JAWAD.OS SSH SERVER                                    │\n")
	fmt.Printf("├────────────────────────────────────────────────────────┤\n")
	fmt.Printf("│ Listening on: %-40s │\n", srv.Addr())
	fmt.Printf("│ Connect with: ssh %-36s │\n", "localhost -p "+port)
	fmt.Printf("│ Press Ctrl+C to shut down gracefully                   │\n")
	fmt.Printf("└────────────────────────────────────────────────────────┘\n")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "SSH server error: %v\n", err)
		}
	}()

	<-done
	fmt.Println("\nShutting down JAWAD.OS SSH server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error during graceful shutdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped successfully.")
}
