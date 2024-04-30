package cmd

import (
	"devOnBoardingUtility/internal/app"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkportsCmd)
	checkportsCmd.Flags().StringP("config", "c", "defaultConfig", "config file path")
	checkportsCmd.Flags().BoolP("open-htmlreport", "o", true, "if its included i will open an html report")

	rootCmd.AddCommand(checkSpecifcPortCmd)
	checkSpecifcPortCmd.Flags().StringP("ip", "i", "localhost", "specify the ip")
	checkSpecifcPortCmd.Flags().IntP("port", "p", 0, "specify the port")
}

var checkportsCmd = &cobra.Command{
	Use:   "checkports",
	Short: "check ports that are needed for dev",
	Long:  "check ports that are needed for dev",
	Run:   runChecker,
}

func runChecker(cmd *cobra.Command, args []string) {
	configLocation, _ := cmd.Flags().GetString("config")
	openHTMLreport, _ := cmd.Flags().GetBool("open-htmlreport")
	app.Start(configLocation, openHTMLreport)
}

var checkSpecifcPortCmd = &cobra.Command{
	Use:   "specifcPort",
	Short: "used for checking a specific port",
	Long: `used for checking a specific port
			like this spc -i <ip> -p <port>
`,
	Aliases: []string{"spc"},
	Run:     tcpconnector.CeckSpecifcPort,
}
