package t4_test

import (
	"maps"
	"t4"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	merged := t4.MergeChannels(ch1, ch2, ch3)

	values := make(map[int]struct{})

	go func() {
		for n := range 30 {
			values[n] = struct{}{}
			switch n % 3 {
			case 0:
				ch1 <- n
			case 1:
				ch2 <- n
			default:
				ch3 <- n
			}
		}
		close(ch1)
		close(ch2)
		close(ch3)
	}()
	numbers := make(map[int]struct{})

	for n := range merged {
		numbers[n] = struct{}{}
	}

	if !maps.Equal(values, numbers) {
		t.Errorf("the return values are different\n%v\n%v", values, numbers)
	}

}
