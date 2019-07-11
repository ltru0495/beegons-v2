package models

type AirQualityObserved struct {
	IdModule string `json:"id" bson:"id"`
	Type     string `json:"type" bson:"type`
	Date     string `json:"date" bson:"date"`
	Value    string `json:"value" bson:"value"`
}

type Parameter struct {
	Value float64 `json:"value"`
}
