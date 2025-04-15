package logger

import (
	"encoding/json"
	"go-eth-auction/shared"
	"sync"
	"time"
)

const MaxLogCountPerEntity = 100

var globalLogger = NewLogger()

func Get() *Logger {
	return globalLogger
}

type logEntry struct {
	Name       string    `json:"name"`
	LogType    string    `json:"type"`
	LogMessage string    `json:"message"`
	At         time.Time `json:"at"`
}
type Logger struct {
	worker *shared.Worker
	lock   *sync.RWMutex
	events map[string][]*logEntry
	logPtr map[string]int
}

func NewLogger() *Logger {
	var l = &Logger{}
	l.events = make(map[string][]*logEntry)
	l.worker = shared.NewWorkerManager(100, 1, nil)
	l.lock = &sync.RWMutex{}

	l.logPtr = make(map[string]int)
	go l.worker.Start()

	return l
}

func (l *Logger) Append(name, logType, logMessage string) {
	entry := &logEntry{
		Name:       name,
		LogType:    logType,
		LogMessage: logMessage,
		At:         time.Now(),
	}
	l.worker.SetWorker(l.logJob)
	l.worker.Enqueue(entry)
}

// this piece of code is not efficient for a production-level
// performant action
func (l *Logger) ListAsJson(name string) ([]byte, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	var logs = make([]*logEntry, 0)
	var ptr = l.getPtr(name, false) - 1
	if ptr < 0 {
		return []byte{}, nil
	}
	for i := ptr; i >= 0; i-- {
		logs = append(logs, l.events[name][i])
	}

	return json.Marshal(logs)
}

func (l *Logger) logJob(workerIndex int, job any) {
	var jobObj, _ = job.(*logEntry)

	l.lock.Lock()
	defer l.lock.Unlock()

	if _, ok := l.events[jobObj.Name]; !ok {
		l.events[jobObj.Name] = make([]*logEntry, MaxLogCountPerEntity)
	}

	var nextPtr = l.getPtr(jobObj.Name, true)
	l.events[jobObj.Name][nextPtr] = jobObj
}

func (l *Logger) getPtr(name string, increment bool) int {
	if v, ok := l.logPtr[name]; ok {
		if increment && v == MaxLogCountPerEntity-1 {
			l.logPtr[name] = 0
		} else if increment {
			l.logPtr[name]++
		}
		return v
	}
	if increment {
		l.logPtr[name] = 1
	}
	return 0
}
