package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache_SetAndGet(t *testing.T) {
	cache := NewInMemoryCache()
	key := "key1"
	value := "value1"
	cache.Set(key, value, 10*time.Minute)

	got, exists := cache.Get(key)
	if !exists || got != value {
		t.Errorf("expected %v, got %v", value, got)
	}
}

func TestInMemoryCache_Expiry(t *testing.T) {
	cache := NewInMemoryCache()
	key := "key1"
	value := "value1"
	cache.Set(key, value, 1*time.Millisecond)

	time.Sleep(2 * time.Millisecond) // Wait for the value to expire

	_, exists := cache.Get(key)
	if exists {
		t.Errorf("expected value to be expired and deleted")
	}
}

func TestInMemoryCache_Delete(t *testing.T) {
	cache := NewInMemoryCache()
	key := "key1"
	value := "value1"
	cache.Set(key, value, 10*time.Minute)
	cache.Delete(key)

	_, exists := cache.Get(key)
	if exists {
		t.Errorf("expected value to be deleted")
	}
}

func TestInMemoryCache_Clear(t *testing.T) {
	cache := NewInMemoryCache()
	cache.Set("key1", "value1", 10*time.Minute)
	cache.Set("key2", "value2", 10*time.Minute)
	cache.Clear()

	if len(cache.data) != 0 {
		t.Errorf("expected cache to be empty, got %d items", len(cache.data))
	}
}
