package data

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

func LoadDataToml(location string) (ProjectData, error) {
	var projectdata ProjectData

	_, err := toml.DecodeFile(location, &projectdata)

	if err != nil {
		log.Print("problem decoding file: ", err)
		return ProjectData{}, err
	}

	return projectdata, nil
}

func PrintReportAsToml(report Report) {
	// Convert the report to a JSON string with indentation
	jsonReport, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Printf("Error marshaling report to JSON: %v", err)
	}

	// Print the JSON string
	fmt.Println(string(jsonReport))
}
