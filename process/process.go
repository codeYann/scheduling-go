package process

import (
	"math/rand"
)

// Create Pid type
type Pid int

// Create process type
type Process struct {
	processId Pid
	time      int
}

// Create process function
func CreateProcess() *Process {
	return &Process{
		processId: Pid(rand.Intn(100)),
		time:      rand.Intn(100) / 1000,
	}
}

// GetProcessId function
func (p *Process) GetProcessId() Pid {
	return p.processId
}

// GetTime function
func (p *Process) GetTime() int {
	return p.time
}
