package app

import (
	"testing"

	"github.com/lexisync/jawad-os/internal/ui"
)

// TestTerminalDimensions verifies that the workstation interface renders cleanly without
// panicking across all required representative terminal sizes: 80x24, 100x30, 120x40, 160x50.
func TestTerminalDimensions(t *testing.T) {
	cfg := newTestConfig()

	dimensions := []struct {
		name   string
		width  int
		height int
	}{
		{"Compact (80x24)", 80, 24},
		{"Medium Workstation (100x30)", 100, 30},
		{"Wide Workstation (120x40)", 120, 40},
		{"Ultra-Wide Workstation (160x50)", 160, 50},
		{"Narrow Window (65x25)", 65, 25},
		{"Constrained Window (40x12)", 40, 12},
	}

	screensToTest := []ui.Screen{
		ui.ScreenDashboard,
		ui.ScreenAbout,
		ui.ScreenProjects,
		ui.ScreenSkills,
		ui.ScreenExperience,
		ui.ScreenInterests,
		ui.ScreenContact,
		ui.ScreenHelp,
	}

	for _, dim := range dimensions {
		t.Run(dim.name, func(t *testing.T) {
			m := NewModel(cfg)
			m.SetSize(dim.width, dim.height)

			for _, sc := range screensToTest {
				m.CurrentScreen = sc
				m.SetSize(dim.width, dim.height)

				view := m.View()
				if view == "" {
					t.Errorf("screen %v rendered empty view at %dx%d", sc, dim.width, dim.height)
				}
			}
		})
	}
}
