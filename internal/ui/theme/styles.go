package theme

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Panel Styles: Structured workstation panels with crisp single borders
	PanelStyle = lipgloss.NewStyle().
			Border(WorkstationPanelBorder).
			BorderForeground(ColorBorder).
			Background(ColorSurface).
			Padding(0, 1)

	PanelActiveStyle = lipgloss.NewStyle().
				Border(WorkstationPanelBorder).
				BorderForeground(ColorAccentAmber).
				Background(ColorSurface).
				Padding(0, 1)

	PanelElevatedStyle = lipgloss.NewStyle().
				Border(WorkstationPanelBorder).
				BorderForeground(ColorBorder).
				Background(ColorSurfaceElevated).
				Padding(0, 1)

	// Header Styles: Bloomberg institutional top bar
	HeaderStyle = lipgloss.NewStyle().
			Background(ColorSurface).
			Foreground(ColorTextPrimary).
			Bold(true).
			Padding(0, 1)

	// Navigation Tickers: Compact numbered modules
	NavTabStyle = lipgloss.NewStyle().
			Foreground(ColorTextSecondary).
			Background(ColorBg).
			Padding(0, 1)

	NavTabActiveStyle = lipgloss.NewStyle().
				Foreground(ColorBg).
				Background(ColorAccentAmber).
				Bold(true).
				Padding(0, 1)

	// Data Table Styles: Compact technical records
	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber).
				Background(ColorSurfaceElevated).
				Padding(0, 1)

	TableRowStyle = lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Padding(0, 1)

	TableRowSelectedStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber).
				Background(ColorSurfaceElevated).
				Padding(0, 1)

	// Command Prompt: Terminal input bar
	CommandBarStyle = lipgloss.NewStyle().
			Border(lipgloss.Border{Top: "─"}).
			BorderForeground(ColorBorder).
			Foreground(ColorTextPrimary).
			Padding(0, 1)

	CommandPromptStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber)

	CommandInputStyle = lipgloss.NewStyle().
				Foreground(ColorTextPrimary)

	// Footer: Contextual keybinding shortcuts
	FooterStyle = lipgloss.NewStyle().
			Border(lipgloss.Border{Top: "─"}).
			BorderForeground(ColorBorder).
			Foreground(ColorTextSecondary).
			Padding(0, 1)

	FooterKeyBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccentAmber)

	FooterKeyDesc = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)

	// Metadata & Tag Pill Badges
	TagStyle = lipgloss.NewStyle().
			Foreground(ColorInfoCyan).
			Background(ColorSurfaceElevated).
			Padding(0, 1).
			MarginRight(1)

	TagAmberStyle = lipgloss.NewStyle().
			Foreground(ColorAccentAmber).
			Background(ColorSurfaceElevated).
			Padding(0, 1).
			MarginRight(1)

	StatusPillActive = lipgloss.NewStyle().
				Foreground(ColorBg).
				Background(ColorSuccessGreen).
				Bold(true).
				Padding(0, 1)

	StatusPillDev = lipgloss.NewStyle().
			Foreground(ColorBg).
			Background(ColorAccentAmber).
			Bold(true).
			Padding(0, 1)
)

// RenderWorkstationPanel wraps content inside a sharp workstation frame with an optional title.
func RenderWorkstationPanel(title string, content string, width, height int, active bool) string {
	style := PanelStyle.Width(width)
	if active {
		style = PanelActiveStyle.Width(width)
	}
	if height > 0 {
		style = style.Height(height)
	}

	if title == "" {
		return style.Render(content)
	}

	headerTag := StylePanelHeading.Render(strings.ToUpper(title))
	dividerLen := width - lipgloss.Width(headerTag) - 4
	if dividerLen < 0 {
		dividerLen = 0
	}
	topDivider := lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat(GlyphHLine, dividerLen))

	headerLine := headerTag + " " + topDivider

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		headerLine,
		content,
	)

	return style.Render(body)
}

// RenderKeyValue formats a technical label and value with aligned columns.
func RenderKeyValue(label, value string, labelWidth int) string {
	lbl := StyleMetaLabel.Width(labelWidth).Render(label)
	val := StyleMetaValue.Render(value)
	return fmt.Sprintf("%s %s", lbl, val)
}

// RenderTags formats multiple items into institutional badges.
func RenderTags(tags []string) string {
	var rendered []string
	for _, t := range tags {
		rendered = append(rendered, TagStyle.Render(t))
	}
	return strings.Join(rendered, " ")
}
