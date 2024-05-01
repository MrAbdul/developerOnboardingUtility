package data

import (
	"log"
	"path/filepath"
)

func LoadData(configLocation string) (ProjectData, error) {
	var projectData ProjectData
	var err error
	if filepath.Ext(configLocation) == ".json" {
		projectData, err = LoadDataJSON(configLocation)
	} else if filepath.Ext(configLocation) == ".toml" {
		projectData, err = LoadDataToml(configLocation)
	} else {
		log.Fatal("Configuration file is not a .json  or .toml file")
	}
	return projectData, err
}
