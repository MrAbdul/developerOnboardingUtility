package cmd

import (
	"devOnBoardingUtility/internal/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkportsCmd)
	checkportsCmd.Flags().StringP("config", "c", "defaultConfig", "config file path")
	checkportsCmd.Flags().BoolP("open-htmlreport", "o", true, "if its included i will open an html report")
}

var checkportsCmd = &cobra.Command{
	Use:   "checkports",
	Short: "check ports that are needed for dev",
	Long:  "check ports that are needed for dev",
	Run:   app.Start,
}
