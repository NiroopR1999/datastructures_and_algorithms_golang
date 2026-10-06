package interview

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
========================================================
QUESTION 13: HOW DO YOU PREVENT GOROUTINE LEAKS?
========================================================

INTERVIEW ANSWER:

"I prevent goroutine leaks by making sure every goroutine
has a clear exit condition.

For goroutines that can block, I use context cancellation,
timeouts, and select statements so they can stop when the
work is no longer needed.

I also use WaitGroup when I need to track goroutine
completion."


========================================================
WHAT IS A GOROUTINE LEAK?
========================================================

A goroutine leak happens when a goroutine keeps running
or stays blocked even though nobody needs it anymore.

For example:
*/

func badExample1() {

	ch := make(chan int)

	go func() {

		// This goroutine waits forever.
		//
		// Nobody sends anything to ch.
		value := <-ch

		fmt.Println(value)
	}()

	// The goroutine is stuck forever.
}

/*
The problem is:

	goroutine
	    |
	    v
	<-ch
	    |
	    v
	WAIT FOREVER

There is no way for this goroutine to exit.


========================================================
SOLUTION 1: GIVE THE GOROUTINE AN EXIT CONDITION
========================================================
*/

func worker12(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			// Context was cancelled.
			// Stop the goroutine.
			return

		default:

			// Do some work.
			fmt.Println("working")

			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
Usage:

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	// Later...
	cancel()

The flow becomes:

	cancel()
	   |
	   v
	ctx.Done()
	   |
	   v
	worker receives cancellation
	   |
	   v
	return
	   |
	   v
	goroutine exits


========================================================
SOLUTION 2: MAKE CHANNEL OPERATIONS CANCELLABLE
========================================================

BAD:

	go func() {

		result := <-resultCh

		fmt.Println(result)
	}()

If nobody sends to resultCh, the goroutine waits forever.

BETTER:
*/

func cancellableReceiver(
	ctx context.Context,
	resultCh <-chan int,
) {

	go func() {

		select {

		case result := <-resultCh:

			fmt.Println("received:", result)

		case <-ctx.Done():

			// Caller no longer needs the result.
			return
		}
	}()
}

/*
Now the goroutine can exit in two ways:

	1. Result arrives
	2. Context is cancelled


========================================================
SOLUTION 3: MAKE CHANNEL SENDS CANCELLABLE
========================================================

This can also cause a goroutine leak:

	go func() {
		resultCh <- result
	}()

If resultCh is unbuffered and nobody receives,
the goroutine blocks forever on the send.

Better:
*/

func cancellableSender(
	ctx context.Context,
	resultCh chan<- int,
) {

	go func() {

		result := 100

		select {

		case resultCh <- result:

			// Result successfully sent.

		case <-ctx.Done():

			// Receiver/caller no longer needs
			// the result.
			return
		}
	}()
}

/*
========================================================
SOLUTION 4: USE TIMEOUTS
========================================================

External operations such as:

	- HTTP calls
	- Database calls
	- RPC calls

should generally have a timeout.

Example:
*/

func callWithTimeout() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	_=ctx

	defer cancel()

	// Pass ctx to an operation that supports
	// context cancellation.
	//
	// example:
	//
	// databaseCall(ctx)
	// httpRequestWithContext(ctx)
}

/*
If the operation takes longer than 2 seconds:

	ctx.Done()
	    |
	    v
	operation can stop
	    |
	    v
	goroutine can exit


========================================================
WAITGROUP VS CONTEXT
========================================================

These solve DIFFERENT problems.

WaitGroup:

	"Has this goroutine finished?"

Context:

	"Should this goroutine stop?"


Example:
*/

func waitGroupExample() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		fmt.Println("doing work")

	}()

	// Wait until the goroutine finishes.
	wg.Wait()
}

/*
Important:

WaitGroup by itself DOES NOT prevent goroutine leaks.

If the goroutine gets stuck forever:

	wg.Wait()

will also wait forever.


========================================================
COMMON CAUSES OF GOROUTINE LEAKS
========================================================

1. Waiting forever on a channel

	go func() {
		<-ch
	}()

If nobody sends → stuck forever.


2. Sending forever to a channel

	go func() {
		ch <- value
	}()

If nobody receives → stuck forever.


3. Infinite loop without an exit condition

	go func() {
		for {
			doWork()
		}
	}()


4. No context cancellation

A request finishes, but a background goroutine
created for that request continues running.


5. External operation without timeout

A network/DB operation can wait indefinitely.


6. Waiting for a channel that is never closed

	for value := range ch {
		process(value)
	}

If nobody closes ch, the range can wait forever.


========================================================
MOST IMPORTANT INTERVIEW RULE
========================================================

Whenever you create a goroutine, immediately ask:

	"HOW DOES THIS GOROUTINE STOP?"

If you cannot answer that clearly,
you may have a goroutine leak.


========================================================
SHORT INTERVIEW ANSWER
========================================================

"I prevent goroutine leaks by giving every goroutine
a clear exit condition. For blocking operations, I use
context cancellation, timeouts, and select so the
goroutine can stop when the work is no longer needed.

I use WaitGroup separately when I need to track whether
goroutines have completed."
*/
