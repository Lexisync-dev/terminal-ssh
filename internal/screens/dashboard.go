package screens

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type DashboardModel struct {
	cfg      *config.PortfolioConfig
	selected int
	width    int
	height   int
}

func NewDashboardModel(cfg *config.PortfolioConfig) DashboardModel {
	return DashboardModel{
		cfg:      cfg,
		selected: 0,
		width:    80,
		height:   24,
	}
}

func (m *DashboardModel) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func (m DashboardModel) SelectedIndex() int {
	return m.selected
}

func (m DashboardModel) SelectedScreen() ui.Screen {
	if m.selected >= 0 && m.selected < len(ui.MenuItems) {
		return ui.MenuItems[m.selected].Screen
	}
	return ui.ScreenDashboard
}

func (m DashboardModel) Update(msg tea.Msg) (DashboardModel, tea.Cmd, *ui.Screen) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			} else {
				m.selected = len(ui.MenuItems) - 1
			}
			return m, nil, nil

		case "down", "j":
			if m.selected < len(ui.MenuItems)-1 {
				m.selected++
			} else {
				m.selected = 0
			}
			return m, nil, nil

		case "enter":
			target := m.SelectedScreen()
			return m, nil, &target

		case "1", "2", "3", "4", "5", "6":
			idx := int(msg.String()[0]-'1')
			if idx >= 0 && idx < len(ui.MenuItems) {
				m.selected = idx
				target := ui.MenuItems[idx].Screen
				return m, nil, &target
			}
		}
	}

	return m, nil, nil
}

func (m DashboardModel) View() string {
	geo := theme.CalculateLayout(m.width, m.height)

	if geo.IsMultiPanel {
		return m.renderWideDashboard(geo)
	}
	return m.renderCompactDashboard(geo)
}

// renderWideDashboard renders the 70/30 dual-panel minimalist dashboard layout.
func (m DashboardModel) renderWideDashboard(geo theme.LayoutGeometry) string {
	mainWidth := geo.MainWidth
	sideWidth := geo.SideWidth

	// ---------------- Left Panel: Identity, Focus & Featured Project ----------------
	identityTitle := theme.StyleIdentity.Bold(true).Render(strings.ToUpper(m.cfg.Profile.Name))
	identityRole := theme.StyleSectionHeading.Render(m.cfg.Profile.Title)
	identityTagline := theme.StyleBodyMuted.Render(m.cfg.Profile.Tagline)

	identitySection := lipgloss.JoinVertical(
		lipgloss.Left,
		identityTitle,
		identityRole,
		"",
		identityTagline,
	)

	// Current Focus areas
	focusHeader := theme.StylePanelHeading.Render("CURRENT FOCUS")
	focusCol1 := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.StyleBody.Render("· Backend engineering"),
		theme.StyleBody.Render("· Workflow automation"),
	)
	focusCol2 := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.StyleBody.Render("· System design & architecture"),
		theme.StyleBody.Render("· Distributed systems"),
	)
	focusGrid := lipgloss.JoinHorizontal(lipgloss.Top, focusCol1, "    ", focusCol2)
	focusSection := lipgloss.JoinVertical(lipgloss.Left, focusHeader, "", focusGrid)

	// Featured Project Section
	var featuredSection string
	if len(m.cfg.Projects) > 0 {
		fp := m.cfg.Projects[0]
		fpStatus := theme.StatusPillDev.Render("[" + fp.Status + "]")
		fpTitle := theme.StyleIdentity.Render(fp.Name)
		fpTop := lipgloss.JoinHorizontal(lipgloss.Center, theme.StylePanelHeading.Render("FEATURED PROJECT"), "  ", fpStatus)

		fpStack := theme.RenderTags(fp.Stack)
		fpDesc := theme.StyleBody.Width(mainWidth - 6).Render(fp.Description)
		fpActionHint := theme.StyleBodyMuted.Render("Select '02 PROJECTS' or enter ':projects' for full technical architecture.")

		featuredSection = lipgloss.JoinVertical(
			lipgloss.Left,
			fpTop,
			"",
			fpTitle,
			"",
			fpStack,
			"",
			fpDesc,
			"",
			fpActionHint,
		)
	}

	leftBody := lipgloss.JoinVertical(
		lipgloss.Left,
		identitySection,
		"\n"+theme.RenderDivider(mainWidth-4),
		"\n"+focusSection,
		"\n"+theme.RenderDivider(mainWidth-4),
		"\n"+featuredSection,
	)

	leftPanel := theme.RenderWorkstationPanel("OVERVIEW", leftBody, mainWidth, 0, false)

	// ---------------- Right Panel: Profile Data & Module Directory ----------------
	profHeader := theme.StylePanelHeading.Render("PROFILE")
	locRow := theme.RenderKeyValue("LOCATION", m.cfg.Profile.Location, 12)
	eduRow := theme.RenderKeyValue("EDUCATION", m.cfg.Profile.Education, 12)
	dirRow := theme.RenderKeyValue("DIRECTION", "Full-Stack / Backend", 12)
	intRow := theme.RenderKeyValue("INTERESTS", "Systems, Automation, Infra", 12)

	profSection := lipgloss.JoinVertical(
		lipgloss.Left,
		profHeader,
		"",
		locRow,
		eduRow,
		dirRow,
		intRow,
	)

	// Module Directory
	dirHeader := theme.StylePanelHeading.Render("DIRECTORY")
	var menuLines []string
	for i, item := range ui.MenuItems {
		isSelected := i == m.selected
		prefix := "  "
		itemNum := theme.StyleMetaLabel.Render("[" + item.Number + "]")
		itemTitle := theme.StyleBody.Width(11).Render(item.Title)
		itemDesc := theme.StyleBodyMuted.Render(item.Subtitle)

		if isSelected {
			prefix = theme.StyleBrand.Render("● ")
			itemNum = theme.StyleBrand.Render("[" + item.Number + "]")
			itemTitle = theme.StyleBrand.Width(11).Render(item.Title)
			itemDesc = theme.StyleBody.Render(item.Subtitle)
		}

		line := fmt.Sprintf("%s%s %s %s", prefix, itemNum, itemTitle, itemDesc)
		menuLines = append(menuLines, line)
	}

	dirMenu := lipgloss.JoinVertical(
		lipgloss.Left,
		dirHeader,
		"",
		strings.Join(menuLines, "\n"),
	)

	rightBody := lipgloss.JoinVertical(
		lipgloss.Left,
		profSection,
		"\n"+theme.RenderDivider(sideWidth-4),
		"\n"+dirMenu,
	)

	rightPanel := theme.RenderWorkstationPanel("PROFILE & DIRECTORY", rightBody, sideWidth, 0, false)

	return theme.RenderDualPane(leftPanel, rightPanel, mainWidth, sideWidth, geo.InnerHeight)
}

