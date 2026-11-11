package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type LayoutDimensions struct {
	TotalWidth  int
	TotalHeight int
	InnerWidth  int
	InnerHeight int
	IsTooSmall  bool
	IsCompact   bool
}

const (
	MinWidth  = 50
	MinHeight = 14
)

// CalculateDimensions computes responsive layout dimensions based on terminal window size.
func CalculateDimensions(width, height int) LayoutDimensions {
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}

	tooSmall := width < MinWidth || height < MinHeight
	compact := width < 80 || height < 25

	// Maximum container width for readability
	innerWidth := width - 4
	if innerWidth > 96 {
		innerWidth = 96
	}
	if innerWidth < 46 {
		innerWidth = 46
	}

	innerHeight := height - 6 // Header (2) + Footer (3) + Padding (1)
	if innerHeight < 8 {
		innerHeight = 8
	}

	return LayoutDimensions{
		TotalWidth:  width,
		TotalHeight: height,
		InnerWidth:  innerWidth,
		InnerHeight: innerHeight,
		IsTooSmall:  tooSmall,
		IsCompact:   compact,
	}
}

// RenderHeader renders the top status bar.
func RenderHeader(width int, currentScreen Screen) string {
	if width < MinWidth {
		return "JAWAD.OS [ONLINE]"
	}

	left := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Render("JAWAD.OS")

	if currentScreen != ScreenDashboard && currentScreen != ScreenBoot {
		separator := lipgloss.NewStyle().Foreground(ColorMuted).Render("  /  ")
		screenName := lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorSecondary).
			Render(strings.ToUpper(currentScreen.String()))
		left = left + separator + screenName
	}

	right := lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.NewStyle().Foreground(ColorMuted).Render("STATUS: "),
		StatusOnline.Render("ONLINE "),
		StatusDot.Render(),
	)

	gap := width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 2 {
		gap = 2
	}

	headerLine := left + strings.Repeat(" ", gap) + right
	divider := lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat("─", width))

	return headerLine + "\n" + divider
}

// RenderFooter renders the bottom action and navigation hints.
func RenderFooter(width int, isCommandMode bool, commandInput string, statusMessage string, currentScreen Screen) string {
	divider := lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat("─", width))

	if isCommandMode {
		prompt := CommandPromptStyle.Render(":")
		input := CommandInputStyle.Render(commandInput)
		cursor := lipgloss.NewStyle().Foreground(ColorHighlight).Render("█")
		hint := lipgloss.NewStyle().Foreground(ColorMuted).Render("  (Press ENTER to run, ESC to cancel)")
		return divider + "\n" + prompt + " " + input + cursor + hint
	}

	if statusMessage != "" {
		msg := lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true).
			Render("› " + statusMessage)
		return divider + "\n" + msg
	}

	// Normal keybindings footer
	var hints []string
	if currentScreen == ScreenDashboard {
		hints = []string{
			KeyBadge.Render("↑/↓") + " " + KeyDesc.Render("Navigate"),
			KeyBadge.Render("ENTER") + " " + KeyDesc.Render("Select"),
			KeyBadge.Render("1-6") + " " + KeyDesc.Render("Quick Jump"),
			KeyBadge.Render(":") + " " + KeyDesc.Render("Command"),
			KeyBadge.Render("?") + " " + KeyDesc.Render("Help"),
			KeyBadge.Render("Q") + " " + KeyDesc.Render("Quit"),
		}
	} else {
		hints = []string{
			KeyBadge.Render("ESC") + " " + KeyDesc.Render("Back"),
			KeyBadge.Render("↑/↓") + " " + KeyDesc.Render("Scroll"),
			KeyBadge.Render(":") + " " + KeyDesc.Render("Command"),
			KeyBadge.Render("?") + " " + KeyDesc.Render("Help"),
			KeyBadge.Render("Q") + " " + KeyDesc.Render("Quit"),
		}
	}

	content := strings.Join(hints, "   ")
	return divider + "\n" + content
}

// RenderSmallScreenWarning displays a helpful message when the terminal window is constrained.
func RenderSmallScreenWarning(width, height int) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorAccent).
		Padding(1, 2).
		Align(lipgloss.Center)

	msg := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().Bold(true).Foreground(ColorAccent).Render("TERMINAL WINDOW TOO SMALL"),
		"",
		lipgloss.NewStyle().Foreground(ColorText).Render("Please expand your terminal window"),
		lipgloss.NewStyle().Foreground(ColorMuted).Render("to experience JAWAD.OS properly."),
		"",
		lipgloss.NewStyle().Foreground(ColorSecondary).Render(fmt.Sprintf("Required: 50x14 | Current: %dx%d", width, height)),
	)

	return box.Render(msg)
}
