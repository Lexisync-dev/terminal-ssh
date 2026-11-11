package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Minimalist Color Palette
	ColorPrimary   = lipgloss.Color("#76B88A") // Soft sage green
	ColorSecondary = lipgloss.Color("#6FAFC7") // Muted cyan
	ColorAccent    = lipgloss.Color("#F0A52B") // Warm amber
	ColorText      = lipgloss.Color("#ECEFF1") // Crisp off-white
	ColorMuted     = lipgloss.Color("#7E8691") // Muted slate
	ColorBorder    = lipgloss.Color("#252A31") // Subtle border
	ColorBg        = lipgloss.Color("#0E1013") // Deep charcoal
	ColorDarkBg    = lipgloss.Color("#0E1013") // Deep charcoal
	ColorHighlight = lipgloss.Color("#F0A52B") // Warm amber accent
	ColorError     = lipgloss.Color("#D4775E") // Soft terracotta
	ColorCodeBg    = lipgloss.Color("#13161B") // Subtle dark surface

	// Core Layout Styles
	DocStyle = lipgloss.NewStyle().
			Padding(0, 1)

	// Header Styles
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorText)

	HeaderBrand = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccent).
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
			Foreground(ColorText).
			MarginBottom(1)

	SectionHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorText)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	TaglineStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Italic(true)

	// Card & Panel Styles
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(0, 1)

	ActivePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorAccent).
				Padding(0, 1)

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
				Foreground(ColorAccent).
				Bold(true)

	SelectedRowStyle = lipgloss.NewStyle().
				Foreground(ColorAccent).
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
		Foreground(ColorSecondary).
		Bold(true)

	StatusBadge = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	// Code & Tag Styles
	TagStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			MarginRight(1)

	// Command Prompt Styles
	CommandPromptStyle = lipgloss.NewStyle().
				Foreground(ColorAccent).
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

