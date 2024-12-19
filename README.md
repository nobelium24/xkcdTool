# xkcdTool

## Overview
xkcdTool is a simple tool designed to fetch and process data from the xkcd webcomic API. It retrieves information about xkcd comics and generates a JSON file containing the data. Additionally, it allows querying the generated JSON file for specific keywords.

## How It Works
1. **Fetching Data**: The tool sends HTTP GET requests to the xkcd API to fetch comic data. It iterates through comic numbers from 0 to 3025, ensuring a 1-second delay between requests to avoid exceeding rate limits.
2. **Processing Data**: The fetched data is decoded from JSON format into Go structs.
3. **Generating JSON File**: The processed data is then marshaled into a formatted JSON string and written to an output file named `output.json`.
4. **Querying Data**: The tool can query the generated JSON file for specific keywords provided as command-line arguments and print the matching comics' details.

## Functions
### GetJSON
This function fetches data from the xkcd API and returns a slice of `XKCDJSON` structs.

### GenerateJSONFile
This function takes a slice of `XKCDJSON` structs, converts it to a formatted JSON string, and writes it to `output.json`.

### Query
This function reads the `output.json` file, searches for comics containing a specified keyword provided as a command-line argument, and prints the details of the matching comics.

## Usage
To use xkcdTool, simply run the program. It will fetch the data and generate the JSON file automatically if it does not already exist. If the file exists, it will query the existing data based on the provided keyword.

## Example
```go
package main

import (
    "log"
    "os"
    "xkcdtool/xkcdtool"
)

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
    }
    xkcdtool.Query()
}
```

To query the data, run the program with a keyword as a command-line argument:
```bash
go run main.go keyword
```

### Dependencies
  - Go standard library packages: `encoding/json`, `log`, `strconv`, `net/http`, `time` , `os`
