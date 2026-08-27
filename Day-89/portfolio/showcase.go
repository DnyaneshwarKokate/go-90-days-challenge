package portfolio

import "fmt"

type ChallengeStats struct {
	DeveloperName string
	TargetRole    string
	TotalDays     int
	CompletedDays int
	TechStack     []string
}

type ChallengePortfolio struct {
	stats ChallengeStats
}

func NewChallengePortfolio(developer string, role string) *ChallengePortfolio {
	return &ChallengePortfolio{
		stats: ChallengeStats{
			DeveloperName: developer,
			TargetRole:    role,
			TotalDays:     90,
			CompletedDays: 90,
			TechStack: []string{
				"Go (Golang)", "Clean Architecture", "Microservices", "gRPC",
				"PostgreSQL / GORM", "Redis", "Kafka", "Docker & Kubernetes",
				"Prometheus Metrics", "OpenTelemetry Tracing", "Saga & CQRS",
			},
		},
	}
}

func (p *ChallengePortfolio) SummaryReport() string {
	completionPct := (float64(p.stats.CompletedDays) / float64(p.stats.TotalDays)) * 100
	return fmt.Sprintf("Developer: %s | Target Role: %s | Challenge Progress: %d/%d (%.2f%%)",
		p.stats.DeveloperName, p.stats.TargetRole, p.stats.CompletedDays, p.stats.TotalDays, completionPct)
}

func (p *ChallengePortfolio) TechStack() []string {
	return p.stats.TechStack
}
