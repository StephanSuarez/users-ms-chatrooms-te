package http

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func SubcribersHandlers(UsersDep *UsersDependencies) {
	// Canal para manejar errores
	errChan := make(chan error)

	// Ejecutar cada suscriptor en su propia goroutine
	go func() {
		errChan <- createUserSubs(os.Stdout, UsersDep)
	}()
	go func() {
		errChan <- deleteUserSubs(os.Stdout, UsersDep)
	}()
	go func() {
		errChan <- seeMsgAmqp()
	}()

	// Manejar errores de las goroutines
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			log.Fatal(err)
		}
	}
}

type Message struct {
	Content string `json:"content"`
}

func seeMsgAmqp() error {
	conn, err := amqp.Dial("amqps://uaixaaut:4Hy7WEJGlEDZUEvDHXL91fSXEiv5gmU9@shrimp.rmq.cloudamqp.com/uaixaaut")
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir el canal: %v", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"microservice_queue", // Nombre de la cola
		false,                // Durabilidad
		false,                // Eliminar cuando no hay consumidores
		false,                // Exclusiva
		false,                // No espera en colas
		nil,                  // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %v", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // Nombre de la cola
		"",     // Consumidor
		true,   // Auto-ack
		false,  // Exclusiva
		false,  // No espera en colas
		false,  // No espera en colas
		nil,    // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var msg Message
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Fatalf("Error al decodificar el mensaje: %v", err)
			}
			log.Printf("Recibido mensaje: %s", msg.Content)

			// Enviar respuesta al API Gateway
			response := Message{Content: "Mensaje recibido correctamente INTO USERS - MS"}
			responseBody, err := json.Marshal(response)
			if err != nil {
				log.Fatalf("Error al serializar la respuesta: %v", err)
			}
			err = ch.Publish(
				"",               // Intercambio
				"response_queue", // Nombre de la cola de respuesta (puedes elegir el nombre que desees)
				false,            // Mandar asincrónicamente
				false,            // Publicar al servidor si no hay consumidores
				amqp.Publishing{
					ContentType: "application/json",
					Body:        responseBody,
				})
			if err != nil {
				log.Fatalf("Error al publicar la respuesta: %v", err)
			}
		}
	}()

	log.Println("Microservicio esperando mensajes...")
	<-forever
	return nil
}
