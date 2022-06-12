package interactive

import (
	"sync"

	"github.com/codeYann/scheduling-go/process"
)

// Round robin struct has a queue of processe, a mutex and quantum time
type RoundRobin struct {
	// Define queue
	queue []process.Process
	// Define mutex
	mutex sync.Mutex

	// Define quantum
	quantum int
}

// Create RoundRobin function
func CreateRoundRobin(quantum int) *RoundRobin {
	return &RoundRobin{
		queue:   make([]process.Process, 0),
		quantum: quantum,
	}
}

// Enqueue adds a process to the queue
func (q *RoundRobin) Enqueue(pcss *process.Process) {
	q.mutex.Lock()
	q.queue = append(q.queue, *pcss)
	q.mutex.Unlock()
}
