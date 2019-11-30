package subscribermanager

import "net"

type Subscriber struct {
	Conn net.Conn
	Addr net.Addr
}
type SubscriberManager struct {
	subsQueue map[string][]Subscriber // Pelo nome da fila sei quem tá inscrito
}

// NewSubscriber ...
func NewSubscriber(conn net.Conn) Subscriber {
	sb := new(Subscriber)
	sb.Conn = conn
	sb.Addr = conn.RemoteAddr()
	return *sb
}

func NewSubscriberManager() SubscriberManager {
	sm := new(SubscriberManager)
	sm.subsQueue = make(map[string][]Subscriber)
	return *sm
}

func (sm *SubscriberManager) SubscriberClient(queueName string, sb Subscriber) {
	sm.subsQueue[queueName] = append(sm.subsQueue[queueName], sb)
}

func (sm *SubscriberManager) Remove(queueName string, sb string) {
	for i := 0; i < len(sm.subsQueue[queueName]); i++ {
		if sm.subsQueue[queueName][i].Addr.String() == sb {
			sm.subsQueue[queueName][i] = sm.subsQueue[queueName][len(sm.subsQueue[queueName])-1]
			sm.subsQueue[queueName] = sm.subsQueue[queueName][:len(sm.subsQueue[queueName])-1]
		}
	}
}

func (sm SubscriberManager) SubscribersInQueue(queueName string) []Subscriber {
	return sm.subsQueue[queueName]
}

func (sm *SubscriberManager) Manager() {

}
