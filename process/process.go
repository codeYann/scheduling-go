package process

import (
	"math/rand"
)

type Pid int

type Process struct {
	processId Pid
}

func (p *Process) CreateProcess() *Process {
	return &Process{
		processId: Pid(rand.Int()),
	}
}

func (p Process) GetPid() Pid {
	return p.processId
}
