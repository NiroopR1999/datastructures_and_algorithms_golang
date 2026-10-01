/*
# Monotonic Increasing Stack

A Monotonic Increasing Stack is a stack that maintains its
elements in increasing order from bottom to top.

Example:

	[2, 5, 7, 10]

	2 < 5 < 7 < 10

The stack follows the LIFO principle for Pop operations.

The important rule is that when a new value is pushed,
all larger values at the top are removed before the
new value is inserted.

Example:

	Current stack:

	[2, 5, 7, 10]

	Push(6)

	10 > 6 -> remove 10
	7 > 6  -> remove 7
	5 < 6  -> stop

	Result:

	[2, 5, 6]

The stack always maintains:

	data[i] <= data[i+1]

for every valid index i.

# Why Elements Are Removed

The purpose of a monotonic stack is to keep only elements
that are still useful for future comparisons.

When a new value is smaller than the element at the top,
the larger top element is removed because the new value
dominates it for the monotonic ordering.

Example:

	[2, 8]

	Push(5)

	8 > 5

	8 is removed.

	Result:

	[2, 5]

# Important Behavior

Unlike a normal stack, Push() can remove existing elements.

For example:

	Push(2)  -> [2]
	Push(10) -> [2, 10]
	Push(5)  -> [2, 5]
	Push(1)  -> [1]
	Push(3)  -> [1, 3]

When 1 is pushed:

	5 > 1  -> remove 5
	2 > 1  -> remove 2

The resulting stack is:

	[1]

# Supported Operations

	Push()     -> Add a value while maintaining increasing order
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element
	IsEmpty()  -> Check whether the stack is empty
	Size()     -> Return the number of elements
	Clear()    -> Remove all elements
	Contains() -> Check whether a value exists
	Display()  -> Print the stack

# Complexity

	Push()     -> O(n) worst case
	Pop()      -> O(1)
	Peek()     -> O(1)
	IsEmpty()  -> O(1)
	Size()     -> O(1)
	Clear()    -> O(1)
	Contains() -> O(n)
	Display()  -> O(n)

Although a single Push() can take O(n), across a sequence
of pushes each element can be removed only once.

Therefore, for a sequence of n Push operations, the
amortized cost is O(1) per operation.

Space Complexity:

	O(n)

where n is the number of elements currently stored.
*/

package stack

import (
	"errors"
	"fmt"
)

var ErrMonotonicIncreasingStackEmpty = errors.New(
	"monotonic increasing stack is empty",
)

// MonotonicIncreasingStack stores elements in increasing
// order from bottom to top.
//
// Example:
//
//	data = [2, 5, 7, 10]
//
//	2 < 5 < 7 < 10
type MonotonicIncreasingStack struct {
	data []int
}

// NewMonotonicIncreasingStack creates an empty
// monotonic increasing stack.
func NewMonotonicIncreasingStack() *MonotonicIncreasingStack {
	return &MonotonicIncreasingStack{
		data: []int{},
	}
}

// Push adds a value while maintaining increasing order.
//
// Any value larger than the new value is removed from
// the top before inserting the new value.
//
// Example:
//
//	data = [2, 5, 7, 10]
//
//	Push(6)
//
//	10 > 6 -> remove 10
//	7 > 6  -> remove 7
//	5 < 6  -> stop
//
//	data = [2, 5, 6]
func (s *MonotonicIncreasingStack) Push(value int) {
	// Remove values that are larger than the new value.
	for len(s.data) > 0 &&
		s.data[len(s.data)-1] > value {

		s.data = s.data[:len(s.data)-1]
	}

	// Add the new value after the stack is restored
	// to increasing order.
	s.data = append(s.data, value)
}

// Pop removes and returns the top element.
//
// Time Complexity: O(1)
func (s *MonotonicIncreasingStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMonotonicIncreasingStackEmpty
	}

	last := len(s.data) - 1
	value := s.data[last]

	s.data = s.data[:last]

	return value, nil
}

// Peek returns the top element without removing it.
//
// Time Complexity: O(1)
func (s *MonotonicIncreasingStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMonotonicIncreasingStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *MonotonicIncreasingStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently
// stored in the stack.
func (s *MonotonicIncreasingStack) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *MonotonicIncreasingStack) Clear() {
	s.data = nil
}

// Contains checks whether a value exists in the stack.
//
// Because the stack is ordered from smallest to largest,
// the search can stop early when the current value becomes
// larger than the target.
//
// Example:
//
//	data = [2, 5, 7, 10]
//
//	Contains(7) -> true
//
//	Contains(6):
//
//	2 < 6
//	5 < 6
//	7 > 6 -> stop
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *MonotonicIncreasingStack) Contains(value int) bool {
	for _, current := range s.data {
		if current == value {
			return true
		}

		// Since the stack is increasing, all following
		// values will be even larger.
		if current > value {
			return false
		}
	}

	return false
}

// Display prints the elements from bottom to top.
func (s *MonotonicIncreasingStack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}