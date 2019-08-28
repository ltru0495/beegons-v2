package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"errors"
	"log"
)

func CheckOrion() {
	orionURL := GetOrionURL()

	_, err := http.Get(orionURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Orion Broker OK at " + orionURL)
}

func PostEntity(entity interface{}) error {
	orionURL := GetOrionURL()
	q := "?options=keyValues"
	url := orionURL + "/entities" + q
	payloadBytes, err := json.Marshal(entity)

	if err != nil {
		return err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	// log.Println(resp)

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(body)
	// s := buf.String()
	// log.Println(s)
	// log.Println(resp.StatusCode)

	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("An error has ocurred")
	}

	defer resp.Body.Close()

	return nil
}

func GetEntities(entityType string, itf interface{}) (err error) {
	orionURL := GetOrionURL()
	url := orionURL + "/entities?options=keyValues&type=" + entityType
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&itf)
	if err != nil {
		return
	}
	return
}

func PostSubscription(data interface{}) error {
	orionURL := GetOrionURL()
	url := orionURL + "/subscriptions"
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

/*
curl -iX POST \
  'http://localhost:1026/v2/subscriptions' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "Notify Cygnus of all context changes",
  "subject": {
    "entities": [
      {
        "idPattern": ""
      }
    ]
  },
  "notification": {
    "http": {
      "url": "http://cygnus:5050/notify"
    },
    "attrsFormat": "legacy"
  },
  "throttling": 5
}'
*/

func GetEntity(entityId string, itf interface{}) (err error) {
	orionURL := GetOrionURL()
	url := orionURL + "/entities/" + entityId + "?options=keyValues"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&itf)
	if err != nil {
		return
	}
	if res.StatusCode == 404 {
		err = errors.New("Not Found")
	}
	return

}
