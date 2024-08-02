package main

import (
	"log"

	"sbgis.com/nwis/things"
)

func main() {
	message, err := things.LoadThings()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(message)
}
