package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type InterestsModel struct {
	cfg      *config.PortfolioConfig
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewInterestsModel(cfg *config.PortfolioConfig) InterestsModel {
	return InterestsModel{
		cfg: cfg,
	}
}

func (m *InterestsModel) SetSize(w, h int) {
	m.width = w
	m.height = h

	geo := theme.CalculateLayout(w, h)
	vpHeight := geo.InnerHeight

	if !m.ready {
		m.viewport = viewport.New(geo.InnerWidth, vpHeight)
		m.ready = true
	} else {
		m.viewport.Width = geo.InnerWidth
		m.viewport.Height = vpHeight
	}

	m.viewport.SetContent(m.renderContent(geo.InnerWidth))
}

func (m InterestsModel) Update(msg tea.Msg) (InterestsModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k", "down", "j", "pgup", "pgdown", "home", "end":
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m InterestsModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Loading editorial entries...")
	}
	return m.viewport.View()
}

func (m InterestsModel) renderContent(width int) string {
	var b strings.Builder

	// Top Editorial Intro
	introText := theme.StyleBodyMuted.Width(width - 6).Render(
		"EDITORIAL // Personal pursuits, physical engineering curiosity, and creative expression beyond software.",
	)
	b.WriteString(theme.RenderWorkstationPanel("BEYOND THE TERMINAL", introText, width, 0, false) + "\n\n")

	// Render each interest item in a workstation panel
	for i, item := range m.cfg.Interests.Items {
		itemNum := strings.ToUpper(theme.StyleBrand.Render(item.Icon + " "))
		tagline := theme.StyleSectionHeading.Render(strings.ToUpper(item.Tagline))
		desc := theme.StyleBody.Width(width - 6).Render(item.Description)

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			tagline,
			"",
			desc,
		)

		panelTitle := strings.ToUpper(item.Name)
		if i == 0 {
			panelTitle = "POETRY // CREATIVE EXPRESSION"
		} else if i == 1 {
			panelTitle = "SUPERBIKES // TWO-WHEELED DYNAMICS"
		} else if i == 2 {
			panelTitle = "MECHANICAL ENGINEERING // PHYSICAL SYSTEMS"
		} else if i == 3 {
			panelTitle = "CURIOSITY // DECONSTRUCTING HOW THINGS WORK"
		}

		panel := theme.RenderWorkstationPanel(itemNum+panelTitle, content, width, 0, false)
		b.WriteString(panel + "\n\n")
	}

	// Personal Reflection Box
	quote := theme.StyleBodyMuted.Italic(true).Width(width - 6).Render(
		"\"The curiosity that drives mechanical tinkering and creative writing is the same mindset that makes for thoughtful software engineering.\"",
	)
	b.WriteString(theme.RenderWorkstationPanel("ENGINEERING REFLECTION", quote, width, 0, false) + "\n")

	return b.String()
}
