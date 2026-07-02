package t3_test

import (
	"fmt"
	"t3"
	"testing"
)

func TestDBPool(t *testing.T) {
	pool := t3.NewDBPool(3)

	for i := range 10 {
		go func(id int) {
			conn, _ := pool.Get()
			defer pool.Release(conn)

			fmt.Printf("Горутина %d: подключение %d получено\n", id, conn.ID)

		}(i)
	}
}
