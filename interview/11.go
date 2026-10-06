// You have 1 million jobs and 20 workers.
// How would you design the worker pool?
//
// INTERVIEW ANSWER:
//
// "I would create a jobs channel and start 20 worker goroutines.
// Each worker continuously receives a job from the channel,
// processes it, and then receives the next job.
//
// I would use a WaitGroup to wait until all 20 workers finish.
// The jobs channel would be closed by the producer after all
// jobs have been submitted."

package interview

import (
	"fmt"
	"sync"
)

func processJob(job int) {
	fmt.Println("processed job no:", job)
}

func question11() {

	// The channel is used to distribute jobs to workers.
	//
	// Buffer size 1000 means we can temporarily hold 1000 jobs
	// without requiring a worker to receive immediately.
	//
	// We DON'T need a buffer of 1 million.
	//
	// A bounded buffer also prevents us from unnecessarily
	// keeping all 1 million jobs in memory.
	jobs := make(chan int)

	var wg sync.WaitGroup

	// Start exactly 20 workers.
	//
	// Each worker keeps taking jobs from the channel until
	// the channel is closed and all jobs have been consumed.
	for i := 1; i <= 20; i++ {

		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for job := range jobs {

				// Each worker processes one job at a time.
				processJob(job)
			}

			// When range finishes, the jobs channel has been
			// closed AND there are no jobs left to consume.
		}(i)
	}

	// Producer sends 1 million jobs.
	//
	// Workers are already running, so they can consume jobs
	// while we are still producing them.
	for i := 1; i <= 1_000_000; i++ {
		jobs <- i
	}

	// IMPORTANT:
	//
	// Closing the channel tells workers:
	//
	// "No more jobs are coming."
	//
	// Workers will finish processing any remaining jobs and
	// then their:
	//
	//     for job := range jobs
	//
	// loop will terminate.
	close(jobs)

	// Wait until all 20 workers have finished.
	wg.Wait()

	fmt.Println("All jobs processed")
}


// Understand the architecture

// Think of it as:

//                  PRODUCER
//                     │
//                     │ 1,000,000 jobs
//                     ↓
//              ┌──────────────┐
//              │ jobs channel │
//              └──────────────┘
//                │ │ │ │ │
//         ┌──────┘ │ │ │ └──────┐
//         ↓        ↓ ↓ ↓        ↓
//      Worker 1 Worker 2 ... Worker 19 Worker 20
//         │        │              │        │
//         ↓        ↓              ↓        ↓
//       Job      Job            Job      Job


// Start 20 workers
//        ↓
// Start producing jobs
//        ↓
// Workers consume jobs concurrently
//        ↓
// Producer finishes
//        ↓
// close(jobs)
//        ↓
// Workers finish remaining jobs
//        ↓
// wg.Wait()


// I would use a worker pool with 20 goroutines consuming from a shared jobs channel. 
// The producer sends the 1 million jobs into the channel, and each worker processes jobs until the channel is closed. 
// I use a WaitGroup to wait for all workers to finish. I would also use a bounded channel rather than buffering all 1 million jobs in memory.