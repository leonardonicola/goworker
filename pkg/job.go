package pkg

import (
	"fmt"
)

// Enum.
type JobProcesses uint8

const (
	Fib = iota
	QuickSort
)

var processName = map[JobProcesses]string{
	Fib:       "fib",
	QuickSort: "quick_sort",
}

type Job[T int | []T] struct {
	Process JobProcesses
	Result  chan<- any // Could be an error or the data
	Data    T
}

func NewJob[T int | []T](process JobProcesses, result chan<- any, data T) (*Job[T], error) {
	switch process {
	case Fib, QuickSort:
		return &Job[T]{
			Process: process,
			Result:  result,
			Data:    data,
		}, nil
	default:
		return nil, fmt.Errorf("Unknown process: %s", processName[process])
	}

}
