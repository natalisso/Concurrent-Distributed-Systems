package brokerproxy

import (
	"NMiddleware/Middleware/Distribution/marshaller"
	"NMiddleware/Middleware/Distribution/miop"
	"NMiddleware/Middleware/Infrastructure/clientrequesthandler"
	"NMiddleware/shared"
	//	"fmt"
)

type BrokerProxy struct {
	queueName string
	crh       clientrequesthandler.ClientRequestHandler
	host      string
	port      int
}

func NewBrokerProxy(qName string, perst bool, myHost string, myPort int) BrokerProxy {
	qmp := new(BrokerProxy)
	qmp.queueName = qName
	qmp.crh = clientrequesthandler.NewClientRequestHandler(shared.N_HOST_MD, shared.N_PORT_MD, perst)
	qmp.host = myHost
	qmp.port = myPort

	return *qmp
}

func (qmp *BrokerProxy) ConnectionBroker() {
	qmp.crh.Connection()
}

func (qmp *BrokerProxy) Exchange_Declare(nameExchange string, typeExchange string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	packet.PacketHeader.Operation = "create_exchange"
	packet.PacketHeader.Exchange_name = nameExchange
	packet.PacketHeader.Exchange_type = typeExchange
	packet.PacketBody.Message = *message

	for true {
		qmp.send(*packet)
		reply := qmp.receive()
		if reply == "exchange created" {
			break
		}
	}
}

func (qmp *BrokerProxy) Basic_Publish(nameExchange string, routingKey string, msg string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.BodyMsg.Body = msg

	packet.PacketHeader.Exchange_name = nameExchange
	packet.PacketHeader.Bind_keys = routingKey
	packet.PacketHeader.Operation = "publish"
	message.HeaderMsg.Life_time = 60000 * 5 // Tempo em segundo
	packet.PacketBody.Message = *message

	for true {
		qmp.send(*packet)
		reply := qmp.receive()
		if reply == "publish received" {
			break
		}
	}
}

func (qmp *BrokerProxy) Queue_Declare(nameQueue string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.HeaderMsg.Destination_queue = nameQueue
	packet.PacketBody.Message = *message
	packet.PacketHeader.Operation = "create_queue"

	for true {
		qmp.send(*packet)
		reply := qmp.receive()
		if reply == "queue created" {
			break
		}
	}
}

func (qmp *BrokerProxy) Queue_Bind(nameExchange string, nameQueue string, routingKey string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.HeaderMsg.Destination_queue = nameQueue

	packet.PacketHeader.Operation = "bind_queue"
	packet.PacketHeader.Exchange_name = nameExchange
	packet.PacketHeader.Bind_keys = routingKey
	packet.PacketBody.Message = *message

	for true {
		qmp.send(*packet)
		reply := qmp.receive()

		if reply == "queue binded" {
			break
		}
	}
}

func (qmp *BrokerProxy) Basic_Consume(nameQueue string) string {
	return qmp.receive()
}

func (qmp *BrokerProxy) send(pckg miop.RequestPacket) {
	marshaller := new(marshaller.Marshaller)
	qmp.crh.Send(marshaller.Marshall(pckg))
}

func (qmp *BrokerProxy) receive() string {
	marshaller := new(marshaller.Marshaller)
	return marshaller.Unmarshall(qmp.crh.Receive()).PacketBody.Message.BodyMsg.Body
}
