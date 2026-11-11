package app

import (
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/screens"
	"github.com/lexisync/jawad-os/internal/ui"
)

// Model represents the top-level Bubble Tea state for JAWAD.OS.
type Model struct {
	Config        *config.PortfolioConfig
	CurrentScreen ui.Screen
	History       []ui.Screen
	Width         int
	Height        int
	Ready         bool

	// Subscreen models
	BootScreen       screens.BootModel
	DashboardScreen  screens.DashboardModel
	AboutScreen      screens.AboutModel
	ProjectsScreen   screens.ProjectsModel
	SkillsScreen     screens.SkillsModel
	ExperienceScreen screens.ExperienceModel
	InterestsScreen  screens.InterestsModel
	ContactScreen    screens.ContactModel

	// Command input state
	CommandMode   bool
	CommandInput  string
	StatusMessage string

	// Modal overlay state (for easter eggs & alerts)
	ModalActive  bool
	ModalTitle   string
	ModalContent string

	// Session metadata
	User string
}

// Option represents a functional configuration option for the Model.
type Option func(*Model)

// WithUser sets the session username (e.g., from SSH).
func WithUser(user string) Option {
	return func(m *Model) {
		m.User = user
	}
}

// WithInitialScreen sets the starting screen (e.g., skipping boot).
func WithInitialScreen(screen ui.Screen) Option {
	return func(m *Model) {
		m.CurrentScreen = screen
	}
}

// NewModel constructs a new application Model.
func NewModel(cfg *config.PortfolioConfig, opts ...Option) Model {
	m := Model{
		Config:           cfg,
		CurrentScreen:    ui.ScreenBoot,
		History:          make([]ui.Screen, 0),
		Width:            80,
		Height:           24,
		BootScreen:       screens.NewBootModel(cfg),
		DashboardScreen:  screens.NewDashboardModel(cfg),
		AboutScreen:      screens.NewAboutModel(cfg),
		ProjectsScreen:   screens.NewProjectsModel(cfg),
		SkillsScreen:     screens.NewSkillsModel(cfg),
		ExperienceScreen: screens.NewExperienceModel(cfg),
		InterestsScreen:  screens.NewInterestsModel(cfg),
		ContactScreen:    screens.NewContactModel(cfg),
		User:             "visitor",
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

// SetSize updates layout dimensions across all sub-screens.
func (m *Model) SetSize(w, h int) {
	m.Width = w
	m.Height = h
	m.Ready = true
	m.BootScreen.SetSize(w, h)
	m.DashboardScreen.SetSize(w, h)
	m.AboutScreen.SetSize(w, h)
	m.ProjectsScreen.SetSize(w, h)
	m.SkillsScreen.SetSize(w, h)
	m.ExperienceScreen.SetSize(w, h)
	m.InterestsScreen.SetSize(w, h)
	m.ContactScreen.SetSize(w, h)
}

// NavigateTo pushes the current screen to history and transitions to the target screen.
func (m *Model) NavigateTo(screen ui.Screen) {
	if m.CurrentScreen != screen {
		m.History = append(m.History, m.CurrentScreen)
		m.CurrentScreen = screen
		// Ensure active screen dimensions are refreshed
		m.SetSize(m.Width, m.Height)
	}
}

// NavigateBack pops the last screen from history or returns to Dashboard.
func (m *Model) NavigateBack() {
	if len(m.History) > 0 {
		prev := m.History[len(m.History)-1]
		m.History = m.History[:len(m.History)-1]
		m.CurrentScreen = prev
		m.SetSize(m.Width, m.Height)
	} else if m.CurrentScreen != ui.ScreenDashboard {
		m.CurrentScreen = ui.ScreenDashboard
		m.SetSize(m.Width, m.Height)
	}
}
