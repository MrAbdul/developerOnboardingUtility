package app

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/dnsCheck"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"fmt"
	"log"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start(configLocation string, openHTMLreport bool) {

	projectData, err := data.LoadData(configLocation)

	if err != nil {
		log.Printf(err.Error())

	}
	fmt.Printf("%+v\n", projectData)
	report := tcpconnector.Run(projectData)
	data.PrintReportAsJSON(report)
	toaddtohosts := dnsCheck.CheckDNS(projectData)
	path, err := data.WriteHTMLFile(data.GenerateHTML(report, toaddtohosts))
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
