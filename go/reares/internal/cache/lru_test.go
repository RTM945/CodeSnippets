package cache

import "testing"

func TestCache(t *testing.T) {
	type CustomKey struct {
		ID int
	}

	type CustomValue struct {
		Data string
	}

	var evictedKey CustomKey
	var evictedVal CustomValue

	onEvict := func(k CustomKey, v CustomValue) {
		evictedKey = k
		evictedVal = v
	}

	lru := New[CustomKey, CustomValue](int64(1), onEvict)

	key1 := CustomKey{ID: 1}
	val1 := CustomValue{Data: "data1"}
	lru.Add(key1, val1)

	// This should evict key1
	key2 := CustomKey{ID: 2}
	val2 := CustomValue{Data: "data2"}
	lru.Add(key2, val2)

	// Check eviction
	if evictedKey.ID != 1 || evictedVal.Data != "data1" {
		t.Fatalf("wrong eviction: got key %v, value %v", evictedKey, evictedVal)
	}

	// key1 should be gone, key2 should be there
	if _, ok := lru.Get(key1); ok {
		t.Fatal("key1 should be evicted")
	}

	if val, ok := lru.Get(key2); !ok || val.Data != "data2" {
		t.Fatalf("key2 should be in cache with value data2, got %v", val)
	}
}
