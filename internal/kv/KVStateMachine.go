package kv

import "github.com/saidatta/RaftGo.git/internal/model"

type KvStateMachine struct {
	kvMap map[string]string
}

func (k *KvStateMachine) Apply(command model.Command) {
	kvCommand := command.(*KvCommand)
	if kvCommand.OpType == SET {
		k.Set(kvCommand.Key, kvCommand.Value)
	} else if kvCommand.OpType == DEL {
		k.Del(kvCommand.Key)
	}
}

func (k *KvStateMachine) Get(key string) interface{} {
	return k.kvMap[key]
}

func (k *KvStateMachine) Set(key string, value interface{}) {
	k.kvMap[key] = value.(string)
}

func (k *KvStateMachine) Del(key string) {
	delete(k.kvMap, key)
}

func NewKvStateMachine() *KvStateMachine {
	return &KvStateMachine{kvMap: make(map[string]string)}
}

func main() {
	stateMachine := &KvStateMachine{kvMap: make(map[string]string)}
	stateMachine.Set("a", "1")
	stateMachine.Set("b", "2")
}
