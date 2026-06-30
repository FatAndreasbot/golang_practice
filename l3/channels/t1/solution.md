```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
)

func main() {
    alreadyStored := make(map[int]struct{})
    capacity := 1000
    doubles := make([]int, 0, capacity)
    for i := 0; i < capacity; i++ {
    		doubles = append(doubles, rand.Intn(10))
    }
    
    uniqueIDs := make(chan int, capacity)
    wg := sync.WaitGroup{}
    for i := 0; i < capacity; i++ {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            if _, ok := alreadyStored[doubles[i]]; !ok {
                alreadyStored[doubles[i]] = struct{}{}
                uniqueIDs <- doubles[i]
            }
        }()
    }
    
    wg.Wait()
    for val := range uniqueIDs {
    	fmt.Println(val)
    }
    
    fmt.Println(uniqueIDs)
}
```
# Замечания
 - Заменить `for i:=0; i<capacity; i++` на 
`for range capacity` и `for i := range capacity`
 - Заменть `wg.Add(1)` на `wg.Go(func...)`
 - Убрал `i:=i`
 - Заменил `wg := sync.WaitGroup{}` на `var wg sync.WaitGroup`
 - 
 # Ошибки
 - Нужен mutex
 - Канал не закрывается

# Исправления

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	capacity := 1000
	doubles := make([]int, 0, capacity)
	for range capacity {
		doubles = append(doubles, rand.Intn(10))
	}

	uniqueIDs := make(chan int, capacity+1)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := range capacity {
		wg.Go(func() {
			mutex.Lock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}
				uniqueIDs <- doubles[i]
			}
			mutex.Unlock()
		})
	}

	wg.Wait()
	close(uniqueIDs)
	for val := range uniqueIDs {
		fmt.Println(val)
	}

	fmt.Println(uniqueIDs)
}

```
