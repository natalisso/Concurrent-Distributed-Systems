package queuemanagerproxy

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/clientrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
)

type QueueManagerProxy struct {
	queueName string
	crh       clientrequesthandler
}

func NewQueueManagerProxy(qName string, perst bool) QueueManagerProxy {
	qmp := new(QueueManagerProxy)
	qmp.queueName = qName
	qmp.crh = clientrequesthandler.NewClientRequestHandler(hared.N_HOST, shared.NAMING_PORT, perst)

	return *qmp
}

func (qmp *QueueManagerProxy) Send(m string) {
	// Cliente (produtor/ consumidor) está enviando uma mensagem pro serviço de mensageria
	marshaller := new(marshaller.Marshaller)
	message := new(miop.Message)
	packet := new(miop.RequestPacket)

	// Configurando a mensagem
	message.HeaderMsg.Destination = qmp.queueName
	message.BodyMsg.Body = m

	// Configurando o pacote
	rpb := new(miop.RequestPacketBody)
	rpb.Parameters = make([]interface{}, 0)
	rpb.Message = message
	packet.PacketBody = rpb
	packet.PacketHeader.Operation = "send"

	crh.Send(marshaller.Marshall(*packet))
}

func (qmp *QueueManagerProxy) Receive() miop.RequestPacket {
	// Cliente (produtor/ consumidor) está recebendo uma mensagem do serviço de mensageria
	crh := clientrequesthandler.NewClientRequestHandler(shared.N_HOST, shared.NAMING_PORT, false)
	marshaller := new(marshaller.Marshaller)

	return marshaller.Unmarshall(crh.Receive()).PacketBody.Message
}
