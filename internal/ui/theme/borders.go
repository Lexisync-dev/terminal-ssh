package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Sharp single-line border definitions matching Bloomberg & engineering workstations.
var (
	// WorkstationPanelBorder uses clean single-line boundaries without rounded corners.
	WorkstationPanelBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	WorkstationHeaderBorder = lipgloss.Border{
		Bottom: "─",
	}

	WorkstationDividerBorder = lipgloss.Border{
		Right: "│",
	}
)

// Box drawing glyph constants for precision tabular layouts
const (
	GlyphHLine     = "─"
	GlyphVLine     = "│"
	GlyphCross     = "┼"
	GlyphTopT      = "┬"
	GlyphBottomT   = "┴"
	GlyphLeftT     = "├"
	GlyphRightT    = "┤"
	GlyphTopLeft   = "┌"
	GlyphTopRight  = "┐"
	GlyphBotLeft   = "└"
	GlyphBotRight  = "┘"
	GlyphBullet    = "▸"
	GlyphCheck     = "✔"
	GlyphLightning = "⚡"
	GlyphDot       = "●"
)

// RenderDivider returns a full-width horizontal divider line.
func RenderDivider(width int) string {
	if width <= 0 {
		return ""
	}
	return lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat(GlyphHLine, width))
}
