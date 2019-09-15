package main

import (
	//"fmt"
	"log"
	"net"

	// "strings"
	// "math/rand"
	// "os"
	"strings"

	//"sync"
	"time"
)

// var users []client

type client struct {
	name string
	conn net.Conn
	addr *net.UDPAddr
}

func checkError(err error) {
	// Verificação de erros durante a troca de menssagens
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
}

func receiveMessage(conn *net.UDPConn) (string, *net.UDPAddr) {
	// Recebe a mensagem do cliente
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	checkError(err)
	//fmt.Println("receiveMessage: ", string(buffer[:n]))

	return string(buffer[:n]), addr
}

func sendMessage(addr *net.UDPAddr, msg string, conn *net.UDPConn) {
	msgToClient := []byte(msg)
	_, err := conn.WriteToUDP(msgToClient, addr)
	checkError(err)
}

func handleConn(conn *net.UDPConn, msg string, addr *net.UDPAddr) {

	//var user client
	i := strings.IndexByte(msg, ' ')
	commandName := msg[:i]
	//fmt.Println("Comando separado", commandName)

	switch commandName {
	case "MSG":
		data := msg[i+1:]
		//fmt.Println("dado da mensagem:", data)
		// Escrevo no arquivo o que foi recebido junto com um formato de tempo
		t := time.Now().UTC()
		msgToClient := t.Format("2006,01,02") + "," + data

		sendMessage(addr, msgToClient, conn)

	case "STOP":
		data := msg[i+1:]
		//fmt.Println("dado da mensagem:", data)

		// Escrevendo no arquivo
		t := time.Now().UTC()
		msgToClient := t.Format("2006,01,02") + "," + data

		sendMessage(addr, msgToClient, conn)

	default:
		errMsg := "Unknown command!\n"
		sendMessage(addr, errMsg, conn)
	}

	return
}

func main() {
	// Inicializa o servidor na porta 4093 e protocolo UDP
	//fmt.Println("Server waiting for connections...")
	port := ":8080"

	service := "localhost" + port
	// Retorna um endereço de udp
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	// Se conecta ao endereço na rede nomeada. (udpAddr)
	// Conecta o servidor
	l, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	// Fecha o socket na saida
	defer l.Close()

	for {
		// ao receber uma nova mensagem cria uma thread para ministrar o que foi recebido
		msgFromClient, addr := receiveMessage(l)
		go handleConn(l, msgFromClient, addr)
	}
}
