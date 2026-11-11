package ssh

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/lexisync/jawad-os/internal/app"
	"github.com/lexisync/jawad-os/internal/config"
)

// SessionHandler returns a Wish Bubble Tea handler that instantiates an isolated,
// independent session model for each connected visitor.
func SessionHandler(cfg *config.PortfolioConfig) bubbletea.Handler {
	return func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
		user := sess.User()
		if user == "" {
			user = "guest"
		}

		m := app.NewModel(
			cfg,
			app.WithUser(user),
		)

		pty, _, active := sess.Pty()
		if active {
			m.SetSize(pty.Window.Width, pty.Window.Height)
		}

		opts := []tea.ProgramOption{
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		}

		return m, opts
	}
}
