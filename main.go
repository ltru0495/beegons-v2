package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beegons/config"
	"github.com/beegons/models"
	"github.com/beegons/routers"
	"github.com/beegons/utils"
)

func main() {
	utils.Init()

	router := routers.InitRoutes()
	config := config.New()
	models.ConnectToDB()
	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	data, _ := models.GetHistoricalData("asd")
	log.Println(data)

	log.Println("Starting server at " + address)
	log.Fatal(server.ListenAndServe())

}
