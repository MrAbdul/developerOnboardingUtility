package data

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GenerateHTML(report Report, toaddtohosts string) string {
	var sb strings.Builder
	sb.WriteString(`
	<html>
	<head>
		<title>Project Report</title>
		<style>
			table {
				border-collapse: collapse;
				width: 100%;
			}
			th, td {
				border: 1px solid black;
				padding: 8px;
				text-align: left;
			}
			tr:nth-child(even) {
				background-color: #f2f2f2;
			}
			tr:hover {
				background-color: #ddd;
			}
		</style>
	</head>
	<body>
		<h1>Project Report (Version ` + report.ReportVersion + `)</h1>
		<table>
			<tr>
				<th>Project Name</th>
				<th>Description</th>
				<th>IP Address</th>
				<th>Port</th>
				<th>Status</th>
			</tr>`)

	for _, project := range report.Projects {
		for _, ipResult := range project.IPResults {
			for _, portResult := range ipResult.PortResults {
				sb.WriteString(`
				<tr>
					<td>` + project.ProjectName + `</td>
					<td>` + project.Description + `</td>
					<td>` + ipResult.IP + `</td>
					<td>` + fmt.Sprintf("%d", portResult.Port) + `</td>
					<td>` + portResult.Status + `</td>
				</tr>`)
			}
		}
	}

	sb.WriteString(`
		</table>
		<h2>Hosts File Entries</h2>
		<pre>` + toaddtohosts + `</pre>
	</body>
	</html>`)

	return sb.String()
}

// Function to write HTML content to a file
func WriteHTMLFile(htmlContent string) (string, error) {
	path, err := getExecutableDirectory()
	if err != nil {
		return "", err
	}
	path = path + string(os.PathSeparator) + "results.html"
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	_, err = file.WriteString(htmlContent)
	return path, err
}

// Function to open an HTML file in the default web browser
func OpenHTMLFile(filePath string) error {
	var cmd *exec.Cmd

	// Determine the appropriate command to open in Chrome
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", "-a", "Google Chrome", filePath)
	case "linux": // Linux
		cmd = exec.Command("google-chrome", filePath) // May also use "chromium-browser"
	case "windows": // Windows
		cmd = exec.Command("cmd", "/c", "start chrome", filePath)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Run() // Execute the command to open the HTML file in Chrome
}

// Function to get the directory path of the current executable
func getExecutableDirectory() (string, error) {
	exePath, err := os.Executable() // Get the path of the current executable
	if err != nil {
		return "", err
	}

	// Resolve to the absolute path (handles symlinks if needed)
	absPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return "", err
	}

	// Get the directory containing the executable
	dirPath := filepath.Dir(absPath)

	return dirPath, nil
}
