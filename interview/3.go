// Write a producer-consumer program using a channel.
//
// Tests:
// - Sending values through a channel
// - Receiving values from a channel
// - Closing a channel
// - Goroutine synchronization using WaitGroup

package interview

import (
	"fmt"
	"sync"
)

// producer generates values and sends them to the channel.
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}

	// The producer closes the channel because it is the sender
	// and knows that no more values will be sent.
	close(ch)
}

// consumer receives and processes values from the channel.
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// range keeps receiving values until the channel is closed.
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
}

func question3() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Producer and consumer run concurrently.
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// Wait for both goroutines to finish.
	wg.Wait()
}

