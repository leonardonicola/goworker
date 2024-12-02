package main

import (
	"fmt"
	"os"

	"github.com/leonardonicola/workers/pkg"
)

func main() {
	// Usage example
	res := make(chan any, 1)
	defer close(res)
	job, err := pkg.NewJob(pkg.Fib, res, 20)

	if err != nil {
		fmt.Printf("Couldn't create a Job: %v\n", err)
		os.Exit(1)
	}

	wp := pkg.NewWorkerPool(2)
	go wp.AssignJob(*job)

	result := <-res
	switch v := result.(type) {
	case int:
		fmt.Printf("Fibonacci result: %d\n", v)
	case error:
		fmt.Printf("Error: %v\n", v)
	default:
		fmt.Printf("Something occurred: %v\n", v)
	}
}
