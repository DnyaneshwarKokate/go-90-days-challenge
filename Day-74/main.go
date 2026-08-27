package main

import (
	"fmt"

	"day74/sharding"
)

func main() {
	fmt.Println("=== Day 74: Database Sharding & Consistent Hashing ===")

	hashRing := sharding.NewConsistentHashRing(5) // 5 virtual nodes per physical shard

	shards := []string{"db-shard-us-east", "db-shard-us-west", "db-shard-eu-central"}
	for _, shard := range shards {
		hashRing.AddNode(shard)
		fmt.Printf("[RING] Added Shard Node: %s (5 Virtual Replicas)\n", shard)
	}

	keys := []string{
		"user_id_1001", "user_id_1002", "order_uuid_abc",
		"payment_id_99", "tenant_xyz_profile",
	}

	fmt.Println("\n--- Consistent Hash Key Mapping ---")
	for _, k := range keys {
		shardNode, _ := hashRing.GetNode(k)
		fmt.Printf("Key '%s' ==> Target Shard: %s\n", k, shardNode)
	}

	fmt.Println("\n--- Simulating Shard Node Removal (db-shard-us-east) ---")
	hashRing.RemoveNode("db-shard-us-east")

	for _, k := range keys {
		shardNode, _ := hashRing.GetNode(k)
		fmt.Printf("Key '%s' ==> Remapped Shard: %s\n", k, shardNode)
	}
}
