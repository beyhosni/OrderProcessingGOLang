package main

import (
	"log"
	"net/http"
	"order-processing-app-go/consumer"
	"order-processing-app-go/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", controller.CreateOrder).Methods("POST")

	go consumer.StartConsumer() // Démarre le consommateur en arrière-plan

	log.Println("Serveur démarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
