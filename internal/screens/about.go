package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type AboutModel struct {
	cfg      *config.PortfolioConfig
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewAboutModel(cfg *config.PortfolioConfig) AboutModel {
	return AboutModel{
		cfg: cfg,
	}
}

func (m *AboutModel) SetSize(w, h int) {
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

func (m AboutModel) Update(msg tea.Msg) (AboutModel, tea.Cmd) {
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

func (m AboutModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Initializing profile...")
	}
	return m.viewport.View()
}

func (m AboutModel) renderContent(width int) string {
	var b strings.Builder

	// Top Section: Personnel Dossier
	col1Block := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.RenderKeyValue("NAME:", m.cfg.Profile.Name, 12),
		theme.RenderKeyValue("STATUS:", m.cfg.Profile.Title, 12),
		theme.RenderKeyValue("LOCATION:", m.cfg.Profile.Location, 12),
	)

	col2Block := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.RenderKeyValue("EDUCATION:", m.cfg.Profile.Education, 12),
		theme.RenderKeyValue("DIRECTION:", m.cfg.Profile.Positioning, 12),
		theme.RenderKeyValue("INTERESTS:", "Systems, Automation, Infra", 12),
	)

	var grid string
	if width >= 80 {
		halfWidth := (width - 8) / 2
		c1 := lipgloss.NewStyle().Width(halfWidth).Render(col1Block)
		c2 := lipgloss.NewStyle().Width(halfWidth).Render(col2Block)
		grid = lipgloss.JoinHorizontal(lipgloss.Top, c1, "  ", c2)
	} else {
		grid = lipgloss.JoinVertical(lipgloss.Left, col1Block, "", col2Block)
	}

	idPanel := theme.RenderWorkstationPanel("PROFILE", grid, width, 0, false)
	b.WriteString(idPanel + "\n\n")

	// Engineering Philosophy
	bioBody := theme.StyleBody.Width(width - 6).Render(m.cfg.Profile.Bio)
	bioPanel := theme.RenderWorkstationPanel("PHILOSOPHY", bioBody, width, 0, false)
	b.WriteString(bioPanel + "\n\n")

	// Active Domains of Exploration
	var areaRows []string
	for _, area := range m.cfg.Direction.Areas {
		title := theme.StyleBrand.Render("· " + strings.ToUpper(area.Title))
		desc := theme.StyleBodyMuted.PaddingLeft(2).Width(width - 8).Render(area.Description)
		areaRows = append(areaRows, fmt.Sprintf("%s\n%s", title, desc))
	}

	directionContent := strings.Join(areaRows, "\n\n")
	dirPanel := theme.RenderWorkstationPanel("AREAS OF EXPLORATION", directionContent, width, 0, false)
	b.WriteString(dirPanel + "\n")

	return b.String()
}
