//with channel
package main

import (
	"fmt"
	"sync"
)

func main() {
	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	// Odd goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= 9; i += 2 {
			<-oddCh

			fmt.Println(i)

			evenCh <- struct{}{}
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {
			<-evenCh

			fmt.Println(i)
oddCh <- struct{}{}
			// if i < 10 {
			// 	oddCh <- struct{}{}
			// }
		}
	}()

	// Give the first turn to odd goroutine
	oddCh <- struct{}{}

	wg.Wait()
}