// Goroutine + WaitGroup
//
// Tests:
// - WaitGroup.Add()
// - WaitGroup.Done()
// - WaitGroup.Wait()
// - Goroutine synchronization
// - Loop variable capture
// - WaitGroup misuse

package interview

import (
	"fmt"
	"sync"
)

func question7_1() {

	// ============================================================
	// Question 7.1
	// What's wrong with this code?
	// Identify all problems.
	// ============================================================

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		go func() {

			fmt.Println(i)

			wg.Done()

		}()

	}

	wg.Wait()

	// PROBLEM 1:
	//
	// wg.Add() was never called.
	//
	// A WaitGroup starts with a counter of 0.
	// Therefore, Wait() does not wait for these goroutines.
	//
	// Correct:
	//
	//     wg.Add(5)
	//
	// before starting the goroutines.
	//
	//
	// PROBLEM 2:
	//
	// wg.Done() is being called even though the counter
	// was never incremented.
	//
	// This can cause:
	//
	//     panic: sync: negative WaitGroup counter
	//
	//
	// PROBLEM 3:
	//
	// The goroutine uses the loop variable i directly.
	//
	// For interview purposes, remember that loop-variable
	// capture has changed with newer Go versions.
	//
	// In older Go versions, all goroutines could observe
	// the same loop variable and potentially print 5.
	//
	// A safe and explicit approach is to pass i as an argument:
	//
	//     go func(i int) {
	//         fmt.Println(i)
	//         wg.Done()
	//     }(i)
	//
	// This gives each goroutine its own value.
}

// ============================================================
// Question 7.2
// Why do we call wg.Add() before starting the goroutine?
// ============================================================

func question7_2() {

	var wg sync.WaitGroup

	// Add to the counter BEFORE starting the goroutine.
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Worker finished")
	}()

	wg.Wait()

	// INTERVIEW ANSWER:
	//
	// "Add() tells the WaitGroup how many goroutines we are
	// waiting for. We call Add() before starting the goroutine
	// so that Wait() cannot observe a counter of zero before
	// the goroutine has incremented the counter."
	//
	// In simple terms:
	//
	//     Add(1)  → "One goroutine needs to finish."
	//     Done()  → "That goroutine has finished."
	//     Wait()  → "Wait until the counter reaches zero."
}

// ============================================================
// Question 7.3
// What happens if wg.Done() is called more times than wg.Add()?
// ============================================================

func question7_3() {

	var wg sync.WaitGroup

	wg.Add(1)

	wg.Done()

	// Counter is now 0.
	//
	// Calling Done() again would make the counter negative:
	//
	// wg.Done()
	//
	// Result:
	//
	//     panic: sync: negative WaitGroup counter
}

// ============================================================
// Question 7.4
// What happens if you forget wg.Done()?
// ============================================================

func question7_4() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		// Work happens here...

		// wg.Done() is missing.
	}()

	// Waits forever because the counter is still 1.
	wg.Wait()

	// INTERVIEW ANSWER:
	//
	// "If Done() is never called, the WaitGroup counter never
	// reaches zero, so Wait() blocks forever."
	//
	// That's why this pattern is safer:
	//
	//     go func() {
	//         defer wg.Done()
	//
	//         // work
	//     }()
}

// ============================================================
// Question 7.5
// Can a WaitGroup be copied after it has started being used?
// Why?
// ============================================================

func question7_5() {

	var wg sync.WaitGroup

	wg.Add(1)

	// DO NOT DO THIS:
	//
	// wg2 := wg
	//
	// A WaitGroup must not be copied after first use.

	wg.Done()

	// INTERVIEW ANSWER:
	//
	// "No. A WaitGroup must not be copied after it has been
	// used. It contains internal synchronization state, and
	// copying that state can lead to incorrect synchronization
	// and undefined behavior."
	//
	// If multiple functions need to operate on the same
	// WaitGroup, pass a pointer:
	//
	//     func worker(wg *sync.WaitGroup)
	//
	// rather than passing the WaitGroup by value.
}

// ============================================================
// Correct version of Question 7.1
// ============================================================

func question7_correct() {

	var wg sync.WaitGroup

	// We are starting 5 goroutines,
	// so the counter starts at 5.
	wg.Add(5)

	for i := 0; i < 5; i++ {

		// Pass i as an argument so this goroutine
		// receives its own copy of the value.
		go func(i int) {
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	// Wait until all 5 goroutines call Done().
	wg.Wait()

	// INTERVIEW ANSWER:
	//
	// "I increment the WaitGroup counter before starting the
	// goroutines. Each goroutine calls Done() when it finishes,
	// and Wait() blocks until the counter reaches zero."
}

// ============================================================
// QUICK REVISION
// ============================================================
//
// WaitGroup has three main operations:
//
//     wg.Add(n)
//         ↓
//     "I am waiting for n goroutines."
//
//
//     wg.Done()
//         ↓
//     "One goroutine has finished."
//     Equivalent to Add(-1).
//
//
//     wg.Wait()
//         ↓
//     "Block until the counter becomes zero."
//
//
// IMPORTANT RULES:
//
// 1. Call Add() BEFORE starting goroutines.
//
// 2. Every Add(1) must eventually have a corresponding Done().
//
// 3. defer wg.Done() is a safe pattern:
//
//        go func() {
//            defer wg.Done()
//            // work
//        }()
//
// 4. Calling Done() too many times can cause:
//
//        panic: sync: negative WaitGroup counter
//
// 5. Forgetting Done() causes Wait() to block forever.
//
// 6. Don't copy a WaitGroup after it has been used.
//
// 7. Pass *sync.WaitGroup when sharing it with functions.
//
// 8. WaitGroup is for synchronization.
//    It does NOT provide communication between goroutines.
//
//    For communication → use channels.
//    For waiting for completion → use WaitGroup.