package controllers

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/beegons/models"
  "github.com/beegons/models/orion"

  "time"
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
  msg.Data[0]["_id"] = time.Now()
  err = models.Insert(msg.Data[0], "data")

}

/*

/*

curl -iX POST \
  'http://localhost:1026/v2/entities?options=keyValues' \
  -H 'Content-Type: application/json' \
  -d '
{
    "id": "urn:ngsi-ld:Module:MOD01",
    "type": "Module",
    "name": "MOD001",
    "mac": "13A434A01B58",
    "controlledProperties": ["temperature", "humidity", "o3", "co"],
    "protocol" :"HTTP",
    "state": "OK"
}'

/*
curl localhost:1026/v2/entities?options=keyValues -s -S -H 'Content-Type: application/json' -d '
{
  "id": "urn:ngsi-ld:AirQualityObserved:MOD01",
  "type": "AirQualityObserved",
  "airQualityIndex"
  "temperature":  23,
  "humidity": 80,
  "co": 100,
  "o3": 5,
  "dateObserved": "2019-07-16T14:00:00Z"
}'

curl 'localhost:1026/v2/entities?options=keyValues&type=AirQualityObserved'


curl -iX POST \
  --url 'http://localhost:1026/v2/subscriptions' \
  --header 'Content-Type: application/json' \
  --data '{
  "description": "AirQualityObserved",
  "subject": {
    "entities": [{"idPattern": ".*","type": "AirQualityObserved"}],
    "condition": {
      "attrs": ["temperature", "humidity", "xo2"]
    }
  },
  "notification": {
    "http": {
      "url": "http://beagons.uni.edu.pe:8081/airQualityObserved"
    },
    "attrsFormat" : "keyValues"
  }
}'
curl localhost:1026/v2/entities/urn:ngsi-ld:AirQualityObserved:Module1/attrs -s -S -H 'Content-Type: application/json' -X PATCH -d '{
      "temperature": {
        "value": '25',
        "type": "Float"
      }
  }'


SIM OF A SUBSCRIPTION
    curl -v -s -S -X POST http://localhost:8081/airQualityObserved \
    --header 'Content-Type: application/json; charset=utf-8' \
    --header 'Accept: application/json' \
    -d  '{
         "data": [
            {
                 "id": "urn:ngsi-ld:AirQualityObserved:Module1","type": "AirQualityObserved",
                 "CO": 40,
                 "humidity": 80,
                 "temperature": 22,
                 "dateObserved": "2019-07-16T14:00:00Z"
            }
         ],
         "subscriptionId": "57458eb60962ef754e7c0998"
     }'
*/
