package cache_test

import (
	"cache"
	"fmt"
	"testing"
	"time"
)

type user struct {
	name string
}

func TestCache(t *testing.T) {
	cache := cache.NewCache()

	cache.Set("age", 42, time.Second)
	cache.Set("user", user{name: "James"}, time.Hour)

	jsonData, err := cache.ToJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonData))

	cache.Delete("age")
	jsonData, err = cache.ToJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonData))

	cache.Clear()
	_, exists := cache.Get("user")
	if exists {
		t.Error("value should be empty")
	}

}
