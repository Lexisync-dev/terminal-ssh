package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// Terminal Color Tokens for Bloomberg-Inspired Engineering Workstation.
// The palette emphasizes information density, visual hierarchy, and institutional precision.
var (
	// Background: Main terminal canvas
	ColorBg = lipgloss.Color("#101214")

	// Surface: Primary panels & information tables
	ColorSurface = lipgloss.Color("#181B1E")

	// Surface Elevated: Secondary panels, active table rows, modals
	ColorSurfaceElevated = lipgloss.Color("#202428")

	// Border: Separators, structural dividers, and grid lines
	ColorBorder = lipgloss.Color("#343A40")

	// Primary Text: Main content, labels, and values
	ColorTextPrimary = lipgloss.Color("#E6E8E9")

	// Secondary Text: Metadata, timestamps, shortcuts, secondary hints
	ColorTextSecondary = lipgloss.Color("#92999F")

	// Accent Amber: Primary highlight, active navigation, tickers, important markers
	ColorAccentAmber = lipgloss.Color("#F0A52B")

	// Success Green: Healthy status, online indicator, completed records
	ColorSuccessGreen = lipgloss.Color("#76B88A")

	// Info Cyan: Technical metadata, architectural tags, stack types
	ColorInfoCyan = lipgloss.Color("#67B8C8")

	// Warning / Error: Alert status, development flags, error indicators
	ColorWarning = lipgloss.Color("#D97757")
)
