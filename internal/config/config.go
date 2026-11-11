package config

import (
	_ "embed"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed default.yaml
var DefaultConfigYAML []byte

type Profile struct {
	Name        string `yaml:"name"`
	Title       string `yaml:"title"`
	Positioning string `yaml:"positioning"`
	Location    string `yaml:"location"`
	Education   string `yaml:"education"`
	Tagline     string `yaml:"tagline"`
	Bio         string `yaml:"bio"`
}

type DirectionArea struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type Direction struct {
	Summary string          `yaml:"summary"`
	Areas   []DirectionArea `yaml:"areas"`
}

type SkillCategory struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Items       []string `yaml:"items"`
}

type SkillsConfig struct {
	Categories []SkillCategory `yaml:"categories"`
}

type Project struct {
	ID                 string   `yaml:"id"`
	Name               string   `yaml:"name"`
	Tagline            string   `yaml:"tagline"`
	Status             string   `yaml:"status"`
	Description        string   `yaml:"description"`
	ConceptualExample  string   `yaml:"conceptual_example"`
	Stack              []string `yaml:"stack"`
	ArchitectureAreas  []string `yaml:"architecture_areas"`
	Decisions          []string `yaml:"decisions"`
	Challenges         []string `yaml:"challenges"`
	Learnings          []string `yaml:"learnings"`
	GitHubURL          string   `yaml:"github_url"`
	LiveURL            string   `yaml:"live_url"`
}

type ExperienceHighlight struct {
	Title       string `yaml:"title"`
	Focus       string `yaml:"focus"`
	Description string `yaml:"description"`
}

type ExperienceConfig struct {
	Summary    string                `yaml:"summary"`
	Status     string                `yaml:"status"`
	Highlights []ExperienceHighlight `yaml:"highlights"`
}

type InterestItem struct {
	Name        string `yaml:"name"`
	Icon        string `yaml:"icon"`
	Tagline     string `yaml:"tagline"`
	Description string `yaml:"description"`
}

type InterestsConfig struct {
	Title    string         `yaml:"title"`
	Subtitle string         `yaml:"subtitle"`
	Items    []InterestItem `yaml:"items"`
}

type ContactConfig struct {
	Prompt   string `yaml:"prompt"`
	GitHub   string `yaml:"github"`
	LinkedIn string `yaml:"linkedin"`
	Email    string `yaml:"email"`
	Website  string `yaml:"website"`
	Location string `yaml:"location"`
}

type PortfolioConfig struct {
	Profile    Profile          `yaml:"profile"`
	Direction  Direction        `yaml:"direction"`
	Skills     SkillsConfig     `yaml:"skills"`
	Projects   []Project        `yaml:"projects"`
	Experience ExperienceConfig `yaml:"experience"`
	Interests  InterestsConfig  `yaml:"interests"`
	Contact    ContactConfig    `yaml:"contact"`
}

// Load reads the configuration from the specified file path.
// If the path is empty or file doesn't exist, it falls back to the embedded default configuration.
func Load(path string) (*PortfolioConfig, error) {
	var data []byte

	if path != "" {
		if content, err := os.ReadFile(path); err == nil {
			data = content
		} else if !os.IsNotExist(err) {
			return nil, fmt.Errorf("reading config file %s: %w", path, err)
		}
	}

	if len(data) == 0 {
		// Fallback to embedded default YAML
		data = DefaultConfigYAML
	}

	var cfg PortfolioConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing portfolio yaml: %w", err)
	}

	return &cfg, nil
}
