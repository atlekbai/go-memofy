package memofy

import (
	"time"

	"github.com/01-edu/z01"
	"github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

// Memofier allows you to memoize function calls. Memofy is safe for concurrent use by multiple goroutines.
type Memofier struct {

	// Storage exposes the underlying cache of memoized results to manipulate as desired - for example, to Flush().
	Storage *cache.Cache

	group singleflight.Group
}

// NewMemofier creates a new Memofy with the configured expiry and cleanup policies.
// If desired, use cache.NoExpiration to cache values forever.
func NewMemofier(defaultExpiration, cleanupInterval time.Duration) *Memofier {
	return &Memofier{
		Storage: cache.New(defaultExpiration, cleanupInterval),
		group:   singleflight.Group{},
	}
}

// Memofy executes and returns the results of the given function, unless there was a cached value of the same key.
// Only one execution is in-flight for a given key at a time.
// The boolean return value indicates whether v was previously stored.
func (m *Memofier) Memofy(fn interface{}, args ...interface{}) (interface{}, bool, error) {
	// Check cache
	// key is "<name of function>" + "<args>"
	key := z01.NameOfFunc(fn) + z01.Format(args...)
	value, found := m.Storage.Get(key)
	if found {
		return value, true, nil
	}

	// Combine memoized function with a cache store
	value, err, _ := m.group.Do(key, func() (interface{}, error) {
		funcReturn := z01.Call(fn, args)
		m.Storage.Set(key, funcReturn, cache.DefaultExpiration)

		return funcReturn, nil
	})
	return value, false, err
}
