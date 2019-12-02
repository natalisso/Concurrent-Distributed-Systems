package broker

import (
	"NMiddleware/Middleware/Distribution/exchange"
	"NMiddleware/Middleware/Distribution/marshaller"
	"NMiddleware/Middleware/Distribution/miop"
	"NMiddleware/Middleware/Distribution/queue"
	"NMiddleware/Middleware/Infrastructure/serverrequesthandler"
	"NMiddleware/Middleware/Infrastructure/subscribermanager"
	"NMiddleware/shared"
	"fmt"
)

type Broker struct {
	Host     string
	port     int
	Exchange map[string]exchange.Exchange
	Queues   map[string]queue.Queue
	sm       subscribermanager.SubscriberManager
	srh      serverrequesthandler.ServerRequestHandler
}

func NewBroker() Broker {
	qm := new(Broker)
	qm.Host = shared.N_HOST_MD
	qm.port = shared.N_PORT_MD
	qm.Queues = make(map[string]queue.Queue) // Inicialmente vazio
	qm.srh = serverrequesthandler.NewServerRequestHandler(shared.N_HOST_MD, shared.N_PORT_MD)
	qm.sm = subscribermanager.NewSubscriberManager()
	qm.Exchange = make(map[string]exchange.Exchange)
	return *qm
}

func (qm *Broker) send(pkg miop.RequestPacket, queueNames []string) {

	marshall := new(marshaller.Marshaller)

	for i := 0; i < len(queueNames); i++ {
		subscribers := qm.sm.SubscribersInQueue(queueNames[i])
		for j := 0; j < len(subscribers); j++ {
			qm.srh.Send(marshall.Marshall(pkg), subscribers[j].Conn, false)
		}
	}
}

func (qm *Broker) sendFor(subscriber subscribermanager.Subscriber, queueName string) {
	marshall := new(marshaller.Marshaller)

	allMessages := qm.Queues[queueName].AllMessages()
	var pkg miop.RequestPacket
	var msg miop.Message

	for i := 0; i < len(allMessages); i++ {
		msg.BodyMsg.Body = allMessages[i]
		pkg.PacketBody.Message = msg
		qm.srh.Send(marshall.Marshall(pkg), subscriber.Conn, false)
	}
}

func (qm *Broker) Manager() {
	for true {
		qm.Receive()
	}
}

func (qm *Broker) Receive() {

	// Recebe o pacote
	marshall := new(marshaller.Marshaller)
	rcvBytes, conn := qm.srh.Receive()
	packetRcv := marshall.Unmarshall(rcvBytes)

	if packetRcv.PacketHeader.Operation == "publish" {
		bindKey := packetRcv.PacketHeader.Bind_keys
		exchangeName := packetRcv.PacketHeader.Exchange_name

		queuesNames := qm.Exchange[exchangeName].FindQueues(bindKey)
		if len(queuesNames) > 0 { // Existe fila que encaixam com esse bindkey
			for i := 0; i < len(queuesNames); i++ {
				queueAux := qm.Queues[queuesNames[i]]
				queueAux.Enqueue(packetRcv.PacketBody.Message.BodyMsg.Body)
				qm.Queues[queuesNames[i]] = queueAux

			}
			qm.send(packetRcv, queuesNames)
		} else {
			fmt.Println("Discarted message!!!!")
		}

		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "publish received"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, true)

	} else if packetRcv.PacketHeader.Operation == "create_exchange" {
		if _, exist := qm.Exchange[packetRcv.PacketHeader.Exchange_name]; !exist {
			qm.Exchange[packetRcv.PacketHeader.Exchange_name] = exchange.NewExchange(packetRcv.PacketHeader.Exchange_type, packetRcv.PacketHeader.Exchange_durable)
		} else {
			fmt.Printf("Exchange: %s already exist!\n", packetRcv.PacketHeader.Exchange_name)
		}

		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "exchange created"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, true)

	} else if packetRcv.PacketHeader.Operation == "create_queue" {
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination_queue
		nExist := true
		for nQueue := range qm.Queues {
			if nQueue == nameQueue {
				nExist = false
			}
		}
		if nExist {
			qm.Queues[nameQueue] = queue.NewQueue()
		}

		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "queue created"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, true)

	} else if packetRcv.PacketHeader.Operation == "bind_queue" {
		nameExg := packetRcv.PacketHeader.Exchange_name
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination_queue
		bindKey := packetRcv.PacketHeader.Bind_keys

		if _, exist := qm.Exchange[nameExg]; exist {
			aux := qm.Exchange[nameExg]
			aux.Bind.BindQueue(nameQueue, bindKey)
			qm.Exchange[nameExg] = aux
			//	fmt.Println("binded queue")

			subs := subscribermanager.NewSubscriber(conn)
			qm.sm.SubscriberClient(nameQueue, subs)
		} else {
			fmt.Printf("Exchange: %s Doesn't exist!!!\n", nameExg)
		}

		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "queue binded"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, false)
	}
}
