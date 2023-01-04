package log

import (
	"fmt"
	"github.com/saidatta/RaftGo.git/internal/model"
	"github.com/saidatta/RaftGo.git/internal/util"
	"strings"
	"sync"
)

type LogModule struct {
	logs    []model.LogEntry
	logLock sync.RWMutex
}

func NewLogModule() *LogModule {
	return &LogModule{
		logs:    make([]model.LogEntry, 0),
		logLock: sync.RWMutex{},
	}
}

func (l *LogModule) indexOf(logIndex int) *model.LogEntry {
	l.logLock.RLock()
	defer l.logLock.RUnlock()

	if len(l.logs) > logIndex {
		return &l.logs[logIndex]
	}

	return nil
}

func (l *LogModule) RemoveFrom(fromIndex int) {
	l.logLock.Lock()
	defer l.logLock.Unlock()

	if len(l.logs) > fromIndex {
		l.logs = l.logs[:fromIndex]
	}
}

func (l *LogModule) AppendAtIndex(index int, entry model.LogEntry) {
	l.logLock.Lock()
	defer l.logLock.Unlock()

	l.logs = append(l.logs[:index], append([]model.LogEntry{entry}, l.logs[index:]...)...)
}

func (l *LogModule) Append(entry model.LogEntry) {
	l.logLock.Lock()
	defer l.logLock.Unlock()

	l.logs = append(l.logs, entry)
}

func (l *LogModule) LastLog() (int, model.LogEntry, error) {
	l.logLock.RLock()
	defer l.logLock.RUnlock()

	if len(l.logs) == 0 {
		return 0, model.LogEntry{}, fmt.Errorf("log is empty")
	} else {
		lastIndex := len(l.logs) - 1
		return lastIndex, l.logs[lastIndex], nil
	}
}

func (l *LogModule) LastLogIndex() int {
	l.logLock.RLock()
	defer l.logLock.RUnlock()

	return len(l.logs) - 1
}

func (l *LogModule) SubLogs(fromIndex, endIndex int) []model.LogEntry {
	l.logLock.RLock()
	defer l.logLock.RUnlock()

	return l.logs[fromIndex : endIndex+1]
}

func (l *LogModule) String() string {
	var b strings.Builder
	b.WriteString("Index\t")
	for i := 0; i < len(l.logs); i++ {
		fmt.Fprintf(&b, "%d\t", i)
	}
	b.WriteString(util.LINE_SEP)

	b.WriteString("Term\t")
	for i := 0; i < len(l.logs); i++ {
		fmt.Fprintf(&b, "%d\t", l.logs[i].Term)
	}
	b.WriteString(util.LINE_SEP)

	b.WriteString("Value\t")
	for i := 0; i < len(l.logs); i++ {
		fmt.Fprintf(&b, "%s\t", l.logs[i].Command)
	}
	b.WriteString(util.LINE_SEP)
	return b.String()
}
