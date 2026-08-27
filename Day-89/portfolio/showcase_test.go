package portfolio_test

import (
	"testing"

	"day89/portfolio"
)

func TestChallengePortfolioSummary(t *testing.T) {
	pf := portfolio.NewChallengePortfolio("Dnyaneshwar Kokate", "Go Backend Engineer")

	report := pf.SummaryReport()
	if report == "" {
		t.Fatalf("Expected non-empty summary report")
	}

	stack := pf.TechStack()
	if len(stack) < 5 {
		t.Fatalf("Expected tech stack highlights")
	}
}
