package cmd

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/dnsCheck"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkDNScmd)
	checkDNScmd.Flags().StringP("config", "c", "defaultConfig", "config file path")

}

var checkDNScmd = &cobra.Command{
	Use:   "checkdns",
	Short: "check dns that are needed for dev",
	Long:  "check dns that are needed for dev",
	Run:   runDNSChecker,
}

func runDNSChecker(cmd *cobra.Command, args []string) {
	configLocation, _ := cmd.Flags().GetString("config")
	loadData, err := data.LoadData(configLocation)
	if err != nil {
		return
	}
	dnsCheck.CheckDNS(loadData)
}
