package t7

type ServerMetric[T comparable] struct {
	Name  string
	Value T
}

const bytesInMegabyte = 1024 * 1024

func Decorate[TInput any, TOutput any](input chan TInput, action func(TInput) TOutput) chan TOutput {
	out := make(chan TOutput)

	go func() {
		defer close(out)
		for in := range input {
			out <- action(in)
		}
	}()

	return out
}

func ServerMetricBytesToMegabytes(inBytes ServerMetric[int]) (inMegabytes ServerMetric[float64]) {
	inMegabytes = ServerMetric[float64]{
		Name:  inBytes.Name,
		Value: float64(inBytes.Value / bytesInMegabyte),
	}
	return
}
