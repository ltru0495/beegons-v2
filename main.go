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

	hub := utils.GetWSHub()
	go hub.Run()
	models.ConnectToDB()
	address := fmt.Sprintf(":%d", config.Server.Port)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Println("Starting server at " + address)
	log.Fatal(server.ListenAndServe())

}
