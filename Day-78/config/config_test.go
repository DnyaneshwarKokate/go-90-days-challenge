package config_test

import (
	"testing"
	"time"

	"day78/config"
)

func TestDynamicConfigServerSetAndWatch(t *testing.T) {
	server := config.NewDynamicConfigServer()

	server.Set("feature.new_ui.enabled", "false")

	watchCh, err := server.Watch("feature.new_ui.enabled")
	if err != nil {
		t.Fatalf("Failed to register watcher: %v", err)
	}

	// Update config key (Triggers Hot Reload Event)
	server.Set("feature.new_ui.enabled", "true")

	select {
	case evt := <-watchCh:
		if evt.OldValue != "false" || evt.NewValue != "true" {
			t.Fatalf("Event value mismatch: %+v", evt)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("Timed out waiting for config change event")
	}

	val, _ := server.Get("feature.new_ui.enabled")
	if val != "true" {
		t.Fatalf("Expected 'true', got %s", val)
	}
}
