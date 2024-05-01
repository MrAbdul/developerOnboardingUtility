package dnsCheck

import (
	"devOnBoardingUtility/internal/pkg/data"
	"fmt"
	"net"
	"strings"
)

func CheckDNS(data data.ProjectData) string {
	// This function will check the DNS of the given domain
	var toAddToHosts strings.Builder
	var ips []net.IP
	var err error
	for _, project := range data.Projects {
		if project.DnsName != "" {
			ips, err = net.LookupIP(project.DnsName)
			fmt.Printf("error is %v\n", err)
			if err != nil && len(ips) == 0 {
				fmt.Printf("No DNS entry found for %s it must be added to hosts file\n", project.DnsName)
				for _, ip := range project.Ips {
					toAddToHosts.WriteString(fmt.Sprintf("%s \t %s\n", project.DnsName, ip))
				}
				continue
			}

		}
	}
	return toAddToHosts.String()

}
