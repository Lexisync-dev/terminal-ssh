package ui

import (
	"strings"
	"testing"
)

func TestCalculateDimensions(t *testing.T) {
	// Standard full terminal
	dim := CalculateDimensions(100, 30)
	if dim.IsTooSmall {
		t.Errorf("expected 100x30 to not be too small")
	}
	if dim.InnerWidth > 96 {
		t.Errorf("expected inner width to be capped at 96, got %d", dim.InnerWidth)
	}

	// Small terminal
	dimSmall := CalculateDimensions(40, 10)
	if !dimSmall.IsTooSmall {
		t.Errorf("expected 40x10 to be flagged as too small")
	}

	// Fallback when 0x0
	dimZero := CalculateDimensions(0, 0)
	if dimZero.TotalWidth != 80 || dimZero.TotalHeight != 24 {
		t.Errorf("expected default 80x24 fallback, got %dx%d", dimZero.TotalWidth, dimZero.TotalHeight)
	}
}

func TestRenderHeader(t *testing.T) {
	hdr := RenderHeader(80, ScreenDashboard)
	if !strings.Contains(hdr, "JAWAD.OS") {
		t.Errorf("expected header to contain 'JAWAD.OS'")
	}
	if !strings.Contains(hdr, "ONLINE") {
		t.Errorf("expected header to contain 'ONLINE'")
	}

	// Non-dashboard screen should show breadcrumb
	hdrAbout := RenderHeader(80, ScreenAbout)
	if !strings.Contains(hdrAbout, "PROFILE") {
		t.Errorf("expected header to contain breadcrumb 'PROFILE'")
	}
}

func TestRenderFooter(t *testing.T) {
	ftr := RenderFooter(80, false, "", "", ScreenDashboard)
	if !strings.Contains(ftr, "Navigate") || !strings.Contains(ftr, "Select") {
		t.Errorf("expected dashboard footer to include basic navigation hints")
	}

	// Command mode footer
	ftrCmd := RenderFooter(80, true, "about", "", ScreenDashboard)
	if !strings.Contains(ftrCmd, ":") || !strings.Contains(ftrCmd, "about") {
		t.Errorf("expected command mode footer to display prompt and input")
	}
}