// renderCompactDashboard renders the single-column vertically stacked layout for narrow terminals.
func (m DashboardModel) renderCompactDashboard(geo theme.LayoutGeometry) string {
	width := geo.InnerWidth

	identityTitle := theme.StyleIdentity.Bold(true).Render(strings.ToUpper(m.cfg.Profile.Name))
	identityRole := theme.StyleSectionHeading.Render(m.cfg.Profile.Title)
	location := theme.StyleBodyMuted.Render(m.cfg.Profile.Location + " · " + m.cfg.Profile.Education)

	identitySection := lipgloss.JoinVertical(
		lipgloss.Left,
		identityTitle,
		identityRole,
		location,
	)

	// Current Focus
	focusList := lipgloss.JoinVertical(
		lipgloss.Left,
		theme.StylePanelHeading.Render("CURRENT FOCUS"),
		"",
		theme.StyleBody.Render("· Backend engineering  · System design"),
		theme.StyleBody.Render("· Workflow automation  · Distributed systems"),
	)

	// Featured Project
	var featuredSection string
	if len(m.cfg.Projects) > 0 {
		fp := m.cfg.Projects[0]
		fpStatus := theme.StatusPillDev.Render("[" + fp.Status + "]")
		featuredSection = lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Center, theme.StylePanelHeading.Render("FEATURED: "+fp.Name), " ", fpStatus),
			"",
			theme.RenderTags(fp.Stack),
			"",
			theme.StyleBody.Width(width-6).Render(fp.Description),
		)
	}

	// Module Navigation
	var menuLines []string
	for i, item := range ui.MenuItems {
		isSelected := i == m.selected
		prefix := "  "
		if isSelected {
			prefix = theme.StyleBrand.Render("● ")
		}
		num := theme.StyleBrand.Render("[" + item.Number + "]")
		title := theme.StyleBody.Width(11).Render(item.Title)
		desc := theme.StyleBodyMuted.Render(item.Subtitle)
		menuLines = append(menuLines, fmt.Sprintf("%s%s %s %s", prefix, num, title, desc))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		identitySection,
		"\n"+theme.RenderDivider(width-4),
		"\n"+focusList,
		"\n"+theme.RenderDivider(width-4),
		"\n"+featuredSection,
		"\n"+theme.RenderDivider(width-4),
		"\n"+theme.StylePanelHeading.Render("SELECT MODULE")+"\n\n"+strings.Join(menuLines, "\n"),
	)

	return theme.RenderWorkstationPanel("OVERVIEW", content, width, 0, false)
}
