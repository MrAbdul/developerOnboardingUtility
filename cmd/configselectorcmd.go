package cmd

import (
	"devOnBoardingUtility/internal/app"
	"devOnBoardingUtility/tui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var selectConfigCmd = &cobra.Command{
	Use:   "select-config",
	Short: "Select a configuration file to process",
	Run: func(cmd *cobra.Command, args []string) {
		initialModel := tui.Model{
			Choices:     []string{},
			Header:      "Select a configuration file to process",
			DisplayLogo: true,
			Selected:    make(map[int]struct{}),
		}

		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".json" || filepath.Ext(path) == ".toml" {
				initialModel.Choices = append(initialModel.Choices, path)
			}
			return nil
		})
		if err != nil {
			fmt.Println("Failed to list JSON files:", err)
			os.Exit(1)
		}

		p := tea.NewProgram(initialModel)
		if first, err := p.Run(); err != nil {
			fmt.Println("Failed to start TUI:", err)
			os.Exit(1)
		} else {
			finalModelTyped := first.(tui.Model)
			selectedFile := finalModelTyped.Choices[finalModelTyped.Cursor]
			secondModel := tui.Model{
				Choices:     []string{"yes", "no"},
				DisplayLogo: false,
				Header:      "you have selected: " + selectedFile + "\n do you want to show report as html?",
				Cursor:      0,
				Selected:    make(map[int]struct{}),
			}
			r := tea.NewProgram(secondModel)
			if second, err := r.Run(); err != nil {
				fmt.Println("Failed to start TUI:", err)
				os.Exit(1)
			} else {
				finalSecondTyped := second.(tui.Model)
				selectedOption := finalSecondTyped.Choices[finalSecondTyped.Cursor]
				var isHtml bool
				switch selectedOption {
				case "yes":
					isHtml = true
				default:
					isHtml = false
				}
				app.Start(selectedFile, isHtml)

			}
			fmt.Println("You selected:", selectedFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectConfigCmd)
}
