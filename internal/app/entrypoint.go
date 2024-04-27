package app

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start(cmd *cobra.Command, args []string) {

	configLocation, _ := cmd.Flags().GetString("config")
	openHTMLreport, _ := cmd.Flags().GetBool("open-htmlreport")
	//saveReportJson, _ := cmd.Flags().GetBool("save-report")
	projectData, err := data.LoadData(configLocation)
	if err != nil {
		log.Printf(err.Error())

	}
	fmt.Printf("%+v\n", projectData)
	report := tcpconnector.Run(projectData)
	data.PrintReportAsJSON(report)
	path, err := data.WriteHTMLFile(data.GenerateHTML(report))
	if err != nil {
		fmt.Printf("Error writing HTML file: %v\n", err)
		return
	}
	if openHTMLreport {
		err = data.OpenHTMLFile(path)
		if err != nil {
			fmt.Printf("Error opening HTML file: %v\n", err)
		}
	}
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
