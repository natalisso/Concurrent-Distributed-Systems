package marshaller

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
	"encoding/json"
	"log"
)

type Marshaller struct{}

// Serializador
func (Marshaller) Marshall(msg miop.RequestPacket) []byte {
	r, err := json.Marshal(msg)

	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

// Deserializador
func (Marshaller) Unmarshall(msg []byte) miop.RequestPacket {
	r := miop.RequestPacket{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}
