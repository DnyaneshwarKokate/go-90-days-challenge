package saga

import (
	"context"
	"fmt"
	"sync"
)

// SagaStep represents a single operation in a distributed transaction,
// paired with its compensating action for rollback handling.
type SagaStep struct {
	Name       string
	Execute    func(ctx context.Context) error
	Compensate func(ctx context.Context) error
}

// SagaResult encapsulates execution metrics and status.
type SagaResult struct {
	SuccessfulSteps []string
	RolledBackSteps []string
	ExecutionError  error
	Success         bool
}

// SagaOrchestrator manages the sequential execution and rollback pipeline.
type SagaOrchestrator struct {
	mu    sync.Mutex
	steps []SagaStep
}

// NewSagaOrchestrator initializes a new orchestrator pipeline.
func NewSagaOrchestrator() *SagaOrchestrator {
	return &SagaOrchestrator{
		steps: make([]SagaStep, 0),
	}
}

// AddStep appends a step to the saga pipeline.
func (o *SagaOrchestrator) AddStep(step SagaStep) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.steps = append(o.steps, step)
}

// Execute runs all registered steps in order.
// If any step fails, compensation actions are triggered in reverse order.
func (o *SagaOrchestrator) Execute(ctx context.Context) SagaResult {
	o.mu.Lock()
	defer o.mu.Unlock()

	result := SagaResult{
		SuccessfulSteps: make([]string, 0),
		RolledBackSteps: make([]string, 0),
		Success:         true,
	}

	executedSteps := make([]SagaStep, 0)

	for _, step := range o.steps {
		select {
		case <-ctx.Done():
			result.Success = false
			result.ExecutionError = ctx.Err()
			o.rollback(ctx, executedSteps, &result)
			return result
		default:
		}

		fmt.Printf("[SAGA ORCHESTRATOR] Executing step: %s\n", step.Name)
		if err := step.Execute(ctx); err != nil {
			fmt.Printf("[SAGA ORCHESTRATOR] Step '%s' failed: %v. Initiating rollback...\n", step.Name, err)
			result.Success = false
			result.ExecutionError = fmt.Errorf("step %s failed: %w", step.Name, err)
			o.rollback(ctx, executedSteps, &result)
			return result
		}

		result.SuccessfulSteps = append(result.SuccessfulSteps, step.Name)
		executedSteps = append(executedSteps, step)
	}

	fmt.Println("[SAGA ORCHESTRATOR] Distributed saga transaction completed successfully.")
	return result
}

func (o *SagaOrchestrator) rollback(ctx context.Context, executed []SagaStep, result *SagaResult) {
	fmt.Println("[SAGA ORCHESTRATOR] Reversing executed steps (LIFO order)...")
	for i := len(executed) - 1; i >= 0; i-- {
		step := executed[i]
		fmt.Printf("[SAGA ORCHESTRATOR] Compensating step: %s\n", step.Name)
		if step.Compensate != nil {
			if err := step.Compensate(ctx); err != nil {
				fmt.Printf("[SAGA ORCHESTRATOR] ERROR: Compensation for step '%s' failed: %v\n", step.Name, err)
			} else {
				result.RolledBackSteps = append(result.RolledBackSteps, step.Name)
			}
		}
	}
}
