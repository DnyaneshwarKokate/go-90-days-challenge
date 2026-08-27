package memory_test

import (
	"testing"

	"day86/memory"
)

func TestStackVsHeapAllocations(t *testing.T) {
	val := memory.StackAllocation(21)
	if val != 42 {
		t.Fatalf("Expected 42, got %d", val)
	}

	escaped := memory.HeapEscapedAllocation(999)
	if escaped == nil || escaped.ID != 999 {
		t.Fatalf("Heap allocation failed: %+v", escaped)
	}

	heapAlloc, _, _ := memory.GetGCMetrics()
	if heapAlloc == 0 {
		t.Fatalf("Expected positive heap allocation count")
	}
}
