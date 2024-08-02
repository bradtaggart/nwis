package main

import (
	"log"

	"sbgis.com/nwis/datastreams"
)

func main() {
	//	message, err := things.LoadThings()

	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}

	message, _ := datastreams.LoadDataStream()

	log.Println(message)
}
