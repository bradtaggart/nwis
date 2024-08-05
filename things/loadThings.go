package things

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LoadThings(url string) (*ThingsList, error) {
	var data ThingsList

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return &data, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return &data, err
	}

	err = json.Unmarshal([]byte(string(body)), &data)
	if err != nil {
		log.Println(err)
		return &data, err
	}

	for i := range data.Value {
		fmt.Println("ID:", data.Value[i].ID, "Name:", data.Value[i].Name, "Description:", data.Value[i].Description)
	}

	fmt.Println("")

	return &data, nil
}
