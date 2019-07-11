package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beegons/models/orion"
)

func OrionSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*******************************")
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg orion.Subscription
	err = json.Unmarshal(b, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range msg.Data[0] {
		fmt.Println(k, v)
	}
}

/*
curl -iX POST \
  --url 'http://localhost:1026/v2/subscriptions' \
  --header 'Content-Type: application/json' \
  --data '{
  "description": "AirQualityObserved",
  "subject": {
    "entities": [{"idPattern": ".*","type": "AirQualityObserved"}],
    "condition": {
      "attrs": ["temperature", "humidity", "CO2"]
    }
  },
  "notification": {
    "http": {
      "url": "http://beagons.uni.edu.pe:8081/airQualityObserved"
    },
    "attrsFormat" : "keyValues"
  }
}'
*/

/*
	curl localhost:1026/v2/entities?options=keyValues -s -S -H 'Content-Type: application/json' -d @- <<EOF
	{
	  "id": "urn:ngsi-ld:AirQualityObserved:Module1",
	  "type": "AirQualityObserved",
	  "temperature":  23,
	  "humidity": 80,
	  "CO": 100
	}
	EOF

*/

/*
curl localhost:1026/v2/entities/urn:ngsi-ld:AirQualityObserved:Module1/attrs -s -S -H 'Content-Type: application/json' -X PATCH -d '{
      "temperature": {
        "value": '25',
        "type": "Float"
      }
  }'

    curl -v -s -S -X POST http://localhost:8081/airQualityObserved \
    --header 'Content-Type: application/json; charset=utf-8' \
    --header 'Accept: application/json' \
    -d  '{
         "data": [
            {
                 "id": "urn:ngsi-ld:AirQualityObserved:Module1","type": "AirQualityObserved",
                 "CO": 40,
                 "humidity": 80,
                 "temperature": 22
            }
         ],
         "subscriptionId": "57458eb60962ef754e7c0998"
     }'



*/
