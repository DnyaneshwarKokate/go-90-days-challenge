package memory

import (
	"runtime"
	"time"
)

type BigDataStruct struct {
	ID        int64
	Buffer    [1024]byte
	CreatedAt time.Time
}

// StackAllocation stays on stack because it does not outlive stack frame.
func StackAllocation(val int) int {
	result := val * 2
	return result
}

// HeapEscapedAllocation escapes to heap because pointer is returned to caller.
func HeapEscapedAllocation(id int64) *BigDataStruct {
	data := &BigDataStruct{
		ID:        id,
		CreatedAt: time.Now(),
	}
	return data
}

// GetGCMetrics reads active runtime memory allocations and GC pause telemetry.
func GetGCMetrics() (heapAlloc uint64, numGC uint32, pauseNs uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc, m.NumGC, m.PauseTotalNs
}
