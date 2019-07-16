package orion

type Subscription struct {
	_Id            string                   `bson:"_id"`
	SubscriptionId string                   `json:"subscriptionId" `
	Data           []map[string]interface{} `json:"data"`
}

// data:[map[CO:100 humidity:80 temperature:25 id:urn:ngsi-ld:AirQualityObserved:Module1 type:AirQualityObserved]]]
