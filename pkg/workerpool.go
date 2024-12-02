package pkg

type WorkerPool[T int] struct {
	available chan *Worker[T]
}

func NewWorkerPool[T int](numOfWorkers uint8) *WorkerPool[T] {
	wp := &WorkerPool[T]{
		available: make(chan *Worker[T], numOfWorkers),
	}

	for i := 0; i < int(numOfWorkers); i++ {
		worker := &Worker[T]{
			ID: uint8(i),
		}
		wp.available <- worker
	}

	return wp
}

func (wp *WorkerPool[T]) AssignJob(job Job[T]) {
	// Get a worker from the available workers
	worker := <-wp.available

	go func() {
		worker.ProcessJob(job)
		wp.available <- worker
	}()
}
