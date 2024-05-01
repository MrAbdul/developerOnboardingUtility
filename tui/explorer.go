package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"strings"
)

type Model struct {
	Choices     []string
	DisplayLogo bool
	Header      string
	Cursor      int
	Selected    map[int]struct{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter", " ":
			return m, tea.Quit
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	//var logo = logoStyle.Render(LogoImg) + "\n" + logoStyle.Render(LogoTextEngilish) + "\n"
	//b.WriteString(logoStyle.Render(LogoImg) + "\n")
	//b.WriteString(logoStyle.Render(LogoTextEngilish) + "\n")
	if m.DisplayLogo {
		b.WriteString(logoStyle.Render(logo) + "\n")
	}
	b.WriteString(titleStyle.Render(m.Header) + "\n\n")

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		// Render the line with or without selection style
		line := fmt.Sprintf("%s %s", cursor, choice)
		if _, selected := m.Selected[i]; selected {
			b.WriteString(selectedItemStyle.Render(line) + "\n")
		} else {
			b.WriteString(itemStyle.Render(line) + "\n")
		}
	}
	return b.String()
}
