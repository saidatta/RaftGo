package server

import (
	"github.com/saidatta/RaftGo.git/internal/fsm"
	"github.com/saidatta/RaftGo.git/internal/log"
	"sync"
	"sync/atomic"
)

type RaftNode struct {
	nodeId           string
	role             RoleType
	currentTerm      int64
	voteFor          *VoteFor
	logModule        *log.LogModule
	stateMachine     fsm.StateMachine
	commitIndex      int64
	lastApplied      int64
	nextIndex        map[string]int
	matchIndex       map[string]int
	leaderId         string
	lastRpcTimestamp int64
	logApplyExecutor *ThreadPoolExecutor
	voteLock         *sync.RWMutex
}

type VoteFor struct {
	nodeId string
	term   int
}

func NewVoteFor(nodeId string, term int) *VoteFor {
	return &VoteFor{
		nodeId: nodeId,
		term:   term,
	}
}

func NewRaftNode(nodeId string, stateMachine fsm.StateMachine) *RaftNode {
	s := &RaftNode{}
	s.nodeId = nodeId
	s.stateMachine = stateMachine
	s.role = FOLLOWER
	atomic.StoreInt64(&s.currentTerm, 0)
	s.logModule = log.NewLogModule()
	atomic.StoreInt64(&s.commitIndex, -1)
	atomic.StoreInt64(&s.lastApplied, -1)
	s.nextIndex = make(map[string]int)
	s.matchIndex = make(map[string]int)
	atomic.StoreInt64(&s.lastRpcTimestamp, 0)
	s.logApplyExecutor = NewThreadPoolExecutor(1, 1, 5, TimeUnitMINUTES, NewLinkedBlockingQueue(1024), NewNamedThreadFactory("RetryAppendEntriesTimer-"), NewCallerRunsPolicy())
	s.voteLock = &sync.RWMutex{}
	return s
}

func (s *RaftNode) CurrentTerm() int64 {
	return s.currentTerm
}

func (s *RaftNode) LastRpcTimestamp() int64 {
	return s.lastRpcTimestamp
}

func (s *RaftNode) VoteFor() *VoteFor {
	s.voteLock.RLock()
	defer s.voteLock.RUnlock()
	return s.voteFor
}

func (s *RaftNode) LogModule() *log.LogModule {
	return s.logModule
}

func (s *RaftNode) StateMachine() fsm.StateMachine {
	return s.stateMachine
}

func (s *RaftNode) CommitIndex() int64 {
	return s.commitIndex
}

func (s *RaftNode) LastApplied() int64 {
	return s.lastApplied
}

func (s *RaftNode) LeaderId() string {
	return s.leaderId
}

func (s *RaftNode) Role() RoleType {
	return s.role
}

func (s *RaftNode) NodeId() string {
	return s.nodeId
}

func (s *RaftNode) SetCurrentTerm(term int64) {
	atomic.StoreInt64(&s.currentTerm, term)
}

func (s *RaftNode) SetVoteFor(voteFor VoteFor) {
	s.voteLock.Lock()
	defer s.voteLock.Unlock()
	s.voteFor = voteFor
}

func (s *RaftNode) setCommitIndex(value int64) {
	atomic.StoreInt64(&s.commitIndex, value)
}

func (s *RaftNode) setLeader(leaderId string) {
	s.leaderId = leaderId
}

func (s *RaftNode) votefor(candidateId string, term int) bool {
	s.voteLock.Lock()
	defer s.voteLock.Unlock()

	success := s.canVoteFor(candidateId, term)

	if success {
		s.voteFor = NewVoteFor(candidateId, term)
	}
	return success
}

/**
CanVoteFor checks whether the Raft node can vote for the given candidate in the given term. It returns true if the node
has not voted before or if the given term is greater than the term of the previously voted candidate. It also returns
true if the given term is equal to the term of the previously voted candidate and the candidate ID is the same as the
previously voted candidate. In all other cases, it returns false. This function ensures that within the same term, a
node can only vote for one candidate.
*/

func (s *RaftNode) canVoteFor(candidateId string, candidateTerm int) bool {
	s.voteLock.RLock()
	defer s.voteLock.RUnlock()

	if s.voteFor == nil {
		return true
	} else if s.voteFor.term < candidateTerm {
		return true
	} else if s.voteFor.term == candidateTerm && s.voteFor.nodeId == candidateId {
		return true
	}

	return false
}
