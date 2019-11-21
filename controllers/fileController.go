package controllers

import (
	"github.com/beegons/models"

	"github.com/gorilla/mux"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	//"log"
	// "encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func GenerateJSON(filename string, data []models.CygnusDocument) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	j, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, j, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func GenerateXML(filename string, data []models.CygnusDocument) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	j, err := xml.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, j, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["moduleid"]
	parameter := vars["parameter"]

	m, err := models.GetModule(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	dataId := "urn:ngsi-ld:DataObserved:" + m.Name
	dataType := m.DataType + "Observed"

	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}

	// TODO
	// start, end <- QUERY
	filename := dataId

	start, err := time.Parse("2006-02-01T15:04Z", vars["start"])
	if err != nil {
		log.Println(err)
	}
	end, err := time.Parse("2006-02-01T15:04Z", vars["end"])

	data, err := models.FilterDataByDate(dataId, dataType, parameter, start, end)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}

	format := vars["format"]
	if format == "csv" {
		filename = filename + ".csv"
		// GenerateCSV(filename, data)
	} else if format == "json" {
		filename = filename + ".json"
		GenerateJSON(filename, data)
	} else if format == "xml" {
		filename = filename + ".xml"
		GenerateXML(filename, data)
	}

	Openfile, err := os.Open(filename)
	defer Openfile.Close() //Close after functio	n return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)

	err = os.Remove(filename)
	if err != nil {
		log.Println(err)
	}

	return
}
