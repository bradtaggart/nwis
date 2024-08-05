package datastreams

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LoadDataStream(url string) error {
	var data ValueList
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return err
	}

	err2 := json.Unmarshal([]byte(string(body)), &data)
	if err2 != nil {
		log.Println(err2)
		return err
	}

	for i := range data.Value {
		fmt.Println("ID:", data.Value[i].ID, "Name:", data.Value[i].Name, "Description:", data.Value[i].Description)
		fmt.Println("UnitOfMeasure:", data.Value[i].UnitOfMeasurement)
		fmt.Println("Observed Area:", data.Value[i].ObservedArea)

		fmt.Println("Thresholds")

		for j := range data.Value[i].Properties.Thresholds {
			fmt.Println("Name:", data.Value[i].Properties.Thresholds[j].Name)
		}
	}

	fmt.Println("")

	return nil
}
