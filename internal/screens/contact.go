package screens

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lexisync/jawad-os/internal/config"
	"github.com/lexisync/jawad-os/internal/ui/theme"
)

type ContactItem struct {
	Label    string
	Value    string
	Protocol string
}

type ContactModel struct {
	cfg      *config.PortfolioConfig
	items    []ContactItem
	selected int
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewContactModel(cfg *config.PortfolioConfig) ContactModel {
	items := make([]ContactItem, 0)
	if cfg.Contact.GitHub != "" {
		items = append(items, ContactItem{Label: "GitHub", Value: cfg.Contact.GitHub, Protocol: "HTTPS / GIT"})
	}
	if cfg.Contact.LinkedIn != "" {
		items = append(items, ContactItem{Label: "LinkedIn", Value: cfg.Contact.LinkedIn, Protocol: "HTTPS"})
	} else {
		items = append(items, ContactItem{Label: "LinkedIn", Value: "Available upon direct request", Protocol: "REQUEST"})
	}
	if cfg.Contact.Email != "" {
		items = append(items, ContactItem{Label: "Email", Value: cfg.Contact.Email, Protocol: "SMTP / RFC5322"})
	} else {
		items = append(items, ContactItem{Label: "Email", Value: "Available upon direct request", Protocol: "REQUEST"})
	}
	if cfg.Contact.Website != "" {
		items = append(items, ContactItem{Label: "Website", Value: cfg.Contact.Website, Protocol: "HTTPS"})
	}
	if cfg.Contact.Location != "" {
		items = append(items, ContactItem{Label: "Location", Value: cfg.Contact.Location, Protocol: "UTC+5 / PK"})
	}

	return ContactModel{
		cfg:      cfg,
		items:    items,
		selected: 0,
	}
}

func (m *ContactModel) SetSize(w, h int) {
	m.width = w
	m.height = h

	geo := theme.CalculateLayout(w, h)
	vpHeight := geo.InnerHeight

	if !m.ready {
		m.viewport = viewport.New(geo.InnerWidth, vpHeight)
		m.ready = true
	} else {
		m.viewport.Width = geo.InnerWidth
		m.viewport.Height = vpHeight
	}

	m.viewport.SetContent(m.renderContent(geo.InnerWidth))
}

type CopyNotificationMsg struct {
	Value string
}

func (m ContactModel) Update(msg tea.Msg) (ContactModel, tea.Cmd, string) {
	var cmd tea.Cmd
	var notification string

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			} else {
				m.selected = len(m.items) - 1
			}
			geo := theme.CalculateLayout(m.width, m.height)
			m.viewport.SetContent(m.renderContent(geo.InnerWidth))
			return m, nil, ""

		case "down", "j":
			if m.selected < len(m.items)-1 {
				m.selected++
			} else {
				m.selected = 0
			}
			geo := theme.CalculateLayout(m.width, m.height)
			m.viewport.SetContent(m.renderContent(geo.InnerWidth))
			return m, nil, ""

		case "enter", "c":
			if len(m.items) > 0 && m.selected < len(m.items) {
				val := m.items[m.selected].Value
				notification = fmt.Sprintf("Copied %s: %s", m.items[m.selected].Label, val)
				copyCmd := copyToClipboardOSC52(val)
				return m, copyCmd, notification
			}
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd, notification
}

func copyToClipboardOSC52(text string) tea.Cmd {
	return func() tea.Msg {
		encoded := base64.StdEncoding.EncodeToString([]byte(text))
		fmt.Printf("\x1b]52;c;%s\x07", encoded)
		return CopyNotificationMsg{Value: text}
	}
}

func (m ContactModel) View() string {
	if !m.ready {
		return theme.StyleBodyMuted.Render("Loading communication console...")
	}
	return m.viewport.View()
}

func (m ContactModel) renderContent(width int) string {
	var b strings.Builder

	// Top Prompt
	promptText := theme.StyleBody.Width(width - 6).Render(
		"COMMUNICATIONS CONSOLE // Open to discussions regarding software engineering roles, backend systems architecture, and technical collaboration.",
	)
	b.WriteString(theme.RenderWorkstationPanel("DISPATCH NOTICE", promptText, width, 0, false) + "\n\n")

	// Channels Data Table
	headerCh := theme.TableHeaderStyle.Width(6).Render("CH")
	headerProto := theme.TableHeaderStyle.Width(16).Render("PROTOCOL")
	headerLabel := theme.TableHeaderStyle.Width(14).Render("CHANNEL")
	headerVal := theme.TableHeaderStyle.Render("ENDPOINT / VALUE")

	headerRow := lipgloss.JoinHorizontal(
		lipgloss.Left,
		headerCh,
		" ",
		headerProto,
		" ",
		headerLabel,
		" ",
		headerVal,
	)

	var rows []string
	rows = append(rows, headerRow)

	for i, item := range m.items {
		isSelected := i == m.selected
		chStr := fmt.Sprintf("[%02d]", i+1)

		chCell := theme.TableRowStyle.Width(6).Render(chStr)
		protoCell := theme.StyleCode.Width(16).Render(item.Protocol)
		labelCell := theme.TableRowStyle.Width(14).Render(item.Label)
		valCell := theme.TableRowStyle.Render(item.Value)

		if isSelected {
			chCell = theme.TableRowSelectedStyle.Width(6).Render(theme.GlyphBullet + " " + chStr)
			protoCell = theme.TableRowSelectedStyle.Width(16).Render(item.Protocol)
			labelCell = theme.TableRowSelectedStyle.Width(14).Render(item.Label)
			valCell = theme.TableRowSelectedStyle.Render(item.Value)
		}

		row := lipgloss.JoinHorizontal(
			lipgloss.Left,
			chCell,
			" ",
			protoCell,
			" ",
			labelCell,
			" ",
			valCell,
		)
		rows = append(rows, row)
	}

	hintText := theme.StyleBodyMuted.Render("\n[Select channel with ↑↓ or j/k, press ENTER to copy value via OSC 52]")
	tableContent := strings.Join(rows, "\n") + "\n" + hintText

	channelsPanel := theme.RenderWorkstationPanel("COMMUNICATION CHANNELS", tableContent, width, 0, true)
	b.WriteString(channelsPanel + "\n\n")

	// Availability & Guidelines
	guidelines := []string{
		"• Candidate Availability: Actively interviewing for Software Engineer / Full-Stack / Backend positions.",
		"• Preferred Topics: System design, distributed workflows, backend API engineering, infrastructure.",
		"• Direct Response: Direct messages on GitHub or direct contact inquiries receive prompt replies.",
	}
	policyText := theme.StyleBodyMuted.Width(width - 6).Render(strings.Join(guidelines, "\n\n"))
	b.WriteString(theme.RenderWorkstationPanel("DISPOSITION & PROTOCOL", policyText, width, 0, false) + "\n")

	return b.String()
}
