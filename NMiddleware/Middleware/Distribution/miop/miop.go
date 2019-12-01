package miop

// ------------------------------

type Message struct {
	HeaderMsg MessageHeader
	BodyMsg   MessageBody
}

type MessageHeader struct {
	Destination_queue string
	Life_time         int
	Content_type      string
}

type MessageBody struct {
	Body string
}

// Mensagem ficará dentro de um packet
type RequestPacket struct {
	PacketHeader RequestPacketHeader
	PacketBody   RequestPacketBody
}

//RequestPacketHeader ...
type RequestPacketHeader struct {
	Operation        string
	Exchange_name    string
	Exchange_type    string
	Exchange_durable bool
	Bind_keys        string
	Mandatory_flag   bool
}

type RequestPacketBody struct {
	Parameters []interface{}
	Message    Message
}

// Enquanto message é usada para comunicação entre as aplicações (produtor/consumidor)
//, packet é usado para comunicação cliente/servidor dos consumidores e produtores (clientes)
// com o servidor de fila (servidor).
