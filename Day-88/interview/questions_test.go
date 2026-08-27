package interview_test

import (
	"testing"

	"day88/interview"
)

func TestInterviewBankRetrieval(t *testing.T) {
	bank := interview.NewInterviewBank()

	all := bank.GetAll()
	if len(all) < 3 {
		t.Fatalf("Expected at least 3 questions in bank, got %d", len(all))
	}

	q1, err := bank.GetByID(1)
	if err != nil || q1.Topic != "Go Scheduler (GMP Model)" {
		t.Fatalf("Failed to retrieve question 1: %v", err)
	}
}
