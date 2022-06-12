package interactive

import (
	"github.com/codeYann/scheduling-go/process"
)

// Define interactive interface for all interactive scheduling algorithms
// with Enqueue, Dequeue, Front, Back, GetSize functions
type Interactive interface {
	Enqueue(pcss *process.Process)
	Dequeue() error
	Front() (*process.Process, error)
	Back() (*process.Process, error)
	GetSize() int
}
