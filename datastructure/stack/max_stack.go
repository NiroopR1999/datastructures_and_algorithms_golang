/*
# Max Stack

A Max Stack is a stack data structure that supports normal
stack operations along with retrieving the maximum element
in constant time.

The stack follows the LIFO principle:

	LIFO = Last In, First Out

Supported operations:

	Push()     -> Add an element to the top
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element without removing it
	GetMax()   -> Return the maximum element in O(1)
	IsEmpty()  -> Check whether the stack is empty
	Size()     -> Return the number of elements
	Clear()    -> Remove all elements
	Contains() -> Check whether an element exists
	Display()  -> Print all elements

# How GetMax Works

The Max Stack maintains two slices:

	data -> stores all stack elements
	max  -> stores the maximum value at each level

Example:

	Push(10)

	data = [10]
	max  = [10]

	Push(20)

	data = [10, 20]
	max  = [10, 20]

	Push(15)

	data = [10, 20, 15]
	max  = [10, 20, 20]

	Push(30)

	data = [10, 20, 15, 30]
	max  = [10, 20, 20, 30]

The last element of max is always the current
maximum of the entire stack.

Therefore:

	GetMax() -> max[len(max)-1]

takes O(1) time.

# Why Do We Store the Maximum at Every Level?

Consider:

	data = [10, 20, 15, 30]
	max  = [10, 20, 20, 30]

If we Pop() 30:

	data = [10, 20, 15]
	max  = [10, 20, 20]

The previous maximum, 20, is immediately available.

Without storing the previous maximum, we would have
to scan the stack again after a Pop().

# Complexity

	Push()     -> O(1)
	Pop()      -> O(1)
	Peek()     -> O(1)
	GetMax()   -> O(1)
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

var ErrMaxStackEmpty = errors.New("max stack is empty")

// MaxStack stores elements and keeps track of the
// maximum value currently present in the stack.
//
// data stores the actual stack elements.
//
// max stores the maximum value corresponding to
// each position in data.
type MaxStack struct {
	data []int
	max  []int
}

// NewMaxStack creates an empty Max Stack.
func NewMaxStack() *MaxStack {
	return &MaxStack{
		data: []int{},
		max:  []int{},
	}
}

// Push adds a value to the top of the stack.
//
// The maximum value for the new position is:
//
//	max(current maximum, new value)
//
// Example:
//
//	data = [10, 20]
//	max  = [10, 20]
//
//	Push(15)
//
//	data = [10, 20, 15]
//	max  = [10, 20, 20]
func (s *MaxStack) Push(value int) {
	s.data = append(s.data, value)

	// If this is the first element, it is automatically
	// the maximum.
	if len(s.max) == 0 {
		s.max = append(s.max, value)
		return
	}

	// Get the maximum value currently present in the stack.
	currentMax := s.max[len(s.max)-1]

	if value > currentMax {
		s.max = append(s.max, value)
	} else {
		// Keep the previous maximum for this level.
		s.max = append(s.max, currentMax)
	}
}

// Pop removes and returns the top element.
//
// The corresponding maximum value is also removed from
// the max slice.
//
// Example:
//
//	data = [10, 20, 15]
//	max  = [10, 20, 20]
//
//	Pop() -> 15
//
//	data = [10, 20]
//	max  = [10, 20]
func (s *MaxStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMaxStackEmpty
	}

	last := len(s.data) - 1

	value := s.data[last]

	s.data = s.data[:last]
	s.max = s.max[:last]

	return value, nil
}

// Peek returns the top element without removing it.
func (s *MaxStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMaxStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// GetMax returns the maximum element currently present
// in the stack.
//
// Because max always stores the current maximum at
// every level, the current maximum is always at the
// last position.
//
// Time Complexity: O(1)
func (s *MaxStack) GetMax() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMaxStackEmpty
	}

	return s.max[len(s.max)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *MaxStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently
// stored in the stack.
func (s *MaxStack) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *MaxStack) Clear() {
	s.data = nil
	s.max = nil
}

// Contains checks whether a value exists in the stack.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *MaxStack) Contains(value int) bool {
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
func (s *MaxStack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}
