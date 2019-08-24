package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"math/rand"
	"time"
)


func handleConnection(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
	for {
		// recebe uma mensagem do cliente
		netData, erro := bufio.NewReader(connection).ReadString('\n')
		// em caso de erro, printa o erro e encerra o processamento
		if erro != nil {
			fmt.Println(erro)
			return
		}

		// fecha a conexão se o cliente enviou um "STOP"
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		// escreve no terminal a mensagem recebida
		fmt.Print("Mensagem recebida from ",connection.RemoteAddr().String(),":", string(netData))

		// converte a mensagem recebida para caixa alta
		result := strings.ToUpper(netData)

		// retorna a serealização da mensagem para o cliente
		connection.Write([]byte(string(result)))
	}
	// fecha a conexão
	connection.Close()
}

func main() {
	// inicializando o servidor na porta 8080 e protocolo TCP 
	fmt.Println("Servidor aguardando conexões...")
	PORT := ":8080"
	l, erro := net.Listen("tcp", PORT)
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer l.Close()

	// lida com a questão da concorrência
	rand.Seed(time.Now().Unix())

	for {
		// aceita conexões
		connection, erro := l.Accept()
		if erro != nil {
			fmt.Println(erro)
			return
		}
		// executa o tratamento de uma requisição via "goroutines"
		go handleConnection(connection)
	}
}

