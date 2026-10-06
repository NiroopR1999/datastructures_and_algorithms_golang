package interview

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
========================================================
QUESTION:
HOW WOULD YOU GRACEFULLY SHUT DOWN 20 WORKER GOROUTINES
WHEN THE APPLICATION RECEIVES SIGTERM?
========================================================

INTERVIEW ANSWER:

"I would listen for SIGTERM, cancel a shared context,
close the jobs channel if I own it, and use a WaitGroup
to wait for all workers to finish their current work
before the application exits."


========================================================
IMPORTANT IDEA
========================================================

SIGTERM
   ↓
Application receives shutdown signal
   ↓
Tell workers to stop
   ↓
Workers finish current jobs
   ↓
Workers exit
   ↓
WaitGroup reaches 0
   ↓
Application exits


========================================================
WORKER
========================================================
*/

func worker19(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		select {

		case <-ctx.Done():

			// Shutdown signal received.
			// Stop accepting new work.
			fmt.Println("worker", id, "shutting down")
			return

		case job, ok := <-jobs:

			if !ok {
				// Jobs channel was closed.
				// No more work will arrive.
				fmt.Println("worker", id, "finished")
				return
			}

			fmt.Println("worker", id, "processing job", job)

			// Simulate work.
			time.Sleep(500 * time.Millisecond)

			fmt.Println("worker", id, "finished job", job)
		}
	}
}

/*
========================================================
APPLICATION
========================================================
*/

func question19() {

	// Context used to tell all workers:
	// "The application is shutting down."
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)

	var wg sync.WaitGroup

	// Start 20 workers.
	for i := 1; i <= 20; i++ {

		wg.Add(1)

		go worker19(ctx, i, jobs, &wg)
	}

	/*
		We listen for SIGTERM.

		SIGTERM is commonly sent when a process is
		being gracefully terminated, for example:

			Kubernetes
			Docker
			Linux service manager
	*/

	signalCh := make(chan os.Signal, 1)

	signal.Notify(
		signalCh,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	// Wait until SIGTERM/SIGINT is received.
	sig := <-signalCh

	fmt.Println("received signal:", sig)

	/*
	====================================================
		START SHUTDOWN
	====================================================
	*/

	// Tell all workers to stop.
	cancel()

	/*
	====================================================
		WAIT FOR WORKERS
	====================================================

	WaitGroup blocks until all 20 workers call:

		wg.Done()

	Only then do we allow the application to exit.
	*/

	wg.Wait()

	fmt.Println("all workers stopped gracefully")
}

/*
========================================================
WHY CONTEXT?
========================================================

All 20 workers share the SAME context:

	worker 1 ─┐
	worker 2  │
	worker 3  │
	...       ├──→ ctx.Done()
	worker 20 ┘

When we call:

	cancel()

ctx.Done() becomes readable.

Every worker notices:

	case <-ctx.Done():

and returns.

We don't have to manually tell each worker:

	"Worker 1 stop"
	"Worker 2 stop"
	...
	"Worker 20 stop"

One cancellation signal reaches all of them.


========================================================
WHY WAITGROUP?
========================================================

Context answers:

	"Should the workers stop?"

WaitGroup answers:

	"Have all workers actually stopped?"

So:

	cancel()
	   ↓
	tell workers to stop

	wg.Wait()
	   ↓
	wait until they actually stop


========================================================
GRACEFUL VS IMMEDIATE SHUTDOWN
========================================================

GRACEFUL:

	SIGTERM
	   ↓
	stop taking new work
	   ↓
	finish current work
	   ↓
	exit


IMMEDIATE:

	SIGTERM
	   ↓
	kill process immediately


Graceful shutdown is preferred when we don't want to
leave work half-completed.


========================================================
IMPORTANT REAL-WORLD DETAIL
========================================================

If `processJob()` performs something that can block
for a long time, it should ALSO accept the context.

For example:

	func processJob(ctx context.Context, job int) error

Then the operation itself can stop when:

	ctx.Done()

is triggered.

Otherwise, the worker may receive the shutdown signal
but remain stuck inside a long-running operation.


========================================================
SHORT INTERVIEW ANSWER
========================================================

"I'd register for SIGTERM using os/signal, cancel a shared
context to notify all 20 workers, and use a WaitGroup to
wait for them to finish and exit. Workers should select
on ctx.Done() so they can stop cleanly, and any long-running
operation should also support context cancellation."
========================================================
*/
