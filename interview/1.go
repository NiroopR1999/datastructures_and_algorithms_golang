// Write a Go program using two goroutines and channels to print odd and even numbers alternately from 1 to 10.

// Requirements:

// One goroutine should print odd numbers: 1, 3, 5, 7, 9.
// Another goroutine should print even numbers: 2, 4, 6, 8, 10.
// The output must always be in order: 1 2 3 4 ... 10.
// Use two unbuffered channels to coordinate whose turn it is.
// Use a sync.WaitGroup to ensure the main function waits for both goroutines to finish.
// Do not use time.Sleep() for synchronization.

package interview

import (
	"fmt"
	"sync"
)

func question1() {

	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	// Odd goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= 9; i += 2 {

			// Wait until even goroutine tells us
			// that it is our turn.
			<-oddChan

			fmt.Println(i)

			// Tell even goroutine that it is its turn.
			evenChan <- struct{}{}
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {

			// Wait until odd goroutine tells us
			// that it is our turn.
			<-evenChan

			fmt.Println(i)

			// Don't send after printing 10,
			// because the odd goroutine has already finished.
			if i < 10 {
				oddChan <- struct{}{}
			}
		}
	}()

	// Give the odd goroutine the first turn.
	oddChan <- struct{}{}

	// Wait for both goroutines to finish.
	wg.Wait()
}