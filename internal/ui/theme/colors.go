package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// Terminal Color Tokens for Minimalist Terminal Interface.
// The palette emphasizes visual restraint, typographic hierarchy, and serene clarity.
var (
	// Background: Deep charcoal canvas
	ColorBg = lipgloss.Color("#0E1013")

	// Surface: Primary panels & subtle content cards
	ColorSurface = lipgloss.Color("#13161B")

	// Surface Elevated: Elevated panels, active rows, focused cards
	ColorSurfaceElevated = lipgloss.Color("#1A1E24")

	// Border: Subtle, low-contrast boundary lines
	ColorBorder = lipgloss.Color("#252A31")

	// Primary Text: Crisp, high-readability content
	ColorTextPrimary = lipgloss.Color("#ECEFF1")

	// Secondary Text: Muted labels, secondary descriptions, shortcuts
	ColorTextSecondary = lipgloss.Color("#7E8691")

	// Accent Amber: Focused highlight, active navigation, precise markers
	ColorAccentAmber = lipgloss.Color("#F0A52B")

	// Success Green: Status indicators, confirmations
	ColorSuccessGreen = lipgloss.Color("#76B88A")

	// Info Cyan: Technical tags and metadata
	ColorInfoCyan = lipgloss.Color("#6FAFC7")

	// Warning / Error: Alert indicators
	ColorWarning = lipgloss.Color("#D4775E")
)

