package theme

import (
	"strings"
	"testing"
)

func TestColorsInitialized(t *testing.T) {
	// Verify hex color values are non-empty
	if string(ColorBg) == "" {
		t.Errorf("ColorBg should not be empty")
	}
	if string(ColorAccentAmber) != "#F0A52B" {
		t.Errorf("expected ColorAccentAmber to be #F0A52B, got %s", ColorAccentAmber)
	}
	if string(ColorSuccessGreen) != "#76B88A" {
		t.Errorf("expected ColorSuccessGreen to be #76B88A, got %s", ColorSuccessGreen)
	}
}

func TestCalculateLayout(t *testing.T) {
	// 1. Standard compact terminal 80x24
	geo80 := CalculateLayout(80, 24)
	if geo80.IsTooSmall {
		t.Errorf("80x24 should not be marked as too small")
	}
	if geo80.IsMultiPanel {
		t.Errorf("80x24 should be single-panel, got multi-panel")
	}
	if geo80.MainWidth <= 0 {
		t.Errorf("80x24 MainWidth should be positive")
	}

	// 2. Wide terminal 100x30 (Multi-panel breakpoint)
	geo100 := CalculateLayout(100, 30)
	if !geo100.IsMultiPanel {
		t.Errorf("100x30 should activate multi-panel layout")
	}
	if geo100.SideWidth <= 0 {
		t.Errorf("100x30 SideWidth should be positive")
	}
	if geo100.MainWidth+geo100.SideWidth >= geo100.TotalWidth {
		t.Errorf("panels total width should fit within terminal width")
	}

	// 3. Ultra-wide terminal 160x50
	geo160 := CalculateLayout(160, 50)
	if !geo160.IsMultiPanel {
		t.Errorf("160x50 should be multi-panel")
	}

	// 4. Constrained terminal (too small)
	geoSmall := CalculateLayout(40, 10)
	if !geoSmall.IsTooSmall {
		t.Errorf("40x10 should be flagged as too small")
	}
}

func TestRenderWorkstationHeader(t *testing.T) {
	hdr := RenderWorkstationHeader(100, "Projects", "recruiter", "12:00:00 UTC")
	if !strings.Contains(hdr, "JAWAD.OS") {
		t.Errorf("expected header to contain 'JAWAD.OS'")
	}
	if !strings.Contains(hdr, "PROJECTS") {
		t.Errorf("expected header to contain 'PROJECTS'")
	}
	if !strings.Contains(hdr, "ONLINE") {
		t.Errorf("expected header to contain 'ONLINE'")
	}
	if !strings.Contains(hdr, "recruiter") {
		t.Errorf("expected header to contain session user 'recruiter'")
	}
}

func TestRenderWorkstationNav(t *testing.T) {
	items := []string{"01 PROFILE", "02 PROJECTS", "03 STACK"}
	nav := RenderWorkstationNav(100, items, 1)

	if !strings.Contains(nav, "01 PROFILE") || !strings.Contains(nav, "02 PROJECTS") {
		t.Errorf("expected nav to contain items")
	}
}

func TestRenderWorkstationCommandBar(t *testing.T) {
	// Idle command bar
	idle := RenderWorkstationCommandBar(80, false, "", "")
	if !strings.Contains(idle, "COMMAND >") {
		t.Errorf("expected command bar to contain 'COMMAND >'")
	}

	// Active command bar
	active := RenderWorkstationCommandBar(80, true, "about", "")
	if !strings.Contains(active, "about") {
		t.Errorf("expected active command bar to contain input text 'about'")
	}

	// Status message in command bar
	msg := RenderWorkstationCommandBar(80, false, "", "Copied to clipboard")
	if !strings.Contains(msg, "Copied to clipboard") {
		t.Errorf("expected command bar to display status message")
	}
}

func TestRenderWorkstationPanel(t *testing.T) {
	panel := RenderWorkstationPanel("TEST PANEL", "panel body content", 50, 10, false)
	if !strings.Contains(panel, "TEST PANEL") {
		t.Errorf("expected panel to contain header")
	}
	if !strings.Contains(panel, "panel body content") {
		t.Errorf("expected panel to contain body")
	}
}

func TestRenderDualPane(t *testing.T) {
	left := "Left Content"
	right := "Right Content"
	dual := RenderDualPane(left, right, 40, 30, 15)

	if !strings.Contains(dual, "Left Content") || !strings.Contains(dual, "Right Content") {
		t.Errorf("expected dual pane to contain both panes")
	}
}
