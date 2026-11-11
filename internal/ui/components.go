package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BrandBanner returns an elegant typography banner for JAWAD.OS.
func BrandBanner(compact bool) string {
	if compact {
		return lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccent).
			Render("╭─ JAWAD.OS ────────────────────────╮")
	}

	banner := `  ██╗ █████╗ ██╗    ██╗ █████╗ ██████╗       ██████╗ ███████╗
  ██║██╔══██╗██║    ██║██╔══██╗██╔══██╗     ██╔═══██╗██╔════╝
  ██║███████║██║ █╗ ██║███████║██║  ██║     ██║   ██║███████╗
████║██╔══██║██║███╗██║██╔══██║██║  ██║     ██║   ██║╚════██║
╚███║██║  ██║╚███╔███╔╝██║  ██║██████╔╝     ╚██████╔╝███████║
 ╚══╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═════╝       ╚═════╝ ╚══════╝`

	return lipgloss.NewStyle().
		Foreground(ColorAccent).
		Bold(true).
		Render(banner)
}

// RenderCard wraps content in a styled panel with an optional header.
func RenderCard(title string, content string, width int, active bool) string {
	style := PanelStyle.Width(width)
	if active {
		style = ActivePanelStyle.Width(width)
	}

	if title == "" {
		return style.Render(content)
	}

	styledTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorSecondary).
		Render(title)

	dividerLen := width - 6
	if dividerLen < 0 {
		dividerLen = 0
	}
	divider := lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat("─", dividerLen))

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		styledTitle,
		divider,
		content,
	)

	return style.Render(body)
}

// RenderKeyValue formats a label and a value cleanly.
func RenderKeyValue(label, value string, labelWidth int) string {
	formattedLabel := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorMuted).
		Width(labelWidth).
		Render(label)

	formattedValue := lipgloss.NewStyle().
		Foreground(ColorText).
		Render(value)

	return fmt.Sprintf("%s %s", formattedLabel, formattedValue)
}

// RenderTags formats a list of strings as subtle pill badges.
func RenderTags(tags []string) string {
	var rendered []string
	for _, tag := range tags {
		rendered = append(rendered, TagStyle.Render(tag))
	}
	return strings.Join(rendered, " ")
}

// RenderBadge creates an accent pill.
func RenderBadge(text string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(ColorDarkBg).
		Background(color).
		Bold(true).
		Padding(0, 1).
		Render(text)
}
