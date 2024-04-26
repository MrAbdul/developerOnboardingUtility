package data

type Project struct {
	ProjectName string   `json:"projectName"`
	Description string   `json:"description"`
	Ips         []string `json:"ips"`   // List of IP addresses
	Ports       []int    `json:"ports"` // List of port numbers
}
type ProjectData struct {
	ListVersion string    `json:"listVersion"`
	Projects    []Project `json:"projects"` // List of projects
}

type PortResult struct {
	Port   int    `json:"port"`
	Status string `json:"status"` // "open" or "closed"
}

type IPResult struct {
	IP          string       `json:"ip"`          // IP address
	PortResults []PortResult `json:"portResults"` // List of port statuses
}

type ProjectResult struct {
	ProjectName string     `json:"projectName"` // Project name
	Description string     `json:"description"` // Project description
	IPResults   []IPResult `json:"ipResults"`   // List of IP results
}

type Report struct {
	ReportVersion string          `json:"reportVersion"` // Version of the report
	Projects      []ProjectResult `json:"projects"`      // List of projects
}
