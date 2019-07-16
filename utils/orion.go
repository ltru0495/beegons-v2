package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"errors"
	"log"
)

const orionURL = "http://localhost:1026/v2/entities"

func PostEntity(entity interface{}) error {

	q := "?options=keyValues"
	payloadBytes, err := json.Marshal(entity)

	if err != nil {
		return err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", orionURL+q, body)
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
	url := orionURL + "?options=keyValues&type=" + entityType
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
	log.Println(itf)
	return
}

func GetEntity(entityId string, itf interface{}) (err error) {
	url := orionURL + "/" + entityId + "?options=keyValues"
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
	log.Println(itf)
	if res.StatusCode == 404 {
		err = errors.New("Not Found")
	}
	return

}

/* [{"id":"000000000000000000000000","type":"Module","controlledProperties":["temperature","humedad","o3"],"mac":"123ADASDA0012","name":"MOD1","protocol":"HTTP","state":"OK"}]
 */
