package consumer

import (
	"encoding/json"
	"log"
	"order-processing-app-go/model"

	"github.com/streadway/amqp"
)

func StartConsumer() {
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

	msgs, err := ch.Consume(
		q.Name, // file d'attente
		"",     // nom du consommateur
		true,   // auto-ack
		false,  // exclusif
		false,  // pas d'attente
		false,  // local
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Échec de la consommation de messages : %s", err)
	}

	go func() {
		for d := range msgs {
			var order model.Order
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("Erreur lors de la désérialisation de la commande : %s", err)
			}
			log.Printf("Commande reçue : %v", order)
		}
	}()
	log.Printf("En attente des messages...")
	select {}
}
