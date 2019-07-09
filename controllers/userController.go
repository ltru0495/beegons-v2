package controllers

import (
	"github.com/beegons/models"
	"github.com/beegons/utils"

	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
		} else {
			user := new(models.User)
			decoder := schema.NewDecoder()
			err = decoder.Decode(user, r.PostForm)
			log.Println(r.PostForm)

			if err != nil {
				log.Println(err)
			} else {
				err = user.Insert()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	utils.RenderTemplate(w, "user_create", nil)
}

func UserTable(w http.ResponseWriter, r *http.Request) {
	users, err := models.AllUsers()
	if err != nil {
		log.Println(err)
	}

	context := make(map[string]interface{})
	context["Users"] = users

	utils.RenderTemplate(w, "user_table", context)
}
