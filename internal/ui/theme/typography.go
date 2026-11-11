package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Typography Hierarchy Levels for the Minimalist Interface.
var (
	// Level 1 — Identity & Brand: Clean emphasis
	StyleIdentity = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorTextPrimary)

	StyleBrand = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccentAmber)

	// Level 2 — Section Headings: Clean, subtle accent
	StyleSectionHeading = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorTextPrimary)

	// Level 3 — Panel Headings: Minimal uppercase labels, clean typography without chunky background blocks
	StylePanelHeading = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber)

	StylePanelHeadingMuted = lipgloss.NewStyle().
				Foreground(ColorTextSecondary)

	// Level 4 — Body: Neutral, legible content
	StyleBody = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)

	StyleBodyMuted = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)

	// Level 5 — Metadata: Muted slate, compact
	StyleMetaLabel = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)

	StyleMetaValue = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)

	StyleCode = lipgloss.NewStyle().
			Foreground(ColorInfoCyan)

	// Level 6 — Status: Semantic colors
	StyleStatusOnline = lipgloss.NewStyle().
				Foreground(ColorSuccessGreen).
				Bold(true)

	StyleStatusWarning = lipgloss.NewStyle().
				Foreground(ColorWarning).
				Bold(true)

	StyleStatusInfo = lipgloss.NewStyle().
			Foreground(ColorInfoCyan).
			Bold(true)
)

// FormatSectionHeader formats a standard section banner.
func FormatSectionHeader(id, title string) string {
	tag := StyleBrand.Render("[" + id + "]")
	name := StyleSectionHeading.Render(strings.ToUpper(title))
	return tag + " " + name
}

