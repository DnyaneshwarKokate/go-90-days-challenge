package main

import (
	"fmt"

	"day89/portfolio"
)

func main() {
	fmt.Println("=== Day 89: Developer Portfolio, Project Showcase & Repository CLI ===")

	pf := portfolio.NewChallengePortfolio("Dnyaneshwar Kokate", "High-Performance Go Backend Engineer")

	fmt.Println("\n--- 🚀 Go 90-Days Challenge Portfolio Statistics ---")
	fmt.Println(pf.SummaryReport())

	fmt.Println("\n--- 🛠️ Mastered Tech Stack Highlights ---")
	for i, tech := range pf.TechStack() {
		fmt.Printf("  %2d. %s\n", i+1, tech)
	}
}
