package brokerproxy

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/clientrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"fmt"
)

type BrokerProxy struct {
	queueName string
	crh       clientrequesthandler.ClientRequestHandler
	host      string // ESSE É DO PUBLISH/SUBSCRIBER
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
	fmt.Println("Connected to Server")
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
		fmt.Printf("Reply: %s\n", reply)
		if reply == "exchange created" {
			break
		}
	}
}

func (qmp *BrokerProxy) basic_Publish(nameExchange string, routingKey string, msg string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.BodyMsg.Body = msg

	packet.PacketHeader.Exchange_name = nameExchange
	packet.PacketHeader.Bind_keys = routingKey
	packet.PacketHeader.Operation = "publish"
	message.HeaderMsg.Life_Time = 60000 * 5 // Tempo em segundo
	packet.PacketBody.Message = *message

	for true {
		qmp.send(*packet)
		reply := qmp.receive()
		fmt.Printf("Reply: %s\n", reply)
		if reply == "publish received" {
			break
		}
	}
}

func (qmp *BrokerProxy) Queue_Declare(nameQueue string) {
	qmp.ConnectionBroker()
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.HeaderMsg.Destination_Queue = nameQueue
	packet.PacketBody.Message = *message
	packet.PacketHeader.Operation = "create_queue"

	for true {
		qmp.send(*packet)
		reply := qmp.receive()
		fmt.Printf("Reply: %s\n", reply)
		if reply == "queue created" {
			break
		}
	}
}

func (qmp *BrokerProxy) Queue_Bind(nameExchange string, nameQueue string, routingKey string) {
	qmp.ConnectionBroker() // essa vai ser salva pra ler dps
	packet := new(miop.RequestPacket)
	message := new(miop.Message)

	message.HeaderMsg.Destination_Queue = nameQueue

	packet.PacketHeader.Operation = "bind_queue"
	packet.PacketHeader.Exchange_name = nameExchange
	packet.PacketHeader.Bind_keys = routingKey
	packet.PacketBody.Message = *message

	for true {
		qmp.send(*packet)
		reply := qmp.receive()

		fmt.Printf("Reply: %s\n", reply)
		if reply == "queue binded" {
			break
		}
	}
}

func (qmp *BrokerProxy) Basic_Consume(nameQueue string) string {
	return qmp.receive()
}

// Cliente (produtor/ consumidor) está enviando uma mensagem pro serviço de mensageria
func (qmp *BrokerProxy) send(pckg miop.RequestPacket) {
	marshaller := new(marshaller.Marshaller)
	qmp.crh.Send(marshaller.Marshall(pckg))
}

// Cliente (produtor/ consumidor) está recebendo uma mensagem do serviço de mensageria
func (qmp *BrokerProxy) receive() string {
	//crh := clientrequesthandler.NewClientRequestHandler(shared.N_HOST, shared.NAMING_PORT, false)
	// srh := serverrequesthandler.NewServerRequestHandler(qmp.host, qmp.port)
	marshaller := new(marshaller.Marshaller)

	return marshaller.Unmarshall(qmp.crh.Receive()).PacketBody.Message.BodyMsg.Body
}
