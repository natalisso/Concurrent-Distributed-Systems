package clientrequesthandler

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

// ClientRequestHandler para produtor e consumidor
type ClientRequestHandler struct {
	hostToConn         string
	portToConn         int
	expectedReply      bool
	clientSocket       net.Conn
	sentMessageSize    int
	receiveMessageSize int
}

func NewClientRequestHandler(host string, port int, expected bool) ClientRequestHandler {
	crh := new(ClientRequestHandler)
	crh.hostToConn = host
	crh.portToConn = port
	crh.expectedReply = expected

	return *crh
}

// Send Envia os bytes
func (crh *ClientRequestHandler) Send(msgToSend []byte) {

	var conn net.Conn
	var err error

	conn, err = net.Dial("tcp", crh.hostToConn+":"+strconv.Itoa(crh.portToConn))
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// Manda o tamanho da mensagem para o servidor
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToSend))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	conn.Write(sizeMsgToServer)

	// Manda mensagem
	_, err = conn.Write(msgToSend)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// Salvo a conexão para poder lê-la depois, caso necessário
	// if crh.expectedReply {
	// 	crh.clientSocket = conn
	// } else {
	// Se não, fecho-a
	conn.Close()
	//	}

	return
}

// Receive recebe os bytes
func (crh *ClientRequestHandler) Receive() []byte {
	sizeMsgFromServer := make([]byte, 4)
	_, err := crh.clientSocket.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.clientSocket.Read(msgFromServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	return msgFromServer
}
