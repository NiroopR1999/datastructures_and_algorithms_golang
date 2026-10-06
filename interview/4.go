// What happens when you send to an unbuffered channel when nobody is receiving?
//
// Test:
// - Unbuffered channel behavior
// - Blocking send
// - Deadlock
//
// Expected behavior:
// The program blocks on the first send because an unbuffered channel
// requires a receiver to be ready before the send can complete.

package interview

func question4() {
	ch := make(chan int)

	// This send blocks because there is no goroutine receiving
	// from ch.
	//
	// Since ch is unbuffered, the value cannot be stored anywhere.
	ch <- 1

	// This line is never reached because the program is
	// already blocked on the send above.
	close(ch)
}

// Sending to an unbuffered channel blocks until another goroutine receives the value.
// Since nobody is receiving here, the goroutine blocks forever and Go reports a deadlock.