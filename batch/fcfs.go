package batch

import (
	"sync"

	"github.com/codeYann/scheduling-go/process"
)

type Qeue struct {
	queue []process.Process
	mutex sync.RWMutex
}

func CreateQueue() *Qeue {
	return &Qeue{
		queue: make([]process.Process, 0),
	}
}

func (q *Qeue) Enqueue(pcss *process.Process) {
	q.mutex.Lock()
	q.queue = append(q.queue, *pcss)
	q.mutex.Unlock()
}

func (q *Qeue) Dequeue() error {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		q.queue = append(q.queue, q.queue[0])
		q.queue = q.queue[1:]
		q.mutex.Unlock()
	}
	return nil
}

func (q *Qeue) Front() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		defer q.mutex.Unlock()
		return &q.queue[0], nil
	}
	return nil, nil
}

func (q *Qeue) Back() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		defer q.mutex.Unlock()
		return &q.queue[len(q.queue)-1], nil
	}
	return nil, nil
}

func (q *Qeue) Size() int {
	return len(q.queue)
}
