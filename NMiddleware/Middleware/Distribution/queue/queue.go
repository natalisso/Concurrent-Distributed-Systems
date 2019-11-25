package queue

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/miop"
)

type Queue struct {
	Queue     []miop.Message
	QueueSize int
}

// NewQueue ...
func NewQueue() Queue {
	queue := new(Queue)
	queue.QueueSize = 0

	return *queue
}

// Enqueue ...
func (q* Queue) Enqueue(msg miop.Message) {
	q.Queue = append(q.Queue, msg)
	q.QueueSize++
}

// Dequeue ...
func (q *Queue) Dequeue() miop.Message {
	msg := q.Queue[q.QueueSize-1]
	q.Queue = q.Queue[:q.QueueSize-1]
	q.QueueSize--

	return msg
}
