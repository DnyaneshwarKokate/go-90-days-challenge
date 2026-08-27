package stream_test

import (
	"testing"

	"day77/stream"
)

func TestEventLogAppendReadAndCompaction(t *testing.T) {
	logStore := stream.NewEventLog()

	logStore.Append("user-100", "CREATED: Alice")
	logStore.Append("user-100", "UPDATED: Alice Smith")
	logStore.Append("user-200", "CREATED: Bob")

	records, err := logStore.ReadFromOffset(0, 10)
	if err != nil || len(records) != 3 {
		t.Fatalf("Expected 3 records before compaction, got %d", len(records))
	}

	// Compact log
	removed := logStore.Compact()
	if removed != 1 {
		t.Fatalf("Expected 1 superseded record removed, got %d", removed)
	}

	compactedRecords, _ := logStore.ReadFromOffset(0, 10)
	if len(compactedRecords) != 2 {
		t.Fatalf("Expected 2 records after compaction, got %d", len(compactedRecords))
	}

	if compactedRecords[0].Value != "UPDATED: Alice Smith" {
		t.Fatalf("Expected latest value for user-100, got %s", compactedRecords[0].Value)
	}
}
