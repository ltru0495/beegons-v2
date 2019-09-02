package controllers

import (
	"github.com/beegons/models"
	"github.com/beegons/utils"

	"log"
	"net/http"
)

func ModuleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			module := new(models.Module)
			err := module.DecodeModuleForm(r)

			err = module.CreateModule()
			if err != nil {
				log.Println(err)
				models.SendUnprocessableEntity(w)
				return
			}

			err = module.CreateDataObserved()
			if err != nil {
				log.Println(err)
				log.Println("Error While Creating data observed")
				models.SendUnprocessableEntity(w)
				return
			}

			err = module.CreateCygnusSubscription()
			if err != nil {
				log.Println(err)
				log.Println("Error While creating subscription")
				models.SendUnprocessableEntity(w)
				return
			}

			err = module.CreateFlinkSubscription()
			if err != nil {
				log.Println(err)
				log.Println("Error While creating subscription")
				models.SendUnprocessableEntity(w)
				return
			}

			res := models.CreateDefaultResponse(w)
			res.Message = "Module has been created"
			res.Data = module
			res.Send()
			return
		}
	}
	utils.RenderTemplate(w, "module_create", nil)
}

func ModuleTable(w http.ResponseWriter, r *http.Request) {
	modules, err := models.GetAllModules()
	if err != nil {
		log.Println(err)
	}

	context := make(map[string]interface{})
	context["Modules"] = modules

	utils.RenderTemplate(w, "module_table", context)
}
