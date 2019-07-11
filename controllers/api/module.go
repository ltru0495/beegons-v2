package api

import (
	"log"
	"net/http"

	"github.com/beegons/models"
	"github.com/beegons/utils"
	"github.com/gorilla/mux"
)

func ModuleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := "urn:ngsi-ld:Module:" + vars["id"]
	module, err := utils.GetEntity(id)
	if err != nil {
		log.Println(err)
		models.SendNotFound(w)
		return
	}
	models.SendData(w, module)
	return
}
