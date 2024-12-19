package xkcdtool

import (
	"encoding/json"
	"log"
	"os"
)

func GenerateJSONFile(data []XKCDJSON) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("An error occurred during JSON conversion: %v", err)
		return
	}
	file, err := os.Create("output.json")
	if err != nil {
		log.Printf("An error occurred during file creation: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("An error occurred during file writing: %v", err)
		return
	}
	log.Println("JSON written to file: output.json")
}
