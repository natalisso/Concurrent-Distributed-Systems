package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}
var NUMCLIENTS = 5

type client struct {
	name   string
	conn   *net.UDPConn
	reader *bufio.Reader
}

func checkError(err error) {
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func getName(user *client) {
	fmt.Printf("Type your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	checkError(err)

	(*user).name = strings.Replace(name, "\n", "", 1)

}

func sendMessage(user *client, msg string) {

	_, err := (*user).conn.Write([]byte(msg))

	checkError(err)
}

func receiveMessage(user *client) {
	msgFromServer := make([]byte, 1024)
	(*user).conn.ReadFromUDP(msgFromServer)

	// Escrevendo a resposta do servidor no terminal
	//fmt.Printf(string(msgFromServer))
}

func runClient(clientName string, server string) {
	RemoteAddr, err := net.ResolveUDPAddr("udp", server)
	checkError(err)

	// Conectando ao servidor
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	checkError(err)
	//fmt.Printf("Connected to server at %s!\n", server)

	user := client{clientName, conn, bufio.NewReader(conn)}

	// Mandando o nome
	sendMessage(&user, "MSG "+user.name+"\n")
	receiveMessage(&user)

	x := float64(NUMCLIENTS)
	for i := 0; i <= 1E4; i++ {
		if i == 1E4 {
			sendMessage(&user, "STOP "+strconv.FormatFloat(x, 'f', 6, 64)+"\n")
			receiveMessage(&user)
		} else {
			time1 := time.Now()
			sendMessage(&user, "MSG "+strconv.FormatFloat(x, 'f', 6, 64)+"\n")
			receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E3
		}
	}
	wg.Done()
}

func main() {
	// Servidor na máquina local na porta 8080 (default)
	server := "localhost:8080"

	// Pego o numero de clients e o endereço ip e a porta do servidor caso tenham sido passados como argumento
	if len(os.Args) == 2 {
		NUMCLIENTS, _ = strconv.Atoi(os.Args[1])
	} else if len(os.Args) == 3 {
		NUMCLIENTS, _ = strconv.Atoi(os.Args[1])
		server = os.Args[2]
	}

	/// Inicializando as threads dos clients
	nome := "nome"
	wg.Add(NUMCLIENTS)
	for i := 0; i < NUMCLIENTS; i++ {
		go runClient(nome, server)
	}
	wg.Wait()
}
