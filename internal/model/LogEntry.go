package model

type Command interface{}

type LogEntry struct {
	Term    int
	Command Command
}

type BaseCommand struct {
	Key string
}

func NewLogEntry(term int, command Command) *LogEntry {
	return &LogEntry{
		Term:    term,
		Command: command,
	}
}
