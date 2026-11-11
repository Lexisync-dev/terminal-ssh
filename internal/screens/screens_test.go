package screens

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lexisync/jawad-os/internal/config"
)

func getTestConfig(t *testing.T) *config.PortfolioConfig {
	cfg, err := config.Load("")
	if err != nil {
		t.Fatalf("failed to load embedded config: %v", err)
	}
	return cfg
}

func TestAboutModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewAboutModel(cfg)
	m.SetSize(80, 80)

	view := m.View()
	if !strings.Contains(view, "Jawad Ahmad") {
		t.Errorf("expected About view to contain 'Jawad Ahmad'")
	}
	if !strings.Contains(view, "Pakistan") {
		t.Errorf("expected About view to contain 'Pakistan'")
	}
	if !strings.Contains(view, "Computer Science student") {
		t.Errorf("expected About view to contain 'Computer Science student'")
	}

	// Test scroll keys
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	_ = updated.View()
}

func TestProjectsModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewProjectsModel(cfg)
	m.SetSize(80, 80)

	view := m.View()
	if !strings.Contains(view, "Workflow Automation Platform") {
		t.Errorf("expected Projects view to contain 'Workflow Automation Platform'")
	}
	if !strings.Contains(view, "In Development") {
		t.Errorf("expected Projects view to contain status 'In Development'")
	}
	if !strings.Contains(view, "Redis") || !strings.Contains(view, "PostgreSQL") {
		t.Errorf("expected Projects view to contain technology stack")
	}

	// Test scroll and navigation keys
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	_ = updated.View()
}

func TestSkillsModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewSkillsModel(cfg)
	m.SetSize(80, 80)

	view := m.View()
	if !strings.Contains(strings.ToUpper(view), "PROGRAMMING & BACKEND") {
		t.Errorf("expected Skills view to contain 'PROGRAMMING & BACKEND'")
	}
	if !strings.Contains(view, "Go") || !strings.Contains(view, "TypeScript") {
		t.Errorf("expected Skills view to contain technologies")
	}
	if !strings.Contains(view, "No arbitrary percentages") {
		t.Errorf("expected Skills view to mention philosophy of no arbitrary percentages")
	}
}

func TestExperienceModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewExperienceModel(cfg)
	m.SetSize(80, 80)

	view := m.View()
	if !strings.Contains(view, "Freelance") {
		t.Errorf("expected Experience view to contain 'Freelance'")
	}
	if !strings.Contains(view, "STATUS:") {
		t.Errorf("expected Experience view to contain status banner")
	}
}

func TestInterestsModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewInterestsModel(cfg)
	m.SetSize(80, 80)

	view := m.View()
	upperView := strings.ToUpper(view)
	if !strings.Contains(upperView, "POETRY") {
		t.Errorf("expected Interests view to contain 'Poetry'")
	}
	if !strings.Contains(upperView, "SUPERBIKES") {
		t.Errorf("expected Interests view to contain 'Superbikes'")
	}
	if !strings.Contains(upperView, "MECHANICAL ENGINEERING") {
		t.Errorf("expected Interests view to contain 'Mechanical Engineering'")
	}
}

func TestContactModel(t *testing.T) {
	cfg := getTestConfig(t)
	m := NewContactModel(cfg)
	m.SetSize(80, 24)

	view := m.View()
	if !strings.Contains(view, "GitHub") {
		t.Errorf("expected Contact view to contain GitHub")
	}

	// Test selecting and copying with Enter
	updated, cmd, notification := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Errorf("expected copy command to be returned on Enter")
	}
	if !strings.Contains(notification, "Copied") {
		t.Errorf("expected notification string to be returned on Enter, got '%s'", notification)
	}

	// Test down navigation
	updated, _, _ = updated.Update(tea.KeyMsg{Type: tea.KeyDown})
	_ = updated.View()
}
