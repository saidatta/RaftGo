package main

import "github.com/saidatta/RaftGo.git/internal/boot"
import "github.com/saidatta/RaftGo.git/internal/kv"

func main() {
	//boot.BootstrapServer{}
	bootstrapServer := new(boot.BootstrapServer)
	bootstrapServer.Start([]string{"127.0.0.1:2001", "127.0.0.1:2001,127.0.0.1:2002,127.0.0.1:2003"}, kv.NewKvStateMachine())
}
