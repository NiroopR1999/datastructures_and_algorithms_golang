package interview

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*
========================================================
QUESTION 14:
A goroutine is blocked forever waiting on a channel.
How would you detect and fix the problem?
========================================================

INTERVIEW ANSWER:

"I would first identify where the goroutine is blocked,
usually by looking at goroutine stack dumps or using
runtime/pprof.

Then I would check why the channel operation can never
complete.

The fix depends on the cause:

	- Make sure someone sends to the channel.
	- Make sure the channel is closed when appropriate.
	- Use context cancellation for operations that
	  should be cancellable.
	- Use a timeout if waiting indefinitely is not valid.


========================================================
EXAMPLE OF THE PROBLEM
========================================================
*/

func badExample() {

	ch := make(chan int)

	go func() {

		fmt.Println("waiting for value...")

		// Goroutine blocks here.
		value := <-ch

		fmt.Println("received:", value)
	}()

	// Nothing ever sends to ch.
	//
	// Therefore the goroutine stays blocked forever.
}

/*
========================================================
HOW DO WE DETECT IT?
========================================================

One way during debugging is to look at the number of
goroutines.

*/

func detectExample() {

	fmt.Println("goroutines:", runtime.NumGoroutine())

	/*
		For deeper investigation, Go provides:

			runtime/pprof

		It can capture goroutine stack traces.

		The stack trace can show something like:

			goroutine 20 [chan receive]:
			    main.badExample.func1()
			        question14.go:XX

		This tells us that the goroutine is stuck waiting
		on a channel receive.
	*/
}

/*
========================================================
FIX 1: MAKE SURE SOMEONE SENDS
========================================================
*/

func fixedExample1() {

	ch := make(chan int)

	go func() {

		// Send a value.
		ch <- 100

	}()

	// Receive the value.
	value := <-ch

	fmt.Println("received:", value)
}

/*
Now the communication has both sides:

	SENDER
	   |
	   v
	ch <- 100
	   |
	   v
	CHANNEL
	   |
	   v
	<-ch
	   |
	   v
	RECEIVER


========================================================
FIX 2: USE CONTEXT CANCELLATION
========================================================

Sometimes we DON'T want to wait forever.

For example, if the operation is no longer needed,
the goroutine should be able to stop.
*/

func fixedExample2() {

	ch := make(chan int)

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	defer cancel()

	go func() {

		select {

		case value := <-ch:

			fmt.Println("received:", value)

		case <-ctx.Done():

			// Context was cancelled.
			// Exit instead of waiting forever.
			fmt.Println("goroutine stopped")
			return
		}

	}()

	// Simulate deciding that the work is no longer needed.
	time.Sleep(100 * time.Millisecond)

	cancel()
}

/*
========================================================
FIX 3: USE A TIMEOUT
========================================================

If we only want to wait for a limited amount of time:
*/

func fixedExample3() {

	ch := make(chan int)

	select {

	case value := <-ch:

		fmt.Println("received:", value)

	case <-time.After(2 * time.Second):

		// Nobody sent a value within 2 seconds.
		fmt.Println("timed out")
	}
}

/*
========================================================
IMPORTANT INTERVIEW POINT
========================================================

Do NOT blindly close the channel just to fix the problem.

The real question is:

	"Who is responsible for sending or closing this channel?"

Usually:

	The sender owns the channel
	and closes it when there are no more values.

Also remember:

	Receiving from a closed channel does NOT block.

But:

	Sending to a closed channel PANICS.


========================================================
DEBUGGING MENTAL MODEL
========================================================

If you see:

	<-ch

ask:

	"Who is going to send to ch?"

If the answer is:

	"Nobody."

then you have found the problem.


If you see:

	ch <- value

ask:

	"Who is going to receive from ch?"

If the answer is:

	"Nobody."

then the goroutine can block forever.


========================================================
SHORT INTERVIEW ANSWER
========================================================

"I'd detect the blocked goroutine using goroutine stack
traces or pprof and identify the channel operation where
it is stuck. Then I'd determine why the corresponding
send or receive can never happen. Depending on the use
case, I'd fix the channel ownership/communication or
make the operation cancellable using context or a
timeout."
*/
