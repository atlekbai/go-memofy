# memofy - cache function calls

Package `memofy` implements a cache for function calls. It can save time when an expensive or I/O bound function is periodically called with the same arguments.

Highly inspired by [lru_cache](https://docs.python.org/3/library/functools.html#functools.lru_cache) Python3 decorator.

Forked from [go-memoize](https://github.com/kofalt/go-memoize) and reimplemented to be more similar to [lru_cache](https://docs.python.org/3/library/functools.html#functools.lru_cache).

## Install

```sh
go get -u github.com/atlekbai/go-memofy
```

## Examples

```go
package main

import (
    "fmt"
    "time"
    "github.com/atlekbai/go-memofy"
)

func main() {
    expensiveSumFunction := func(a, b int) int {
		time.Sleep(2 * time.Second)
		return a + b
    }

    // Cache expensive calls in memory for 90 seconds, purging old entries every 10 minutes.
	cache := memofy.NewMemofier(90*time.Second, 10*time.Minute)

    // first call will take 2 seconds and save to cache computed values
    result, cached, err := cache.Memofy(expensiveSumFunction, 5, 7)
    // second call will return cached value
    result, cached, err = cache.Memofy(expensiveSumFunction, 5, 7)

    // result returns []interface{}, because function can return multiple values
    returnValues := result.([]interface{})

    // [12]
    fmt.Println(returnValues)

    // cast to int to retrieve function return value
    sumValue := returnValues[0].(int)
}
```
