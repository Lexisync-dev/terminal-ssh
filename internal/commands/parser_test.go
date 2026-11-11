package commands

import (
	"strings"
	"testing"

	"github.com/lexisync/jawad-os/internal/ui"
)

func TestParseNavigation(t *testing.T) {
	cases := []struct {
		input    string
		expected ui.Screen
	}{
		{"about", ui.ScreenAbout},
		{"profile", ui.ScreenAbout},
		{"projects", ui.ScreenProjects},
		{"skills", ui.ScreenSkills},
		{"stack", ui.ScreenSkills},
		{"experience", ui.ScreenExperience},
		{"journey", ui.ScreenExperience},
		{"interests", ui.ScreenInterests},
		{"contact", ui.ScreenContact},
		{"help", ui.ScreenHelp},
		{"home", ui.ScreenDashboard},
	}

	for _, c := range cases {
		cmd := Parse(c.input, "testuser", 80)
		if cmd.Type != CmdNavigate {
			t.Errorf("expected CmdNavigate for input '%s', got %v", c.input, cmd.Type)
		}
		if cmd.TargetScreen != c.expected {
			t.Errorf("expected target screen %v for input '%s', got %v", c.expected, c.input, cmd.TargetScreen)
		}
	}
}

func TestParseSubcommands(t *testing.T) {
	cmd := Parse("projects workflow", "testuser", 80)
	if cmd.Type != CmdNavigate || cmd.TargetScreen != ui.ScreenProjects {
		t.Fatalf("expected navigate to projects")
	}
	if cmd.Subcommand != "workflow" {
		t.Errorf("expected subcommand 'workflow', got '%s'", cmd.Subcommand)
	}
}

func TestParseEasterEggs(t *testing.T) {
	// Neofetch
	neo := Parse("neofetch", "recruiter", 80)
	if neo.Type != CmdModal {
		t.Errorf("expected CmdModal for neofetch, got %v", neo.Type)
	}
	if !strings.Contains(neo.Output, "JAWAD.OS") {
		t.Errorf("expected neofetch to contain 'JAWAD.OS'")
	}

	// Sudo hire jawad
	sudo := Parse("sudo hire jawad", "recruiter", 80)
	if sudo.Type != CmdModal {
		t.Errorf("expected CmdModal for sudo hire jawad, got %v", sudo.Type)
	}
	if !strings.Contains(sudo.Output, "HIRING PROTOCOL INITIATED") {
		t.Errorf("expected sudo hire output to contain hiring protocol headline")
	}

	// Secret developer message
	sec := Parse("secret", "recruiter", 80)
	if sec.Type != CmdModal {
		t.Errorf("expected CmdModal for secret, got %v", sec.Type)
	}
	if !strings.Contains(sec.Output, "Jawad Ahmad") {
		t.Errorf("expected secret note to contain author name")
	}

	// Whoami
	who := Parse("whoami", "alex", 80)
	if who.Type != CmdOutput {
		t.Errorf("expected CmdOutput for whoami, got %v", who.Type)
	}
	if !strings.Contains(who.StatusMessage, "alex") {
		t.Errorf("expected whoami message to contain user 'alex'")
	}

	// Date
	dt := Parse("date", "alex", 80)
	if dt.Type != CmdOutput {
		t.Errorf("expected CmdOutput for date, got %v", dt.Type)
	}
	if !strings.Contains(dt.StatusMessage, "System Time:") {
		t.Errorf("expected date message to contain 'System Time:'")
	}
}

func TestParseTypoSuggestion(t *testing.T) {
	res := Parse("projcts", "user", 80)
	if res.Type != CmdUnknown {
		t.Errorf("expected CmdUnknown for typo, got %v", res.Type)
	}
	if !strings.Contains(res.StatusMessage, "Did you mean 'projects'?") {
		t.Errorf("expected suggestion 'projects', got '%s'", res.StatusMessage)
	}
}

func TestParseExit(t *testing.T) {
	for _, exitCmd := range []string{"exit", "quit", "q"} {
		res := Parse(exitCmd, "user", 80)
		if res.Type != CmdQuit {
			t.Errorf("expected CmdQuit for '%s', got %v", exitCmd, res.Type)
		}
	}
}
