// context package usage

package interview

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ------------------------------------------------------------
// 1. context.WithCancel
// ------------------------------------------------------------

// worker demonstrates manual cancellation using context.WithCancel.
//
// The worker keeps running until the parent context is cancelled.
func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return

		case <-ticker.C:
			fmt.Println("Processing...")
		}
	}
}

// question21_1 demonstrates context.WithCancel.
//
// cancel() sends the cancellation signal.
// wg.Wait() waits until the worker has actually stopped.
func question21_1() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	// Let the worker run for 2 seconds.
	time.Sleep(2 * time.Second)

	// Tell the worker to stop.
	cancel()

	// Wait until the worker actually exits.
	wg.Wait()
}

// ------------------------------------------------------------
// 2. context.WithTimeout
// ------------------------------------------------------------

// question21_2 demonstrates automatic cancellation using
// context.WithTimeout.
//
// The context is automatically cancelled after 2 seconds.
func question21_2() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	// Wait until the worker notices the timeout and exits.
	wg.Wait()
}

// ------------------------------------------------------------
// 3. HTTP request context
// ------------------------------------------------------------

// handler demonstrates how to obtain the context associated
// with an incoming HTTP request.
//
// If the client disconnects or the request is cancelled,
// r.Context() is cancelled as well.
func handler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	// result, err := doDatabaseWork(ctx)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Fprintln(w, result)
}

// ------------------------------------------------------------
// 4. Context propagation
// ------------------------------------------------------------

// Context should flow down the call chain:
//
// HTTP request
//      ↓
//   handler
//      ↓
//   service
//      ↓
// repository
//      ↓
//     DB
//
// The same context is passed through every layer.

// Example:
//
// func handler(w http.ResponseWriter, r *http.Request) {
//     ctx := r.Context()
//     err := service(ctx)
// }
//
// func service(ctx context.Context) error {
//     return repository(ctx)
// }
//
// func repository(ctx context.Context) error {
//     return dbOperation(ctx)
// }

// ------------------------------------------------------------
// 5. Database context
// ------------------------------------------------------------

// getUser demonstrates passing a context to database/sql.
//
// ExecContext allows the database operation to observe
// cancellation and deadlines from the context.
func getUser(ctx context.Context, db *sql.DB, id int) error {
	_, err := db.ExecContext(
		ctx,
		"SELECT * FROM users WHERE id = $1",
		id,
	)

	return err
}

// ------------------------------------------------------------
// 6. Context timeout + database
// ------------------------------------------------------------

// getUserWithTimeout creates a child context with a 2-second
// timeout.
//
// If the database operation takes longer than 2 seconds,
// the context is cancelled and the DB operation can stop.
func getUserWithTimeout(
	ctx context.Context,
	db *sql.DB,
	id int,
) error {
	ctx, cancel := context.WithTimeout(
		ctx,
		2*time.Second,
	)
	defer cancel()

	_, err := db.ExecContext(
		ctx,
		"SELECT * FROM users WHERE id = $1",
		id,
	)

	return err
}

// ------------------------------------------------------------
// 7. Context propagation example
// ------------------------------------------------------------

func service(ctx context.Context) error {
	return repository(ctx)
}

func repository(ctx context.Context) error {
	return dbOperation(ctx)
}

func dbOperation(ctx context.Context) error {
	// Example database operation.
	//
	// In real code, pass ctx to the database driver:
	//
	// db.QueryContext(ctx, ...)
	// db.ExecContext(ctx, ...)
	// db.BeginTx(ctx, ...)

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-time.After(500 * time.Millisecond):
		return nil
	}
}

// ------------------------------------------------------------
// 8. BAD: Breaking context propagation
// ------------------------------------------------------------

