package model

import "fmt"

type NodeEvent struct {
	NodeID    string
	EventType int
}

type NodeCommand struct {
	NodeEvent
}

type NodeEventType int

const (
	ADD = iota
	REMOVE
)

func BuildAdd(nodeID string) NodeCommand {
	return NodeCommand{
		NodeEvent{
			NodeID:    nodeID,
			EventType: ADD,
		},
	}
}

func BuildRemove(nodeID string) NodeCommand {
	return NodeCommand{
		NodeEvent{
			NodeID:    nodeID,
			EventType: REMOVE,
		},
	}
}

func (nc NodeCommand) String() string {
	return fmt.Sprintf("%d:%s", nc.EventType, nc.NodeID)
}
