// What happens when you receive from an empty buffered channel?
//
// Test:
// - Receiving from a buffered channel
// - Blocking behavior

package interview

import "fmt"

func question5() {
	ch := make(chan int, 5)

	// The channel has a capacity of 5,
	// but currently contains no values.
	//
	// Since there is nothing to receive,
	// this operation blocks.
	val := <-ch

	// This line is never reached because the receive
	// above remains blocked.
	fmt.Println(val)
}

// Receiving from an empty buffered channel blocks until a value is available. 
// The buffer being non-zero only determines how many values can be stored; 
// it doesn't mean a receive can happen when the channel is empty."