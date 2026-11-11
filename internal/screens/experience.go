package screens

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type ExperienceModel struct {
	cfg      *config.PortfolioConfig
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewExperienceModel(cfg *config.PortfolioConfig) ExperienceModel {
	return ExperienceModel{
		cfg: cfg,
	}
}

func (m *ExperienceModel) SetSize(w, h int) {
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

func (m ExperienceModel) Update(msg tea.Msg) (ExperienceModel, tea.Cmd) {
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

func (m ExperienceModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Loading engineering ledger...")
	}
	return m.viewport.View()
}

func (m ExperienceModel) renderContent(width int) string {
	var b strings.Builder

	// Current Candidate Status Banner
	statusTag := theme.StatusPillActive.Render("STATUS: " + strings.ToUpper(m.cfg.Experience.Status))
	summaryText := theme.StyleBody.Width(width - 6).Render(m.cfg.Experience.Summary)

	bannerContent := lipgloss.JoinVertical(
		lipgloss.Left,
		statusTag,
		"",
		summaryText,
	)
	b.WriteString(theme.RenderWorkstationPanel("CANDIDATE DISPOSITION", bannerContent, width, 0, false) + "\n\n")

	// Chronological Engineering Ledger
	milestones := []struct {
		Tag   string
		Title string
		Body  string
	}{
		{
			Tag:   "[MILESTONE 01 // ACADEMIC]",
			Title: "Computer Science Education",
			Body:  "Undergraduate computer science student in Pakistan. Rigorous focus on algorithms, data structures, operating systems, database internals, and distributed architecture fundamentals.",
		},
		{
			Tag:   "[MILESTONE 02 // APPLIED]",
			Title: "Independent Freelance Software Projects",
			Body:  "Built and shipped freelance software solutions. Handled end-to-end development lifecycles, user requirements, database design, and shipping working web services.",
		},
		{
			Tag:   "[MILESTONE 03 // SYSTEMS EXPLORATION]",
			Title: "Self-Directed Systems Building & Tooling",
			Body:  "Invested deep focus into systems architecture, building workflow automation engines, evaluating queue patterns (Redis), and developing terminal operating environments in Go.",
		},
		{
			Tag:   "[MILESTONE 04 // OBJECTIVE]",
			Title: "Pursuing First Formal Software Engineering Role",
			Body:  "Actively seeking full-stack and backend software engineering opportunities. Ready to contribute production code, understand existing architectures, and grow under senior engineering mentorship.",
		},
	}

	var ledgerLines []string
	for _, ms := range milestones {
		tag := theme.StyleBrand.Render(ms.Tag)
		title := theme.StyleIdentity.Render(ms.Title)
		body := theme.StyleBodyMuted.Width(width - 8).Render(ms.Body)

		entry := lipgloss.JoinVertical(
			lipgloss.Left,
			tag+" "+title,
			body,
		)
		ledgerLines = append(ledgerLines, entry)
	}

	ledgerContent := strings.Join(ledgerLines, "\n\n"+theme.RenderDivider(width-6)+"\n\n")
	b.WriteString(theme.RenderWorkstationPanel("CHRONOLOGICAL ENGINEERING LEDGER", ledgerContent, width, 0, false) + "\n\n")

	// Candidate Capabilities & Attributes
	attributes := []string{
		"• High Intrinsic Curiosity: Deep desire to understand how systems operate under the surface.",
		"• Fast Execution & Iteration: Learning by building, breaking, and methodically improving solutions.",
		"• Pragmatic Architecture: Preference for clean, maintainable modular design over accidental complexity.",
		"• Growth Mindset: Eagerness to absorb feedback and adapt to team engineering standards.",
	}
	attrContent := theme.StyleBody.Width(width - 6).Render(strings.Join(attributes, "\n\n"))
	b.WriteString(theme.RenderWorkstationPanel("TEAM CAPABILITIES & ATTRIBUTES", attrContent, width, 0, false) + "\n")

	return b.String()
}
