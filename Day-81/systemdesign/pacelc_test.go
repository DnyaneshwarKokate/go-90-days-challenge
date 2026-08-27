package systemdesign_test

import (
	"testing"

	"day81/systemdesign"
)

func TestPACELCClusterTradeoffs(t *testing.T) {
	// Test CP/EC Mode (Rejects writes during partition)
	cpCluster := systemdesign.NewDistributedCluster(systemdesign.ModeCP_EC)
	cpCluster.AddNode("node-1")
	cpCluster.AddNode("node-2")

	cpCluster.SetPartition(true)
	errCP := cpCluster.Write("setting", "val1")
	if errCP == nil {
		t.Fatalf("CP mode should reject writes during network partition")
	}

	// Test AP/EL Mode (Accepts writes during partition)
	apCluster := systemdesign.NewDistributedCluster(systemdesign.ModeAP_EL)
	apCluster.AddNode("node-1")
	apCluster.SetPartition(true)

	errAP := apCluster.Write("setting", "val2")
	if errAP != nil {
		t.Fatalf("AP mode should accept write during partition, got %v", errAP)
	}

	val, ver, _ := apCluster.Read("node-1")
	if val != "val2" || ver != 1 {
		t.Fatalf("Expected val2 version 1, got %s ver %d", val, ver)
	}
}
