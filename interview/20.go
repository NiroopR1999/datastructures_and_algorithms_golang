package interview

import (
	"context"
	"fmt"
	"time"
)

/*
========================================================
QUESTION 20:
YOU HAVE A WORKER PROCESSING A DATABASE REQUEST.
THE HTTP REQUEST IS CANCELLED.
HOW DO YOU STOP THE DATABASE OPERATION TOO?
========================================================

INTERVIEW ANSWER:

"I would pass the HTTP request's context down to the
worker20 and then to the database operation.

When the HTTP request is cancelled, its context is
cancelled as well. The database operation can then stop
by observing ctx.Done(), or the database driver can
handle the cancellation directly."


========================================================
THE FLOW
========================================================

HTTP Request
     |
     | r.Context()
     v
   Worker
     |
     | ctx
     v
Database Operation
     |
     v
HTTP request cancelled
     |
     v
ctx.Done()
     |
     v
Database operation cancelled


========================================================
WHY CONTEXT?
========================================================

Context allows cancellation/deadlines to travel through
different layers of an application.

For example:

	HTTP Handler
	    ↓
	Service
	    ↓
	Repository
	    ↓
	Database


Instead of creating a separate cancellation mechanism
for every layer, we pass the same context through them.


========================================================
HTTP HANDLER
========================================================
*/

func handler20(ctx context.Context) {

	// Pass the request context to the worker20.
	worker20(ctx)
}

/*
In a real HTTP handler20:

	func handler20(w http.ResponseWriter, r *http.Request) {
		worker20(r.Context())
	}

`r.Context()` is automatically cancelled when the
client disconnects or the request is otherwise cancelled.


========================================================
WORKER
========================================================
*/

func worker20(ctx context.Context) {

	fmt.Println("worker20 started")

	err := databaseOperation(ctx)

	if err != nil {
		fmt.Println("database operation stopped:", err)
		return
	}

	fmt.Println("database operation completed")
}

/*
========================================================
DATABASE OPERATION
========================================================
*/

func databaseOperation(ctx context.Context) error {

	fmt.Println("database operation started")

	// Simulate a database operation that takes 5 seconds.
	operation := time.NewTimer(5 * time.Second)

	defer operation.Stop()

	select {

	case <-operation.C:

		// Database operation completed normally.
		return nil

	case <-ctx.Done():

		// HTTP request was cancelled.
		//
		// Stop the operation and return the
		// cancellation error.
		return ctx.Err()
	}
}

/*
========================================================
SIMULATING HTTP REQUEST CANCELLATION
========================================================
*/

func question20() {

	// Create a context that we can cancel manually.
	ctx, cancel := context.WithCancel(context.Background())

	// Start worker20.
	go worker20(ctx)

	// Simulate the HTTP client disconnecting
	// after 1 second.
	time.Sleep(1 * time.Second)

	fmt.Println("HTTP request cancelled")

	cancel()

	// Give the worker20 time to observe cancellation.
	time.Sleep(100 * time.Millisecond)
}

/*
========================================================
WHAT HAPPENS?
========================================================

Initially:

	HTTP request
	    |
	    v
	    ctx
	    |
	    v
	Worker
	    |
	    v
	Database
	    |
	    | waiting...
	    |
	    v
	5 seconds


After 1 second:

	Client disconnects
	      |
	      v
	   ctx cancelled
	      |
	      v
	   ctx.Done()
	      |
	      v
	Database operation
	      |
	      v
	    STOP


The database operation does NOT have to wait
the full 5 seconds.


========================================================
DEADLINE / TIMEOUT
========================================================

Context can also automatically cancel an operation
after a specific amount of time.

Example:

*/

func databaseWithTimeout() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	err := databaseOperation(ctx)

	if err != nil {
		fmt.Println("operation failed:", err)
	}
}

/*
The database operation has a maximum lifetime of
2 seconds.

	ctx
	 |
	 +---- 2 seconds ----+
	 |                   |
	 v                   v
	database          cancelled
	operation


========================================================
REAL DATABASE CODE
========================================================

With a database driver that supports context, we don't
usually manually implement the select shown above.

For example, with database/sql:

*/

func queryDatabase(ctx context.Context) error {

	/*
		rows, err := db.QueryContext(
			ctx,
			"SELECT * FROM users WHERE id = $1",
			123,
		)

		if err != nil {
			return err
		}

		defer rows.Close()
	*/

	// QueryContext receives ctx.
	//
	// If ctx is cancelled, the database driver can
	// cancel/interrupt the database operation.

	return nil
}

/*
========================================================
CONTEXT METHODS TO REMEMBER
========================================================

context.Background()

	Root context.
	Usually used at the beginning of a request/application.

context.WithCancel(ctx)

	Manually cancel the context.

context.WithTimeout(ctx, duration)

	Automatically cancel after a duration.

context.WithDeadline(ctx, time)

	Automatically cancel at a specific time.

ctx.Done()

	Channel that becomes readable when context is cancelled.

ctx.Err()

	Tells you why the context was cancelled.

	Possible values include:

		context.Canceled
		context.DeadlineExceeded

========================================================
IMPORTANT RULE
========================================================

DO NOT create a new unrelated context inside the worker20.

BAD:

	func worker20(parentCtx context.Context) {
		ctx := context.Background()

		databaseOperation(ctx)
	}

Now cancellation of the HTTP request does NOT reach
the database operation.

GOOD:

	func worker20(ctx context.Context) {
		databaseOperation(ctx)
	}

Pass the existing context downward.

========================================================
CONTEXT FLOW
========================================================

	HTTP Request
	      |
	      | r.Context()
	      v
	   Handler
	      |
	      v
	    Service
	      |
	      v
	   Repository
	      |
	      v
	   Database

	One context flows downward.

========================================================
SHORT INTERVIEW ANSWER
========================================================

"I would pass `r.Context()` from the HTTP handler20 through
the service and repository layers into the database call.
When the HTTP request is cancelled, that context is
cancelled, and a context-aware database driver can cancel
the query as well. I can also use `context.WithTimeout`
when the database operation needs its own deadline."

========================================================
EASY WAY TO REMEMBER
========================================================

Context =

	"Tell all the work below me when
	 I am no longer interested."

HTTP cancelled

	↓

Context cancelled

	↓

Worker stops

	↓

Database operation stops
*/
func main() {
	question20()
}
