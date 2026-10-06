package interview

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
========================================================
QUESTION 17:
WHAT HAPPENS WHEN YOU START 100,000 GOROUTINES?
IS THAT EQUIVALENT TO CREATING 100,000 OS THREADS?
========================================================

INTERVIEW ANSWER:

"No. 100,000 goroutines do NOT mean 100,000 OS threads.

Goroutines are lightweight and are managed by the Go
runtime scheduler. The runtime multiplexes many goroutines
onto a much smaller number of OS threads.

The number of goroutines and the number of OS threads
are therefore different things."


========================================================
EXAMPLE: 100,000 GOROUTINES
========================================================
*/

func question17() {

	var wg sync.WaitGroup

	wg.Add(100_000)

	for i := 0; i < 100_000; i++ {

		go func(id int) {
			defer wg.Done()

			// Simulate some work.
			time.Sleep(time.Second)

		}(i)
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()
}

/*
At this point, we could have approximately:

	100,000 goroutines

But NOT:

	100,000 OS threads


========================================================
WHAT ACTUALLY HAPPENS?
========================================================

Think of it like this:

	100,000 Goroutines
	        |
	        v
	+-------------------+
	| Go Runtime         |
	| Scheduler          |
	+-------------------+
	        |
	        v
	  OS Threads
	    /   |   \
	   /    |    \
	  v     v     v
	 CPU   CPU   CPU


The Go scheduler decides which goroutine should
run on which OS thread.


========================================================
WHY CAN GO HANDLE SO MANY GOROUTINES?
========================================================

A goroutine starts with a relatively small stack,
and its stack can grow/shrink as needed.

OS threads generally have significantly more
overhead.

Therefore creating:

	100,000 goroutines

is usually much cheaper than creating:

	100,000 OS threads.


========================================================
IMPORTANT:
100,000 GOROUTINES CAN STILL BE A PROBLEM
========================================================

"Lightweight" does NOT mean "free".

Every goroutine still consumes memory and runtime
resources.

For example:

	go func() {
		for {
			// Never exits
		}
	}()

If you accidentally create thousands of these,
you can eventually exhaust memory/resources.


So the real question is:

	"Do these goroutines have useful work
	 and a clear lifecycle?"


========================================================
GOMAXPROCS IS DIFFERENT
========================================================

Suppose:

	runtime.GOMAXPROCS(4)

and:

	100,000 goroutines

This does NOT mean only 4 goroutines exist.

It means up to 4 logical processors can execute
Go code simultaneously.

Conceptually:

	100,000 goroutines
	        |
	        v
	   Go Scheduler
	        |
	  +-----+-----+-----+
	  |     |     |     |
	 CPU1  CPU2  CPU3  CPU4

The remaining goroutines wait to be scheduled
or may be blocked on I/O, timers, channels, etc.


========================================================
GOMAXPROCS vs GOROUTINES
========================================================

Goroutines:

	"How many units of work/tasks do I have?"

GOMAXPROCS:

	"How many logical processors can execute
	 Go code simultaneously?"


Example:

	100,000 goroutines
	GOMAXPROCS = 4

means:

	100,000 concurrent goroutines

but at most:

	4 logical processors executing Go code
	simultaneously.


========================================================
CONCURRENCY VS PARALLELISM
========================================================

100,000 goroutines give you the ability to
structure a huge amount of CONCURRENT work.

They do NOT imply:

	100,000 things executing simultaneously.


Actual parallel execution is limited by the
available logical processors / GOMAXPROCS.


========================================================
COMMON INTERVIEW TRAP
========================================================

QUESTION:

	"If I create 100,000 goroutines,
	 do I create 100,000 threads?"

WRONG:

	"Yes, each goroutine creates a thread."

CORRECT:

	"No. Goroutines are user-space managed by the
	 Go runtime and are multiplexed onto OS threads."


========================================================
SHORT INTERVIEW ANSWER
========================================================

"100,000 goroutines do not create 100,000 OS threads.
Goroutines are lightweight units managed by the Go
runtime scheduler, which multiplexes them onto a much
smaller number of OS threads. The number of goroutines
controls how much concurrent work we can have, while
GOMAXPROCS determines how many logical processors can
execute Go code in parallel."


========================================================
EASY WAY TO REMEMBER
========================================================

	Goroutines
	    ↓
	LOTS of lightweight tasks

	Go Scheduler
	    ↓
	decides what runs

	OS Threads
	    ↓
	fewer execution resources

	CPU Cores
	    ↓
	actual parallel execution
========================================================
*/

func callQ17() {
	question17()
}