// Thread-Safe Stack
//
// A Thread-Safe Stack protects its internal data so that multiple goroutines
// can safely access the stack at the same time.
//
// WHY DO WE NEED A MUTEX?
//
// A stack contains shared mutable state. For example, Push modifies the slice
// while Pop reads and modifies the same slice.
//
// If multiple goroutines perform these operations at the same time, their
// operations can overlap and cause a data race or inconsistent state.
//
// sync.Mutex provides mutual exclusion:
//
//   Lock()   -> only one goroutine can enter the critical section.
//   Unlock() -> allows another waiting goroutine to enter.
//
// The mutex protects both the stack data and its size.
//
// TIME COMPLEXITY:
//
// Push()    -> O(1) amortized
// Pop()     -> O(1)
// Peek()    -> O(1)
// IsEmpty() -> O(1)
// Size()    -> O(1)
// Contains() -> O(n)
// Clear()   -> O(1)
//
// The mutex itself does not change the algorithmic complexity.
// It only provides safe access when multiple goroutines use the stack.

package stack

import (
	"errors"
	"fmt"
	"sync"
)

type ThreadSafeStack struct {
	mu   sync.Mutex
	data []int
	size int
}

var ErrThreadSafeStackEmpty = errors.New("thread-safe stack is empty")

// NewThreadSafeStack creates an empty thread-safe stack.
func NewThreadSafeStack() *ThreadSafeStack {
	return &ThreadSafeStack{
		data: []int{},
		size: 0,
	}
}

// Push adds a value to the top of the stack.
func (s *ThreadSafeStack) Push(value int) {
	// Lock before modifying shared state.
	//
	// This prevents another goroutine from modifying the stack
	// at the same time.
	s.mu.Lock()

	// Always unlock before returning from the method.
	defer s.mu.Unlock()

	// Add the new value at the end of the slice.
	// The last element represents the top of the stack.
	s.data = append(s.data, value)

	// Keep the size consistent with the data slice.
	s.size++
}

// Pop removes and returns the top value.
func (s *ThreadSafeStack) Pop() (int, error) {
	// Lock because Pop reads and modifies shared state.
	s.mu.Lock()
	defer s.mu.Unlock()

	// We cannot remove an element from an empty stack.
	if s.size == 0 {
		return 0, ErrThreadSafeStackEmpty
	}

	// The last element is the top of the stack.
	value := s.data[s.size-1]

	// Clear the removed value before shrinking the slice.
	//
	// This is not strictly necessary for int values, but it makes
	// the removal explicit and is useful when the element type later
	// becomes a pointer or contains references.
	s.data[s.size-1] = 0

	// Remove the last element.
	s.data = s.data[:s.size-1]

	// Update the number of elements.
	s.size--

	return value, nil
}

// Peek returns the top value without removing it.
func (s *ThreadSafeStack) Peek() (int, error) {
	// Lock because another goroutine could modify the stack
	// while we are reading its top element.
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size == 0 {
		return 0, ErrThreadSafeStackEmpty
	}

	return s.data[s.size-1], nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *ThreadSafeStack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.size == 0
}

// Size returns the number of elements in the stack.
func (s *ThreadSafeStack) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.size
}

// Contains checks whether the stack contains the given value.
func (s *ThreadSafeStack) Contains(value int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// The lock must remain held for the entire traversal.
	//
	// Otherwise another goroutine could modify the slice while
	// this method is reading it.
	for _, item := range s.data {
		if item == value {
			return true
		}
	}

	return false
}

// Clear removes all elements from the stack.
func (s *ThreadSafeStack) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Remove all elements and release the slice contents.
	s.data = []int{}
	s.size = 0
}

// Display prints the stack from bottom to top.
func (s *ThreadSafeStack) Display() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}

// ToSlice returns a copy of the stack.
//
// A copy is important here.
//
// Returning s.data directly would allow the caller to modify the
// underlying slice without acquiring the mutex, defeating the purpose
// of making the stack thread-safe.
func (s *ThreadSafeStack) ToSlice() []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]int, len(s.data))
	copy(result, s.data)

	return result
}