// Go Channels — Interview Questions
//
// Topics:
// - Buffered vs unbuffered channels
// - Blocking behavior
// - Closing channels
// - Receiving from closed channels
// - Sending to closed channels
// - Channel ownership
// - struct{}{} as a signal

package interview

import "fmt"

// ============================================================
// Question 6.1
// What's the difference between these?
//
//     make(chan int)
//     make(chan int, 5)
// ============================================================

func question6_1() {

	// Unbuffered channel.
	//
	// Capacity = 0.
	// There is no space to store values.
	ch1 := make(chan int)

	// Buffered channel.
	//
	// Capacity = 5.
	// It can store up to 5 values before a sender blocks.
	ch2 := make(chan int, 5)

	_ = ch1
	_ = ch2

	// INTERVIEW ANSWER:
	//
	// "An unbuffered channel has zero capacity, so the sender
	// and receiver have to synchronize directly. A send blocks
	// until another goroutine receives the value.
	//
	// A buffered channel has a fixed capacity, so the sender
	// can send values without a receiver being immediately ready,
	// until the buffer becomes full."

	// KEY POINT:
	//
	// Unbuffered:
	//     send blocks until receive
	//
	// Buffered:
	//     send blocks only when buffer is full
}

// ============================================================
// Question 6.2
// What happens here?
// Why doesn't "Hello" get printed?
// ============================================================

func question6_2() {

	ch := make(chan int)

	// This blocks because ch is unbuffered
	// and there is no receiver.
	ch <- 10

	// This line is never reached.
	fmt.Println("Hello")

	// INTERVIEW ANSWER:
	//
	// "The channel is unbuffered, so the send blocks until
	// another goroutine receives the value. Since there is
	// no receiver, the goroutine blocks forever and 'Hello'
	// is never printed. Go eventually reports a deadlock."

	// KEY POINT:
	//
	// Unbuffered channel + send + no receiver
	//                 =
	//               BLOCKS
}

// ============================================================
// Question 6.3
// What happens if you close a channel and then receive from it?
// ============================================================

func question6_3() {

	ch := make(chan int, 1)

	ch <- 10
	close(ch)

	// There is still a value inside the buffer,
	// so we can receive it normally.
	val := <-ch
	fmt.Println(val) // 10

	// The channel is closed AND empty now.
	//
	// Receiving does NOT block.
	// It immediately returns the zero value of int.
	val = <-ch
	fmt.Println(val) // 0

	// INTERVIEW ANSWER:
	//
	// "Receiving from a closed channel is allowed.
	// If there are buffered values remaining, we receive them
	// normally. Once the channel is closed and empty, receiving
	// immediately returns the zero value."

	// KEY POINT:
	//
	// Closed channel:
	//
	//     values remaining → receive normally
	//     empty            → zero value immediately
}

// ============================================================
// Question 6.4
// How do you distinguish a real zero value from
// a zero value caused by a closed channel?
// ============================================================

func question6_4() {

	ch := make(chan int)

	close(ch)

	val, ok := <-ch

	fmt.Println(val) // 0
	fmt.Println(ok)  // false

	// INTERVIEW ANSWER:
	//
	// "We can use the two-value receive syntax.
	// The second value tells us whether the receive got
	// an actual value from the channel.
	//
	// If ok is false, the channel is closed and empty."

	// KEY POINT:
	//
	//     val, ok := <-ch
	//
	//     ok == true  → received a value
	//     ok == false → channel is closed and empty
}

// ============================================================
// Question 6.5
// What happens if you close a channel and then send to it?
// ============================================================

func question6_5() {

	ch := make(chan int)

	close(ch)

	// PANIC:
	//
	// panic: send on closed channel
	ch <- 10

	// INTERVIEW ANSWER:
	//
	// "Sending to a closed channel causes a panic.
	// Once a channel is closed, no more values can be sent
	// to it."

	// KEY POINT:
	//
	//     send to open channel   → allowed
	//     send to closed channel → PANIC
}

// ============================================================
// Question 6.6
// Who should normally close a channel — sender or receiver?
// Why?
// ============================================================

func question6_6() {

	ch := make(chan int)

	// The sender knows when there are no more values to send.
	go func() {
		ch <- 10
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

	// INTERVIEW ANSWER:
	//
	// "Normally, the sender should close the channel because
	// the sender knows when it has finished sending values.
	//
	// The receiver should not close the channel because it
	// usually doesn't know whether another value will be sent."

	// KEY POINT:
	//
	// Sender → sends values → knows when finished → closes
	//
	// Receiver → receives values → normally does NOT close
	//
	// Also remember:
	//
	// Closing a channel means:
	// "No more values will be sent."
	//
	// It does NOT mean:
	// "The channel is deleted."
}

// ============================================================
// Question 6.7
// What does this do?
//
//     ch <- struct{}{}
// ============================================================

func question6_7() {

	ch := make(chan struct{})

	go func() {
		// Send an empty value as a signal.
		ch <- struct{}{}
	}()

	// Wait for the signal.
	<-ch

	fmt.Println("Signal received")

	// INTERVIEW ANSWER:
	//
	// "`struct{}{}` is an empty struct. It contains no data.
	// Sending it through the channel means:
	//
	//     'Something happened.'
	//
	// We are communicating an event or signal rather than
	// sending actual data."

	// KEY POINT:
	//
	//     ch <- struct{}{}
	//
	// means:
	//
	//     "Send a signal."
}

// ============================================================
// Question 6.8
// Why use struct{}{} instead of bool?
// ============================================================

func question6_8() {

	// We only care about whether the event happened.
	// We don't need to send true or false.
	signal := make(chan struct{})

	go func() {
		// Signal that the work is complete.
		signal <- struct{}{}
	}()

	// Wait for the signal.
	<-signal

	fmt.Println("Done")

	// INTERVIEW ANSWER:
	//
	// "We use struct{}{} when we only need to communicate
	// that an event happened and don't need to send any data.
	//
	// An empty struct has no fields, so it carries no meaningful
	// data. It is commonly used for signaling because we only
	// care about the event itself, not a value like true or false."

	// KEY POINT:
	//
	// bool:
	//     carries data → true / false
	//
	// struct{}:
	//     carries no data → only a signal
}

// ============================================================
// QUICK REVISION
// ============================================================
//
// 1. Unbuffered channel
//
//    make(chan int)
//
//    Send blocks until a receiver is ready.
//
//
// 2. Buffered channel
//
//    make(chan int, 5)
//
//    Can hold 5 values.
//    Send blocks when the buffer is full.
//
//
// 3. Receive from empty channel
//
//    Blocks until a value becomes available.
//
//
// 4. Receive from closed channel
//
//    Allowed.
//
//    If values remain → receive them.
//    If empty → zero value immediately.
//
//
// 5. Send to closed channel
//
//    PANIC.
//
//
// 6. Who closes?
//
//    Normally the sender.
//
//    Reason: sender knows when no more values will be sent.
//
//
// 7. range over channel
//
//    for value := range ch
//
//    Keeps receiving until the channel is closed.
//
//
// 8. struct{}{}
//
//    Empty value used when we only need to send a signal
//    and don't need to send actual data.
