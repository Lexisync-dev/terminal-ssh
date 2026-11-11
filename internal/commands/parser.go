package commands

import (
	"fmt"
	"strings"

	"github.com/lexisync/jawad-os/internal/ui"
)

// Parse converts a user raw string into a structured ParsedCommand.
func Parse(input string, user string, width int) ParsedCommand {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return ParsedCommand{Type: CmdNone}
	}

	// Normalize input
	lowered := strings.ToLower(trimmed)

	// Check multi-word commands first
	if lowered == "sudo hire jawad" || lowered == "sudo hire" || lowered == "hire jawad" {
		return ParsedCommand{
			Type:   CmdModal,
			Title:  "RECRUITER ACCESS",
			Output: SudoHireOutput(width),
		}
	}

	if lowered == "cat secret.txt" || lowered == "cat developer.txt" || lowered == "cat secret" {
		return ParsedCommand{
			Type:   CmdModal,
			Title:  "DEVELOPER NOTE",
			Output: SecretDeveloperMessage(width),
		}
	}

	parts := strings.Fields(lowered)
	root := parts[0]
	var sub string
	if len(parts) > 1 {
		sub = parts[1]
	}

	switch root {
	case "exit", "quit", "q":
		return ParsedCommand{Type: CmdQuit}

	case "home", "dashboard", "main":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenDashboard,
		}

	case "about", "profile", "bio", "who":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenAbout,
		}

	case "projects", "proj":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenProjects,
			Subcommand:   sub,
		}

	case "skills", "tech", "stack":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenSkills,
		}

	case "experience", "journey", "exp":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenExperience,
		}

	case "interests", "beyond", "passions":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenInterests,
		}

	case "contact", "email", "reach":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenContact,
		}

	case "help", "?", "man":
		return ParsedCommand{
			Type:         CmdNavigate,
			TargetScreen: ui.ScreenHelp,
		}

	case "clear", "cls":
		return ParsedCommand{Type: CmdClear}

	case "whoami", "id":
		msg := fmt.Sprintf("visitor (%s) • role: guest explorer • platform: JAWAD.OS", user)
		return ParsedCommand{
			Type:          CmdOutput,
			StatusMessage: msg,
		}

	case "date", "time":
		return ParsedCommand{
			Type:          CmdOutput,
			StatusMessage: DateOutput(),
		}

	case "neofetch", "fastfetch", "sysinfo":
		return ParsedCommand{
			Type:   CmdModal,
			Title:  "NEOFETCH",
			Output: NeofetchOutput(user, width),
		}

	case "secret", "developer":
		return ParsedCommand{
			Type:   CmdModal,
			Title:  "DEVELOPER NOTE",
			Output: SecretDeveloperMessage(width),
		}

	default:
		// Attempt closest match suggestion
		suggestion := suggestCommand(root)
		msg := fmt.Sprintf("Unknown command '%s'.", root)
		if suggestion != "" {
			msg += fmt.Sprintf(" Did you mean '%s'?", suggestion)
		} else {
			msg += " Type 'help' for available commands."
		}

		return ParsedCommand{
			Type:          CmdUnknown,
			StatusMessage: msg,
		}
	}
}

// suggestCommand offers a helpful fuzzy hint for common typos.
func suggestCommand(typo string) string {
	candidates := []string{
		"about", "projects", "skills", "experience", "interests",
		"contact", "help", "neofetch", "whoami", "clear", "home", "exit",
	}

	bestMatch := ""
	minDist := 3 // max typo tolerance

	for _, cand := range candidates {
		d := levenshteinDistance(typo, cand)
		if d < minDist {
			minDist = d
			bestMatch = cand
		}
	}

	return bestMatch
}

// levenshteinDistance calculates edit distance between two strings.
func levenshteinDistance(a, b string) int {
	la := len(a)
	lb := len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	prev := make([]int, lb+1)
	curr := make([]int, lb+1)

	for j := 0; j <= lb; j++ {
		prev[j] = j
	}

	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			curr[j] = min(curr[j-1]+1, min(prev[j]+1, prev[j-1]+cost))
		}
		copy(prev, curr)
	}

	return prev[lb]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
