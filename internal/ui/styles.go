package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Color Palette
	ColorPrimary   = lipgloss.Color("#3FB950") // Muted Green
	ColorSecondary = lipgloss.Color("#38BDF8") // Cyan
	ColorAccent    = lipgloss.Color("#F59E0B") // Amber
	ColorText      = lipgloss.Color("#E6EDF3") // Soft White
	ColorMuted     = lipgloss.Color("#8B949E") // Muted Gray
	ColorBorder    = lipgloss.Color("#30363D") // Slate Border
	ColorBg        = lipgloss.Color("#0D1117") // Dark Terminal Bg
	ColorDarkBg    = lipgloss.Color("#0D1117") // Deep dark background
	ColorHighlight = lipgloss.Color("#56D364") // Bright Green Accent
	ColorError     = lipgloss.Color("#F85149") // Red
	ColorCodeBg    = lipgloss.Color("#161B22") // Subtle card bg

	// Core Layout Styles
	DocStyle = lipgloss.NewStyle().
			Padding(0, 1)

	// Header Styles
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorText)

	HeaderBrand = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			SetString("JAWAD.OS")

	StatusOnline = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	StatusDot = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			SetString("●")

	// Section Titles & Subtitles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorSecondary).
			MarginBottom(1)

	SectionHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorSecondary)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	TaglineStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Italic(true)

	// Card & Panel Styles
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2)

	ActivePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorSecondary).
				Padding(1, 2)

	// Navigation & List Item Styles
	ItemNumberStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	ItemTitleStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Bold(true)

	ItemDescStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	SelectedRowStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	UnselectedRowStyle = lipgloss.NewStyle().
				Foreground(ColorText)

	// Footer & Keybinding Styles
	FooterStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			PaddingTop(1)

	KeyBadge = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	KeyDesc = lipgloss.NewStyle().
		Foreground(ColorMuted)

	// Status / Badge Styles
	Badge = lipgloss.NewStyle().
		Foreground(ColorDarkBg).
		Background(ColorSecondary).
		Padding(0, 1).
		Bold(true)

	StatusBadge = lipgloss.NewStyle().
			Foreground(ColorDarkBg).
			Background(ColorAccent).
			Padding(0, 1).
			Bold(true)

	// Code & Tag Styles
	TagStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Background(ColorCodeBg).
			Padding(0, 1).
			MarginRight(1)

	// Command Prompt Styles
	CommandPromptStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary).
				Bold(true)

	CommandInputStyle = lipgloss.NewStyle().
				Foreground(ColorText)

	CommandLineBox = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, false, false, false).
			BorderForeground(ColorBorder).
			Padding(0, 1)

	// Error & Warning
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)
)
