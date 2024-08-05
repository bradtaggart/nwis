package locations

type Location struct {
	SelfLink       string      `json:"@iot.selfLink"`
	ID             string      `json:"@iot.id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	EncodingType   string      `json:"encodingType"`
	Location       GeoLocation `json:"location"`
	Properties     Properties  `json:"properties"`
	HistoricalLink string      `json:"HistoricalLocations@iot.navigationLink"`
	ThingsLink     string      `json:"Things@iot.navigationLink"`
}

// GeoLocation represents the location's geographical details.
type GeoLocation struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Properties represents various properties associated with the location.
type Properties struct {
	Type                     string  `json:"type"`
	COMID                    string  `json:"comid"`
	Measure                  float64 `json:"measure"`
	COMIDURL                 string  `json:"comidURL"`
	ReachCode                string  `json:"reachcode"`
	MainstemURL              string  `json:"mainstemURL"`
	MainstemNameAtOutlet     string  `json:"mainstemNameAtOutlet"`
	DownstreamMainstemURL    string  `json:"downstreamMainstemURL"`
	MainstemNameAtOutletGNIS int     `json:"mainstemNameAtOutletGNIS"`
}

type ValueList struct {
	Value []Location `json:"value"`
}
