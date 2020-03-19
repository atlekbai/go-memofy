package memofy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {

	// some expensive function
	expensiveSumFunction := func(a, b int) int {
		time.Sleep(2 * time.Second)
		return a + b
	}
	cache := NewMemofier(90*time.Second, 10*time.Minute)

	// first call should return cached = false,
	// because it is not yet cached
	result, cached, err := cache.Memofy(expensiveSumFunction, 5, 7)
	results := result.([]interface{})

	assert.Equal(t, cached, false, "==")
	assert.Equal(t, results[0].(int), 5+7, "==")
	assert.Nil(t, err)

	// second call should return cached = true,
	// because it was computer before
	result, cached, err = cache.Memofy(expensiveSumFunction, 5, 7)
	results = result.([]interface{})

	assert.Equal(t, cached, true, "==")
	assert.Equal(t, results[0].(int), 5+7, "==")
	assert.Nil(t, err)

}
