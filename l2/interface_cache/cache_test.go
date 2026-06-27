package cache_test

import (
	C "cache"
	"fmt"
	"testing"
	"time"
)

type user struct {
	Name string `json:"name"`
}

func TestCache(t *testing.T) {
	cache := C.NewCache()

	cache.Set("age", 42, time.Second)
	cache.Set("user", user{Name: "James"}, time.Hour)

	jsonData, err := cache.ToJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonData))
	if string(jsonData) != "{\"age\":42,\"user\":{\"name\":\"James\"}}" {
		t.Error("got wrong json data")
	}

	time.Sleep(time.Second * 2)
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
	cache.Set("string", "a string", time.Hour)

	_, err = C.GetAs[string](&cache, "string")
	if err != nil {
		t.Error(err)
	}
	_, err = C.GetAs[int](&cache, "string")
	if err == nil {
		t.Error("i did not got the error")
	}
}
