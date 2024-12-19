package main

import (
	"log"
	"os"
	"xkcdtool/xkcdtool"
)

// func main() {
// 	data := xkcdtool.GetJSON()
// 	if data == nil {
// 		log.Println("No data fetched.")
// 		return
// 	}
// 	xkcdtool.GenerateJSONFile(data)
// }

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func main() {
	fileName := "output.json"
	if !fileExists(fileName) {
		data := xkcdtool.GetJSON()
		if data == nil {
			log.Println("No data fetched.")
			return
		}
		xkcdtool.GenerateJSONFile(data)
		xkcdtool.Query()
	} else {
		xkcdtool.Query()
	}
}
