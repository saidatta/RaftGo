package fsm

import (
	"github.com/saidatta/RaftGo.git/internal/model"
)

type StateMachine interface {
	Apply(logEntry model.LogEntry)
	ApplyCommand(command model.Command)
	Get(key string) interface{}
	Set(key string, value interface{})
	Del(key string)
}

type nodeStateMachine struct{}

func (n *nodeStateMachine) Apply(logEntry model.LogEntry) {
	if _, ok := logEntry.Command.(model.NodeCommand); ok {
		NodeStateMachine.GetInstance().Apply(logEntry.Command)
	} else {
		n.ApplyCommand(logEntry.Command)
	}
}

func (n *nodeStateMachine) ApplyCommand(command Command) {
	// TODO: implement this method
}

func (n *nodeStateMachine) Get(key string) interface{} {
	// TODO: implement this method
	return nil
}

func (n *nodeStateMachine) Set(key string, value interface{}) {
	// TODO: implement this method
}

func (n *nodeStateMachine) Del(key string) {
	// TODO: implement this method
}

func NewStateMachine() StateMachine {
	return &nodeStateMachine{}
}
