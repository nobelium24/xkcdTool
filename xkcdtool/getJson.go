package xkcdtool

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetJSON() []XKCDJSON {
	var responses []XKCDJSON
	for i := 0; i <= 3025; i++ {
		urlIndex := strconv.Itoa(i)
		url := "https://xkcd.com/" + urlIndex + "/info.0.json"
		log.Print(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("An error occurred creating request: %v", err)
			continue
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("An error occurred during request: %v", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to fetch comic, status code: %d", resp.StatusCode)
			continue
		}
		var comic XKCDJSON
		err = json.NewDecoder(resp.Body).Decode(&comic)
		if err != nil {
			log.Printf("An error occurred during JSON decoding: %v", err)
			continue
		}
		comic.Url = url
		responses = append(responses, comic)
		time.Sleep(1 * time.Second)
	}
	return responses
}
