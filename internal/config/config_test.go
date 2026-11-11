package config

import (
	"testing"
)

func TestLoadEmbeddedDefault(t *testing.T) {
	cfg, err := Load("")
	if err != nil {
		t.Fatalf("expected nil error loading embedded config, got: %v", err)
	}

	if cfg.Profile.Name != "Jawad Ahmad" {
		t.Errorf("expected Profile.Name to be 'Jawad Ahmad', got '%s'", cfg.Profile.Name)
	}

	if len(cfg.Skills.Categories) == 0 {
		t.Errorf("expected skill categories to be populated")
	}

	if len(cfg.Projects) == 0 {
		t.Errorf("expected projects to be populated")
	}

	if cfg.Projects[0].Name != "Workflow Automation Platform" {
		t.Errorf("expected first project to be Workflow Automation Platform, got '%s'", cfg.Projects[0].Name)
	}
}

func TestLoadFromFile(t *testing.T) {
	cfg, err := Load("../../configs/portfolio.yaml")
	if err != nil {
		t.Fatalf("expected nil error loading from file, got: %v", err)
	}

	if cfg.Profile.Name != "Jawad Ahmad" {
		t.Errorf("expected Profile.Name to be 'Jawad Ahmad', got '%s'", cfg.Profile.Name)
	}
}
