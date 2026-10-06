package interview

import (
	"fmt"
	"runtime"
	"sync"
)

/*
========================================================
QUESTION 16:
CAN GO RUN MULTIPLE GOROUTINES IN PARALLEL?
WHAT DETERMINES THAT?
========================================================

INTERVIEW ANSWER:

"Yes. Go can run multiple goroutines in parallel.

Whether they actually execute in parallel depends mainly
on the number of logical CPUs available to the Go runtime,
which is controlled by GOMAXPROCS.

Goroutines are concurrent by default, but they execute
in parallel when multiple logical processors are available."


========================================================
CONCURRENCY VS PARALLELISM
========================================================

Suppose we have:

	Goroutine A
	Goroutine B

With only 1 logical processor:

	CPU
	 |
	 +--> A
	 |
	 +--> B
	 |
	 +--> A
	 |
	 +--> B

They are CONCURRENT, but not executing at exactly
the same time.

With 2 logical processors:

	CPU 1              CPU 2
	  |                  |
	  +--> A             +--> B

Now A and B can execute in PARALLEL.


========================================================
WHAT DETERMINES PARALLEL EXECUTION?
========================================================

The main factor is:

	runtime.GOMAXPROCS

It determines how many logical processors can
execute Go code simultaneously.

You can check the current value:

*/

func question161() {

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("Available CPUs:", runtime.NumCPU())
}

/*
For example, if:

	runtime.NumCPU() = 8
	runtime.GOMAXPROCS(0) = 8

Go can potentially execute Go code on 8 logical
processors simultaneously.


========================================================
EXAMPLE
========================================================
*/

func example() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		fmt.Println("Goroutine A running")
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Goroutine B running")
	}()

	wg.Wait()
}

/*
A and B are CONCURRENT because they are both goroutines.

If GOMAXPROCS allows multiple logical processors,
the runtime can execute A and B in PARALLEL.


========================================================
IMPORTANT: GOROUTINE != CPU THREAD
========================================================

Do NOT say:

	"Each goroutine gets its own CPU thread."

That's incorrect.

Go has its own scheduler.

Conceptually:

	Goroutines
	    ↓
	Go Scheduler
	    ↓
	OS Threads
	    ↓
	CPU Cores / Logical CPUs


Many goroutines can be multiplexed onto fewer
OS threads.


========================================================
GOMAXPROCS EXAMPLE
========================================================

You can explicitly limit Go to one logical processor:

	runtime.GOMAXPROCS(1)

Now:

	Goroutine A
	Goroutine B
	Goroutine C

can still be CONCURRENT,

but only one can execute Go code at a time.


If:

	runtime.GOMAXPROCS(4)

then up to 4 logical processors can execute
Go code simultaneously.


========================================================
IMPORTANT INTERVIEW POINT
========================================================

GOMAXPROCS is NOT:

	"Maximum number of goroutines."

You can have:

	100,000 goroutines

with:

	GOMAXPROCS = 4

The 100,000 goroutines are scheduled across
the available logical processors.


========================================================
WHAT ABOUT I/O?
========================================================

Goroutines are especially useful for I/O-bound work.

For example:

	Goroutine A → waiting for database
	Goroutine B → making HTTP request
	Goroutine C → processing another request

While A is waiting, the scheduler can run B or C.

So even with one logical processor, many goroutines
can make progress concurrently.


========================================================
SHORT INTERVIEW ANSWER
========================================================

"Yes. Go can execute goroutines in parallel. Goroutines
are scheduled by the Go runtime onto OS threads, and
GOMAXPROCS determines how many logical processors can
execute Go code simultaneously. So concurrency is
provided by goroutines, while actual parallelism depends
on the available processors and GOMAXPROCS."


========================================================
EASY WAY TO REMEMBER
========================================================

	Goroutines
	    ↓
	CONCURRENCY

	Multiple logical processors
	    +
	GOMAXPROCS
	    ↓
	PARALLELISM
========================================================
*/

func question16() {
	question16()
	example()
}