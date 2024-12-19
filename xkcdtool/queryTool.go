package xkcdtool

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func Query() {
	keyWord := os.Args[1]
	jsonData, err := os.Open("output.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	var comics []XKCDJSON
	err = json.NewDecoder(jsonData).Decode(&comics)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	var searchedComics []XKCDJSON

	for i := 0; i < len(comics); i++ {
		comic := comics[i]

		title := comic.Title
		safeTitle := comic.SafeTitle
		transacript := comic.Transcript
		alt := comic.Alt

		if strings.Contains(title, keyWord) || strings.Contains(safeTitle, keyWord) || strings.Contains(transacript, keyWord) || strings.Contains(alt, keyWord) {
			searchedComics = append(searchedComics, comic)
		}
	}
	if len(searchedComics) == 0 {
		fmt.Printf("Search item %s can't be found in the offline index\n", keyWord)
		return
	}

	for _, item := range searchedComics {
		fmt.Println("Month Num Link Year News SafeTitle Transcript Alt Img TItle Day Url")
		fmt.Printf("%s, %d, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n",
			item.Month, item.Num, item.Link, item.Year, item.News,
			item.SafeTitle, item.Transcript, item.Alt, item.Img, item.Title, item.Day, item.Url)
	}
}
