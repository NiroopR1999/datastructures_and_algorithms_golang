/*
# Min Stack

A Min Stack is a stack data structure that supports normal
stack operations along with retrieving the minimum element
in constant time.

The stack follows the LIFO principle:

	LIFO = Last In, First Out

Supported operations:

	Push()     -> Add an element to the top
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element without removing it
	GetMin()   -> Return the minimum element in O(1)
	IsEmpty()  -> Check whether the stack is empty
	Size()     -> Return the number of elements
	Clear()    -> Remove all elements
	Contains() -> Check whether an element exists
	Display()  -> Print all elements

# How GetMin Works

The Min Stack maintains two slices:

	data -> stores all stack elements
	min  -> stores the minimum value at each level

Example:

	Push(10)

	data = [10]
	min  = [10]

	Push(5)

	data = [10, 5]
	min  = [10, 5]

	Push(8)

	data = [10, 5, 8]
	min  = [10, 5, 5]

	Push(3)

	data = [10, 5, 8, 3]
	min  = [10, 5, 5, 3]

The last element of `min` is always the current
minimum of the entire stack.

Therefore:

	GetMin() -> min[len(min)-1]

takes O(1) time.

# Why Do We Store the Minimum at Every Level?

Consider:

	data = [10, 5, 8, 3]
	min  = [10, 5, 5, 3]

If we Pop() 3:

	data = [10, 5, 8]
	min  = [10, 5, 5]

The previous minimum, 5, is immediately available.

If we did not store the previous minimum, we would
have to scan the stack again after every Pop().

# Complexity

	Push()     -> O(1)
	Pop()      -> O(1)
	Peek()     -> O(1)
	GetMin()   -> O(1)
	IsEmpty()  -> O(1)
	Size()     -> O(1)
	Clear()    -> O(1)
	Contains() -> O(n)
	Display()  -> O(n)

Space Complexity:

	O(n)

where n is the number of elements in the stack.
*/

package stack

import (
	"errors"
	"fmt"
)

var ErrMinStackEmpty = errors.New("min stack is empty")

// MinStack stores elements and keeps track of the
// minimum value currently present in the stack.
//
// data stores the actual stack elements.
//
// min stores the minimum value corresponding to
// each position in data.
type MinStack struct {
	data []int
	min  []int
}

// NewMinStack creates an empty Min Stack.
func NewMinStack() *MinStack {
	return &MinStack{
		data: []int{},
		min:  []int{},
	}
}

// Push adds a value to the top of the stack.
//
// The minimum value for the new position is:
//
//	min(current minimum, new value)
//
// Example:
//
//	data = [10, 5]
//	min  = [10, 5]
//
//	Push(8)
//
//	data = [10, 5, 8]
//	min  = [10, 5, 5]
func (s *MinStack) Push(value int) {
	s.data = append(s.data, value)

	// If this is the first element, it is automatically
	// the minimum.
	if len(s.min) == 0 {
		s.min = append(s.min, value)
		return
	}

	// Get the minimum value currently present in the stack.
	currentMin := s.min[len(s.min)-1]

	if value < currentMin {
		s.min = append(s.min, value)
	} else {
		// Keep the previous minimum for this level.
		s.min = append(s.min, currentMin)
	}
}

// Pop removes and returns the top element.
//
// The corresponding minimum value is also removed from
// the min slice.
//
// Example:
//
//	data = [10, 5, 8]
//	min  = [10, 5, 5]
//
//	Pop() -> 8
//
//	data = [10, 5]
//	min  = [10, 5]
func (s *MinStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinStackEmpty
	}

	last := len(s.data) - 1

	value := s.data[last]

	s.data = s.data[:last]
	s.min = s.min[:last]

	return value, nil
}

// Peek returns the top element without removing it.
func (s *MinStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// GetMin returns the minimum element currently present
// in the stack.
//
// Because min always stores the current minimum at
// every level, the current minimum is always at the
// last position.
//
// Time Complexity: O(1)
func (s *MinStack) GetMin() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinStackEmpty
	}

	return s.min[len(s.min)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *MinStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently
// stored in the stack.
func (s *MinStack) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *MinStack) Clear() {
	s.data = nil
	s.min = nil
}

// Contains checks whether a value exists in the stack.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *MinStack) Contains(value int) bool {
	left := 0
	right := len(s.data) - 1

	for left <= right {
		if s.data[left] == value || s.data[right] == value {
			return true
		}

		left++
		right--
	}

	return false
}

// Display prints the elements from bottom to top.
func (s *MinStack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}

