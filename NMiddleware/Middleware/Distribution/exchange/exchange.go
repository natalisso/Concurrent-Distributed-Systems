package exchange

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/bind"
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
)

type Exchange struct {
	Bind bind.Bind
}

func NewExchange() Exchange {
	ex := new(Exchange)
	ex.Bind = bind.NewBind()
	return *ex
}

func (ex *Exchange) findQueue(pkt miop.RequestPacket) []string {
	var queues []string
	// OBS: TEM QUE AJEITAR LOGO AS INFORMAÇÕES DO PACOTE E DAS MENSAGENS
	// TO USANDO OQ TÁ, MAS TEM QUE MUDAR DPS I GUESS
	if pkt.PacketHeader.Operation == "direct" {
		queues = append(queues, pkt.PacketBody.Message.HeaderMsg.Destination)
	} else if pkt.PacketHeader.Operation == "topic" {
		queues = append(queues, ex.Bind.SearchQueue(pkt.PacketBody.Message.HeaderMsg.Destination)...)
	}

	return queues
}
