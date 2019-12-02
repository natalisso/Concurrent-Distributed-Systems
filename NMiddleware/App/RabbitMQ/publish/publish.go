package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

/*IMPORTANTE:
Necessário inicializar o servidor de mensageria antes:
	sudo service rabbitmq-server start
Pra parar:
	sudo service rabbitmq-server stop
*/

type Request struct {
	Header        string
	RequestNumber int
}

func randFloats(min, max float64) float64 {
	res := min + rand.Float64()*(max-min)
	return res
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	// Abre arquivo de saida
	nameDataBase := "./dataBaseFanout" + ".csv"
	// fmt.Println("ARQ = ",nameDataBase)

	dataBase, err := os.Create(nameDataBase)
	if err != nil {
		panic(err)
	}

	defer dataBase.Close()

	if _, err := dataBase.Write([]byte("data\n")); err != nil {
		panic(err)
	}

	// Conecta ao servidor de mensageria
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err, "Não foi possível se conectar ao servidor de mensageria")
	defer connection.Close()

	// Cria o canal
	ch, err := connection.Channel()
	checkError(err, "Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()
	ch.ExchangeDeclare("Fanout-X", "fanout", false, false, false, false, nil)
	fmt.Scanln()

	// prepara request
	msgRequest := "Toggle\n"
	msgRequestBytes, err := json.Marshal(msgRequest)
	checkError(err, "Falha ao serializar a mensagem")

	// publica request

	for i := 0; i < 10000; i++ {
		t1 := time.Now()

		err = ch.Publish("Fanout-X", "Key1", false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})

		t2 := time.Now()
		checkError(err, "Falha ao enviar a mensagem para o servidor de mensageria")
		deltaTime := float64(t2.Sub(t1).Nanoseconds()) / 1E6
		if _, err := dataBase.Write([]byte(strconv.FormatFloat(deltaTime, 'f', 6, 64) + "\n")); err != nil {
			panic(err)
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(randFloats(8, 12)) * time.Millisecond)
	}

}
