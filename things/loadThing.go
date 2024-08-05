package things

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LoadThing(url string) (*Thing, error) {
	var data Thing

	log.Println(url)
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

	err2 := json.Unmarshal([]byte(string(body)), &data)
	if err2 != nil {
		log.Println(err2)
		return &data, err
	}

	fmt.Println("ID:", data.ID, "Name:", data.Name, "Description:", data.Description)

	fmt.Println("")

	return &data, nil
}
