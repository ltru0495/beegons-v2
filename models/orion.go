package models

import (
	"time"
)

// Subscriptions
type Payload struct {
	Description  string       `json:"description"`
	Subject      Subject      `json:"subject"`
	Notification Notification `json:"notification"`
	Throttling   int          `json:"throttling"`
}
type Entity struct {
	Id string `json:"id"`
}
type Subject struct {
	Entities []Entity `json:"entities"`
}
type HTTP struct {
	URL string `json:"url"`
}
type Notification struct {
	HTTP        HTTP   `json:"http"`
	AttrsFormat string `json:"attrsFormat"`
}

type CygnusDocument struct {
	RecvTime  time.Time `json:"recvTime" bson:"recvTime"`
	AttrValue string    `json:"attrValue" bson:"attrValue"`
}

/*	"_id" : ObjectId("5d65e81731389c001002afd8"),
	"recvTime" : ISODate("2019-08-28T02:33:59.969Z"),
	"attrName" : "temperature",
	"attrType" : "Number",
	"attrValue" : "29"
*/

type Subscription struct {
	Data           []map[string]interface{} `json:"data"`
	SubscriptionId string                   `json:"subscriptionId"`
}

type Alert struct {
	Id           string                 `json:"id"`
	Condition    string                 `json:"condition"`
	RefModule    string                 `json:"refModule"`
	DateObserved string                 `json:"time"`
	Parameters   map[string]interface{} `json:"parameters"`
}
