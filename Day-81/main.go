package main

import (
	"fmt"

	"day81/systemdesign"
)

func main() {
	fmt.Println("=== Day 81: High-Availability System Design (CAP & PACELC Theorem) ===")

	// 1. CP/EC Cluster (Strong Consistency Mode)
	cpCluster := systemdesign.NewDistributedCluster(systemdesign.ModeCP_EC)
	cpCluster.AddNode("node-cp-1")
	cpCluster.AddNode("node-cp-2")

	fmt.Println("\n--- 1. CP/EC Cluster under Network Partition ---")
	cpCluster.SetPartition(true)
	errCP := cpCluster.Write("db.config", "strict_value")
	fmt.Printf("[CP/EC OUTCOME] Write Status: %v\n", errCP)

	// 2. AP/EL Cluster (High Availability Mode)
	apCluster := systemdesign.NewDistributedCluster(systemdesign.ModeAP_EL)
	apCluster.AddNode("node-ap-1")

	fmt.Println("\n--- 2. AP/EL Cluster under Network Partition ---")
	apCluster.SetPartition(true)
	_ = apCluster.Write("db.config", "eventual_value")
	val, ver, _ := apCluster.Read("node-ap-1")
	fmt.Printf("[AP/EL OUTCOME] Write Status: Success | Value: '%s' | Version: %d\n", val, ver)
}
