package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/beegons/models"

	"log"
)

const orionURL = "http://localhost:1026/v2/entities"

func PostEntity(entity interface{}) error {

	q := "?options=keyValues"
	payloadBytes, err := json.Marshal(entity)
	log.Println(string(payloadBytes))

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

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(body)
	// s := buf.String()

	log.Println(resp)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func GetEntities(entityType string) (modules []models.Module, err error) {
	url := "http://localhost:1026/v2/entities?options=keyValues&type=" + entityType
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&modules)
	if err != nil {
		return
	}
	log.Println(modules)

	return

}

/* [{"id":"000000000000000000000000","type":"Module","controlledProperties":["temperature","humedad","o3"],"mac":"123ADASDA0012","name":"MOD1","protocol":"HTTP","state":"OK"}]
 */
