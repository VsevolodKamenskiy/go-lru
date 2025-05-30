# LRU Cache implementation

An efficient LRU cache implementation in Go with O(1) time complexity for all operations.

## Features
- Constant time Add, Get, and Remove operations
- Automatic eviction of least recently used items
- Thread-safe for concurrent usage

## Usage
```go
import "github.com/VsevolodKamenskiy/go-lru"

func main() {
    cache := lru.NewLRUCache(100)
    cache.Add("key1", "value1")
    value, ok := cache.Get("key1")
    cache.Remove("key1")
}