package server

type RoleType int

const (
	FOLLOWER RoleType = iota
	CANDIDATE
	LEADER
)