// Never replace the incoming context with context.Background()
// inside a lower layer.
//
// BAD:
//
// func service(ctx context.Context) error {
//     newCtx := context.Background()
//     return repository(newCtx)
// }
//
// This breaks the cancellation/deadline chain.
//
// If the HTTP request is cancelled:
//
// HTTP request
//      ↓
// handler
//      ↓
// service
//      X  cancellation chain broken
//      ↓
// context.Background()
//      ↓
// repository

// ------------------------------------------------------------
// 9. context.WithDeadline
// ------------------------------------------------------------

// question21_3 demonstrates context.WithDeadline.
//
// WithDeadline uses an absolute point in time instead of
// a duration.
func question21_3() {
	deadline := time.Now().Add(5 * time.Second)

	ctx, cancel := context.WithDeadline(
		context.Background(),
		deadline,
	)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Context stopped:", ctx.Err())

	case <-time.After(10 * time.Second):
		fmt.Println("Work completed")
	}
}

// ------------------------------------------------------------
// 10. ctx.Done()
// ------------------------------------------------------------

// ctx.Done() returns a receive-only channel:
//
// <-chan struct{}
//
// When the context is cancelled, this channel is closed.
//
// Therefore:
//
// select {
// case <-ctx.Done():
//     fmt.Println("Cancelled")
// }
//
// means:
//
// "Wait until the context is cancelled."

func waitForCancellation(ctx context.Context) {
	<-ctx.Done()

	fmt.Println("Context cancelled:", ctx.Err())
}

// ------------------------------------------------------------
// 11. ctx.Err()
// ------------------------------------------------------------

// ctx.Err() tells us WHY the context stopped.
//
// WithCancel:
//
//     context canceled
//
// WithTimeout / WithDeadline:
//
//     context deadline exceeded

func checkContextError(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("Context stopped:", err)
	}
}

// ------------------------------------------------------------
// 12. context.WithValue
// ------------------------------------------------------------

// Context values should be used only for request-scoped metadata,
// such as tracing information or request IDs.
//
// Do NOT use context as a general-purpose map.

// Prefer a custom key type instead of a plain string.

type contextKey string

const userIDKey contextKey = "userID"

func contextValueExample() {
	ctx := context.WithValue(
		context.Background(),
		userIDKey,
		123,
	)

	userID := ctx.Value(userIDKey)

	fmt.Println("User ID:", userID)
}

// Bad:
//
// ctx = context.WithValue(ctx, "name", "Parker")
// ctx = context.WithValue(ctx, "age", 27)
// ctx = context.WithValue(ctx, "country", "India")
//
// Normal business data should be passed explicitly as function
// parameters instead.

// ------------------------------------------------------------
// Interview answer
// ------------------------------------------------------------

// If the interviewer asks:
//
// "What is context.Context in Go?"
//
// Answer:
//
// "Context is used to propagate cancellation, deadlines, timeouts,
// and request-scoped values across a call chain. For example,
// an HTTP request's context can be passed from the handler to the
// service and database layer, so if the client disconnects or the
// request times out, downstream operations can stop as well."

// Important methods:
//
// context.Context
//
// ├── Done()        → tells us cancellation happened
// ├── Err()         → tells us why it stopped
// ├── WithCancel()  → manually cancel
// ├── WithTimeout() → automatically cancel after a duration
// ├── WithDeadline()→ automatically cancel at a specific time
// └── WithValue()   → store request-scoped metadata
//
// Context flows DOWN the call chain:
//
// handler
//    ↓
// service
//    ↓
// repository
//    ↓
// database
//
// ------------------------------------------------------------
// Key interview concept
// ------------------------------------------------------------
//
// Context does NOT magically stop a function.
//
// It carries the cancellation signal.
//
// The operation must cooperate by checking the context:
//
// select {
// case <-ctx.Done():
//     return ctx.Err()
// }
//
// Or pass the context to an API that understands it:
//
// db.QueryContext(ctx, ...)
// db.ExecContext(ctx, ...)
// http.NewRequestWithContext(ctx, ...)
//
// ------------------------------------------------------------