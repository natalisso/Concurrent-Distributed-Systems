package marshaller

import (
	"encoding/json"
	"log"
	"Middleware/Distribution/miop"
)

type Marshaller struct{}

// Serializador
func (Marshaller) Marshall(msg miop.Packet) []byte {
	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

//  Deserializador
func (Marshaller) Unmarshall(msg []byte) miop.Packet {
	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}


