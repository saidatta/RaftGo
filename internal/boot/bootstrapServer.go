package boot

import (
	"fmt"
	"github.com/saidatta/RaftGo.git/internal/conf"
	"os"
	"runtime"
)

type BootstrapServer struct{}

func (bs *BootstrapServer) Start(args []string, stateMachine raft.StateMachine) {
	if len(args) != 2 {
		fmt.Println("Usage: ")
		fmt.Println("\tjava [bootstrapClass] [nodeId] [clusterNodes...]")
		fmt.Println("Example: java org.raft.BootStrap 127.0.0.1:2001 127.0.0.1:2001,127.0.0.1:2002,127.0.0.1:2003")
		os.Exit(1)
	}

	raftConfig := new(conf.RaftConfig)
	raftConfig.NodeId = args[0]
	raftConfig.ClusterNodes = args[1]
	raftNodeServer := raft.NewNodeServer(raftConfig.NodeId, stateMachine)
	raftNodeServer.InitConfig(raftConfig)

	runtime.SetFinalizer(raftNodeServer, func(raftNodeServer *raft.NodeServer) {
		raftNodeServer.Close()
	})
	raftNodeServer.Start()
}
