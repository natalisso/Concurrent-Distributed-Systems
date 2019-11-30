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
	clientConn         net.Conn
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

func (crh *ClientRequestHandler) Connection() {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", crh.hostToConn+":"+strconv.Itoa(crh.portToConn))
		if err == nil {
			//log.Fatalf("CRH erro no Dial:: %s", err)
			break
		}
	}
	crh.clientConn = conn
}

// Send Envia os bytes
func (crh *ClientRequestHandler) Send(msgToSend []byte) {

	// Manda o tamanho da mensagem para o servidor
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToSend))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	_, err := crh.clientConn.Write(sizeMsgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// Manda mensagem
	_, err = crh.clientConn.Write(msgToSend)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// Salvo a conexão para poder lê-la depois, caso necessário
	// if crh.expectedReply {
	// 	crh.clientSocket = conn
	//
}

// Receive recebe os bytes
func (crh *ClientRequestHandler) Receive() []byte {
	sizeMsgFromServer := make([]byte, 4)
	_, err := crh.clientConn.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = crh.clientConn.Read(msgFromServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	return msgFromServer
}
