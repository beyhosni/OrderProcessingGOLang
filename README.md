# OrderProcessingGOLang
OrderProcessingGOLang


Démarrer l'application
Lancer RabbitMQ sur votre machine.

Démarrer l'application Go :

bash
Copy code
go run main.go
Envoyer une requête POST pour créer une commande :
bash
Copy code
# Send a POST request to the local server to create a new order

## Request

### URL

http://localhost:8080/orders

### Method

POST

### Headers

Content-Type: application/json

### Body

```json
{
    "orderId": "1001",
    "customerName": "Hosni",
    "product": "Bey",
    "quantity": 1
}

