package pokecache

import (
	"testing"
	"time"
)

func TestPokeCache(t *testing.T) {
	testcases := []struct {
		data     map[string][]byte
		duration time.Duration
	}{
		{
			data: map[string][]byte{
				"oi":     []byte("oi"),
				"a":      []byte("a"),
				"xalala": []byte("xalala"),
				"xilili": []byte("xilili"),
			},
			duration: 100 * time.Millisecond,
		},
	}
	for _, tc := range testcases {
		cache := NewCache(tc.duration)
		for k, v := range tc.data {
			cache.Add(k, v)
		}
		for k := range tc.data {
			_, ok := cache.Get(k)
			if !ok {
				t.Errorf("expected data for key %v to be stored, but it was not", k)
			}
		}
		time.Sleep(2 * tc.duration)
		for k := range tc.data {
			_, ok := cache.Get(k)
			if ok {
				t.Errorf("expected data for key %v to be expired, but it was still present", k)
			}
		}
	}
}
