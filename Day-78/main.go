package main

import (
	"fmt"
	"time"

	"day78/config"
)

func main() {
	fmt.Println("=== Day 78: Distributed Configuration Server & Hot Reloading ===")

	configServer := config.NewDynamicConfigServer()

	// Initial configuration
	configServer.Set("db.max_connections", "10")
	configServer.Set("auth.jwt_ttl_minutes", "60")

	// Subscribe to live updates for db.max_connections
	watchCh, _ := configServer.Watch("db.max_connections")

	// Background worker simulating application listening to hot reloads
	go func() {
		for evt := range watchCh {
			fmt.Printf("  ==> [HOT RELOAD WORKER] Config '%s' changed: '%s' -> '%s' (Time: %s)\n",
				evt.Key, evt.OldValue, evt.NewValue, evt.Timestamp.Format("15:04:05.000"))
		}
	}()

	fmt.Println("\n--- Updating Configuration Values Live ---")
	time.Sleep(50 * time.Millisecond)
	configServer.Set("db.max_connections", "50")

	time.Sleep(50 * time.Millisecond)
	configServer.Set("db.max_connections", "100")

	time.Sleep(100 * time.Millisecond)

	val, _ := configServer.Get("db.max_connections")
	fmt.Printf("\nFinal In-Memory Config Value for 'db.max_connections': %s\n", val)
}
