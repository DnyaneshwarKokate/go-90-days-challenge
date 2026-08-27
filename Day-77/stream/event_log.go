package stream

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// LogRecord represents an append-only commit log record.
type LogRecord struct {
	Offset    int64
	Key       string
	Value     string
	Timestamp time.Time
}

// EventLog represents an in-memory event stream log with compaction support.
type EventLog struct {
	mu          sync.RWMutex
	nextOffset  int64
	records     []LogRecord
}

// NewEventLog initializes a commit log.
func NewEventLog() *EventLog {
	return &EventLog{
		records: make([]LogRecord, 0),
	}
}

// Append appends a record to the log stream and assigns an auto-incrementing offset.
func (l *EventLog) Append(key string, value string) LogRecord {
	l.mu.Lock()
	defer l.mu.Unlock()

	offset := atomic.AddInt64(&l.nextOffset, 1) - 1
	record := LogRecord{
		Offset:    offset,
		Key:       key,
		Value:     value,
		Timestamp: time.Now(),
	}

	l.records = append(l.records, record)
	return record
}

// ReadFromOffset retrieves up to limit records starting from startOffset.
func (l *EventLog) ReadFromOffset(startOffset int64, limit int) ([]LogRecord, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if startOffset < 0 {
		return nil, errors.New("invalid negative offset")
	}

	results := make([]LogRecord, 0)
	for _, rec := range l.records {
		if rec.Offset >= startOffset {
			results = append(results, rec)
			if len(results) >= limit {
				break
			}
		}
	}
	return results, nil
}

// Compact retains only the most recent LogRecord for each unique Key.
func (l *EventLog) Compact() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	latestByKey := make(map[string]LogRecord)
	for _, rec := range l.records {
		latestByKey[rec.Key] = rec
	}

	compacted := make([]LogRecord, 0, len(latestByKey))
	for _, rec := range l.records {
		if latest, exists := latestByKey[rec.Key]; exists && latest.Offset == rec.Offset {
			compacted = append(compacted, rec)
		}
	}

	removedCount := len(l.records) - len(compacted)
	l.records = compacted
	return removedCount
}
