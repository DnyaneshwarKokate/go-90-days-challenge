package bench

import (
	"fmt"
	"strings"
)

func ConcatPlus(items []string) string {
	var s string
	for _, item := range items {
		s += item
	}
	return s
}

func ConcatSprintf(items []string) string {
	var s string
	for _, item := range items {
		s = fmt.Sprintf("%s%s", s, item)
	}
	return s
}

func ConcatBuilder(items []string) string {
	var builder strings.Builder
	// Pre-allocate memory capacity to eliminate re-allocations
	totalLen := 0
	for _, item := range items {
		totalLen += len(item)
	}
	builder.Grow(totalLen)

	for _, item := range items {
		builder.WriteString(item)
	}
	return builder.String()
}
