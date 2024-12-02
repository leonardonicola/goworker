package pkg

import (
	"fmt"
)

type Worker[T int] struct {
	ID uint8
}

// Example of process that a worker could do.
func (wk Worker[T]) Fib(n int) int {
	if n <= 1 {
		return n
	}

	return wk.Fib(n-1) + wk.Fib(n-2)
}

func (wk Worker[T]) ProcessJob(job Job[T]) {
	switch job.Process {
	case Fib:
		// Type assertion to ensure job data type is correct before processing it.
		if num, ok := any(job.Data).(int); ok {
			result := wk.Fib(num)
			// Send the result back to the client via channels.
			job.Result <- result
		} else {
			job.Result <- fmt.Errorf("Invalid data type for Fibonacci calculation: expected int")
		}
		// Just in case...
	default:
		job.Result <- fmt.Errorf("Unknown process: %s", processName[job.Process])
	}
}
