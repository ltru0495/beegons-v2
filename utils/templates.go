package utils

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	rnd = renderer.New(
		renderer.Options{
			ParseGlobPattern: "./templates/*.html",
		},
	)
}

func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	log.Println("Rendering", filename)
	err := rnd.HTML(w, http.StatusOK, filename, data)
	if err != nil {
		log.Println(err)
		rnd.HTML(w, http.StatusOK, "error.html", nil)
	}
}
