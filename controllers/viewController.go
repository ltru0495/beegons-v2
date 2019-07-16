package controllers

import (
	// "fmt"
	// "log"
	"net/http"
	// "strings"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/schema"
	// "encoding/json"
	"github.com/beegons/models"
	"github.com/beegons/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	utils.RenderTemplate(w, "index", context)
}

func ChartsRealTime(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})

	modules, err := models.GetAllModules()
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	context["Modules"] = modules
	utils.RenderTemplate(w, "charts_realtime", context)
}
