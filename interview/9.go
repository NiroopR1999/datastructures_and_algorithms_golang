// Go select
//
// Tests:
// - Receiving from multiple channels
// - Blocking behavior
// - Multiple ready cases
// - default case
// - Non-blocking channel operations
//
// IMPORTANT IDEA:
//
// select lets a goroutine wait on multiple channel operations
// at the same time.
//
// It is similar to switch, but instead of checking values,
// select checks which channel operation is ready.

package interview

// import "fmt"

func question9() {

	// ============================================================
	// Question 9.1
	// What does select do in Go?
	// ============================================================

	// Example:
	//
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	//
	// select {
	// case val := <-ch1:
	//     fmt.Println("Received from ch1:", val)
	//
	// case val := <-ch2:
	//     fmt.Println("Received from ch2:", val)
	// }

	// INTERVIEW ANSWER:
	//
	// "select lets a goroutine wait on multiple channel
	// operations. It executes the case whose channel operation
	// is ready.
	//
	// If none of the cases is ready and there is no default,
	// select blocks until at least one operation becomes ready."

	// Think:
	//
	//     select
	//       │
	//       ├── ch1 ready? → use ch1
	//       │
	//       └── ch2 ready? → use ch2


	// =====================================


	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
		// defer wg.Done()

		// Both channels are unbuffered.
		//
		// This send blocks until the select below
		// receives from ch1.
		// ch1 <- 10

		// After ch1 <- 10 succeeds, the goroutine continues.
		//
		// This send now blocks until something receives
		// from ch2.
		// ch2 <- 20
	// }()

	// select waits until at least one of these channel
	// operations is ready.
	//
	// Since the goroutine sends to ch1 first, ch1 will
	// normally become ready first.
	//
	// The receive from ch1 synchronizes with:
	//
	//     ch1 <- 10
	//
	// and prints:
	//
	//     received from channel ch1 10
	//
	// IMPORTANT:
	//
	// select executes ONLY ONE case and then exits.
	// It does NOT continue waiting for ch2.
	// select {
	// case val := <-ch1:
	// 	fmt.Println("received from channel ch1", val)

	// case val := <-ch2:
	// 	fmt.Println("received from channel ch2", val)
	// }

	// At this point, select has finished.
	//
	// But the goroutine may now be stuck here:
	//
	//     ch2 <- 20
	//
	// because there is no longer a receiver waiting on ch2.
	//
	// Therefore, the goroutine cannot reach wg.Done().
	//
	// wg.Wait() waits for the WaitGroup counter to become 0,
	// but it remains 1.
	//
	// Result:
	//
	//     DEADLOCK
	// wg.Wait()
	// ========================================

	// ============================================================
	// Question 9.2
	// What happens if multiple channels are ready at the same time?
	// ============================================================

	// INTERVIEW ANSWER:
	//
	// "If multiple cases are ready at the same time, select
	// chooses one of them pseudo-randomly."
	//
	// Therefore, you should NOT assume that the first case
	// in the select will always execute.
	//
	// Example:
	//
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func() {
	// 	ch1 <- 10
	// }()

	// go func() {
	// 	ch2 <- 20
	// }()

	// select {
	// case val := <-ch1:
	// 	fmt.Println("received from channel ch1", val)
	// case val := <-ch2:
	// 	fmt.Println("received from channel ch2", val)
	// }
	//
	// If both ch1 and ch2 are ready, either case can be selected.


	// ============================================================
	// Question 9.3
	// What's the purpose of default in a select?
	// ============================================================

	// Example:
	//
	// select {
	// case msg := <-ch:
	//     fmt.Println(msg)
	//
	// default:
	//     fmt.Println("No message")
	// }

	// INTERVIEW ANSWER:
	//
	// "default makes the select non-blocking.
	//
	// If a channel operation is ready, that case executes.
	// If no channel operation is ready, default executes
	// immediately instead of waiting."
	//
	// Without default:
	//
	//     select {
	//     case msg := <-ch:
	//         fmt.Println(msg)
	//     }
	//
	// The goroutine blocks until ch has a value.
	//
	// With default:
	//
	//     select {
	//     case msg := <-ch:
	//         fmt.Println(msg)
	//     default:
	//         fmt.Println("No message")
	//     }
	//
	// The goroutine does NOT wait.


	// ============================================================
	// Question 9.4
	// What's wrong with this?
	//
	//     select {
	//     case msg := <-ch:
	//         fmt.Println(msg)
	//     default:
	//         fmt.Println("No message")
	//     }
	// ============================================================

	// NOTHING IS INHERENTLY WRONG.
	//
	// It is useful when you intentionally want a
	// non-blocking receive.
	//
	// However, you need to understand that it does NOT wait.
	//
	// If ch doesn't currently have a value, default executes
	// immediately.
	//
	// INTERVIEW ANSWER:
	//
	// "There is nothing wrong with this select. The important
	// point is that default makes it non-blocking. If ch isn't
	// ready at that exact moment, we immediately execute default.
	// So we should use it only when we don't want to wait."


	// ============================================================
	// Question 9.5
	// When would this be useful?
	// ============================================================

	// A common use case is checking whether work or a signal
	// is available without blocking.
	//
	// Example:
	//
	// select {
	// case job := <-jobCh:
	//     process(job)
	//
	// default:
	//     fmt.Println("No job available right now")
	// }
	//
	// INTERVIEW ANSWER:
	//
	// "default is useful when we want to perform a
	// non-blocking channel operation. For example, we might
	// check whether a job is available without making the
	// goroutine wait."


	// ============================================================
	// Question 9.6
	// What happens if there is no default and no channel
	// operation is ready?
	// ============================================================

	// Example:
	//
	// ch := make(chan int)
	//
	// select {
	// case val := <-ch:
	//     fmt.Println(val)
	// }

	// INTERVIEW ANSWER:
	//
	// "The select blocks until one of its channel operations
	// becomes ready."
	//
	// Since ch is unbuffered and nobody sends to it,
	// this select blocks forever.


	// ============================================================
	// QUICK REVISION
	// ============================================================
	//
	// select:
	//
	//     "Wait for one of these channel operations."
	//
	//
	// One case ready:
	//
	//     → that case executes.
	//
	//
	// Multiple cases ready:
	//
	//     → one is chosen pseudo-randomly.
	//
	//
	// No case ready + no default:
	//
	//     → blocks.
	//
	//
	// No case ready + default:
	//
	//     → default executes immediately.
	//
	//
	// Therefore:
	//
	//     select + no default
	//         = can block
	//
	//     select + default
	//         = non-blocking
}

// select
//   │
//   ├── channel A ready? ──→ execute A
//   │
//   ├── channel B ready? ──→ execute B
//   │
//   └── nothing ready?
//           │
//           ├── default exists → execute default immediately
//           │
//           └── no default     → BLOCK

// select does not necessarily execute the first ready case.
// If multiple cases are ready, Go chooses one pseudo-randomly.
