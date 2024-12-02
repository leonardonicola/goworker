package pkg

type WorkerPool[T int] struct {
	available chan *Worker[T]
}

func NewWorkerPool[T int](numOfWorkers uint8) *WorkerPool[T] {
	wp := &WorkerPool[T]{
		// Create buffered channel as the pool of workers
		available: make(chan *Worker[T], numOfWorkers),
	}

	for i := 0; i < int(numOfWorkers); i++ {
		worker := &Worker[T]{
			ID: uint8(i),
		}

		// Add created worker to pool.
		wp.available <- worker
	}

	return wp
}

func (wp *WorkerPool[T]) AssignJob(job Job[T]) {
	// Bother a worker from the available ones.
	worker := <-wp.available

	// Process job concurrently.
	go func() {
		worker.ProcessJob(job)

		// Add the worker back to the pool when job finished.
		wp.available <- worker
	}()
}
