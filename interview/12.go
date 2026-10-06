// You have an API receiving 10,000 concurrent requests,
// but only want 20 expensive operations running at once.
//
// INTERVIEW QUESTION:
//
// "How would you limit the expensive operation to 20 concurrent
// executions?"
//
// INTERVIEW ANSWER:
//
// "I would use a buffered channel as a semaphore with a capacity
// of 20. Each request acquires a slot before starting the expensive
// operation and releases the slot when the operation finishes.
//
// This allows 10,000 requests to exist concurrently, while only
// 20 can execute the expensive operation at any given time."

package interview

import (
	"fmt"
	"sync"
	"time"
)
func expensiveOperation(requestID int) {
	fmt.Println("processing request:", requestID)

	// Simulate expensive work.
	time.Sleep(time.Second)

	fmt.Println("finished request:", requestID)
}

func question12() {

	// Channel capacity = 20.
	//
	// We use the channel as a semaphore.
	//
	// At most 20 goroutines can enter
	// expensiveOperation() at the same time.
	sem := make(chan struct{}, 20)

	var wg sync.WaitGroup

	// Simulate 10,000 concurrent API requests.
	for requestID := 1; requestID <= 10_000; requestID++ {

		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// ACQUIRE ONE PERMIT
			//
			// We put one value into the channel.
			//
			// If the channel has space:
			//     -> send succeeds
			//     -> this request can continue
			//
			// If the channel already contains 20 values:
			//     -> channel is full
			//     -> this goroutine blocks here
			sem <- struct{}{}

			// RELEASE THE PERMIT
			//
			// When this request finishes, remove its
			// value from the channel.
			//
			// This creates space for another waiting request.
			defer func() {
				<-sem
			}()

			// Only 20 requests can reach this point
			// at the same time.
			expensiveOperation(id)

		}(requestID)
	}

	// Wait until all 10,000 requests are completed.
	wg.Wait()

	fmt.Println("All requests completed")
}

//                  10,000 requests
//                        │
//           ┌────────────┴────────────┐
//           ↓                         ↓
//    20 spaces available       Requests waiting
//    ┌─────────────────┐        ┌──────────────┐
//    │ □ □ □ ... □ □ □ │        │ 21, 22, ...  │
//    │     20 slots    │        │    waiting   │
//    └─────────────────┘        └──────────────┘
//           │
//           ↓
//    expensiveOperation()

// Request 1  → takes slot → expensive operation
// Request 2  → takes slot → expensive operation
// ...
// Request 20 → takes slot → expensive operation

// Request 21 → BLOCKS
// Request 22 → BLOCKS

// Worker pool:

// 10,000 jobs
//      ↓
//  jobs channel
//      ↓
// 20 workers

// You have 20 fixed workers processing queued jobs.

// Semaphore:

// 10,000 concurrent requests
//      ↓
// 20 permits
//      ↓
// expensive operation

// The requests already exist; you're simply limiting how many may enter the expensive section simultaneously.


// I can use a buffered channel as a semaphore. 
// With a capacity of 20, each request acquires a 
// permit before the expensive operation and 
// releases it afterward, so even if 10,000 
// requests arrive concurrently, only 20 can execute the expensive operation at once.




// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

/*
QUESTION:

Suppose 10,000 API requests arrive concurrently,
but an expensive operation should run only 20 times
concurrently.

How do we solve this?

ANSWER:

Use a buffered channel as a semaphore.
*/

// Capacity = 20.
//
// This means at most 20 goroutines can be inside
// the expensive operation at the same time.
// var sem = make(chan struct{}, 20)

// func expensiveOperation() {
// 	time.Sleep(time.Second)
// }

// func handler(w http.ResponseWriter, r *http.Request) {

// 	// ACQUIRE a permit.
// 	//
// 	// If fewer than 20 requests are currently
// 	// doing the expensive operation, this succeeds.
// 	//
// 	// If 20 requests are already inside,
// 	// this blocks until one finishes.
// 	sem <- struct{}{}

// 	// RELEASE the permit when this request finishes.
// 	defer func() {
// 		<-sem
// 	}()

// 	expensiveOperation()

// 	fmt.Fprintln(w, "done")
// }

// func main() {
// 	http.HandleFunc("/process", handler)

// 	http.ListenAndServe(":8080", nil)
// }

/*
========================================================
HOW DOES THIS WORK WITH 10,000 API REQUESTS?
========================================================

10,000 requests
      |
      v
   handlers
      |
      v
+----------------+
| Semaphore = 20 |
+----------------+
      |
      +----> Request 1  \
      +----> Request 2   |
      +----> Request 3   |
      ...                | ---> expensiveOperation()
      +----> Request 20  /
      
Requests 21 - 10,000 wait at:

    sem <- struct{}{}

When one of the first 20 finishes:

    <-sem

a waiting request can enter.


========================================================
DOES THIS MAKE THE SERVER STATEFUL?
========================================================

NO.

A stateless API means the server does not depend on
stored user/request/session state between requests.

The semaphore is NOT storing:

    - user data
    - session data
    - request data
    - business data

It is only temporary process-level concurrency control.

It answers:

    "How many expensive operations are running RIGHT NOW?"


========================================================
IMPORTANT: MULTIPLE SERVER INSTANCES
========================================================

Suppose we have:

    Load Balancer
       |
       +---- Server A -> semaphore = 20
       |
       +---- Server B -> semaphore = 20
       |
       +---- Server C -> semaphore = 20

Each server has its OWN semaphore.

Therefore:

    Server A -> max 20
    Server B -> max 20
    Server C -> max 20

Potential total = 60


So:

    In-memory semaphore
            |
            +--> per-instance limit


If we need:

    "ONLY 20 expensive operations across
     the ENTIRE system"

then an in-memory Go semaphore is NOT enough.

We need shared coordination, such as:

    - distributed semaphore
    - shared datastore mechanism
    - message queue + controlled consumers


========================================================
INTERVIEW ANSWER
========================================================

"For a single server instance, I can use a buffered
channel as a semaphore with capacity 20. Each API
request acquires a permit before the expensive operation
and releases it afterward.

This doesn't make the API stateful because the semaphore
is only temporary process-level concurrency control; it
doesn't store user or session state.

If the service is horizontally scaled and I need a global
limit of 20 across all instances, I would use shared
coordination or a queue instead of an in-memory semaphore."


========================================================
KEY DIFFERENCE
========================================================

Semaphore:

    "How many can run at the same time?"

WaitGroup:

    "Have they all finished?"

Statelessness:

    "Does this server need to remember
     request/user state between requests?"

These are three different concepts.
*/