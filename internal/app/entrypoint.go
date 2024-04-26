package app

import (
	"devOnBoardingUtility/internal/pkg/data"
	"devOnBoardingUtility/internal/pkg/tcpconnector"
	"fmt"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start() {
	projectData := data.LoadData()
	fmt.Printf("%+v\n", projectData)
	report := tcpconnector.Run(projectData)
	data.PrintReportAsJSON(report)
	err, path := data.WriteHTMLFile(data.GenerateHTML(report))
	if err != nil {
		return
	}
	data.OpenHTMLFile(path)

}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
