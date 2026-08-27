package cqrs_test

import (
	"testing"

	"day64/cqrs"
)

func TestCQRSWriteAndReadFlow(t *testing.T) {
	writeStore := cqrs.NewWriteStore()
	readStore := cqrs.NewReadStore()
	service := cqrs.NewCQRSService(writeStore, readStore)

	// 1. Send Write Command
	cmd := cqrs.CreateAccountCommand{
		ID:      "ACC-701",
		Owner:   "Diana",
		Balance: 500.00,
	}

	if err := service.HandleCreateAccount(cmd); err != nil {
		t.Fatalf("Failed to execute create account command: %v", err)
	}

	// 2. Query Read Store View
	view, err := service.GetAccountView("ACC-701")
	if err != nil {
		t.Fatalf("Failed to query read store: %v", err)
	}

	if view.Owner != "Diana" || view.CurrentBalance != 500.00 {
		t.Fatalf("Read view mismatch: %+v", view)
	}

	// 3. Send Deposit Command
	depositCmd := cqrs.DepositMoneyCommand{
		ID:     "ACC-701",
		Amount: 250.00,
	}

	if err := service.HandleDepositMoney(depositCmd); err != nil {
		t.Fatalf("Failed to deposit money: %v", err)
	}

	// 4. Verify Updated Read View
	updatedView, err := service.GetAccountView("ACC-701")
	if err != nil {
		t.Fatalf("Failed to query updated read store: %v", err)
	}

	if updatedView.CurrentBalance != 750.00 || updatedView.TotalDeposits != 1 {
		t.Fatalf("Expected updated balance 750.00 and 1 deposit, got balance %.2f, deposits %d",
			updatedView.CurrentBalance, updatedView.TotalDeposits)
	}
}
