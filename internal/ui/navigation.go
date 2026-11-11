package ui

type Screen int

const (
	ScreenBoot Screen = iota
	ScreenDashboard
	ScreenAbout
	ScreenProjects
	ScreenSkills
	ScreenExperience
	ScreenInterests
	ScreenContact
	ScreenHelp
)

func (s Screen) String() string {
	switch s {
	case ScreenBoot:
		return "BOOT"
	case ScreenDashboard:
		return "WORKSTATION"
	case ScreenAbout:
		return "PROFILE"
	case ScreenProjects:
		return "PROJECTS"
	case ScreenSkills:
		return "STACK"
	case ScreenExperience:
		return "JOURNEY"
	case ScreenInterests:
		return "INTERESTS"
	case ScreenContact:
		return "CONTACT"
	case ScreenHelp:
		return "HELP"
	default:
		return "UNKNOWN"
	}
}

type MenuItem struct {
	ID       string
	Number   string
	Title    string
	Subtitle string
	Screen   Screen
}

var MenuItems = []MenuItem{
	{
		ID:       "about",
		Number:   "01",
		Title:    "PROFILE",
		Subtitle: "Who am I?",
		Screen:   ScreenAbout,
	},
	{
		ID:       "projects",
		Number:   "02",
		Title:    "PROJECTS",
		Subtitle: "Things I've built",
		Screen:   ScreenProjects,
	},
	{
		ID:       "skills",
		Number:   "03",
		Title:    "STACK",
		Subtitle: "Technologies I explore",
		Screen:   ScreenSkills,
	},
	{
		ID:       "experience",
		Number:   "04",
		Title:    "JOURNEY",
		Subtitle: "My journey",
		Screen:   ScreenExperience,
	},
	{
		ID:       "interests",
		Number:   "05",
		Title:    "INTERESTS",
		Subtitle: "Beyond programming",
		Screen:   ScreenInterests,
	},
	{
		ID:       "contact",
		Number:   "06",
		Title:    "CONTACT",
		Subtitle: "Let's connect",
		Screen:   ScreenContact,
	},
}

// WorkstationNavLabels returns the formatted ticker labels for the workstation navigation bar.
func WorkstationNavLabels() []string {
	labels := make([]string, len(MenuItems))
	for i, item := range MenuItems {
		labels[i] = item.Number + " " + item.Title
	}
	return labels
}

// ScreenToNavIndex maps a screen to its corresponding 0-indexed navigation bar position.
func ScreenToNavIndex(s Screen) int {
	switch s {
	case ScreenAbout:
		return 0
	case ScreenProjects:
		return 1
	case ScreenSkills:
		return 2
	case ScreenExperience:
		return 3
	case ScreenInterests:
		return 4
	case ScreenContact:
		return 5
	default:
		return -1
	}
}
