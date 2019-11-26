package queuemanager

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/marshaller"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/queue"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Infrastructure/clientrequesthandler"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
)

// Isso vai estar no serviço de mensageria
// QueueManager -> Gerenciador das filas
type QueueManager struct {
	Host   string
	port   int
	Queues map[string]queue.Queue
}

func NewQueueManager(sHost string, iPort int) QueueManager {
	qm := new(QueueManager)
	qm.Host = sHost
	qm.port = iPort
	qm.Queues = make(map[string]queue.Queue) // Inicialmente vazio

	return *qm
}

// o serviço de mensageria está enviando uma mensagem
// vai utilizar o server request handler
func (qm *QueueManager) Send(msg string) {

}

// serviço de mensageria está recebendo uma mensagem
// vai utilizar o client request handler
func (qm *QueueManager) Receive() {

	crh := clientrequesthandler.NewClientRequestHandler(shared.N_HOST, shared.NAMING_PORT)
	marshall := new(marshaller.Marshaller)

	rcvBytes := crh.Receive()
	packetRcv := marshall.Unmarshall(rcvBytes)

	nameQueue := packetRcv.PacketBody.Message.HeaderMsg.Destination

	// VERIFICAR SE FUNCIONA
	queueAux := qm.Queues[nameQueue]
	queueAux.Enqueue(packetRcv.PacketBody.Message)

}
