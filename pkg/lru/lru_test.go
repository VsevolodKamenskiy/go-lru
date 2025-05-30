package lru

import "testing"

func TestEviction(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Add("key1", "val1")
	cache.Add("key2", "val2")
	cache.Add("key3", "val3")

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should be evicted")
	}

	if _, ok := cache.Get("key2"); !ok {
		t.Error("key2 should be present")
	}

	if _, ok := cache.Get("key3"); !ok {
		t.Error("key3 should be present")
	}

	// Check eviction when updating
	cache.Add("key2", "new_val") // Updating, eviction shouldn't be caused
	cache.Add("key4", "val4")    // Key3 should be evicted

	if _, ok := cache.Get("key3"); ok {
		t.Error("key3 should be evicted")
	}
	if val, _ := cache.Get("key2"); val != "new_val" {
		t.Error("key2 should be updated")
	}
}

func TestZeroCapacity(t *testing.T) {
	cache := NewLRUCache(0)

	// Element shouldn't be added
	if cache.Add("key1", "val1") {
		t.Error("Add should fail for zero capacity cache")
	}

}

func TestSingleElement(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Add("key1", "val1")
	cache.Add("key2", "val2") // Should evict key1

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should be evicted")
	}
	if _, ok := cache.Get("key2"); !ok {
		t.Error("key2 should be present")
	}

	// Updating element
	cache.Add("key2", "new_val")
	if val, _ := cache.Get("key2"); val != "new_val" {
		t.Error("key2 should be updated")
	}
}
