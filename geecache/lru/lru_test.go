package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Size() int {
	return len(d)
}

func TestCache_Get(t *testing.T) {
	lruCache := New(int64(0), nil)
	lruCache.Add("key1", String("1234"))
	if v, ok := lruCache.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lruCache.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	cap := len(k1 + k2 + v1 + v2)
	lruCache := New(int64(cap), nil)
	lruCache.Add(k1, String(v1))
	lruCache.Add(k2, String(v2))
	lruCache.Add(k3, String(v3))

	if _, ok := lruCache.Get("key1"); ok || lruCache.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestCache_OnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lruCache := New(int64(10), callback)
	lruCache.Add("key1", String("123456"))
	lruCache.Add("k2", String("k2"))
	lruCache.Add("k3", String("k3"))
	lruCache.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}
	
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
