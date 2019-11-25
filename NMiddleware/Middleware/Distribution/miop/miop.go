package miop

type Packet struct {
	Hdr Header
	Bd  Body
}

type Header struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        int
}

type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

type RequestHeader struct {
	Context          string
	RequestId        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type RequestBody struct {
	Body []interface{}
}

type ReplyHeader struct {
	Context   string
	RequestId int
	Status    int
}

type ReplyBody struct {
	OperationResult interface{}
}

// ------------------------------

type Message struct {
	HeaderMsg MessageHeader
	BodyMsg   MessageBody
}

type MessageHeader struct {
	// Nome da fila onde a mensagem será armazenada
	Destination string
}

type MessageBody struct {
	Body string
}
// Mensagem ficará dentro de um packet

type RequestPacket struct {
	PacketHeader RequestPacketHeader
	PacketBody   RequestPacketBody
}

type RequestPacketHeader struct {
	Operation string
}

type RequestPacketBody struct {
	Parameters []interface{}
	Message Message
}
// Enquanto message é usada para comunicação entre as aplicações (produtor/consumidor)
//, packet é usado para comunicação cliente/servidor dos consumidores e produtores (clientes)
// com o servidor de fila (servidor). 
