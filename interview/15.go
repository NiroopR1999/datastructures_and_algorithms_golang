package interview

import (
	"fmt"
	"sync"
	"time"
)

/*
========================================================
QUESTION 15:
WHAT'S THE DIFFERENCE BETWEEN CONCURRENCY
AND PARALLELISM?
========================================================

INTERVIEW ANSWER:

"Concurrency means multiple tasks can make progress
independently, while parallelism means multiple tasks
are actually executing at the same time.

Concurrency is about STRUCTURING multiple tasks.
Parallelism is about EXECUTING multiple tasks
simultaneously."


========================================================
CONCURRENCY
========================================================

Concurrency does NOT necessarily mean tasks execute
at the exact same time.

Example:

	Task A: ████      ████
	Task B:     ████      ████

The CPU can switch between tasks.

In Go, goroutines give us concurrency.

*/

func concurrencyExample() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Println("Task A:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 3; i++ {
			fmt.Println("Task B:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}

/*
The scheduler may execute:

	Task A
	Task B
	Task A
	Task B
	Task A
	Task B

The tasks are CONCURRENT because both can make progress
independently.

They don't necessarily need two CPU cores.


========================================================
PARALLELISM
========================================================

Parallelism means two tasks are executing at the
same time on different CPU cores.

For example, with 2 CPU cores:

	CPU CORE 1              CPU CORE 2

	Task A █████████         Task B █████████
	       ↑                       ↑
	   executing               executing
	   simultaneously

This is actual parallel execution.


========================================================
GO EXAMPLE
========================================================

Go can use multiple CPU cores to execute goroutines
in parallel.

For example:

*/

func parallelismExample() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		// Task A
		for i := 0; i < 1_000_000; i++ {
		}

		fmt.Println("Task A finished")
	}()

	go func() {
		defer wg.Done()

		// Task B
		for i := 0; i < 1_000_000; i++ {
		}

		fmt.Println("Task B finished")
	}()

	wg.Wait()
}

/*
If Go has multiple CPUs available, Task A and Task B
may execute simultaneously on different CPU cores.


========================================================
THE IMPORTANT DIFFERENCE
========================================================

CONCURRENCY:

	Multiple tasks are in progress.

	Focus:
	"How do I structure multiple tasks?"

PARALLELISM:

	Multiple tasks execute at the same time.

	Focus:
	"How do I execute multiple tasks simultaneously?"


========================================================
REAL-WORLD ANALOGY
========================================================

CONCURRENCY:

	One chef handling:

		- cooking
		- preparing vegetables
		- checking the oven

The chef switches between tasks.

PARALLELISM:

	Two chefs:

		Chef 1 → cooking
		Chef 2 → preparing vegetables

Both work at the same time.


========================================================
IMPORTANT GO POINT
========================================================

Goroutines provide CONCURRENCY.

Whether they execute in PARALLEL depends on the
available CPU resources and Go's scheduler/runtime.

You can control the maximum number of CPUs Go uses with:

	runtime.GOMAXPROCS()

For example:

	runtime.GOMAXPROCS(2)

This allows Go to execute Go code simultaneously
on up to 2 logical processors.


========================================================
INTERVIEW ONE-LINER
========================================================

"Concurrency is about dealing with multiple tasks at
once, while parallelism is about executing multiple
tasks at the same time."


========================================================
EASY WAY TO REMEMBER
========================================================

	CONCURRENCY
	    ↓
	"Multiple things in progress"

	PARALLELISM
	    ↓
	"Multiple things executing simultaneously"
*/
func question15() {
	concurrencyExample()
	parallelismExample()
}