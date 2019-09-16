package main

import (
	"github.com/streadway/amqp"
	"encoding/json"
	"fmt"
	"log"
)

const SAMPLE_SIZE = 1000

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
	// Conecta ao servidor de mensageria
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err,"Não foi possível se conectar ao servidor de mensageria")
	defer connection.Close()

	// Cria o canal
	ch, err := connection.Channel()
	checkError(err,"Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// Declara as filas
	requestQueue, err := ch.QueueDeclare(
		"request", false, false, false, false, nil, )
	checkError(err,"Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare(
		"response", false, false, false, false, nil, )
	checkError(err,"Não foi possível criar a fila no servidor de mensageria")

	// Cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil, )
	checkError(err,"Falha ao registrar o consumidor servidor de mensageria")

	//start := time.Now()
	for i := 0; i < SAMPLE_SIZE; i++{

		//t1 := time.Now()

		// prepara request
		msgRequest := Request{Header:"Request",RequestNumber:i}
		msgRequestBytes,err := json.Marshal(msgRequest)
		checkError(err,"Falha ao serializar a mensagem")

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes,})
		checkError(err,"Falha ao enviar a mensagem para o servidor de mensageria")

		// recebe resposta
		msgRet := <- msgsFromServer

		fmt.Println(string(msgRet.Body))
		//t2 := time.Now()
		//x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		//fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}
