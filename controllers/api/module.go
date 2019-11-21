package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/gorilla/mux"
)

func Module(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["moduleid"]
	module, err := models.GetModule(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, module)
	return
}

func ModuleParameters(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["moduleid"]
	data, err := models.GetDataObserved(id)
	var params []string

	for k := range data.Parameters {
		params = append(params, k)
	}

	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, params)
	return
}

func ModuleRealTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["moduleid"]

	d, err := models.GetDataObserved(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	dataId := d.Id
	dataType := d.DataType

	var content []models.ApiData
	for parameter := range d.Parameters {
		var apiData models.ApiData

		data, err := models.GetLastData(dataId, dataType, parameter, 10)
		if err != nil {
			log.Println(err)
			models.SendNotFound(w)
			return
		}
		apiData.Data = data
		apiData.Id = dataId
		apiData.Type = dataType
		apiData.Parameter = parameter
		content = append(content, apiData)
	}

	models.SendData(w, content)
	return
}

func GetModules(w http.ResponseWriter, r *http.Request) {
	modules, err := models.GetAllModules()
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, modules)
	return
}

func GetModulesWithData(w http.ResponseWriter, r *http.Request) {
	modules, err := models.GetAllModules()
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}

	var md []models.ModuleAndData
	for _, mod := range modules {
		data, err := models.GetDataObserved(mod.Id)
		if err != nil {
			log.Println(err)
			models.SendNotFound(w)
			return
		}
		md = append(md, models.ModuleAndData{Module: mod, Data: data})
	}
	models.SendData(w, md)
	return
}
