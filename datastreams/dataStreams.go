package datastreams

type DataStream struct {
	SelfLink             string            `json:"@iot.selfLink"`
	ID                   string            `json:"@iot.id"`
	Name                 string            `json:"name"`
	Description          string            `json:"description"`
	ObservationType      string            `json:"observationType"`
	UnitOfMeasurement    UnitOfMeasurement `json:"unitOfMeasurement"`
	ObservedArea         ObservedArea      `json:"observedArea"`
	PhenomenonTime       string            `json:"phenomenonTime"`
	Properties           Properties        `json:"properties"`
	ObservedPropertyLink string            `json:"ObservedProperty@iot.navigationLink"`
	SensorLink           string            `json:"Sensor@iot.navigationLink"`
	ThingLink            string            `json:"Thing@iot.navigationLink"`
	ObservationsLink     string            `json:"Observations@iot.navigationLink"`
}

type UnitOfMeasurement struct {
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Definition string `json:"definition"`
}

type ObservedArea struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Thresholds            []Threshold `json:"Thresholds"`
	ParameterCode         string      `json:"ParameterCode"`
	StatisticCode         string      `json:"StatisticCode"`
	WebDescription        string      `json:"WebDescription"`
	SubLocationIdentifier string      `json:"SubLocationIdentifier"`
}

type Threshold struct {
	Name          string   `json:"Name"`
	Type          string   `json:"Type"`
	Periods       []Period `json:"Periods"`
	ReferenceCode string   `json:"ReferenceCode"`
}

type Period struct {
	EndTime                        string      `json:"EndTime"`
	StartTime                      string      `json:"StartTime"`
	SuppressData                   bool        `json:"SuppressData"`
	ReferenceValue                 float64     `json:"ReferenceValue"`
	ReferenceValueToTriggerDisplay interface{} `json:"ReferenceValueToTriggerDisplay"`
}

type ValueList struct {
	Value []DataStream `json:"value"`
}
