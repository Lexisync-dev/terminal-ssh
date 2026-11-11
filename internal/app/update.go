package app

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lexisync/jawad-os/internal/commands"
	"github.com/lexisync/jawad-os/internal/screens"
	"github.com/lexisync/jawad-os/internal/ui"
)

type ClearStatusMsg struct{}

func clearStatusAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return ClearStatusMsg{}
	})
}

// Init initializes the Bubble Tea program.
func (m Model) Init() tea.Cmd {
	if m.CurrentScreen == ui.ScreenBoot {
		return m.BootScreen.Init()
	}
	return nil
}

// Update handles incoming messages and dispatches state updates.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil

	case ClearStatusMsg:
		m.StatusMessage = ""
		return m, nil

	case screens.CopyNotificationMsg:
		m.StatusMessage = "Copied to clipboard: " + msg.Value
		return m, clearStatusAfter(3 * time.Second)

	case screens.BootTickMsg:
		if m.CurrentScreen == ui.ScreenBoot {
			var cmd tea.Cmd
			var done bool
			m.BootScreen, cmd, done = m.BootScreen.Update(msg)
			if done {
				m.CurrentScreen = ui.ScreenDashboard
				m.SetSize(m.Width, m.Height)
			}
			return m, cmd
		}

	case tea.MouseMsg:
		// Support terminal mouse wheel scrolling across viewports
		if msg.Action == tea.MouseActionPress {
			if msg.Button == tea.MouseButtonWheelUp {
				return m.Update(tea.KeyMsg{Type: tea.KeyUp})
			} else if msg.Button == tea.MouseButtonWheelDown {
				return m.Update(tea.KeyMsg{Type: tea.KeyDown})
			}
		}

	case tea.KeyMsg:
		// Global exit on Ctrl+C
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		// Dismiss modal overlays (e.g., neofetch, easter eggs)
		if m.ModalActive {
			m.ModalActive = false
			m.ModalContent = ""
			m.ModalTitle = ""
			return m, nil
		}

		// Handle active command mode input
		if m.CommandMode {
			return m.handleCommandInput(msg)
		}

		// Global keybindings (when NOT in command mode)
		switch msg.String() {
		case ":":
			if m.CurrentScreen != ui.ScreenBoot {
				m.CommandMode = true
				m.CommandInput = ""
				return m, nil
			}

		case "?", "f1":
			if m.CurrentScreen != ui.ScreenBoot {
				if m.CurrentScreen == ui.ScreenHelp {
					m.NavigateBack()
				} else {
					m.NavigateTo(ui.ScreenHelp)
				}
				return m, nil
			}

		case "esc":
			if m.CurrentScreen != ui.ScreenDashboard && m.CurrentScreen != ui.ScreenBoot {
				m.NavigateBack()
				return m, nil
			}

		case "q":
			if m.CurrentScreen != ui.ScreenBoot {
				return m, tea.Quit
			}
		}
	}

	// Dispatch to active screen update logic
	switch m.CurrentScreen {
	case ui.ScreenBoot:
		var cmd tea.Cmd
		var done bool
		m.BootScreen, cmd, done = m.BootScreen.Update(msg)
		if done {
			m.CurrentScreen = ui.ScreenDashboard
			m.SetSize(m.Width, m.Height)
		}
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenDashboard:
		var cmd tea.Cmd
		var targetScreen *ui.Screen
		m.DashboardScreen, cmd, targetScreen = m.DashboardScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		if targetScreen != nil {
			m.NavigateTo(*targetScreen)
		}

	case ui.ScreenAbout:
		var cmd tea.Cmd
		m.AboutScreen, cmd = m.AboutScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenProjects:
		var cmd tea.Cmd
		m.ProjectsScreen, cmd = m.ProjectsScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenSkills:
		var cmd tea.Cmd
		m.SkillsScreen, cmd = m.SkillsScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenExperience:
		var cmd tea.Cmd
		m.ExperienceScreen, cmd = m.ExperienceScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenInterests:
		var cmd tea.Cmd
		m.InterestsScreen, cmd = m.InterestsScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

	case ui.ScreenContact:
		var cmd tea.Cmd
		var notification string
		m.ContactScreen, cmd, notification = m.ContactScreen.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		if notification != "" {
			m.StatusMessage = notification
			cmds = append(cmds, clearStatusAfter(3*time.Second))
		}
	}

	return m, tea.Batch(cmds...)
}

// handleCommandInput processes keystrokes while the command prompt is active.
func (m Model) handleCommandInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEsc:
		m.CommandMode = false
		m.CommandInput = ""
		return m, nil

	case tea.KeyEnter:
		m.CommandMode = false
		cmdText := strings.TrimSpace(m.CommandInput)
		m.CommandInput = ""
		return m.executeCommand(cmdText)

	case tea.KeyBackspace:
		if len(m.CommandInput) > 0 {
			m.CommandInput = m.CommandInput[:len(m.CommandInput)-1]
		}
		return m, nil

	case tea.KeyRunes, tea.KeySpace:
		m.CommandInput += msg.String()
		return m, nil
	}

	return m, nil
}

// executeCommand runs commands typed into the command mode bar.
func (m Model) executeCommand(cmd string) (tea.Model, tea.Cmd) {
	dim := ui.CalculateDimensions(m.Width, m.Height)
	parsed := commands.Parse(cmd, m.User, dim.InnerWidth)

	switch parsed.Type {
	case commands.CmdNone:
		return m, nil

	case commands.CmdQuit:
		return m, tea.Quit

	case commands.CmdClear:
		m.StatusMessage = ""
		return m, nil

	case commands.CmdNavigate:
		m.NavigateTo(parsed.TargetScreen)
		return m, nil

	case commands.CmdOutput:
		m.StatusMessage = parsed.StatusMessage
		return m, clearStatusAfter(4 * time.Second)

	case commands.CmdModal:
		m.ModalActive = true
		m.ModalTitle = parsed.Title
		m.ModalContent = parsed.Output
		return m, nil

	case commands.CmdUnknown:
		m.StatusMessage = parsed.StatusMessage
		return m, clearStatusAfter(4 * time.Second)
	}

	return m, nil
}
