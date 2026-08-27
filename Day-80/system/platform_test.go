package system_test

import (
	"context"
	"testing"

	"day80/system"
)

func TestMicroservicesPlatformOrderSubmission(t *testing.T) {
	platform := system.NewMicroservicesPlatform(1)

	req := system.PlatformOrderRequest{
		Customer: "Grace Hopper",
		Amount:   1500.00,
	}

	resp, err := platform.SubmitOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("SubmitOrder failed: %v", err)
	}

	if resp.ID == 0 || resp.ShardNode == "" || resp.Status != "PROCESSED" {
		t.Fatalf("Invalid platform response: %+v", resp)
	}

	if platform.TotalProcessed() != 1 {
		t.Fatalf("Expected 1 processed order, got %d", platform.TotalProcessed())
	}
}
