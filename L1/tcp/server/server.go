package main

import (
	"net"
	"bufio"
	"log"
	//"fmt"
	//"strings"
	"time"
)

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
	msg, err := (*user).reader.ReadString('\n')
	checkError(err)

	return msg
}

func sendMessage(user *client, msg string){
	t := time.Now().UTC()
	(*user).conn.Write([]byte(t.Format("2006,01,02") + "," + msg))	
}

func handleConn(conn net.Conn){
	// Cria um objeto do tipo client
	user := client{"",conn,bufio.NewReader(conn)}
	
	stp := true
	for stp{
		commandName, err := user.reader.ReadString(' ')
		checkError(err)
		
		switch commandName {
		case "MSG ":
			// Recebe um request e retorna o processamento
			msg := receiveMessage(&user)
			sendMessage(&user,msg)
		case "STOP ":
			// Recebe um request e retorna o processamento
			msg := receiveMessage(&user)
			sendMessage(&user,msg)
			// Fecha a conexão e encerra o hadle desse client
			conn.Close()
			stp = false
		default:
			errMsg := "Unknown command!\n"
			sendMessage(&user,errMsg)
			stp = false
		}
	}
	return
}

func main() {
	// Inicializa o servidor na porta 4093 e protocolo TCP 
	// fmt.Println("Server waiting for connections...")
	port := ":4093"
	l, err1 := net.Listen("tcp", port)
	checkError(err1)

	// Fecha o socket no final da execução
	defer l.Close()

	for {
		// Aceita conexões e inicia a rotina de tratamento (goroutine)
		conn, err2 := l.Accept()
		checkError(err2)

		// Lida com a requisição do client em uma goroutine
		go handleConn(conn)
	}
}
