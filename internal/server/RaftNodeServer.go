package server

import "log"

type IService interface {
	start()
	close()
}

type RaftNodeServer struct {
	node          *RaftNode
	requestVote   IHandler
	appendEntries IHandler
	election      *ElectionService
	appendSender  *AppendEntriesSender
	httpServer    *HttpNettyServer
	nodeManager   *NodeManageService
}

func NewRaftNodeServer(nodeId string, stateMachine StateMachine) *RaftNodeServer {
	s := &RaftNodeServer{}
	s.node = NewRaftNode(nodeId, stateMachine)
	s.requestVote = NewRequestVoteHandler(s.node, s)
	s.appendEntries = NewAppendEntriesHandler(s.node, s)
	s.election = NewElectionService(s.node, s)
	s.appendSender = NewAppendEntriesSender(s.node, s)
	s.httpServer = NewHttpNettyServer(s)
	s.nodeManager = NewNodeManageService(s)
	return s
}

func (s *RaftNodeServer) Start() {
	log.Printf("Node %s start...", s.node.GetNodeId())
	s.election.Start()
	s.httpServer.Start()
}

func (s *RaftNodeServer) Close() {
	log.Printf("Node %s stop...", s.node.GetNodeId())
	s.election.Close()
	s.appendSender.Close()
	s.nodeManager.Close()
	s.httpServer.Destroy()
	s.node.SaveSnapshot()
}
