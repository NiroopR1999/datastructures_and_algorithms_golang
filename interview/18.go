package interview

import (
	"fmt"
	"sync"
)

/*
========================================================
QUESTION 18:
WHY CAN THIS CODE PRODUCE UNEXPECTED RESULTS?
========================================================

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

========================================================
THE PROBLEM
========================================================

The goroutine is using the loop variable `i` from the
surrounding scope.

The goroutine does NOT receive a copy of `i` when it is
created.

It accesses the variable that the loop is modifying.

The goroutine may execute later, after the loop has
already changed `i`.

========================================================
IMPORTANT MODERN GO NOTE
========================================================

Starting with Go 1.22, loop variables declared by the
loop are created separately for each iteration.

Therefore, with modern Go versions, this code does NOT
have the old classic loop-variable capture bug.

However, this is still an important interview question
because:

	1. Older Go versions had the bug.
	2. You may encounter older codebases.
	3. The underlying concept of closure variable capture
	   is still important.

========================================================
OLD GO BEHAVIOR
========================================================

In older Go versions, all goroutines could capture the
same `i`.

For example, the loop could finish:

	i = 10

before the goroutines actually execute.

Then several goroutines could print:

	10
	10
	10
	...

This is why the old code could produce unexpected results.


========================================================
SAFE / EXPLICIT SOLUTION
========================================================

Pass `i` as an argument to the goroutine.

*/

func question18() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func(i int) {

			defer wg.Done()

			// `i` is now this goroutine's own
			// parameter value.
			fmt.Println(i)

		}(i)
	}

	wg.Wait()
}

/*
========================================================
WHY DOES THIS FIX IT?
========================================================

Look at:

	go func(i int) {
		fmt.Println(i)
	}(i)

The final `(i)` means:

	"Take the current value of i
	 and pass it into this function."

Example:

Iteration 0:

	go func(i int) {...}(0)

Iteration 1:

	go func(i int) {...}(1)

Iteration 2:

	go func(i int) {...}(2)

...

Each goroutine receives its own parameter value.


========================================================
CLOSURE CONCEPT
========================================================

A closure is a function that can access variables
from its surrounding scope.

Example:

*/

func closureExample() {

	x := 10

	go func() {

		// This function can access `x`
		// from the surrounding scope.
		fmt.Println(x)

	}()
}

/*
The goroutine is therefore "capturing" x.

This is useful, but you need to understand which
variables the goroutine is accessing and how their
lifetime/value changes.


========================================================
INTERVIEW ANSWER
========================================================

"The goroutine function is a closure and accesses the
loop variable `i` from the surrounding scope. In older
Go versions, all goroutines could capture the same loop
variable, so they might observe a later value such as 10.

The explicit and interview-safe approach is to pass the
loop variable as a function argument:

	go func(i int) {
		fmt.Println(i)
	}(i)

That gives each goroutine its own parameter value.

In Go 1.22+, loop variables declared by the loop have
per-iteration scope, so the classic bug is fixed for
this form of loop."


========================================================
IMPORTANT:
WHY USE WAITGROUP HERE?
========================================================

Without:

	wg.Wait()

the main goroutine may finish before the goroutines
get a chance to print.

So there are actually TWO separate concepts:

	1. Loop-variable capture
	   → Which value does the goroutine see?

	2. WaitGroup
	   → Does main wait for the goroutines to finish?


========================================================
EASY WAY TO REMEMBER
========================================================

OLD PROBLEM:

	go func() {
		fmt.Println(i)
	}()

	↓
	goroutine accesses outer `i`

SAFE EXPLICIT VERSION:

	go func(i int) {
		fmt.Println(i)
	}(i)

	↓
	current `i` is passed as an argument


INTERVIEW RULE:

	"When a goroutine uses a loop variable,
	be aware of closure capture.
	Passing the variable as an argument makes
	the intended value explicit."
*/

func call18() {
	question18()
}