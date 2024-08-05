package main

import (
	"log"

	"sbgis.com/nwis/datastreams"
	"sbgis.com/nwis/things"

	"sbgis.com/nwis/locations"
)

func main() {
	var thing *things.Thing
	var err error
	var message string

	thing, err = things.LoadThing("https://labs.waterdata.usgs.gov/sta/v1.1/Things('USGS-09380000')")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("getting locations....")
	err = locations.LoadLocations(thing.LocationsLink)
	if err != nil {
		log.Println(err)
		return
	}
	err = datastreams.LoadDataStream(thing.DatastreamsLink)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(message)
}
