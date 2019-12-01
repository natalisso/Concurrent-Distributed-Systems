package serverrequesthandler

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	serverHost string
	serverPort int
	//	ConnServer     []net.Conn
	ListenerServer net.Listener
}

func NewServerRequestHandler(host string, port int) ServerRequestHandler {
	srh := new(ServerRequestHandler)
	srh.serverHost = host
	srh.serverPort = port

	// Cria o listener
	var err error
	srh.ListenerServer, err = net.Listen("tcp", srh.serverHost+":"+strconv.Itoa(srh.serverPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	return *srh
}

func (srh *ServerRequestHandler) Send(msgToClient []byte, conn net.Conn, close bool) {

	// Envia o tamanho da mensagem
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err := conn.Write(size)
	if err != nil {
		log.Fatalf("SRH escrita:: %s", err)
	}

	// Envia mensagem
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("SRH escrita:: %s", err)
	}

	if close == true {
		conn.Close()
	}
}

func (srh *ServerRequestHandler) Receive() ([]byte, net.Conn) {

	// Aceita conexão
	conn, err := srh.ListenerServer.Accept()
	if err != nil {
		log.Fatalf("SRH leitura:: %s", err)
	}

	fmt.Println("Aceitei conexao")

	// Recebe tamanho da mensagem
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH leitura:: %s", err)
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// Recebe mensagem
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msg, conn
}

func (srh *ServerRequestHandler) CloseSRH() {
	srh.ListenerServer.Close()
}
