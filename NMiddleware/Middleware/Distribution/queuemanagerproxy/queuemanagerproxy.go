package queuemanagerproxy

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/clientrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/serverrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
)

type QueueManagerProxy struct {
	queueName string
	crh       clientrequesthandler.ClientRequestHandler
	host      string // ESSE É DO PUBLISH/SUBSCRIBER
	port      int
}

func NewQueueManagerProxy(qName string, perst bool, myHost string, myPort int) QueueManagerProxy {
	qmp := new(QueueManagerProxy)
	qmp.queueName = qName
	qmp.crh = clientrequesthandler.NewClientRequestHandler(shared.N_HOST, shared.NAMING_PORT, perst)
	qmp.host = myHost
	qmp.port = myPort

	return *qmp
}

// Cliente (produtor/ consumidor) está enviando uma mensagem pro serviço de mensageria
func (qmp *QueueManagerProxy) Send(m string, operation string) {
	marshaller := new(marshaller.Marshaller)
	message := new(miop.Message)
	packet := new(miop.RequestPacket)

	// Configurando a mensagem
	message.HeaderMsg.Destination = qmp.queueName
	message.BodyMsg.Body = m

	// Configurando o pacote
	rpb := new(miop.RequestPacketBody)
	rpb.Parameters = make([]interface{}, 0)
	rpb.Parameters = append(rpb.Parameters, qmp.host, qmp.port) // PARA O S.M CONSEGUIR ME ENVIAR A MENSAGEM CASO EU SEJA UM SUBSCRIBER

	rpb.Message = *message
	packet.PacketBody = *rpb
	packet.PacketHeader.Operation = operation

	qmp.crh.Send(marshaller.Marshall(*packet))
}

// Cliente (produtor/ consumidor) está recebendo uma mensagem do serviço de mensageria
func (qmp *QueueManagerProxy) Receive() string {
	//crh := clientrequesthandler.NewClientRequestHandler(shared.N_HOST, shared.NAMING_PORT, false)
	srh := serverrequesthandler.NewServerRequestHandler(qmp.host, qmp.port)
	marshaller := new(marshaller.Marshaller)

	return marshaller.Unmarshall(srh.Receive()).PacketBody.Message.BodyMsg.Body
}
