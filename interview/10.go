// Mutex vs Channels
//
// package interview

package interview

import (
	"fmt"
	"sync"
)

func question10() {

	// ============================================================
	// QUESTION 10.1
	//
	// When would you use a sync.Mutex instead of a channel?
	// ============================================================

	// INTERVIEW ANSWER:
	//
	// "I would use a Mutex when multiple goroutines need to access
	// or modify the same shared state and I simply need to protect
	// that state from concurrent access."
	//
	// A channel is more appropriate when goroutines need to
	// communicate, send data, transfer ownership, or coordinate work.
	//
	// Simple rule:
	//
	//     Shared state → Mutex
	//     Communication → Channel
	//
	// Example:
	//
	//     counter++
	//
	// This is shared state, so a Mutex is usually the simpler
	// and more natural solution.


	// ============================================================
	// QUESTION 10.2
	//
	// Protect this counter from concurrent access:
	//
	//     counter := 0
	//
	//     for i := 0; i < 100; i++ {
	//
	//         go func() {
	//
	//             counter++
	//
	//         }()
	//
	//     }
	// ============================================================

	// PROBLEM:
	//
	// counter++ is NOT one indivisible operation.
	//
	// Conceptually, it is:
	//
	//     1. Read counter
	//     2. Add 1
	//     3. Write counter
	//
	// Two goroutines can do this at the same time:
	//
	//     Goroutine 1: reads 0
	//     Goroutine 2: reads 0
	//
	//     Goroutine 1: writes 1
	//     Goroutine 2: writes 1
	//
	// Expected:
	//
	//     2
	//
	// Actual:
	//
	//     1
	//
	// This is a DATA RACE.
	//
	// We need to protect the critical section.


	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {

		go func() {
			defer wg.Done()

			// Lock before accessing the shared variable.
			//
			// If another goroutine already holds the lock,
			// this goroutine waits here.
			mu.Lock()

			counter++

			// Unlock so another goroutine can access counter.
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)

	// INTERVIEW ANSWER:
	//
	// "The Mutex ensures that only one goroutine can execute
	// counter++ at a time, preventing concurrent read/write
	// access to the shared variable."


	// ============================================================
	// QUESTION 10.3
	//
	// Implement the same counter using a channel instead
	// of a Mutex.
	// ============================================================

	// A channel can be used to ensure that only one goroutine
	// owns/accesses the counter at a time.
	//
	// Here we use a channel as a "token".
	//
	// Since the channel has capacity 1, only one token exists.
	//
	// Whoever receives the token gets permission to modify
	// the counter.
	//
	// After modifying it, the goroutine puts the token back.


	counter = 0

	// Buffered channel with capacity 1.
	//
	// Think of the single value inside this channel as:
	//
	//     "permission to access counter"
	//
	token := make(chan struct{}, 1)

	// Put the initial token into the channel.
	token <- struct{}{}

	wg = sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {

		go func() {
			defer wg.Done()

			// Receive the token.
			//
			// If another goroutine currently has the token,
			// this goroutine blocks here.
			<-token

			// Only the goroutine holding the token modifies
			// the counter.
			counter++

			// Return the token so another goroutine can
			// access the counter.
			token <- struct{}{}
		}()
	}

	wg.Wait()

	fmt.Println("Counter using channel:", counter)


	// ============================================================
	// WHY WOULD YOU CHOOSE MUTEX HERE?
	// ============================================================

	// INTERVIEW ANSWER:
	//
	// "Although a channel can be used to protect the counter,
	// I would choose a Mutex here because the problem is simply
	// protecting shared memory.
	//
	// There is no data that needs to be communicated between
	// goroutines. The Mutex directly expresses the intention:
	// only one goroutine can access this critical section at
	// a time."
	//
	// A channel solution works, but it introduces communication
	// just to provide mutual exclusion.
	//
	// Therefore:
	//
	//     Mutex → protect shared state
	//     Channel → communicate / transfer data / coordinate work


	// ============================================================
	// IMPORTANT INTERVIEW DISTINCTION
	// ============================================================

	// Mutex:
	//
	//     "Don't let two goroutines access this shared state
	//      at the same time."
	//
	//
	// Channel:
	//
	//     "Send this value/event/work from one goroutine
	//      to another."
	//
	//
	// Don't say:
	//
	//     "Channels are always better than Mutexes."
	//
	// or:
	//
	//     "Mutexes are always better than channels."
	//
	// The choice depends on what problem you are solving.


	// ============================================================
	// QUICK INTERVIEW ANSWER
	// ============================================================
	//
	// Q: Why did you choose a Mutex instead of a channel?
	//
	// A:
	//
	// "Because I'm protecting shared state rather than
	// communicating data between goroutines. A Mutex is the
	// simpler and more direct synchronization primitive for
	// protecting the critical section."
}