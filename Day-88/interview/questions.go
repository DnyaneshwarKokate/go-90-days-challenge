package interview

import "fmt"

type InterviewQuestion struct {
	ID        int
	Topic     string
	Question  string
	Answer    string
}

type InterviewBank struct {
	questions []InterviewQuestion
}

func NewInterviewBank() *InterviewBank {
	bank := &InterviewBank{
		questions: make([]InterviewQuestion, 0),
	}
	bank.loadQuestions()
	return bank
}

func (b *InterviewBank) loadQuestions() {
	b.questions = append(b.questions,
		InterviewQuestion{
			ID:       1,
			Topic:    "Go Scheduler (GMP Model)",
			Question: "Explain the Go GMP Scheduler model.",
			Answer:   "G (Goroutine), M (OS Thread), P (Logical Processor / Context). The scheduler maps N Goroutines to M OS Threads using P Processors (GOMAXPROCS) with work-stealing queues.",
		},
		InterviewQuestion{
			ID:       2,
			Topic:    "Memory & Garbage Collection",
			Question: "How does Go's Garbage Collector work?",
			Answer:   "Go uses a concurrent, tri-color mark-and-sweep garbage collector with minimal STW (Stop-The-World) pauses.",
		},
		InterviewQuestion{
			ID:       3,
			Topic:    "Channels & Concurrency",
			Question: "What happens when reading from or writing to a closed channel?",
			Answer:   "Reading from a closed channel yields the zero-value immediately with ok=false. Writing to or closing an already closed channel causes a runtime panic.",
		},
	)
}

func (b *InterviewBank) GetAll() []InterviewQuestion {
	return b.questions
}

func (b *InterviewBank) GetByID(id int) (*InterviewQuestion, error) {
	for _, q := range b.questions {
		if q.ID == id {
			return &q, nil
		}
	}
	return nil, fmt.Errorf("question ID %d not found", id)
}
