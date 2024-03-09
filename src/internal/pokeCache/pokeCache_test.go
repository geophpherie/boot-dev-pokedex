package pokeCache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = time.Duration(5 * time.Second)

	testCases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://myfakesite.com",
			val: []byte("datadatadatadata"),
		},
		{
			key: "http://myfakesite.com",
			val: []byte("data2data2data2data2"),
		},
	}

	for u, c := range testCases {
		t.Run(fmt.Sprintf("Test case %v", u), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find %v", c.key)
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("value didn't match")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const staleTime = 5 * time.Millisecond
	// const sleepTime = staleTime + 5*time.Millisecond
	const sleepTime = 50 * time.Millisecond

	cache := NewCache(staleTime)

	cacheData := struct {
		key string
		val []byte
	}{
		key: "http://myfakesite.com",
		val: []byte("my fake data"),
	}

	cache.Add(cacheData.key, cacheData.val)
	_, ok := cache.Get(cacheData.key)
	if !ok {
		t.Errorf("expected to find %v", cacheData.key)
		return
	}

	time.Sleep(sleepTime)
	_, ok = cache.Get(cacheData.key)
	if ok {
		t.Errorf("expected to not find %v", cacheData.key)
		return
	}

}
