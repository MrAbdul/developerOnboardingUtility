package app

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"fmt"
	"log"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start() {
	projectData, err := data.LoadData()
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
	err = data.OpenHTMLFile(path)
	if err != nil {
		fmt.Printf("Error opening HTML file: %v\n", err)
	}

}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
