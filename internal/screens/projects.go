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

type ProjectsModel struct {
	cfg          *config.PortfolioConfig
	selectedProj int
	viewport     viewport.Model
	ready        bool
	width        int
	height       int
}

func NewProjectsModel(cfg *config.PortfolioConfig) ProjectsModel {
	return ProjectsModel{
		cfg:          cfg,
		selectedProj: 0,
	}
}

func (m *ProjectsModel) SetSize(w, h int) {
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

func (m ProjectsModel) Update(msg tea.Msg) (ProjectsModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k", "down", "j", "pgup", "pgdown", "home", "end":
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd

		case "tab", "n":
			if len(m.cfg.Projects) > 1 {
				m.selectedProj = (m.selectedProj + 1) % len(m.cfg.Projects)
				geo := theme.CalculateLayout(m.width, m.height)
				m.viewport.SetContent(m.renderContent(geo.InnerWidth))
				m.viewport.GotoTop()
			}
			return m, nil

		case "shift+tab", "p":
			if len(m.cfg.Projects) > 1 {
				m.selectedProj = (m.selectedProj - 1 + len(m.cfg.Projects)) % len(m.cfg.Projects)
				geo := theme.CalculateLayout(m.width, m.height)
				m.viewport.SetContent(m.renderContent(geo.InnerWidth))
				m.viewport.GotoTop()
			}
			return m, nil
		}
	}
	return m, nil
}

func (m ProjectsModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Loading technical records...")
	}
	return m.viewport.View()
}

func (m ProjectsModel) renderContent(width int) string {
	if len(m.cfg.Projects) == 0 {
		return theme.StyleBodyMuted.Render("No technical project records found.")
	}

	var b strings.Builder

	// ---------------- Top Table: Technical Records Ledger ----------------
	tableHeaderID := theme.TableHeaderStyle.Width(6).Render("ID")
	tableHeaderName := theme.TableHeaderStyle.Width(28).Render("PROJECT")
	tableHeaderStatus := theme.TableHeaderStyle.Width(16).Render("STATUS")
	tableHeaderTech := theme.TableHeaderStyle.Render("TECHNOLOGIES")

	tableHeaderRow := lipgloss.JoinHorizontal(
		lipgloss.Left,
		tableHeaderID,
		" ",
		tableHeaderName,
		" ",
		tableHeaderStatus,
		" ",
		tableHeaderTech,
	)

	var tableRows []string
	tableRows = append(tableRows, tableHeaderRow)

	for i, p := range m.cfg.Projects {
		isSelected := i == m.selectedProj
		rowIDStr := fmt.Sprintf("%03d", i+1)
		stackSummary := strings.Join(p.Stack, ", ")

		idCell := theme.TableRowStyle.Width(6).Render(rowIDStr)
		nameCell := theme.TableRowStyle.Width(28).Render(p.Name)
		statusCell := theme.StatusPillDev.Width(16).Render(strings.ToUpper(p.Status))
		techCell := theme.StyleCode.Render(stackSummary)

		if isSelected {
			idCell = theme.TableRowSelectedStyle.Width(6).Render(theme.GlyphBullet + " " + rowIDStr)
			nameCell = theme.TableRowSelectedStyle.Width(28).Render(p.Name)
		}

		row := lipgloss.JoinHorizontal(
			lipgloss.Left,
			idCell,
			" ",
			nameCell,
			" ",
			statusCell,
			" ",
			techCell,
		)
		tableRows = append(tableRows, row)
	}

	tablePanel := theme.RenderWorkstationPanel("TECHNICAL RECORDS LEDGER", strings.Join(tableRows, "\n"), width, 0, false)
	b.WriteString(tablePanel + "\n\n")

	// ---------------- Bottom: Selected Project Technical Dossier ----------------
	proj := m.cfg.Projects[m.selectedProj]

	// 1. Overview & Problem Definition
	descText := theme.StyleBody.Width(width - 6).Render(proj.Description)
	pipeText := lipgloss.NewStyle().
		Foreground(theme.ColorInfoCyan).
		Render("Conceptual Flow: [GitHub Issue] ──▶ [Trigger Intake] ──▶ [DAG Engine] ──▶ [AI Summary] ──▶ [Slack Webhook]")

	overviewBlock := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.RenderKeyValue("NAME:", proj.Name, 12),
		theme.RenderKeyValue("STATUS:", proj.Status, 12),
		theme.RenderKeyValue("STACK:", strings.Join(proj.Stack, " • "), 12),
		"",
		theme.StylePanelHeading.Render("PROBLEM BEING SOLVED"),
		descText,
		"",
		pipeText,
	)
	b.WriteString(theme.RenderWorkstationPanel("PROJECT SPECIFICATION & ARCHITECTURE", overviewBlock, width, 0, false) + "\n\n")

	// 2. Engineering Architecture Areas
	var archRows []string
	for _, area := range proj.ArchitectureAreas {
		archRows = append(archRows, theme.StyleBody.Render("  "+theme.GlyphBullet+" "+area))
	}
	b.WriteString(theme.RenderWorkstationPanel("ENGINEERING SUBSYSTEMS", strings.Join(archRows, "\n"), width, 0, false) + "\n\n")

	// 3. Key Implementation Decisions
	var decisionRows []string
	for _, dec := range proj.Decisions {
		marker := theme.StyleBrand.Render(theme.GlyphBullet + " ")
		text := theme.StyleBody.Width(width - 8).Render(dec)
		decisionRows = append(decisionRows, marker+text)
	}
	b.WriteString(theme.RenderWorkstationPanel("KEY ARCHITECTURAL DECISIONS", strings.Join(decisionRows, "\n\n"), width, 0, false) + "\n\n")

	// 4. Engineering Challenges
	var challengeRows []string
	for _, ch := range proj.Challenges {
		marker := theme.StyleStatusWarning.Render(theme.GlyphLightning + " ")
		text := theme.StyleBody.Width(width - 8).Render(ch)
		challengeRows = append(challengeRows, marker+text)
	}
	b.WriteString(theme.RenderWorkstationPanel("ENGINEERING CHALLENGES & IDEMPOTENCY", strings.Join(challengeRows, "\n\n"), width, 0, false) + "\n\n")

	// 5. What I Learned & Takeaways
	var learningRows []string
	for _, lr := range proj.Learnings {
		marker := theme.StyleStatusOnline.Render(theme.GlyphCheck + " ")
		text := theme.StyleBody.Width(width - 8).Render(lr)
		learningRows = append(learningRows, marker+text)
	}
	b.WriteString(theme.RenderWorkstationPanel("ENGINEERING TAKEAWAYS", strings.Join(learningRows, "\n\n"), width, 0, false) + "\n\n")

	// 6. External Links (if provided)
	if proj.GitHubURL != "" || proj.LiveURL != "" {
		var links []string
		if proj.GitHubURL != "" {
			links = append(links, theme.RenderKeyValue("REPOSITORY:", proj.GitHubURL, 14))
		}
		if proj.LiveURL != "" {
			links = append(links, theme.RenderKeyValue("DEPLOYMENT:", proj.LiveURL, 14))
		}
		b.WriteString(theme.RenderWorkstationPanel("PROJECT ENDPOINTS", strings.Join(links, "\n"), width, 0, false) + "\n")
	}

	return b.String()
}
