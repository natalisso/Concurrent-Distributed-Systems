package broker

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/exchange"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/queue"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/serverrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/subscribermanager"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"fmt"
)

// Isso vai estar no serviço de mensageria
// Broker -> Gerenciador das filas
type Broker struct {
	Host     string
	port     int
	Exchange map[string]exchange.Exchange
	Queues   map[string]queue.Queue
	sm       subscribermanager.SubscriberManager
	srh      serverrequesthandler.ServerRequestHandler
}

// Vou ter varios Exchanges no broker, cada exchange tem seu proprio tipo
// e seus próprios binds

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

// o serviço de mensageria está enviando uma mensagem
// vai utilizar o client request handler (vai se conectar com um dos extremos para enviar a mensagem)
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

	// POSSÍVEL MUTEX AQUI
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

// serviço de mensageria está recebendo uma mensagem
// vai utilizar o server request handler (sempre atv)
func (qm *Broker) Receive() {

	// Recebe o pacote
	marshall := new(marshaller.Marshaller)
	//fmt.Println("Esperando mensagem")
	rcvBytes, conn := qm.srh.Receive()
	packetRcv := marshall.Unmarshall(rcvBytes)
	fmt.Printf("Chegou: %s; bindkey: %s\n", packetRcv.PacketBody.Message.BodyMsg.Body, packetRcv.PacketHeader.Bind_keys)

	if packetRcv.PacketHeader.Operation == "publish" {
		bindKey := packetRcv.PacketHeader.Bind_keys
		exchangeName := packetRcv.PacketHeader.Exchange_name

		queuesNames := qm.Exchange[exchangeName].FindQueues(bindKey)
		fmt.Printf("tamanho queuenames: %d\n", len(queuesNames))
		if len(queuesNames) > 0 { // Existe fila que encaixam com esse bindkey
			for i := 0; i < len(queuesNames); i++ {
				// BOTAR UM MUTEX AQUI!!!
				fmt.Printf("mensagens na fila %s, antes do enqueue: %d\n", queuesNames[i], len(qm.Queues[queuesNames[i]].Queue))
				queueAux := qm.Queues[queuesNames[i]]
				queueAux.Enqueue(packetRcv.PacketBody.Message.BodyMsg.Body)
				qm.Queues[queuesNames[i]] = queueAux

				fmt.Printf("mensagens na fila %s, dps do enqueue: %d\n", queuesNames[i], len(qm.Queues[queuesNames[i]].Queue))
			}
			fmt.Printf("VOU ENVIAR PRO CLIENTE\n")
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
			fmt.Printf("Created exchange: %s", packetRcv.PacketHeader.Exchange_name)
		} else {
			fmt.Printf("Exchange: %s already exist!\n", packetRcv.PacketHeader.Exchange_name)
		}

		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "exchange created"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, true)

	} else if packetRcv.PacketHeader.Operation == "create_queue" {
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.DestinationQueue
		nExist := true
		for nQueue := range qm.Queues {
			if nQueue == nameQueue {
				nExist = false
			}
		}
		if nExist {
			// MUTEX AQUI
			qm.Queues[nameQueue] = queue.NewQueue()
		}
		pckgReply := new(miop.RequestPacket)
		pckgReply.PacketBody.Message.BodyMsg.Body = "queue created"
		qm.srh.Send(marshall.Marshall(*pckgReply), conn, true)

	} else if packetRcv.PacketHeader.Operation == "bind_queue" {
		nameExg := packetRcv.PacketHeader.Exchange_name
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.DestinationQueue
		bindKey := packetRcv.PacketHeader.Bind_keys

		if _, exist := qm.Exchange[nameExg]; exist {
			aux := qm.Exchange[nameExg]
			aux.Bind.BindQueue(nameQueue, bindKey)
			qm.Exchange[nameExg] = aux
			fmt.Println("binded queue")

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
