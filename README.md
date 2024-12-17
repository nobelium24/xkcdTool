# xkcdTool

## Overview
xkcdTool is a simple tool designed to fetch and process data from the xkcd webcomic API. It retrieves information about xkcd comics and generates a JSON file containing the data.

## How It Works
1. **Fetching Data**: The tool sends HTTP GET requests to the xkcd API to fetch comic data. It iterates through comic numbers from 0 to 3025, ensuring a 1-second delay between requests to avoid exceeding rate limits.
2. **Processing Data**: The fetched data is decoded from JSON format into Go structs.
3. **Generating JSON File**: The processed data is then marshaled into a formatted JSON string and written to an output file named `output.json`.

## Functions
### GetJSON
This function fetches data from the xkcd API and returns a slice of `XKCDJSON` structs.

### GenerateJSONFile
This function takes a slice of `XKCDJSON` structs, converts it to a formatted JSON string, and writes it to `output.json`.

## Usage
To use xkcdTool, simply run the program. It will fetch the data and generate the JSON file automatically.

## Example
```go
package main

import (
    "xkcdtool"
)

func main() {
    data := xkcdtool.GetJSON()
    xkcdtool.GenerateJSONFile(data)
}
```
### Dependencies
  - Go standard library packages: `encoding/json`, `log`, `strconv`, `net/http`, `time` , `os`
