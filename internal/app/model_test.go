package app

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/screens"
	"github.com/lexisync/jawad-os/internal/ui"
)

func newTestConfig() *config.PortfolioConfig {
	cfg, _ := config.Load("")
	return cfg
}

func TestModelNavigation(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg, WithInitialScreen(ui.ScreenDashboard))

	if m.CurrentScreen != ui.ScreenDashboard {
		t.Fatalf("expected initial screen to be Dashboard, got %v", m.CurrentScreen)
	}

	// Navigate to About
	m.NavigateTo(ui.ScreenAbout)
	if m.CurrentScreen != ui.ScreenAbout {
		t.Fatalf("expected current screen to be About, got %v", m.CurrentScreen)
	}
	if len(m.History) != 1 || m.History[0] != ui.ScreenDashboard {
		t.Fatalf("expected history to contain [Dashboard], got %v", m.History)
	}

	// Navigate to Projects
	m.NavigateTo(ui.ScreenProjects)
	if m.CurrentScreen != ui.ScreenProjects {
		t.Fatalf("expected current screen to be Projects, got %v", m.CurrentScreen)
	}
	if len(m.History) != 2 {
		t.Fatalf("expected history length 2, got %d", len(m.History))
	}

	// Navigate Back
	m.NavigateBack()
	if m.CurrentScreen != ui.ScreenAbout {
		t.Fatalf("expected back navigation to return to About, got %v", m.CurrentScreen)
	}

	// Navigate Back again to Dashboard
	m.NavigateBack()
	if m.CurrentScreen != ui.ScreenDashboard {
		t.Fatalf("expected back navigation to return to Dashboard, got %v", m.CurrentScreen)
	}

	// Back when at root should stay on Dashboard
	m.NavigateBack()
	if m.CurrentScreen != ui.ScreenDashboard {
		t.Fatalf("expected staying on Dashboard when history is empty, got %v", m.CurrentScreen)
	}
}

func TestModelCommandMode(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg, WithInitialScreen(ui.ScreenDashboard))

	// Press ":" to enter command mode
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{':'}})
	model := updated.(Model)
	if !model.CommandMode {
		t.Errorf("expected CommandMode to be active after pressing ':'")
	}

	// Type "about"
	for _, ch := range "about" {
		updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{ch}})
		model = updated.(Model)
	}

	if model.CommandInput != "about" {
		t.Errorf("expected CommandInput to be 'about', got '%s'", model.CommandInput)
	}

	// Press Enter
	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model = updated.(Model)

	if model.CommandMode {
		t.Errorf("expected CommandMode to be inactive after Enter")
	}
	if model.CurrentScreen != ui.ScreenAbout {
		t.Errorf("expected command 'about' to navigate to ScreenAbout, got %v", model.CurrentScreen)
	}
}

func TestModelBootTransition(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg)

	if m.CurrentScreen != ui.ScreenBoot {
		t.Fatalf("expected default initial screen to be Boot, got %v", m.CurrentScreen)
	}

	// Pressing Enter during boot skips directly to Dashboard
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model := updated.(Model)

	if model.CurrentScreen != ui.ScreenDashboard {
		t.Fatalf("expected pressing enter to transition to Dashboard, got %v", model.CurrentScreen)
	}
}

func TestBootTicks(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg)

	// Step through boot initialization ticks
	for i := 0; i < 6; i++ {
		updated, _ := m.Update(screens.BootTickMsg{})
		m = updated.(Model)
	}

	if !m.BootScreen.IsCompleted() {
		t.Fatalf("expected boot sequence steps to be completed")
	}

	// Pressing Enter when boot steps finish transitions to Dashboard
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(Model)
	if m.CurrentScreen != ui.ScreenDashboard {
		t.Fatalf("expected pressing enter after boot completion to transition to Dashboard, got %v", m.CurrentScreen)
	}
}

func TestScreenViews(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg, WithInitialScreen(ui.ScreenDashboard))
	m.SetSize(80, 60)

	// 1. Dashboard View
	view := m.View()
	if !strings.Contains(view, "JAWAD AHMAD") {
		t.Errorf("expected Dashboard view to contain 'JAWAD AHMAD'")
	}

	// 2. About View
	m.NavigateTo(ui.ScreenAbout)
	view = m.View()
	if !strings.Contains(view, "PROFILE") || !strings.Contains(view, "Jawad Ahmad") {
		t.Errorf("expected About view to contain 'PROFILE' and 'Jawad Ahmad'")
	}

	// 3. Projects View
	m.NavigateTo(ui.ScreenProjects)
	view = m.View()
	if !strings.Contains(view, "PROJECTS") || !strings.Contains(view, "Workflow Automation") {
		t.Errorf("expected Projects view to contain 'PROJECTS' and 'Workflow Automation'")
	}

	// 4. Skills View
	m.NavigateTo(ui.ScreenSkills)
	view = m.View()
	if !strings.Contains(view, "STACK") || !strings.Contains(strings.ToUpper(view), "PROGRAMMING & BACKEND") {
		t.Errorf("expected Skills view to contain 'STACK' and 'PROGRAMMING & BACKEND'")
	}

	// 5. Experience View
	m.NavigateTo(ui.ScreenExperience)
	view = m.View()
	if !strings.Contains(view, "JOURNEY") || !strings.Contains(view, "Freelance") {
		t.Errorf("expected Experience view to contain 'JOURNEY' and 'Freelance'")
	}

	// 6. Interests View
	m.NavigateTo(ui.ScreenInterests)
	view = m.View()
	if !strings.Contains(view, "INTERESTS") || !strings.Contains(strings.ToUpper(view), "POETRY") {
		t.Errorf("expected Interests view to contain 'INTERESTS' and 'POETRY'")
	}

	// 7. Contact View
	m.NavigateTo(ui.ScreenContact)
	view = m.View()
	if !strings.Contains(view, "CONTACT") || !strings.Contains(view, "GitHub") {
		t.Errorf("expected Contact view to contain 'CONTACT' and 'GitHub'")
	}

	// 8. Help View
	m.NavigateTo(ui.ScreenHelp)
	view = m.View()
	if !strings.Contains(view, "HELP") || !strings.Contains(view, "KEYBOARD & MOUSE SHORTCUTS") {
		t.Errorf("expected Help view to contain 'HELP'")
	}
}

func TestEasterEggModal(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg, WithInitialScreen(ui.ScreenDashboard))
	m.SetSize(80, 24)

	// Trigger neofetch via command mode
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{':'}})
	model := updated.(Model)
	for _, ch := range "neofetch" {
		updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{ch}})
		model = updated.(Model)
	}
	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model = updated.(Model)

	if !model.ModalActive {
		t.Fatalf("expected ModalActive to be true after running :neofetch")
	}

	view := model.View()
	if !strings.Contains(view, "JAWAD.OS SYSTEM TELEMETRY") {
		t.Errorf("expected modal view to display system telemetry")
	}

	// Any key dismisses the modal
	updated, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	model = updated.(Model)
	if model.ModalActive {
		t.Errorf("expected ModalActive to be false after pressing ESC")
	}
}

func TestMouseScrolling(t *testing.T) {
	cfg := newTestConfig()
	m := NewModel(cfg, WithInitialScreen(ui.ScreenAbout))
	m.SetSize(80, 24)

	// Simulate mouse wheel down
	mouseMsg := tea.MouseMsg{
		Button: tea.MouseButtonWheelDown,
		Action: tea.MouseActionPress,
	}
	updated, _ := m.Update(mouseMsg)
	_ = updated.View()
}
