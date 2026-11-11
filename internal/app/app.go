package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lexisync/jawad-os/internal/config"
)

// NewProgram creates a new Bubble Tea program instance for JAWAD.OS.
func NewProgram(cfg *config.PortfolioConfig, modelOpts []Option, teaOpts ...tea.ProgramOption) *tea.Program {
	m := NewModel(cfg, modelOpts...)
	defaultOpts := []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
	opts := append(defaultOpts, teaOpts...)
	return tea.NewProgram(m, opts...)
}
