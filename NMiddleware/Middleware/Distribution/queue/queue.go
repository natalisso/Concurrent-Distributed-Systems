package queue

type Queue struct {
	Queue     []string
	QueueSize int
}

// NewQueue ...
func NewQueue() Queue {
	queue := new(Queue)
	queue.QueueSize = 0

	return *queue
}

// Enqueue ...
func (q *Queue) Enqueue(msg string) {
	q.Queue = append(q.Queue, msg)
	q.QueueSize++
}

// Dequeue ...
func (q *Queue) Dequeue() string {
	msg := q.Queue[q.QueueSize-1]
	q.Queue = q.Queue[:q.QueueSize-1]
	q.QueueSize--

	return msg
}
