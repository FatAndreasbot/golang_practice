```go
package main

import (
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	println(int(time.Since(timeStart).Seconds()))
}
```
# Ответ
6 секунд
# Объяснение
После создания первого воркера исполнение остановилось,
пока он не закончил работу, и не отправил 
свое значение, и поэтому второй пока не создался
# Исправление
```go

// TODO fan-in solution
func main() {
	timeStart := time.Now()
	ch1 := worker()
    ch2 := worker()

    _, _ = <-ch1, <-ch2
	println(int(time.Since(timeStart).Seconds()))
}
```
