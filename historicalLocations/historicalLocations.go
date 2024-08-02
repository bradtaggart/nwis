package historicallocations

type HistoricalLocation struct {
	SelfLink      string `json:"@iot.selfLink"`
	ID            string `json:"@iot.id"`
	Time          string `json:"time"`
	ThingLink     string `json:"Thing@iot.navigationLink"`
	LocationsLink string `json:"Locations@iot.navigationLink"`
}

type HistoricalLocationList struct {
	Value []HistoricalLocation `json:"value"`
}
