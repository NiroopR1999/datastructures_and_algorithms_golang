// Go Deadlock Questions
//
// Tests:
// - Unbuffered channel blocking
// - Sender/receiver synchronization
// - Multiple goroutines waiting on each other
// - Identifying the exact point where a deadlock occurs

package interview

import "fmt"

// ============================================================
// Question 8.1
// Find the deadlock.
//
//     ch := make(chan int)
//
//     ch <- 10
//
//     go func() {
//         fmt.Println(<-ch)
//     }()
// ============================================================

func question8_1() {

	ch := make(chan int)

	// DEADLOCK HERE ❌
	//
	// ch is unbuffered.
	//
	// Sending 10 requires another goroutine to receive it.
	// But the receiver goroutine hasn't been started yet.
	ch <- 10

	// This goroutine is never created because the line above
	// blocks forever.
	go func() {
		fmt.Println(<-ch)
	}()

	// INTERVIEW ANSWER:
	//
	// "The deadlock occurs at ch <- 10.
	// Since ch is unbuffered, the send blocks until another
	// goroutine receives the value.
	//
	// The receiving goroutine is created only after the send,
	// but the program is already blocked on the send, so the
	// receiver is never started."

	// KEY POINT:
	//
	// Wrong:
	//
	//     send
	//     start receiver
	//
	// Correct:
	//
	//     start receiver
	//     send
}

// ============================================================
// Question 8.2
// Does this code deadlock?
//
//     ch := make(chan int)
//
//     go func() {
//         ch <- 10
//         ch <- 20
//     }()
//
//     fmt.Println(<-ch)
// ============================================================

func question8_2() {

	ch := make(chan int)

	go func() {

		// This send succeeds because main is receiving.
		ch <- 10

		// This send blocks because there is currently
		// no receiver waiting for another value.
		ch <- 20
	}()

	// Receive the first value.
	fmt.Println(<-ch)

	// OUTPUT:
	//
	//     10
	//
	// The producer goroutine is then blocked on ch <- 20.
	//
	// IMPORTANT:
	//
	// This is NOT a deadlock if main returns immediately
	// after receiving 10.
	//
	// The main goroutine exits, and the program terminates.
	//
	// However, if main also waited for the producer using
	// a WaitGroup, then it would deadlock because the producer
	// would remain blocked trying to send 20.

	// INTERVIEW ANSWER:
	//
	// "The first send of 10 succeeds because the main goroutine
	// is receiving it. The second send of 20 blocks because
	// there is no second receive.
	//
	// If the main goroutine waits for the producer to finish,
	// that creates a deadlock. If main simply exits, the program
	// terminates while the producer is still blocked."

	// KEY POINT:
	//
	// Every send on an unbuffered channel needs a corresponding
	// receive.
	//
	//     ch <- 10  ↔  <-ch
	//     ch <- 20  ↔  <-ch
}

// ============================================================
// Question 8.3
// Why does this deadlock?
//
//     ch1 := make(chan int)
//     ch2 := make(chan int)
//
//     go func() {
//         ch1 <- 10
//         fmt.Println(<-ch2)
//     }()
//
//     go func() {
//         ch2 <- 20
//         fmt.Println(<-ch1)
//     }()
// ============================================================

func question8_3() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		// BLOCKS HERE ❌
		//
		// ch1 is unbuffered.
		// This goroutine is waiting for somebody to receive 10.
		ch1 <- 10

		// This line is never reached.
		fmt.Println(<-ch2)
	}()

	go func() {

		// BLOCKS HERE ❌
		//
		// ch2 is unbuffered.
		// This goroutine is waiting for somebody to receive 20.
		ch2 <- 20

		// This line is also never reached.
		fmt.Println(<-ch1)
	}()

	// INTERVIEW ANSWER:
	//
	// "Both goroutines are blocked trying to send on separate
	// unbuffered channels.
	//
	// The first goroutine is waiting for somebody to receive
	// from ch1.
	//
	// The second goroutine is waiting for somebody to receive
	// from ch2.
	//
	// But both goroutines are stuck on their sends, so neither
	// reaches the receive operation."

	// The situation is:
	//
	//     Goroutine 1                 Goroutine 2
	//
	//     ch1 <- 10  ← BLOCKED       ch2 <- 20  ← BLOCKED
	//          │                           │
	//          │                           │
	//          ↓                           ↓
	//     waiting for                  waiting for
	//     ch1 receiver                 ch2 receiver
	//
	// Neither can continue.
	//
	// Therefore:
	//
	//     DEADLOCK ❌
}

// ============================================================
// QUICK REVISION
// ============================================================
//
// Deadlock #1:
//
//     ch <- 10
//     go receiver()
//
// The sender blocks before the receiver goroutine is created.
//
//
// Deadlock #2:
//
//     goroutine:
//         ch <- 10
//         ch <- 20
//
//     main:
//         <-ch
//
// First send succeeds.
// Second send blocks because there is no second receive.
//
// If main waits for the sender → DEADLOCK.
//
//
// Deadlock #3:
//
//     Goroutine 1:
//         ch1 <- 10
//         <-ch2
//
//     Goroutine 2:
//         ch2 <- 20
//         <-ch1
//
// Both goroutines block on their first send.
// Neither reaches its receive.
//
//
// GENERAL RULE:
//
// For an unbuffered channel:
//
//     SEND  ↔  RECEIVE
//
// A send needs a receiver.
// A receive needs a sender.
//
// If every goroutine is waiting for another goroutine
// and nobody can make progress → DEADLOCK.