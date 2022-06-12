package batch

import (
	"sort"
	"sync"

	"github.com/codeYann/scheduling-go/process"
)

// Create SFJ struct with a queue of processes and a mutex
type SFJ struct {
	queue []process.Process
	mutex sync.Mutex
}

// Use Batch interface to implement SFJ enqueue
func (q *SFJ) Enqueue(pcss *process.Process) {
	q.mutex.Lock()
	q.queue = append(q.queue, *pcss)

	// Sort q.queue to put the process with the lowest time in the front
	sort.Slice(q.queue, func(i, j int) bool {
		return q.queue[i].GetTime() < q.queue[j].GetTime()
	})

	q.mutex.Unlock()
}

// Use Batch interface to implement SFJ dequeue
func (q *SFJ) Dequeue() error {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		q.queue = append(q.queue, q.queue[0])
		q.queue = q.queue[1:]
		q.mutex.Unlock()
	}
	return nil
}

// Use Batch interface to implement SFJ front
func (q *SFJ) Front() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		defer q.mutex.Unlock()
		return &q.queue[0], nil
	}
	return nil, nil
}

// Use Batch interface to implement SFJ back
func (q *SFJ) Back() (*process.Process, error) {
	if len(q.queue) > 0 {
		q.mutex.Lock()
		defer q.mutex.Unlock()
		return &q.queue[len(q.queue)-1], nil
	}
	return nil, nil
}

// Use Batch interface to implement SFJ get size
func (q *SFJ) GetSize() int {
	return len(q.queue)
}

// Create SFJ batch scheduler
func CreateSFJ() *SFJ {
	return &SFJ{
		queue: make([]process.Process, 0),
	}
}
