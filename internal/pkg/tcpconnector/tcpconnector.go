package tcpconnector

import (
	"devOnBoardingUtility/internal/pkg/data"
	"fmt"
	"net"
	"sync"
	"time"
)

const timeout = 3 * time.Second

func TestPort(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		//fmt.Printf(err.Error())
		return false // Failed to connect (port is closed or filtered)
	}

	err = conn.Close()
	if err != nil {
		fmt.Printf(err.Error())
		panic(0)
	}
	return true // Successfully connected (port is open)
}
func getNeededChannelSize(d data.ProjectData) int {
	var size int
	for _, project := range d.Projects {
		size += len(project.Ips) * len(project.Ports)
	}
	return size
}
func Run(d data.ProjectData) data.Report {
	report := data.Report{ReportVersion: "1.0"}

	var wg sync.WaitGroup
	resultsChan := make(chan data.IPResult, getNeededChannelSize(d)) // Buffer size can be adjusted

	// Launch goroutines to test each IP/port combination
	for _, project := range d.Projects {
		for _, ip := range project.Ips {
			wg.Add(1)
			go func(ip string, ports []int) {
				defer wg.Done()
				ipResult := data.IPResult{IP: ip}

				for _, port := range ports {
					portStatus := "closed"
					if TestPort(ip, port) {
						portStatus = "open"
						fmt.Printf("Port %d is open on IP %s\n", port, ip)
					} else {
						fmt.Printf("Port %d is not open on IP %s\n", port, ip)
					}
					ipResult.PortResults = append(ipResult.PortResults, data.PortResult{
						Port:   port,
						Status: portStatus,
					})
				}

				// Send the completed IP result to the results channel
				resultsChan <- ipResult
			}(ip, project.Ports)
		}
	}

	// Wait for all goroutines to complete, then close the results channel
	wg.Wait()
	close(resultsChan)

	// Create a map to associate projects with their IP results
	projectMap := make(map[string]*data.ProjectResult)

	// Process the results from the channel
	for ipResult := range resultsChan {
		// Find the corresponding project for this IP address
		for _, project := range d.Projects {
			if contains(project.Ips, ipResult.IP) { // Helper function to check if IP is in the project
				projectName := project.ProjectName

				// If this project hasn't been added to the map, create a new ProjectResult
				if _, exists := projectMap[projectName]; !exists {
					projectMap[projectName] = &data.ProjectResult{
						ProjectName: project.ProjectName,
						Description: project.Description,
						IPResults:   []data.IPResult{},
					}
				}

				// Add the IPResult to the correct project in the map
				projectMap[projectName].IPResults = append(projectMap[projectName].IPResults, ipResult)
				break
			}
		}
	}

	// Convert the projectMap to a list and add it to the report
	for _, projectResult := range projectMap {
		report.Projects = append(report.Projects, *projectResult)
	}

	return report
}

func contains(ips []string, ip string) bool {
	for _, item := range ips {
		if item == ip {
			return true
		}
	}
	return false
}
