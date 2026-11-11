package theme

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// LayoutGeometry stores calculated viewport dimensions and panel boundaries.
type LayoutGeometry struct {
	TotalWidth   int
	TotalHeight  int
	InnerWidth   int
	InnerHeight  int
	MainWidth    int
	SideWidth    int
	IsMultiPanel bool
	IsTooSmall   bool
}

// CalculateLayout computes responsive dimensions based on terminal window geometry.
func CalculateLayout(width, height int) LayoutGeometry {
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}

	tooSmall := width < MinTerminalWidth || height < MinTerminalHeight

	// Usable inner width
	innerWidth := width - 2
	if innerWidth < MinTerminalWidth-2 {
		innerWidth = MinTerminalWidth - 2
	}

	// Workstation chrome: Header (1) + Divider (1) + Nav (1) + Divider (1) + Command Bar (1) + Footer (1) = 6 rows
	innerHeight := height - 6
	if innerHeight < 8 {
		innerHeight = 8
	}

	isMulti := width >= MultiPanelBreakpointWidth
	var mainW, sideW int

	if isMulti {
		// ~68% main workspace, ~32% contextual telemetry
		mainW = int(float64(innerWidth) * MainPanelRatio)
		sideW = innerWidth - mainW - 1 // 1 column for divider/spacing
		if sideW < 26 {
			sideW = 26
			mainW = innerWidth - sideW - 1
		}
	} else {
		mainW = innerWidth
		sideW = 0
	}

	return LayoutGeometry{
		TotalWidth:   width,
		TotalHeight:  height,
		InnerWidth:   innerWidth,
		InnerHeight:  innerHeight,
		MainWidth:    mainW,
		SideWidth:    sideW,
		IsMultiPanel: isMulti,
		IsTooSmall:   tooSmall,
	}
}

// RenderWorkstationHeader produces the clean top status bar.
func RenderWorkstationHeader(width int, sectionName string, sessionUser string, timestamp string) string {
	leftBrand := StyleBrand.Render("JAWAD.OS")

	var left string
	if sectionName != "" && sectionName != "WORKSTATION" {
		secTag := StyleSectionHeading.Render(strings.ToUpper(sectionName))
		left = leftBrand + StyleBodyMuted.Render("  /  ") + secTag
	} else {
		left = leftBrand + StyleBodyMuted.Render("  /  ") + StyleIdentity.Render("WORKSTATION")
	}

	statusDot := StyleStatusOnline.Render("ONLINE " + GlyphDot)
	rightParts := []string{statusDot}

	if sessionUser != "" {
		userTag := StyleBodyMuted.Render("user: ") + StyleMetaValue.Render(sessionUser)
		rightParts = append(rightParts, userTag)
	}

	if timestamp != "" {
		timeTag := StyleBodyMuted.Render(timestamp)
		rightParts = append(rightParts, timeTag)
	}

	right := strings.Join(rightParts, StyleBodyMuted.Render("  ·  "))

	gap := width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 2 {
		gap = 2
	}

	line := left + strings.Repeat(" ", gap) + right
	divider := RenderDivider(width)

	return line + "\n" + divider
}

// RenderWorkstationNav renders the minimalist navigation bar.
func RenderWorkstationNav(width int, items []string, activeIndex int) string {
	var rendered []string
	for i, item := range items {
		if i == activeIndex {
			rendered = append(rendered, NavTabActiveStyle.Render("● "+item))
		} else {
			rendered = append(rendered, NavTabStyle.Render("  "+item))
		}
	}

	navLine := strings.Join(rendered, " ")
	divider := RenderDivider(width)

	return navLine + "\n" + divider
}

// RenderWorkstationCommandBar renders the minimalist command bar.
func RenderWorkstationCommandBar(width int, isCommandMode bool, commandInput string, statusMessage string) string {
	prompt := CommandPromptStyle.Render("COMMAND > ")

	if isCommandMode {
		input := CommandInputStyle.Render(commandInput)
		cursor := StyleBrand.Render("▎")
		hint := StyleBodyMuted.Render("   (enter to execute, esc to cancel)")
		line := prompt + input + cursor + hint
		return RenderDivider(width) + "\n" + line
	}

	if statusMessage != "" {
		msg := StyleSectionHeading.Render("› " + statusMessage)
		return RenderDivider(width) + "\n" + msg
	}

	idleHint := StyleBodyMuted.Render("type ':' for command mode, '?' for help")
	return RenderDivider(width) + "\n" + prompt + idleHint
}

// RenderWorkstationFooter renders the keybinding shortcuts row.
func RenderWorkstationFooter(width int, shortcuts [][]string) string {
	var badges []string
	for _, sc := range shortcuts {
		k := FooterKeyBadge.Render(sc[0])
		d := FooterKeyDesc.Render(sc[1])
		badges = append(badges, fmt.Sprintf("%s %s", k, d))
	}

	line := strings.Join(badges, StyleBodyMuted.Render("  ·  "))
	return RenderDivider(width) + "\n" + line
}

// RenderDualPane joins left and right panels horizontally.
func RenderDualPane(left, right string, leftWidth, rightWidth, height int) string {
	leftStyled := lipgloss.NewStyle().
		Width(leftWidth).
		Render(left)

	rightStyled := lipgloss.NewStyle().
		Width(rightWidth).
		Render(right)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftStyled, " ", rightStyled)
}

// RenderConstrainedWindowWarning displays a structured alert when terminal dimensions are below minimum.
func RenderConstrainedWindowWarning(width, height int) string {
	box := lipgloss.NewStyle().
		Border(WorkstationPanelBorder).
		BorderForeground(ColorWarning).
		Padding(1, 2).
		Align(lipgloss.Center)

	title := StyleStatusWarning.Render("TERMINAL WINDOW CONSTRAINED")
	msg := StyleBody.Render("Please expand your terminal window to experience JAWAD.OS.")
	req := StyleMetaLabel.Render(fmt.Sprintf("Required: %dx%d | Current: %dx%d", MinTerminalWidth, MinTerminalHeight, width, height))

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		msg,
		"",
		req,
	)

	return box.Render(content)
}
