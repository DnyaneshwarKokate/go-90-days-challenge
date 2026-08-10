package logger

import (
	"fmt"
	"log"
	"sync"

	"day-28/domain"
)

// ConsoleLogger implements domain.Logger using Go standard log library.
type ConsoleLogger struct {
	prefix string
}

// NewConsoleLogger returns a new ConsoleLogger instance.
func NewConsoleLogger(prefix string) domain.Logger {
	return &ConsoleLogger{prefix: prefix}
}

func (l *ConsoleLogger) Info(msg string, args ...interface{}) {
	formatted := fmt.Sprintf(msg, args...)
	log.Printf("[INFO] [%s] %s", l.prefix, formatted)
}

func (l *ConsoleLogger) Error(msg string, args ...interface{}) {
	formatted := fmt.Sprintf(msg, args...)
	log.Printf("[ERROR] [%s] %s", l.prefix, formatted)
}

// MockLogger implements domain.Logger and retains logs in memory for assertions during unit testing.
type MockLogger struct {
	mu   sync.Mutex
	Logs []string
}

// NewMockLogger returns a thread-safe MockLogger.
func NewMockLogger() *MockLogger {
	return &MockLogger{Logs: make([]string, 0)}
}

func (m *MockLogger) Info(msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Logs = append(m.Logs, fmt.Sprintf("[INFO] "+msg, args...))
}

func (m *MockLogger) Error(msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Logs = append(m.Logs, fmt.Sprintf("[ERROR] "+msg, args...))
}
