package kv

import "fmt"

type KvOpType int

const (
	GET KvOpType = iota
	SET
	DEL
)

type KvCommand struct {
	Key    string
	OpType KvOpType
	Value  string
}

func NewKvCommand(opType KvOpType, key, value string) *KvCommand {
	return &KvCommand{
		Key:    key,
		OpType: opType,
		Value:  value,
	}
}

func BuildGet(key string) *KvCommand {
	return NewKvCommand(GET, key, "")
}

func BuildSet(key, value string) *KvCommand {
	return NewKvCommand(SET, key, value)
}

func BuildDel(key string) *KvCommand {
	return NewKvCommand(DEL, key, "")
}

func (c *KvCommand) String() string {
	return fmt.Sprintf("%d:%s=>%s", c.OpType, c.Key, c.Value)
}
