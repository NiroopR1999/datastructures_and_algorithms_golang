/*
# Min-Max Stack

A Min-Max Stack is a stack data structure that supports
normal stack operations along with retrieving both the
minimum and maximum elements in constant time.

The stack follows the LIFO principle:

	LIFO = Last In, First Out

Supported operations:

	Push()     -> Add an element to the top
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element without removing it
	GetMin()   -> Return the minimum element in O(1)
	GetMax()   -> Return the maximum element in O(1)
	IsEmpty()  -> Check whether the stack is empty
	Size()     -> Return the number of elements
	Clear()    -> Remove all elements
	Contains() -> Check whether an element exists
	Display()  -> Print all elements

# How It Works

The stack maintains three slices:

	data -> stores the actual stack elements
	min  -> stores the minimum value at each level
	max  -> stores the maximum value at each level

Example:

	Push(10)

	data = [10]
	min  = [10]
	max  = [10]

	Push(5)

	data = [10, 5]
	min  = [10, 5]
	max  = [10, 10]

	Push(8)

	data = [10, 5, 8]
	min  = [10, 5, 5]
	max  = [10, 10, 10]

	Push(20)

	data = [10, 5, 8, 20]
	min  = [10, 5, 5, 5]
	max  = [10, 10, 10, 20]

At every level:

	min[i] = minimum value from data[0] through data[i]

	max[i] = maximum value from data[0] through data[i]

Therefore, the last elements always contain the
current minimum and maximum.

	GetMin() -> min[len(min)-1]
	GetMax() -> max[len(max)-1]

Both operations take O(1) time.

# Why Store Minimum and Maximum at Every Level?

Consider:

	data = [10, 5, 8, 20]
	min  = [10, 5, 5, 5]
	max  = [10, 10, 10, 20]

If Pop() removes 20:

	data = [10, 5, 8]
	min  = [10, 5, 5]
	max  = [10, 10, 10]

The previous minimum and maximum are immediately
available.

No traversal of the stack is required.

# Complexity

	Push()     -> O(1)
	Pop()      -> O(1)
	Peek()     -> O(1)
	GetMin()   -> O(1)
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

var ErrMinMaxStackEmpty = errors.New("min-max stack is empty")

// MinMaxStack stores elements and maintains the
// minimum and maximum values at every stack level.
//
// data stores the actual stack elements.
//
// min stores the minimum value up to each position.
//
// max stores the maximum value up to each position.
type MinMaxStack struct {
	data []int
	min  []int
	max  []int
}

// NewMinMaxStack creates an empty Min-Max Stack.
func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		data: []int{},
		min:  []int{},
		max:  []int{},
	}
}

// Push adds a value to the top of the stack.
//
// The minimum and maximum values for the new level
// are calculated using the previous minimum and maximum.
//
// Example:
//
//	data = [10, 5]
//	min  = [10, 5]
//	max  = [10, 10]
//
//	Push(8)
//
//	data = [10, 5, 8]
//	min  = [10, 5, 5]
//	max  = [10, 10, 10]
func (s *MinMaxStack) Push(value int) {
	s.data = append(s.data, value)

	// First element is both the minimum and maximum.
	if len(s.min) == 0 {
		s.min = append(s.min, value)
		s.max = append(s.max, value)
		return
	}

	currentMin := s.min[len(s.min)-1]
	currentMax := s.max[len(s.max)-1]

	// Store the minimum for the new level.
	if value < currentMin {
		s.min = append(s.min, value)
	} else {
		s.min = append(s.min, currentMin)
	}

	// Store the maximum for the new level.
	if value > currentMax {
		s.max = append(s.max, value)
	} else {
		s.max = append(s.max, currentMax)
	}
}

// Pop removes and returns the top element.
//
// The corresponding minimum and maximum values are
// removed from their respective tracking slices.
//
// Example:
//
//	data = [10, 5, 8]
//	min  = [10, 5, 5]
//	max  = [10, 10, 10]
//
//	Pop() -> 8
//
//	data = [10, 5]
//	min  = [10, 5]
//	max  = [10, 10]
func (s *MinMaxStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinMaxStackEmpty
	}

	last := len(s.data) - 1

	value := s.data[last]

	s.data = s.data[:last]
	s.min = s.min[:last]
	s.max = s.max[:last]

	return value, nil
}

// Peek returns the top element without removing it.
func (s *MinMaxStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinMaxStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// GetMin returns the minimum element currently present
// in the stack.
//
// The current minimum is always stored at the last
// position of the min slice.
//
// Time Complexity: O(1)
func (s *MinMaxStack) GetMin() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinMaxStackEmpty
	}

	return s.min[len(s.min)-1], nil
}

// GetMax returns the maximum element currently present
// in the stack.
//
// The current maximum is always stored at the last
// position of the max slice.
//
// Time Complexity: O(1)
func (s *MinMaxStack) GetMax() (int, error) {
	if s.IsEmpty() {
		return 0, ErrMinMaxStackEmpty
	}

	return s.max[len(s.max)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *MinMaxStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently
// stored in the stack.
func (s *MinMaxStack) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *MinMaxStack) Clear() {
	s.data = nil
	s.min = nil
	s.max = nil
}

// Contains checks whether a value exists in the stack.
//
// The search checks both ends of the stack simultaneously.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *MinMaxStack) Contains(value int) bool {
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
func (s *MinMaxStack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}