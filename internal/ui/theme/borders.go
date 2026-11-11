package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Minimalist single-line and rounded border definitions.
var (
	// WorkstationPanelBorder uses clean rounded boundaries for a modern, minimalist frame.
	WorkstationPanelBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	WorkstationHeaderBorder = lipgloss.Border{
		Bottom: "─",
	}

	WorkstationDividerBorder = lipgloss.Border{
		Right: "│",
	}
)

// Box drawing and minimalist typographic glyph constants
const (
	GlyphHLine     = "─"
	GlyphVLine     = "│"
	GlyphCross     = "┼"
	GlyphTopT      = "┬"
	GlyphBottomT   = "┴"
	GlyphLeftT     = "├"
	GlyphRightT    = "┤"
	GlyphTopLeft   = "╭"
	GlyphTopRight  = "╮"
	GlyphBotLeft   = "╰"
	GlyphBotRight  = "╯"
	GlyphBullet    = "·"
	GlyphCheck     = "✓"
	GlyphLightning = "⚡"
	GlyphDot       = "●"
)

// RenderDivider returns a full-width subtle horizontal divider line.
func RenderDivider(width int) string {
	if width <= 0 {
		return ""
	}
	return lipgloss.NewStyle().
		Foreground(ColorBorder).
		Render(strings.Repeat(GlyphHLine, width))
}

