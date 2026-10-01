/*
Package stack implements a Stack data structure using a Go slice.

# Stack - Slice Based Implementation

A Stack is a linear data structure that follows the LIFO principle:

	LIFO = Last In, First Out

The element inserted last is the first element removed.

Example:

	Push(10)
	Push(20)
	Push(30)

	Stack:
		30 <- Top
		20
		10

	Pop() -> 30
	Pop() -> 20
	Pop() -> 10

This implementation uses a Go slice as the underlying storage.

The end of the slice represents the top of the stack:

	data = [10, 20, 30]
	                ^
	                Top

Why use a slice?

Go slices provide dynamic sizing and efficient append operations.
Adding and removing elements from the end of a slice is O(1)
amortized time.

Supported operations:

	Push()     -> Add an element to the top
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element without removing it
	IsEmpty()  -> Check whether the stack is empty
	Size()     -> Return the number of elements
	Clear()    -> Remove all elements
	Contains() -> Check whether an element exists
	ToSlice()  -> Return a copy of the stack as a slice
	Display()  -> Print the stack

Time Complexity:

	Push()     -> O(1) amortized
	Pop()      -> O(1)
	Peek()     -> O(1)
	IsEmpty()  -> O(1)
	Size()     -> O(1)
	Clear()    -> O(1)
	Contains() -> O(n)
	ToSlice()  -> O(n)
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

// ErrStackIsEmpty is returned when an operation requires
// an element from an empty stack.
var ErrStackIsEmpty = errors.New("stack is empty")

// Stack represents a LIFO stack implemented using a Go slice.
//
// The last element of data is always considered the top
// of the stack.
type Stack struct {
	data []int
}

// NewStack creates and returns an empty Stack.
func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

// Push adds a value to the top of the stack.
//
// Since the top of the stack is represented by the end
// of the slice, append() adds the new value to the top.
//
// Example:
//
//	Before: [10, 20]
//	Push(30)
//	After:  [10, 20, 30]
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element of the stack.
//
// Stack follows LIFO:
//
//	[10, 20, 30]
//	         ^
//	       Top
//
// Pop() returns 30 and changes the stack to:
//
//	[10, 20]
//
// Time Complexity: O(1)
func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrStackIsEmpty
	}

	// The last element is the top of the stack.
	last := len(s.data) - 1
	pop := s.data[last]

	// Remove the last element from the slice.
	s.data = s.data[:last]

	return pop, nil
}

// Peek returns the top element without removing it.
//
// Example:
//
//	Stack: [10, 20, 30]
//
//	Peek() -> 30
//
// The stack remains:
//
//	[10, 20, 30]
func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrStackIsEmpty
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

// Display prints all elements in the stack from bottom to top.
//
// Example:
//
//	Stack: [10, 20, 30]
//
//	Output:
//
//	10 20 30
func (s *Stack) Display() {
	for _, value := range s.data {
		fmt.Print(value, " ")
	}

	fmt.Println()
}

// Clear removes all elements from the stack.
//
// Setting data to nil releases the slice reference,
// allowing the underlying array to be garbage collected
// if there are no other references to it.
func (s *Stack) Clear() {
	s.data = nil
}

// Contains checks whether the given value exists in the stack.
//
// This implementation searches from both ends simultaneously.
//
// Example:
//
//	Stack:
//
//	[10, 20, 30, 40, 50]
//
//	Searching for 50:
//
//	Left  -> 10
//	Right -> 50  <- found
//
// Instead of always scanning from the beginning, we check
// both the left and right sides on every iteration.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *Stack) Contains(val int) bool {
	left := 0
	right := len(s.data) - 1

	for left <= right {
		// Check both ends of the remaining search range.
		if s.data[left] == val || s.data[right] == val {
			return true
		}

		left++
		right--
	}

	return false
}

// ToSlice returns a copy of the stack's underlying data.
//
// A copy is returned instead of s.data directly so that
// callers cannot accidentally modify the internal stack.
//
// Example:
//
//	stack:  [10, 20, 30]
//	result: [10, 20, 30]
//
// Modifying result does not modify the stack.
func (s *Stack) ToSlice() []int {
	result := make([]int, len(s.data))
	copy(result, s.data)

	return result
}