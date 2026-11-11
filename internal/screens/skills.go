package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type SkillsModel struct {
	cfg      *config.PortfolioConfig
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewSkillsModel(cfg *config.PortfolioConfig) SkillsModel {
	return SkillsModel{
		cfg: cfg,
	}
}

func (m *SkillsModel) SetSize(w, h int) {
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

func (m SkillsModel) Update(msg tea.Msg) (SkillsModel, tea.Cmd) {
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

func (m SkillsModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Loading technical stack taxonomy...")
	}
	return m.viewport.View()
}

func (m SkillsModel) renderContent(width int) string {
	var b strings.Builder

	// Top Taxonomy Notice
	noticeText := theme.StyleBodyMuted.Width(width - 6).Render(
		"SYS.TAXONOMY // Active working stack and engineering topics under exploration. No arbitrary percentages.",
	)
	noticePanel := theme.RenderWorkstationPanel("TAXONOMY POLICY", noticeText, width, 0, false)
	b.WriteString(noticePanel + "\n\n")

	// Render Each Category as a Precision Workstation Panel
	for _, category := range m.cfg.Skills.Categories {
		desc := theme.StyleBodyMuted.Width(width - 6).Render(category.Description)
		tags := theme.RenderTags(category.Items)

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			desc,
			"",
			tags,
		)

		panel := theme.RenderWorkstationPanel("CATEGORY // "+strings.ToUpper(category.Name), content, width, 0, false)
		b.WriteString(panel + "\n\n")
	}

	return b.String()
}
