package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	"github.com/beegons/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSAlert(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket: New client connected.")
	hub := utils.GetWSAlertHub()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &utils.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.WritePump()
}

func WSData(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket: New client connected.")
	hub := utils.GetWSDataHub()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &utils.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.WritePump()
}

func WSPS(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket: New client connected.")
	hub := utils.GetWSPSHub()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &utils.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.WritePump()
}
