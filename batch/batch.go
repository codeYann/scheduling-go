package batch

import (
	"github.com/codeYann/scheduling-go/process"
)

// Create interface for batch scheduling algorithms
// with Enqueue, Dequeue, Front, Back, GetSize functions
type Batch interface {
	Enqueue(pcss *process.Process)
	Dequeue() error
	Front() (*process.Process, error)
	Back() (*process.Process, error)
	GetSize() int
}
