package tui

import "github.com/charmbracelet/lipgloss"

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("16")).
	Background(lipgloss.Color("124")).
	Padding(0, 1) // Top and bottom padding 0, left and right padding 1
var logoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")). // White
	TabWidth(4)
var itemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#333333")).
	Padding(0, 0, 0, 1)

var selectedItemStyle = itemStyle.Copy().
	Foreground(lipgloss.Color("#333333")).
	Background(lipgloss.Color("#FFFFFF"))

var inheaderSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#333333")).
	Background(lipgloss.Color("#FFFFFF"))
