package serverrequesthandler

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	serverHost   string
	serverPort   int
	serverSocket net.Conn
	ln           net.Listener
}

func NewServerRequestHandler(host string, port int) ServerRequestHandler {
	srh := new(ServerRequestHandler)
	srh.serverHost = host
	srh.serverPort = port

	return *srh
}

func (srh *ServerRequestHandler) Send(msgToClient []byte) {

	// Envia o tamanho da mensagem
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err := srh.serverSocket.Write(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// Envia mensagem
	_, err = srh.serverSocket.Write(msgToClient)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// Fecha conexão
	srh.serverSocket.Close()
	srh.ln.Close()
}

func (srh *ServerRequestHandler) Receive() []byte {

	// Cria o llistener
	var err error
	srh.ln, err = net.Listen("tcp", srh.serverHost+":"+strconv.Itoa(srh.serverPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// Aceita conexão
	srh.serverSocket, err = srh.ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// Recebe tamanho da mensagem
	size := make([]byte, 4)
	_, err = srh.serverSocket.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// Recebe mensagem
	msg := make([]byte, sizeInt)
	_, err = srh.serverSocket.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msg
}
