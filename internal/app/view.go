package app

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/commands"
	"github.com/lexisync/jawad-os/internal/ui"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

// View renders the terminal user interface according to active state and dimensions.
func (m Model) View() string {
	geo := theme.CalculateLayout(m.Width, m.Height)

	// Guard against unreadably constrained terminal dimensions
	if geo.IsTooSmall {
		return theme.RenderConstrainedWindowWarning(m.Width, m.Height)
	}

	// Boot screen has its own minimal sequence layout
	if m.CurrentScreen == ui.ScreenBoot {
		return m.BootScreen.View()
	}

	header := theme.RenderWorkstationHeader(geo.TotalWidth, m.CurrentScreen.String(), m.User, "v1.0")
	nav := theme.RenderWorkstationNav(geo.TotalWidth, ui.WorkstationNavLabels(), ui.ScreenToNavIndex(m.CurrentScreen))
	commandBar := theme.RenderWorkstationCommandBar(geo.TotalWidth, m.CommandMode, m.CommandInput, m.StatusMessage)

	// Shortcuts for the footer
	var shortcuts [][]string
	if m.CurrentScreen == ui.ScreenDashboard {
		shortcuts = [][]string{
			{"F1/?", "HELP"},
			{"↑↓", "NAVIGATE"},
			{"ENTER", "SELECT"},
			{"1-6", "JUMP"},
			{":", "COMMAND"},
			{"ESC", "BACK"},
			{"Q", "EXIT"},
		}
	} else if m.CurrentScreen == ui.ScreenHelp {
		shortcuts = [][]string{
			{"ESC", "CLOSE"},
			{":", "COMMAND"},
			{"Q", "EXIT"},
		}
	} else {
		shortcuts = [][]string{
			{"F1/?", "HELP"},
			{"↑↓", "SCROLL"},
			{"ESC", "BACK"},
			{":", "COMMAND"},
			{"Q", "EXIT"},
		}
	}
	footer := theme.RenderWorkstationFooter(geo.TotalWidth, shortcuts)

	// Modal overlay takes center stage if active (for easter eggs & system alerts)
	if m.ModalActive {
		dismissHint := theme.StyleBodyMuted.Render("\n  [Press any key or ESC to return]")

		modalBox := theme.RenderWorkstationPanel(m.ModalTitle, m.ModalContent+dismissHint, geo.InnerWidth, 0, true)

		content := lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			nav,
			"\n"+modalBox,
			"\n"+commandBar,
			footer,
		)
		return content
	}

	var body string
	switch m.CurrentScreen {
	case ui.ScreenDashboard:
		body = m.DashboardScreen.View()

	case ui.ScreenAbout:
		body = m.AboutScreen.View()

	case ui.ScreenProjects:
		body = m.ProjectsScreen.View()

	case ui.ScreenSkills:
		body = m.SkillsScreen.View()

	case ui.ScreenExperience:
		body = m.ExperienceScreen.View()

	case ui.ScreenInterests:
		body = m.InterestsScreen.View()

	case ui.ScreenContact:
		body = m.ContactScreen.View()

	case ui.ScreenHelp:
		body = m.renderHelpScreen(geo.InnerWidth)

	default:
		body = m.renderPlaceholderScreen(geo.InnerWidth)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		nav,
		"\n"+body,
		"\n"+commandBar,
		footer,
	)

	return content
}

func (m Model) renderHelpScreen(width int) string {
	var b strings.Builder

	keys := [][]string{
		{"↑ / k", "Navigate up in menus / scroll viewport"},
		{"↓ / j", "Navigate down in menus / scroll viewport"},
		{"PgUp / PgDn", "Fast scroll up/down in long modules"},
		{"Mouse Wheel", "Scroll active screen viewport"},
		{"ENTER", "Select item / Copy contact value / Confirm"},
		{"1 - 6", "Quick jump to corresponding module"},
		{":", "Enter command mode (type commands like 'projects', 'help')"},
		{"ESC", "Return to previous screen / Cancel command mode / Close modal"},
		{"F1 / ?", "Toggle this help screen"},
		{"q / Ctrl+C", "Quit or disconnect from JAWAD.OS"},
	}

	b.WriteString(theme.StyleSectionHeading.Render("KEYBOARD & MOUSE SHORTCUTS:") + "\n\n")
	for _, k := range keys {
		keyBadge := theme.FooterKeyBadge.Width(16).Render(k[0])
		desc := theme.StyleBody.Render(k[1])
		b.WriteString("  " + keyBadge + " " + desc + "\n")
	}

	b.WriteString("\n" + theme.StyleSectionHeading.Render("COMMAND MODE REFERENCE (Press ':' to open):") + "\n\n")

	for _, c := range commands.AvailableCommands {
		cmdName := theme.StyleBrand.Width(20).Render(":" + c.Name)
		cmdDesc := theme.StyleBodyMuted.Render(c.Description)
		b.WriteString("  " + cmdName + " " + cmdDesc + "\n")
	}

	return theme.RenderWorkstationPanel("HELP & WORKSTATION REFERENCE", b.String(), width, 0, false)
}

func (m Model) renderPlaceholderScreen(width int) string {
	screenName := m.CurrentScreen.String()
	title := theme.StyleSectionHeading.Render("MODULE: " + strings.ToUpper(screenName))
	desc := theme.StyleBody.Render("Active screen view.")
	hint := theme.StyleBodyMuted.Render("Press ESC to return to Dashboard.")

	content := lipgloss.JoinVertical(lipgloss.Left, title, "", desc, "", hint)
	return theme.RenderWorkstationPanel("MODULE VIEW", content, width, 0, false)
}
