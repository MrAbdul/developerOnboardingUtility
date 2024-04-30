package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"strings"
)

type firstScreenModel struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func initialScreen1Model() firstScreenModel {
	return firstScreenModel{
		// Our to-do list is a grocery list
		Choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}
func (m firstScreenModel) Init() tea.Cmd {
	return nil
}

func (m firstScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m firstScreenModel) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("What Would you like to do?") + "\n\n")

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
