package requestor

import (
	"shared"
	"Middleware/Distribution/marshaller"
	"Middleware/Infrastructure/crh"
	"Middleware/Distribution/miop"
	"Middleware/aux"
)

type Requestor struct{}

func (Requestor) Invoke(inv aux.Invocation) interface{} {
	marshallerInst := marshaller.Marshaller{}
	crhInst := crh.CRH{ServerHost:inv.Host,ServerPort:inv.Port}

	// Cria a mensagem a ser transmitida
	reqHeader := miop.RequestHeader{Context:"Context",RequestId:1000,ResponseExpected:true, ObjectKey:inv.id, Operation:inv.Request.Op}
	reqBody := miop.RequestBody{Body:inv.Request.Params}
	header := miop.Header{Magic:"MIOP",Version:"1.0",ByteOrder:true,MessageType:shared.MIOP_REQUEST}
	body := miop.Body{ReqHeader:reqHeader,ReqBody:reqBody}
	miopPacketRequest := miop.Packet{Hdr:header,Bd:body}

	// Serializa a mensagem de request
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)

	// Envia a mensagem e Recebe a mensagem de resposta
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes)
	
	// Deserializa a mensagem de resposta
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	// Envia a respota ao client proxy
	r := miopPacketReply.Bd.RepBody.OperationResult

	return r
}




