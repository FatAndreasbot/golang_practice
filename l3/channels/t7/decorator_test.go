package t7_test

import (
	"fmt"
	"math/rand"
	"slices"
	"t7"
	"testing"
)

func transformSliceForCompare[T any, R comparable](input []T, transform func(T) R) []R {
	result := make([]R, 0, len(input))

	for _, value := range input {
		result = append(result, transform(value))
	}

	return result
}

func fromMetricToComparable[T comparable](metric t7.ServerMetric[T]) T {
	return metric.Value
}

func TestTransform(t *testing.T) {
	inputCount := 5
	inputs := make([]t7.ServerMetric[int], 0, inputCount)
	expected := make([]t7.ServerMetric[float64], 0, inputCount)
	result := make([]t7.ServerMetric[float64], 0, inputCount)

	input := make(chan t7.ServerMetric[int])
	output := t7.Decorate(input, t7.ServerMetricBytesToMegabytes)

	for range inputCount {
		inBytes := t7.ServerMetric[int]{
			Name:  fmt.Sprintf("server n%d", inputCount),
			Value: (rand.Intn(200) + 100) * rand.Intn(100) * 1000,
		}
		inMegabytes := t7.ServerMetricBytesToMegabytes(inBytes)
		inputs = append(inputs, inBytes)
		expected = append(expected, inMegabytes)
	}

	go func() {
		defer close(input)
		for _, value := range inputs {
			input <- value
		}
	}()

	for inMegabytes := range output {
		result = append(result, inMegabytes)
	}

	if slices.Compare(
		transformSliceForCompare(expected, fromMetricToComparable),
		transformSliceForCompare(result, fromMetricToComparable),
	) != 0 {
		t.Error("got wrong values")
	}

}
