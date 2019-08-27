package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"math/rand"
	"time"
	"os"
	"sync"
)

var mutex = &sync.Mutex{}

type client struct {
    name string
	conn  net.Conn
	reader *bufio.Reader
}

func checkError(err error) {
	// Verificação de erros durante a troca de menssagens
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return 
	}
}

func receiveMessage(user *client) string {
	// Recebe o nome do client
	name, err := (*user).reader.ReadString('\n')
	checkError(err)

	(*user).name = strings.Replace(name, "\n", "", 1)

	return name
	// return "Hello " + (*user).name + "!\n"
}

func sendMessage(user *client, msg string){
	(*user).conn.Write([]byte(string(msg)))	
}

func handleConn(conn net.Conn){
	log.Printf("Serving %s\n", conn.RemoteAddr().String())
	user := client{"",conn,bufio.NewReader(conn)}

	// open output file
	nameDataBase := "dataBase" + conn.RemoteAddr().String() + ".txt"
    dataBase, err := os.Create(nameDataBase)
    if err != nil {
		panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
		if err := dataBase.Close(); err != nil {
			panic(err)
        }
		}()
		
		stp := true
		for stp{
			commandName, err := user.reader.ReadString(' ')
			checkError(err)
			
			fmt.Println("cmd = ",commandName)
			switch commandName {
			case "MSG ":
				data := receiveMessage(&user)
				t := time.Now().UTC()
				if _, err := dataBase.Write([]byte(t.Format("2006-01-02 15:04:05") + " -> " + data)); err != nil {
					panic(err)
				}
				logMsg := "Data Stored!\n"
				sendMessage(&user,logMsg)
			case "STOP ":
				stopSign := "STOP\n"
				sendMessage(&user,stopSign)
				conn.Close()
				stp = false
			default:
				errMsg := "Unknown command!\n"
				sendMessage(&user,errMsg)
				stp = false
			}
	}
	fmt.Println("out loop!")
	return
}

func main() {
	// Inicializa o servidor na porta 8080 e protocolo TCP 
	fmt.Println("Server waiting for connections...")
	port := ":8080"
	l, err1 := net.Listen("tcp", port)
	checkError(err1)

	// Fecha o socket na saida
	defer l.Close()

	// Lida com a questão da concorrência
	rand.Seed(time.Now().Unix())

	for {
		// Aceita conexões e inicia a rotina de tratamento (goroutine)
		conn, err2 := l.Accept()
		checkError(err2)

		// fmt.Println(STOPSERVER)
		go handleConn(conn)
	}
}