package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const fileName = "fileName.json"

func LoadData() ProjectData {
	file, err := os.Open(fileName)
	if err != nil {
		log.Print("problem opening file: ", err)
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print("problem closing file: ", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	var data ProjectData
	err = decoder.Decode(&data)
	if err != nil {
		log.Print("problem decoding file: ", err)
		log.Fatal(err)
	}

	return data
}

func PrintReportAsJSON(report Report) {
	// Convert the report to a JSON string with indentation
	jsonReport, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling report to JSON: %v", err)
	}

	// Print the JSON string
	fmt.Println(string(jsonReport))
}
