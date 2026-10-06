// Print 1–10 using a goroutine and a channel.
// Tests: goroutines, channels, buffered vs unbuffered channels, and synchronization.

package interview

import (
	"fmt"
	"sync"
)

func question2() {

	// ============================================================
	// Approach 1: Buffered Channel
	// ============================================================
	//
	// A buffered channel can store values without a receiver
	// being ready immediately.
	//
	// Buffer size = 10, so all 10 values can be sent before
	// the receiver goroutine starts.
	//
	// If the buffer were smaller than 10, the sender would block
	// once the buffer became full.

	ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		ch <- i
	}

	close(ch)

	go func() {
		defer wg.Done()

		for num := range ch {
			fmt.Println(num)
		}
	}()

	wg.Wait()

	// ============================================================
	// Approach 2: Unbuffered Channel
	// ============================================================
	//
	// An unbuffered channel has no storage.
	//
	// Every send:
	//
	//     ch <- value
	//
	// waits until another goroutine receives that value.
	//
	// Therefore, the receiver goroutine MUST start before
	// the sender starts sending values.

	// ch := make(chan int)
	//
	// var wg sync.WaitGroup
	// wg.Add(1)
	//
	// go func() {
	// 	defer wg.Done()
	//
	// 	for num := range ch {
	// 		fmt.Println(num)
	// 	}
	// }()
	//
	// for i := 1; i <= 10; i++ {
	// 	ch <- i
	// }
	//
	// // Closing tells the receiver that no more values
	// // will be sent.
	// close(ch)
	//
	// wg.Wait()
}
