package invoker

import (
	"App/impl"
	"Middleware/Distribution/marshaller"
	"Middleware/Distribution/miop"
	"Middleware/Infrastructure/srh"
	"shared"
	"Middleware/Distribution/lifecycleManagement"
)

type DataBankInvoker struct{}

func NewDataBankInvoker() DataBankInvoker {
	p := new(DataBankInvoker)

	return *p
}

func (DataBankInvoker) Invoke() {
	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: shared.CALCULATOR_PORT}
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)
	
	//var lifecycle lifecycleManagement

	
	// Cria o objeto remoto
	dataBankImpl := impl.DataBank{}

	// Loop de inversão de controle
	for {
		// Invoca o server request handler para receber uma mensagem com a invocação do cliente (request)
		rcvMsgBytes := srhImpl.Receive()

		// Deserializa a mensagem de request
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)

		// Acessa a operação requisitada pelo cliente na mensagem
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// Decide qual  método invocar (demultiplexação) e invoca-o no objeto remoto
		switch operation {
		case "Save":
			_p1 := string(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := string(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			_p3 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = dataBankImpl.Save(_p1, _p2, _p3)
		case "Search":
			_p1 := string(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			replParams[0] = dataBankImpl.Search(_p1)
		}

		// Recebe o resultado da invocação
		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// Serializa a mensagem de resposta
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// Invoca o server request handler pra enviar a mensagem de resposta ao cliente
		srhImpl.Send(msgToClientBytes)
	}
}
