package subscribermanager

type Subscriber struct {
	Host string
	Port int
}

type SubscriberManager struct {
	queue map[string][]Subscriber // Pelo nome da fila sei quem tá inscrito
}

// NewSubscriber ...
func NewSubscriber(host string, port int) Subscriber {
	sb := new(Subscriber)
	sb.Host = host
	sb.Port = port
	return *sb
}

func NewSubscriberManager() SubscriberManager {
	sm := new(SubscriberManager)
	sm.queue = make(map[string][]Subscriber, 0)
	return *sm
}

func (sm *SubscriberManager) Save(queue string, sb Subscriber) {
	sm.queue[queue] = append(sm.queue[queue], sb)
}

func (sm *SubscriberManager) Remove(queue string, sb Subscriber) {
	for i := 0; i < len(sm.queue[queue]); i++ {
		if sm.queue[queue][i].Host == sb.Host && sm.queue[queue][i].Port == sb.Port {
			sm.queue[queue][i] = sm.queue[queue][len(sm.queue[queue])-1]
			sm.queue[queue] = sm.queue[queue][:len(sm.queue[queue])-1]
		}
	}
}

func (sm *SubscriberManager) Manager() {

}
