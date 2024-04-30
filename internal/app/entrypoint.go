package app

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"fmt"
	"log"
	"path/filepath"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start(configLocation string, openHTMLreport bool) {

	var projectData data.ProjectData
	var err error
	if filepath.Ext(configLocation) == ".json" {
		projectData, err = data.LoadData(configLocation)
	} else if filepath.Ext(configLocation) == ".toml" {
		projectData, err = data.LoadDataToml(configLocation)
	} else {
		log.Fatal("Configuration file is not a .json  or .toml file")
	}
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
