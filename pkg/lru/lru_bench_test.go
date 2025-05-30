package lru

import (
	"sync"
	"testing"
)

func BenchmarkConcurrentAdd(b *testing.B) {
	cache := NewLRUCache(1000)
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Add(string(rune(i)), "value")
		}(i)
	}
	wg.Wait()
}

func BenchmarkConcurrentGet(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := 0; i < 1000; i++ {
		cache.Add(string(rune(i)), "value")
	}

	b.ResetTimer()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Get(string(rune(i % 1000)))
		}(i)
	}
	wg.Wait()
}

func BenchmarkConcurrentMixed(b *testing.B) {
	cache := NewLRUCache(1000)
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(3)
		go func(i int) {
			defer wg.Done()
			cache.Add(string(rune(i)), "value")
		}(i)
		go func(i int) {
			defer wg.Done()
			cache.Get(string(rune(i % 1000)))
		}(i)
		go func(i int) {
			defer wg.Done()
			cache.Remove(string(rune(i % 500)))
		}(i)
	}
	wg.Wait()
}
