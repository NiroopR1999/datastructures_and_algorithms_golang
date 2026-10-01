/*
# Monotonic Decreasing Stack

A Monotonic Decreasing Stack is a stack that maintains its
elements in decreasing order from bottom to top.

Example:

	[12, 10, 7, 3]

	12 > 10 > 7 > 3

The stack follows the LIFO principle for Pop operations.

The important rule is that when a new value is pushed,
all smaller values at the top are removed before the
new value is inserted.

Example:

	Current stack:

	[12, 10, 7, 3]

	Push(8)

	3 < 8 -> remove 3
	7 < 8 -> remove 7
	10 > 8 -> stop

	Result:

	[12, 10, 8]

The stack always maintains:

	data[i] >= data[i+1]

for every valid index i.

# Why Elements Are Removed

The purpose of a monotonic stack is to keep only elements
that are still useful for future comparisons.

When a new value is greater than the element at the top,
the smaller top element is removed because the new value
dominates it for the monotonic ordering.

Example:

	[10, 5]

	Push(8)

	5 < 8

	5 is removed.

	Result:

	[10, 8]

# Important Behavior

Unlike a normal stack, Push() can remove existing elements.

For example:

	Push(10) -> [10]
	Push(2)  -> [10, 2]
	Push(9)  -> [10, 9]
	Push(12) -> [12]
	Push(10) -> [12, 10]

When 12 is pushed:

	9 < 12  -> remove 9
	10 < 12 -> remove 10

The resulting stack is:

	[12]

# Supported Operations

	Push()     -> Add a value while maintaining decreasing order
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

var ErrMonotonicDecreasingStackEmpty = errors.New(
	"monotonic decreasing stack is empty",
)

// MonotonicDecreasingStack stores elements in decreasing
// order from bottom to top.
//
// Example:
//
//	data = [12, 10, 7, 3]
//
//	12 > 10 > 7 > 3
type MonotonicDecreasingStack struct {
	data []int
}

// NewMonotonicDecreasingStack creates an empty
// monotonic decreasing stack.
func NewMonotonicDecreasingStack() *MonotonicDecreasingStack {
	return &MonotonicDecreasingStack{
		data: []int{},
	}
}

// Push adds a value while maintaining decreasing order.
//
// Any value smaller than the new value is removed from
// the top before inserting the new value.
//
// Example:
//
//	data = [12, 10, 7, 3]
//
//	Push(8)
//
//	3 < 8 -> remove 3
//	7 < 8 -> remove 7
//	10 > 8 -> stop
//
//	data = [12, 10, 8]
func (s *MonotonicDecreasingStack) Push(value int) {
	// Remove values that are smaller than the new value.
	for len(s.data) > 0 &&
		s.data[len(s.data)-1] < value {

		s.data = s.data[:len(s.data)-1]
	}

	// Add the new value after the stack is restored
	// to decreasing order.
	s.data = append(s.data, value)
}

// Pop removes and returns the top element.
//
// Time Complexity: O(1)
func (s *MonotonicDecreasingStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMonotonicDecreasingStackEmpty
	}

	last := len(s.data) - 1
	value := s.data[last]

	s.data = s.data[:last]

	return value, nil
}

// Peek returns the top element without removing it.
//
// Time Complexity: O(1)
func (s *MonotonicDecreasingStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMonotonicDecreasingStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *MonotonicDecreasingStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently
// stored in the stack.
func (s *MonotonicDecreasingStack) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *MonotonicDecreasingStack) Clear() {
	s.data = nil
}

// Contains checks whether a value exists in the stack.
//
// Because the stack is ordered from largest to smallest,
// the search can stop early when the current value becomes
// smaller than the target.
//
// Example:
//
//	data = [12, 10, 8, 5, 2]
//
//	Contains(8) -> true
//
//	Contains(7):
//
//	12 > 7
//	10 > 7
//	8 > 7
//	5 < 7 -> stop
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *MonotonicDecreasingStack) Contains(value int) bool {
	for _, current := range s.data {
		if current == value {
			return true
		}

		// Since the stack is decreasing, all following
		// values will be even smaller.
		if current < value {
			return false
		}
	}

	return false
}

// Display prints the elements from bottom to top.
func (s *MonotonicDecreasingStack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}