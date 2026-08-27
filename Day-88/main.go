package main

import (
	"fmt"

	"day88/interview"
)

func main() {
	fmt.Println("=== Day 88: Go Backend Engineering Interview Preparation & Systems Q&A ===")

	bank := interview.NewInterviewBank()

	fmt.Println("\n--- High-Frequency Go Technical Interview Questions ---")
	for _, q := range bank.GetAll() {
		fmt.Printf("\n[Q%d - %s]\n  Question: %s\n  Answer:   %s\n",
			q.ID, q.Topic, q.Question, q.Answer)
	}
}
