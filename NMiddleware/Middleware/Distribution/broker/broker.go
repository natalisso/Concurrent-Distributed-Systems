package broker

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/exchange"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/queue"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/serverrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/subscribermanager"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
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

func NewBroker(sHost string, iPort int) Broker {
	qm := new(Broker)
	qm.Host = sHost
	qm.port = iPort
	qm.Queues = make(map[string]queue.Queue) // Inicialmente vazio
	qm.srh = serverrequesthandler.NewServerRequestHandler(shared.N_HOST, shared.NAMING_PORT)
	qm.sm = subscribermanager.NewSubscriberManager()
	qm.Exchange = make(map[string]exchange.Exchange)
	return *qm
}

// o serviço de mensageria está enviando uma mensagem
// vai utilizar o client request handler (vai se conectar com um dos extremos para enviar a mensagem)
func (qm *Broker) Send(msg string) {

}

// serviço de mensageria está recebendo uma mensagem
// vai utilizar o server request handler (sempre atv)
func (qm *Broker) Receive() {

	// Recebe o pacote
	marshall := new(marshaller.Marshaller)
	rcvBytes := qm.srh.Receive()
	packetRcv := marshall.Unmarshall(rcvBytes)

	if packetRcv.PacketHeader.Operation == "publish" {
		bindKey := packetRcv.PacketHeader.Bind_keys
		exchangeName := packetRcv.PacketHeader.Exchange_name

		queuesNames := qm.Exchange[exchangeName].FindQueues(bindKey)

		for i := 0; i < len(queuesNames); i++ {
			// BOTAR UM MUTEX AQUI!!!
			queueAux := qm.Queues[queuesNames[i]]
			queueAux.Enqueue(packetRcv.PacketBody.Message.BodyMsg.Body)
		}

	} else if packetRcv.PacketHeader.Operation == "createExchange" {
		qm.Exchange[packetRcv.PacketHeader.Exchange_name] = exchange.NewExchange(packetRcv.PacketHeader.Exchange_type, packetRcv.PacketHeader.Exchange_durable)
	}
}

// nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination
// host := packetRcv.PacketBody.Parameters[0] // IP string
// port := packetRcv.PacketBody.Parameters[1] // PORT int

// sb := subscribermanager.NewSubscriber(host.(string), port.(int))

// qm.sm.Save(nameQueue, sb)
