package queue

type Queue struct {
	Queue     []string
	QueueSize int
}

func NewQueue() Queue {
	queue := new(Queue)
	queue.QueueSize = 0

	return *queue
}

func (q *Queue) Enqueue(msg string) {
	q.Queue = append(q.Queue, msg)
	q.QueueSize++
}

func (q *Queue) Dequeue() string {
	msg := q.Queue[q.QueueSize-1]
	q.Queue = q.Queue[:q.QueueSize-1]
	q.QueueSize--

	return msg
}

func (q Queue) AllMessages() []string {
	msgs := make([]string, 0)
	for i := 0; i < q.QueueSize; i++ {
		msgs = append(msgs, q.Queue[i])
	}
	return msgs
}
