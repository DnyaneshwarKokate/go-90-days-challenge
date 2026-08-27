package sharding_test

import (
	"testing"

	"day74/sharding"
)

func TestConsistentHashingDistributionAndRemoval(t *testing.T) {
	ring := sharding.NewConsistentHashRing(3)

	ring.AddNode("shard-db-1")
	ring.AddNode("shard-db-2")
	ring.AddNode("shard-db-3")

	nodeA, errA := ring.GetNode("user_account_9001")
	nodeB, errB := ring.GetNode("user_account_9002")

	if errA != nil || errB != nil {
		t.Fatalf("GetNode failed: %v, %v", errA, errB)
	}

	if nodeA == "" || nodeB == "" {
		t.Fatalf("Expected non-empty node assignment")
	}

	// Remove shard-db-1
	ring.RemoveNode("shard-db-1")

	nodeAfter, errAfter := ring.GetNode("user_account_9001")
	if errAfter != nil {
		t.Fatalf("GetNode after removal failed: %v", errAfter)
	}

	if nodeAfter == "shard-db-1" {
		t.Fatalf("Key mapped to removed node shard-db-1!")
	}
}
