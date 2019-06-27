package controllers

import (
	// "fmt"
	// "log"
	"net/http"
	// "strings"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/schema"

	"github.com/beegons/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "index", nil)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user_create", nil)
}
