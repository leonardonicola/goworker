package pkg

import (
	"fmt"
)

type Worker[T int] struct {
	ID uint8
}

func (wk Worker[T]) Fib(n int) int {
	if n <= 1 {
		return n
	}

	return wk.Fib(n-1) + wk.Fib(n-2)
}

func (wk Worker[T]) ProcessJob(job Job[T]) {
	switch job.Process {
	case Fib:
		// Type assertion to ensure Data is uint
		if num, ok := any(job.Data).(int); ok {
			result := wk.Fib(num)
			job.Result <- result
		} else {
			job.Result <- fmt.Errorf("invalid data type for Fibonacci calculation: expected uint")
		}
	default:
		job.Result <- fmt.Errorf("Unknown process: %s", processName[job.Process])
	}
}
