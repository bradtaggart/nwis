package things

type Thing struct {
	Selflink    string `json:"@iot.selfLink"`
	ID          string `json:"@iot.id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Properties  struct {
		State                        string  `json:"state"`
		County                       string  `json:"county"`
		Country                      string  `json:"country"`
		MapScale                     int     `json:"mapScale"`
		StateCode                    string  `json:"stateCode"`
		AgencyCode                   string  `json:"agencyCode"`
		CountryFIPS                  string  `json:"countryfips"`
		HydrologicUnit               string  `json:"hydrologicUnit"`
		DecimalLatitude              float64 `json:"decimalLatitude"`
		DecimalLongitude             float64 `json:"decimalLongitude"`
		MonitoringLocationUrl        string  `json:"monitoringLocationUrl"`
		MonitoringLocationName       string  `json:"monitoringLocationName"`
		MonitoringLocationType       string  `json:"monitoringLocationType"`
		MonitoringLocationNumber     string  `json:"monitoringLocationNumber"`
		LocationHUCTwelveDigitCode   string  `json:"locationHUCTwelveDigitCode"`
		DecimalLatitudeStandardized  float64 `json:"decimalLatitudeStandardized"`
		DecimalLongitudeStandardized float64 `json:"decimalLongitudeStandardized"`
	} `json:"properties"`
	LocationLink   string `json:"Locations@iot.navigationLink"`
	HistoricalLink string `json:"HistoricalLocations@iot.navigationLink"`
	DatastreamLink string `json:"Datastreams@iot.navigationLink"`
}

type ThingsList struct {
	Value []Thing `json:"value"`
}
