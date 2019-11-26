package queuemanager

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/queue"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/serverrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/subscribermanager"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
)

// Isso vai estar no serviço de mensageria
// QueueManager -> Gerenciador das filas
type QueueManager struct {
	Host   string
	port   int
	Queues map[string]queue.Queue
	srh    serverrequesthandler.ServerRequestHandler
	sm     subscribermanager.SubscriberManager
}

func NewQueueManager(sHost string, iPort int) QueueManager {
	qm := new(QueueManager)
	qm.Host = sHost
	qm.port = iPort
	qm.Queues = make(map[string]queue.Queue) // Inicialmente vazio
	qm.srh = serverrequesthandler.NewServerRequestHandler(shared.N_HOST, shared.NAMING_PORT)
	qm.sm = subscribermanager.NewSubscriberManager()
	return *qm
}

// o serviço de mensageria está enviando uma mensagem
// vai utilizar o client request handler (vai se conectar com um dos extremos para enviar a mensagem)
func (qm *QueueManager) Send(msg string) {

}

// serviço de mensageria está recebendo uma mensagem
// vai utilizar o server request handler (sempre atv)
func (qm *QueueManager) Receive() {

	marshall := new(marshaller.Marshaller)

	rcvBytes := qm.srh.Receive()
	packetRcv := marshall.Unmarshall(rcvBytes)

	if packetRcv.PacketHeader.Operation == "send" {
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination

		// VERIFICAR SE FUNCIONA
		queueAux := qm.Queues[nameQueue]
		queueAux.Enqueue(packetRcv.PacketBody.Message.BodyMsg.Body)
	} else if packetRcv.PacketHeader.Operation == "subscribe" {
		nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination
		host := packetRcv.PacketBody.Parameters[0] // IP string
		port := packetRcv.PacketBody.Parameters[1] // PORT int

		sb := subscribermanager.NewSubscriber(host.(string), port.(int))

		qm.sm.Save(nameQueue, sb)
	}
}
