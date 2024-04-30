package data

type Project struct {
	ProjectName string   `json:"projectName" toml:"projectName"`
	Description string   `json:"description" toml:"description"`
	Ips         []string `json:"ips" toml:"ips"`     // List of IP addresses
	Ports       []int    `json:"ports" toml:"ports"` // List of port numbers
}
type ProjectData struct {
	ListVersion string    `json:"listVersion" toml:"listVersion"`
	Projects    []Project `json:"projects" toml:"projects"` // List of projects
}

type PortResult struct {
	Port   int    `json:"port" toml:"port"`
	Status string `json:"status" toml:"status"` // "open" or "closed"
}

type IPResult struct {
	IP          string       `json:"ip" toml:"ip"`                   // IP address
	PortResults []PortResult `json:"portResults" toml:"portResults"` // List of port statuses
}

type ProjectResult struct {
	ProjectName string     `json:"portResults" toml:"portResults"` // Project name
	Description string     `json:"description" toml:"description"` // Project description
	IPResults   []IPResult `json:"ipResults" toml:"ipResults"`     // List of IP results
}

type Report struct {
	ReportVersion string          `json:"reportVersion" toml:"reportVersion"` // Version of the report
	Projects      []ProjectResult `json:"projects" toml:"projects"`           // List of projects
}
