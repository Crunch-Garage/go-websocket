package main

import (
	controller "Crunch-Garage/go-websocket/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go To websockets")
	handleRoutes()
}

func handleRoutes() {

	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/ws", controller.WsEndPoint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
