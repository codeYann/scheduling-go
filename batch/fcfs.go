package batch

import (
	"sync"

	"github.com/codeYann/scheduling-go/process"
)

// Create fcfs algorhtim for batch scheduling
func CreateFCFS() *FCFS {
	return &FCFS{
		queue: make([]process.Process, 0),
	}
}

// FCFS is a first come first serve algorithm for batch scheduling
type FCFS struct {
	queue []process.Process
	sync.Mutex
}

// Enqueue adds a process to the queue
func (q *FCFS) Enqueue(pcss *process.Process) {
	q.Lock()
	q.queue = append(q.queue, *pcss)
	q.Unlock()
}

// Dequeue removes the first process from the queue
func (q *FCFS) Dequeue() error {
	if len(q.queue) > 0 {
		q.Lock()
		q.queue = append(q.queue, q.queue[0])
		q.queue = q.queue[1:]
		q.Unlock()
	}
	return nil
}

// Front returns the first process in the queue
func (q *FCFS) Front() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.Lock()
		defer q.Unlock()
		return &q.queue[0], nil
	}
	return nil, nil
}

// Back returns the last process in the queue
func (q *FCFS) Back() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.Lock()
		defer q.Unlock()
		return &q.queue[len(q.queue)-1], nil
	}
	return nil, nil
}

// GetSize returns the size of the queue
func (q *FCFS) GetSize() int {
	return len(q.queue)
}

// GetQueue returns the queue
func (q *FCFS) GetQueue() []process.Process {
	return q.queue
}
