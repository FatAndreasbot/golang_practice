package main

import (
	"fmt"
	"obj_cache"
	"time"
)

func main() {
	cache := obj_cache.NewObjectCache(5 * time.Second)

	// Добавляем данные в кэш
	cache.Set("user:1", map[string]string{"name": "Alice", "role": "admin"})
	cache.Set("user:2", map[string]string{"name": "Bob", "role": "user"})
	cache.Set("user:3", "asdasdasd")
	cache.Set("user:4", []string{"asdasd", "1231231", "zxczxc"})

	// Получаем объект
	if user, found := cache.Get("user:1"); found {
		fmt.Println("Найден:", user)
	}

	// Выводим JSON
	jsonData, _ := cache.ToJSON()
	fmt.Println("Кэш в JSON:", string(jsonData))

	// Ждём истечения TTL и проверяем снова
	time.Sleep(6 * time.Second)
	_, found := cache.Get("user:1")
	fmt.Println("После TTL, user:1 найден?", found)
}
