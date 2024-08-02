package things

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LoadThings() (string, error) {
	var data ThingsList
	resp, err := http.Get("https://labs.waterdata.usgs.gov/sta/v1.1/Things")

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return "", err
	}

	err2 := json.Unmarshal([]byte(string(body)), &data)
	if err2 != nil {
		log.Println(err2)
		return "", err
	}

	for i := range data.Value {
		fmt.Println("ID:", data.Value[i].ID, "Name:", data.Value[i].Name, "Descritption:", data.Value[i].Description)
	}

	fmt.Println("")

	return "ok", nil
}
