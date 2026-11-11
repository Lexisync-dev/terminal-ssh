package screens

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui"
)

type BootTickMsg struct{}

type BootModel struct {
	cfg        *config.PortfolioConfig
	step       int
	totalSteps int
	completed  bool
	width      int
	height     int
}

func NewBootModel(cfg *config.PortfolioConfig) BootModel {
	return BootModel{
		cfg:        cfg,
		step:       0,
		totalSteps: 5,
		completed:  false,
		width:      80,
		height:     24,
	}
}

func (m *BootModel) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func (m BootModel) Init() tea.Cmd {
	return tickBoot()
}

func tickBoot() tea.Cmd {
	return tea.Tick(time.Millisecond*260, func(t time.Time) tea.Msg {
		return BootTickMsg{}
	})
}

func (m BootModel) Update(msg tea.Msg) (BootModel, tea.Cmd, bool) {
	switch msg := msg.(type) {
	case BootTickMsg:
		if m.step < m.totalSteps {
			m.step++
			return m, tickBoot(), false
		}
		m.completed = true
		return m, nil, false

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "space", "esc", "q":
			// If already completed or user pressed a key to skip, transition to dashboard
			return m, nil, true
		default:
			if m.completed {
				return m, nil, true
			}
		}
	}

	return m, nil, false
}

func (m BootModel) View() string {
	var b strings.Builder

	// Header banner
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(ui.ColorAccent).
		Render("  JAWAD.OS")
	b.WriteString("\n" + header + "\n\n")

	// Steps sequence
	lines := []string{
		"  Initializing personal environment...",
		"  Loading profile............. " + statusOK(m.step >= 1),
		"  Loading projects............ " + statusOK(m.step >= 2),
		"  Loading experience.......... " + statusOK(m.step >= 3),
		"  Initializing interface...... " + statusOK(m.step >= 4),
	}

	for i := 0; i <= m.step && i < len(lines); i++ {
		b.WriteString(lines[i] + "\n")
	}

	// If sequence has finished all steps
	if m.step >= m.totalSteps {
		welcome := lipgloss.NewStyle().
			Bold(true).
			Foreground(ui.ColorText).
			Render("\n  Welcome to JAWAD.OS\n")

		tagline := lipgloss.NewStyle().
			Foreground(ui.ColorMuted).
			Render("  A developer exploring systems,\n  building software, and learning how things work.\n")

		prompt := lipgloss.NewStyle().
			Foreground(ui.ColorAccent).
			Bold(true).
			Render("  Press ENTER to continue.")

		skipHint := lipgloss.NewStyle().
			Foreground(ui.ColorMuted).
			Render("  (or press any key to enter)")

		b.WriteString(welcome)
		b.WriteString(tagline)
		b.WriteString("\n" + prompt + "\n" + skipHint + "\n")
	} else {
		skip := lipgloss.NewStyle().
			Foreground(ui.ColorMuted).
			Render("\n  [Press ENTER or SPACE to skip boot sequence]")
		b.WriteString(skip + "\n")
	}

	content := b.String()
	// Place neatly inside available window height
	topPadding := (m.height - 16) / 2
	if topPadding < 1 {
		topPadding = 1
	}

	return strings.Repeat("\n", topPadding) + content
}

func statusOK(done bool) string {
	if done {
		return lipgloss.NewStyle().Bold(true).Foreground(ui.ColorPrimary).Render("OK")
	}
	return lipgloss.NewStyle().Foreground(ui.ColorMuted).Render("..")
}

func (m BootModel) IsCompleted() bool {
	return m.completed
}
