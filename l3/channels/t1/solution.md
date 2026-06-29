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
# Ошибки
строка 
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
    for i := 0; i < capacity; i++ {
    		doubles = append(doubles, rand.Intn(10))
    }
    
    uniqueIDs := make(chan int, capacity)
    wg := sync.WaitGroup{}
    var mu sync.Mutex 
    
    for _, val := range doubles {
        wg.Go(func() {
      			mu.Lock()
            if _, ok := alreadyStored[val]; !ok {
                alreadyStored[val] = struct{}{}
                uniqueIDs <- val
            }
            mu.Unlock()
        })
    }
    
    wg.Wait()
    for val := range uniqueIDs {
    	fmt.Println(val)
    }
    
    fmt.Println(uniqueIDs)
}
```
