package main

import (
	"encoding/json"
	"fmt"

	"day64/cqrs"
)

func main() {
	fmt.Println("=== Day 64: CQRS (Command Query Responsibility Segregation) ===")

	writeStore := cqrs.NewWriteStore()
	readStore := cqrs.NewReadStore()
	cqrsService := cqrs.NewCQRSService(writeStore, readStore)

	fmt.Println("\n--- Executing Write Commands (Commands) ---")
	// Create Account Command
	createCmd := cqrs.CreateAccountCommand{
		ID:      "ACC-1001",
		Owner:   "Eve Online",
		Balance: 1250.00,
	}

	if err := cqrsService.HandleCreateAccount(createCmd); err != nil {
		fmt.Printf("Error creating account: %v\n", err)
		return
	}
	fmt.Println("[COMMAND] Account ACC-1001 created with initial balance $1250.00")

	// Deposit Money Command 1
	deposit1 := cqrs.DepositMoneyCommand{ID: "ACC-1001", Amount: 300.00}
	_ = cqrsService.HandleDepositMoney(deposit1)
	fmt.Println("[COMMAND] Deposited $300.00 into ACC-1001")

	// Deposit Money Command 2
	deposit2 := cqrs.DepositMoneyCommand{ID: "ACC-1001", Amount: 450.00}
	_ = cqrsService.HandleDepositMoney(deposit2)
	fmt.Println("[COMMAND] Deposited $450.00 into ACC-1001")

	fmt.Println("\n--- Querying Optimized Read Store (Queries) ---")
	readView, err := cqrsService.GetAccountView("ACC-1001")
	if err != nil {
		fmt.Printf("Error querying read view: %v\n", err)
		return
	}

	jsonBytes, _ := json.MarshalIndent(readView, "", "  ")
	fmt.Printf("[QUERY RESULT] Formatted Read View Projection:\n%s\n", string(jsonBytes))
}
