package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"order-processing-app-go/model"
	"order-processing-app-go/producer"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Printf("Erreur lors de la création de la commande : %s", err)
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	producer.SendOrder(order)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Commande envoyée avec succès",
	})
}
