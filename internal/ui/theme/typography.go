package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Typography Hierarchy Levels for the Workstation Interface.
var (
	// Level 1 — Identity & Brand: Strong emphasis, neutral foreground
	StyleIdentity = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorTextPrimary)

	StyleBrand = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccentAmber)

	// Level 2 — Section Headings: Uppercase, Amber foreground, compact
	StyleSectionHeading = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber)

	// Level 3 — Panel Headings: Smaller uppercase labels, subtle surface contrast
	StylePanelHeading = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorAccentAmber).
				Background(ColorSurfaceElevated).
				Padding(0, 1)

	StylePanelHeadingMuted = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorTextSecondary).
				Background(ColorSurface).
				Padding(0, 1)

	// Level 4 — Body: Neutral, readable content
	StyleBody = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)

	StyleBodyMuted = lipgloss.NewStyle().
			Foreground(ColorTextSecondary)

	// Level 5 — Metadata: Muted slate, compact
	StyleMetaLabel = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorTextSecondary)

	StyleMetaValue = lipgloss.NewStyle().
			Foreground(ColorTextPrimary)

	StyleCode = lipgloss.NewStyle().
			Foreground(ColorInfoCyan).
			Background(ColorSurfaceElevated).
			Padding(0, 1)

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

// FormatSectionHeader formats a standard workstation section banner.
func FormatSectionHeader(id, title string) string {
	tag := StyleBrand.Render("[" + id + "]")
	name := StyleSectionHeading.Render(strings.ToUpper(title))
	return tag + " " + name
}
