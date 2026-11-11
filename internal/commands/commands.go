package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/ui"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

// CommandType defines the classification of a parsed command.
type CommandType int

const (
	CmdNone CommandType = iota
	CmdNavigate
	CmdQuit
	CmdClear
	CmdHelp
	CmdOutput
	CmdModal
	CmdUnknown
)

// ParsedCommand represents the outcome of parsing a visitor's command input.
type ParsedCommand struct {
	Type          CommandType
	TargetScreen  ui.Screen
	Subcommand    string
	Title         string
	Output        string
	StatusMessage string
}

// AvailableCommands lists all recognized commands for help and suggestions.
var AvailableCommands = []struct {
	Name        string
	Args        string
	Description string
}{
	{"about", "", "Jump to Profile dossier (bio, philosophy & direction)"},
	{"profile", "", "Alias for 'about' (Screen 01)"},
	{"projects", "[list|workflow]", "Open projects explorer or view a project"},
	{"skills", "", "Open categorized technology stack matrix"},
	{"stack", "", "Alias for 'skills' (Screen 03)"},
	{"experience", "", "Open engineering journey and milestones"},
	{"journey", "", "Alias for 'experience' (Screen 04)"},
	{"interests", "", "Open 'Beyond the Terminal' personal interests"},
	{"contact", "", "Open contact console & clipboard copy helper"},
	{"home", "", "Return to the main workstation dashboard"},
	{"help", "", "Display help and command reference"},
	{"clear", "", "Clear the current status bar message"},
	{"whoami", "", "Print active session identity and permissions"},
	{"date", "", "Display current system date and time"},
	{"neofetch", "", "Display JAWAD.OS system diagnostics & telemetry"},
	{"sudo hire jawad", "", "Execute high-priority hiring protocol"},
	{"secret", "", "Read a personal message from the developer"},
	{"exit", "", "Disconnect or quit JAWAD.OS (alias: quit, q)"},
}

// NeofetchOutput generates a stylized system info display.
func NeofetchOutput(user string, width int) string {
	asciiLogo := `    _                       
  _| | __ ___      ____ _   
 / _` + "`" + ` |/ _` + "`" + ` \ \ /\ / / _` + "`" + ` |  
| (_| | (_| |\ V  V / (_| |  
 \__,_|\__,_| \_/\_/ \__,_|  `

	styledLogo := lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.ColorAccentAmber).
		Render(asciiLogo)

	sep := theme.RenderDivider(36)

	specs := []string{
		theme.StyleBrand.Render("jawad@JAWAD.OS // WORKSTATION"),
		sep,
		theme.RenderKeyValue("OS:", "JAWAD.OS v1.0 (Terminal Workstation)", 14),
		theme.RenderKeyValue("RUNTIME:", "Golang 1.27+ (Charm Ecosystem)", 14),
		theme.RenderKeyValue("KERNEL:", "Bubble Tea + Lip Gloss + Wish SSH", 14),
		theme.RenderKeyValue("ARCH:", "Unidirectional Model-View-Update (Elm)", 14),
		theme.RenderKeyValue("SHELL:", "jawad-cmd / wish-ssh", 14),
		theme.RenderKeyValue("USER:", user, 14),
		theme.RenderKeyValue("TERMINAL:", "xterm-256color // OSC 52 Enabled", 14),
		theme.RenderKeyValue("MEMORY:", "Lightweight Goroutine Concurrency", 14),
		theme.RenderKeyValue("STATUS:", theme.StyleStatusOnline.Render("ONLINE "+theme.GlyphDot+" Exploring Systems"), 14),
	}

	specBlock := strings.Join(specs, "\n")

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		styledLogo,
		"    ",
		specBlock,
	)

	return theme.RenderWorkstationPanel("JAWAD.OS SYSTEM TELEMETRY", content, width, 0, true)
}

// SudoHireOutput generates an entertaining recruiter easter egg.
func SudoHireOutput(width int) string {
	prompt := lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.ColorAccentAmber).
		Render("[sudo] authenticating engineering recruiter: **********")

	granted := theme.StyleStatusOnline.Render("Authentication successful. Clearance Level 1 Granted.")

	headline := theme.StyleIdentity.Render("HIRING PROTOCOL INITIATED")

	body := theme.StyleBody.Width(width - 6).Render(
		"Thank you for considering me for your engineering team!\n\n" +
			"I am an aspiring Full-Stack Software Engineer who loves building real-world software, " +
			"exploring distributed backend architectures, and understanding systems beneath the surface.\n\n" +
			"Let's build something impactful together. Type ':contact' to reach out directly.",
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		prompt,
		"",
		granted,
		"",
		headline,
		"",
		body,
	)

	return theme.RenderWorkstationPanel("RECRUITER CLEARANCE // CONFIDENTIAL", content, width, 0, true)
}

// SecretDeveloperMessage generates a personal note from Jawad.
func SecretDeveloperMessage(width int) string {
	title := theme.StyleIdentity.Render("A NOTE FROM JAWAD AHMAD")

	body := theme.StyleBody.Width(width - 6).Render(
		"Hey there! Thanks for taking the time to explore JAWAD.OS.\n\n" +
			"I built this portfolio using Go and Charm's Bubble Tea, Lip Gloss, and Wish libraries because I wanted my portfolio " +
			"to reflect what I actually care about: terminal interfaces, backend foundations, clean code, and building things from first principles.\n\n" +
			"Whether you're a recruiter, an engineering manager, or a fellow builder, I'd love to hear your thoughts, feedback, or just talk about software.\n\n" +
			"— Jawad Ahmad",
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		body,
	)

	return theme.RenderWorkstationPanel("BUILDER'S LOG // DEVELOPER NOTE", content, width, 0, true)
}

// DateOutput returns formatted local timestamp.
func DateOutput() string {
	now := time.Now()
	return fmt.Sprintf("System Time: %s (%s)", now.Format("Monday, 02 Jan 2006 15:04:05 MST"), now.Format(time.RFC3339))
}
