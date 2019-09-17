package main

import (
	"github.com/streadway/amqp"
	"encoding/json"
	"strconv"
	// "fmt"
	"log"
)

type Request struct {
	Header string
	RequestNumber int
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Necessário inicializar o servidor de mensageria antes:
	// sudo service rabbitmq-server start
	// Pra parar:
	// sudo service rabbitmq-server stop

	// Conecta ao servidor de mensageria
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err,"Não foi possível se conectar ao servidor de mensageria")
	defer connection.Close()

	// Cria o canal
	ch, err := connection.Channel()
	checkError(err,"Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// Declaração de filas
	requestQueue, err := ch.QueueDeclare("request", false, false, false,
		false, nil, )
	checkError(err,"Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare("response", false, false, false,
		false, nil, )
	checkError(err,"Não foi possível criar a fila no servidor de mensageria")

	// Prepara o recebimento de mensagens do cliente
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil, )
	checkError(err,"Falha ao registrar o consumidor do servidor de mensageria")

	log.Println("Servidor pronto...")
	for d := range msgsFromClient {
		// Recebe request
		msgRequest := Request{}
		err := json.Unmarshal(d.Body, &msgRequest)
		checkError(err,"Falha ao desserializar a mensagem")

		// fmt.Println(msgRequest)

		// Processa request
		r := "Here is your answer for the request " + strconv.Itoa(msgRequest.RequestNumber) + "!"

		// Prepara resposta
		replyMsg := r
		replyMsgBytes, err := json.Marshal(replyMsg)
		checkError(err,"Não foi possível criar a fila no servidor de mensageria")
		if err != nil {
			log.Fatalf("%s: %s", "Falha ao serializar mensagem", err)
		}

		// Publica resposta
		err = ch.Publish("", replyQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes),})
		checkError(err,"Falha ao enviar a mensagem para o servidor de mensageria")
	}
}
