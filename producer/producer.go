package producer

import (
	"encoding/json"
	"log"
	"order-processing-app-go/model"

	"github.com/streadway/amqp"
)

// Fonction pour envoyer une commande à RabbitMQ
func SendOrder(order model.Order) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Échec de la connexion à RabbitMQ : %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Échec de l'ouverture du canal : %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orderQueue", // nom de la file d'attente
		false,        // durable
		false,        // effacé si inactif
		false,        // exclusif
		false,        // pas d'attente
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Échec de la déclaration de la file d'attente : %s", err)
	}

	body, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Erreur lors de la sérialisation de la commande : %s", err)
	}

	err = ch.Publish(
		"",     // échange
		q.Name, // clé de routage
		false,  // obligatoire
		false,  // immédiat
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la commande : %s", err)
	}

	log.Printf("Commande envoyée : %s", order.OrderID)
}
