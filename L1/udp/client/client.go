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

func receiveMessage(user *client) string {
	msgFromServer := make([]byte, 1024)
	(*user).conn.ReadFromUDP(msgFromServer)

	return string(msgFromServer)
	// Escrevendo a resposta do servidor no terminal
	//fmt.Printf(string(msgFromServer))
}

func runClient(clientName string, server string, i int) {
	RemoteAddr, err := net.ResolveUDPAddr("udp", server)
	checkError(err)

	// Conectando ao servidor
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	checkError(err)
	//fmt.Printf("Connected to server at %s!\n", server)

	user := client{clientName, conn, bufio.NewReader(conn)}

	// open output file
	nameDataBase := "./data_bases/dataBase" + clientName + "(" + strconv.Itoa(NUMCLIENTS) + ")"+ ".csv"
	dataBase, err := os.Create(nameDataBase)
	checkError(err)

	
	x := float64(NUMCLIENTS)
	for i := 0; i < 1E4; i++ {
		if i == 1E4 -1 {
			time1 := time.Now()
			sendMessage(&user, "STOP "+strconv.FormatFloat(x, 'f', 6, 64)+"\n")
			msgFromServer := receiveMessage(&user)
			// fmt.Println(msgFromServer)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6
			if _, err := dataBase.Write([]byte(msgFromServer)); err != nil {
				panic(err)
			}

		}else if i == 0{
			time1 := time.Now()
			sendMessage(&user, "MSG "+"data\n")
			msgFromServer := receiveMessage(&user)
			// fmt.Println(msgFromServer)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6

			if _, err := dataBase.Write([]byte(msgFromServer)); err != nil {
				panic(err)
			}
		} else {
			time1 := time.Now()
			sendMessage(&user, "MSG "+strconv.FormatFloat(x, 'f', 6, 64)+"\n")
			msgFromServer := receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6

			if _, err := dataBase.Write([]byte(msgFromServer)); err != nil {
				panic(err)
			}
		}
	}
	wg.Done()
}

func main() {
	// Servidor na máquina local na porta 4093 (default)
	server := "localhost:8080"

	// Pego o numero de clients e o endereço ip e a porta do servidor caso tenham sido passados como argumento
	if len(os.Args) == 2 {
		NUMCLIENTS, _ = strconv.Atoi(os.Args[1])
	} else if len(os.Args) == 3 {
		NUMCLIENTS, _ = strconv.Atoi(os.Args[1])
		server = os.Args[2]
	}

	/// Inicializando as threads dos clients
	wg.Add(NUMCLIENTS)
	for i := 0; i < NUMCLIENTS; i++ {
		nome := strconv.Itoa(i+1)
		go runClient(nome, server, i)
	}
	wg.Wait()
}
