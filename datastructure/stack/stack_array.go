/*
# Array Stack

A Stack is a linear data structure that follows the LIFO principle:

	LIFO = Last In, First Out

The element inserted last is the first element removed.

This implementation uses a fixed-size array to store elements.

Example:

	Stack capacity = 5

	[10, 20, 30, _, _]
	          ^
	         Top

	Push(40)

	[10, 20, 30, 40, _]
	              ^
	             Top

The stack has a fixed capacity. Once all positions are occupied,
no more elements can be added.

Stack Overflow:
	Occurs when Push is called on a full stack.

Stack Underflow:
	Occurs when Pop or Peek is called on an empty stack.

The stack uses a `top` variable to track the next available
position in the array.

Example:

	data = [10, 20, 30, _, _]
	top  = 3

The valid elements are:

	data[0] = 10
	data[1] = 20
	data[2] = 30

The next element will be inserted at data[3].

Supported operations:

	Push()     -> Add an element to the top
	Pop()      -> Remove and return the top element
	Peek()     -> Return the top element without removing it
	IsEmpty()  -> Check whether the stack is empty
	IsFull()   -> Check whether the stack is full
	Size()     -> Return the number of elements
	Capacity() -> Return the maximum capacity
	Clear()    -> Remove all elements
	Contains() -> Check whether an element exists
	Display()  -> Print all elements

Time Complexity:

	Push()     -> O(1)
	Pop()      -> O(1)
	Peek()     -> O(1)
	IsEmpty()  -> O(1)
	IsFull()   -> O(1)
	Size()     -> O(1)
	Capacity() -> O(1)
	Clear()    -> O(1)
	Contains() -> O(n)
	Display()  -> O(n)

Space Complexity:

	O(n)

where n is the fixed capacity of the stack.
*/

package stack

import (
	"errors"
	"fmt"
)

var (
	// ErrArrayStackEmpty is returned when Pop or Peek
	// is called on an empty stack.
	ErrArrayStackEmpty = errors.New("stack is empty")

	// ErrArrayStackFull is returned when Push is called
	// on a full stack.
	ErrArrayStackFull = errors.New("stack is full")

	// ErrInvalidStackCapacity is returned when a stack
	// is created with zero or negative capacity.
	ErrInvalidStackCapacity = errors.New(
		"stack capacity must be greater than zero",
	)
)

// ArrayStack represents a fixed-size stack.
//
// `data` stores the stack elements.
//
// `top` represents the next available position
// in the array.
//
// Example:
//
//	data = [10, 20, 30, 0, 0]
//	top  = 3
//
// The current stack is:
//
//	[10, 20, 30]
//	          ^
//	         Top
type ArrayStack struct {
	data []int
	top  int
}

// NewArrayStack creates an empty stack with the given capacity.
//
// Example:
//
//	NewArrayStack(5)
//
//	data = [0, 0, 0, 0, 0]
//	top  = 0
func NewArrayStack(capacity int) (*ArrayStack, error) {
	if capacity <= 0 {
		return nil, ErrInvalidStackCapacity
	}

	return &ArrayStack{
		data: make([]int, capacity),
		top:  0,
	}, nil
}

// Push adds a value to the top of the stack.
//
// Before:
//
//	data = [10, 20, 30, _, _]
//	top  = 3
//
// Push(40)
//
// After:
//
//	data = [10, 20, 30, 40, _]
//	top  = 4
//
// Time Complexity: O(1)
func (s *ArrayStack) Push(value int) error {
	if s.IsFull() {
		return ErrArrayStackFull
	}

	s.data[s.top] = value
	s.top++

	return nil
}

// Pop removes and returns the top element.
//
// Before:
//
//	data = [10, 20, 30, _, _]
//	top  = 3
//
// Pop() returns 30.
//
// After:
//
//	data = [10, 20, _, _, _]
//	top  = 2
//
// Time Complexity: O(1)
func (s *ArrayStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrArrayStackEmpty
	}

	s.top--

	value := s.data[s.top]

	// Reset the removed position.
	s.data[s.top] = 0

	return value, nil
}

// Peek returns the top element without removing it.
//
// Example:
//
//	data = [10, 20, 30, _, _]
//	top  = 3
//
//	Peek() -> 30
func (s *ArrayStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrArrayStackEmpty
	}

	return s.data[s.top-1], nil
}

// IsEmpty returns true when the stack contains no elements.
func (s *ArrayStack) IsEmpty() bool {
	return s.top == 0
}

// IsFull returns true when all positions in the array
// are occupied.
func (s *ArrayStack) IsFull() bool {
	return s.top == len(s.data)
}

// Size returns the number of elements currently
// stored in the stack.
func (s *ArrayStack) Size() int {
	return s.top
}

// Capacity returns the maximum number of elements
// the stack can store.
func (s *ArrayStack) Capacity() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
//
// The underlying array is retained and can be reused.
func (s *ArrayStack) Clear() {
	s.top = 0
}

// Contains checks whether a value exists in the stack.
//
// Only indexes from 0 through top-1 contain valid
// stack elements.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *ArrayStack) Contains(value int) bool {
	for i := 0; i < s.top; i++ {
		if s.data[i] == value {
			return true
		}
	}

	return false
}

// Display prints all elements from bottom to top.
//
// Example:
//
//	10 20 30
func (s *ArrayStack) Display() {
	for i := 0; i < s.top; i++ {
		fmt.Print(s.data[i], " ")
	}

	fmt.Println()
}