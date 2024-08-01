package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"./types/things"
)

func main() {
	var data things.ThingsList
	resp, err := http.Get("https://labs.waterdata.usgs.gov/sta/v1.1/Things")

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	err2 := json.Unmarshal([]byte(string(body)), &data)
	if err2 != nil {
		log.Println(err2)
		return
	}

}
