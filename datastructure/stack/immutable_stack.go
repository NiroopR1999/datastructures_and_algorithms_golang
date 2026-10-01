/*
# Immutable Stack

An Immutable Stack is a stack whose existing state cannot be modified.

The stack follows the LIFO principle:
	LIFO = Last In, First Out

The important property of an immutable stack is that operations do not
change the existing stack.

Instead:

	Push() -> returns a new stack containing the new element.
	Pop()  -> returns a new stack with the top element removed.

This means an older version of the stack remains unchanged.

Example:

	stack1 := [10, 20, 30]

	stack2 := stack1.Push(40)

	stack1 is still:
	[10, 20, 30]

	stack2 is:
	[10, 20, 30, 40]

This is useful when multiple parts of a program need to safely keep
different versions of the same stack.

## Operations

	Push(value)
		Adds a value and returns a new stack.
		The original stack is unchanged.

	Pop()
		Removes the top value and returns a new stack.
		The original stack is unchanged.

	Peek()
		Returns the top value without modifying the stack.

	IsEmpty()
		Checks whether the stack contains no elements.

	Size()
		Returns the number of elements.

	Contains(value)
		Checks whether a value exists in the stack.

	Clear()
		Returns a new empty stack.
		The original stack remains unchanged.

	ToSlice()
		Returns a copy of the stack data.

## Why use an immutable stack?

Immutability provides predictable state.

If we create multiple versions:

	stack1 := [10, 20]
	stack2 := stack1.Push(30)
	stack3 := stack2.Push(40)

Then all three versions remain available:

	stack1 = [10, 20]
	stack2 = [10, 20, 30]
	stack3 = [10, 20, 30, 40]

Changing stack3 does not change stack1 or stack2.

## Complexity

	Push:     O(n)
	Pop:      O(n)
	Peek:     O(1)
	IsEmpty:  O(1)
	Size:     O(1)
	Contains: O(n)
	Clear:    O(1)
	ToSlice:  O(n)

Space:

	O(n) for each stack state.

Note:
This implementation prioritizes clarity and guarantees that the
underlying stack data is not mutated.
*/

package stack

import "errors"

var ErrImmutableStackEmpty = errors.New("immutable stack is empty")

// ImmutableStack represents a stack whose existing state cannot be modified.
type ImmutableStack struct {
	data []int
}

// NewImmutableStack creates an empty immutable stack.
func NewImmutableStack() *ImmutableStack {
	return &ImmutableStack{
		data: []int{},
	}
}

// Push returns a new stack containing value.
//
// Why create a new slice?
// The original stack must remain unchanged.
func (s *ImmutableStack) Push(value int) *ImmutableStack {
	newData := make([]int, len(s.data)+1)

	// Copy the existing elements into the new stack.
	copy(newData, s.data)

	// Add the new value at the top.
	newData[len(s.data)] = value

	return &ImmutableStack{
		data: newData,
	}
}

// Pop returns the top value and a new stack with that value removed.
//
// Why return a new stack?
// Removing an element must not modify the existing stack.
func (s *ImmutableStack) Pop() (int, *ImmutableStack, error) {
	if len(s.data) == 0 {
		return 0, s, ErrImmutableStackEmpty
	}

	topIndex := len(s.data) - 1
	value := s.data[topIndex]

	newData := make([]int, topIndex)

	// Copy everything except the old top element.
	copy(newData, s.data[:topIndex])

	return value, &ImmutableStack{
		data: newData,
	}, nil
}

// Peek returns the top element without modifying the stack.
func (s *ImmutableStack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrImmutableStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *ImmutableStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *ImmutableStack) Size() int {
	return len(s.data)
}

// Contains checks whether value exists in the stack.
func (s *ImmutableStack) Contains(value int) bool {
	left := 0
	right := len(s.data) - 1

	// Check from both ends.
	// This can find the value without always scanning from one side.
	for left <= right {
		if s.data[left] == value || s.data[right] == value {
			return true
		}

		left++
		right--
	}

	return false
}

// Clear returns a new empty stack.
//
// Why return a new stack?
// The current stack must remain unchanged.
func (s *ImmutableStack) Clear() *ImmutableStack {
	return NewImmutableStack()
}

// ToSlice returns a copy of the stack data.
//
// Why return a copy?
// Returning s.data directly would allow the caller to modify the
// underlying slice and break immutability.
func (s *ImmutableStack) ToSlice() []int {
	result := make([]int, len(s.data))
	copy(result, s.data)

	return result
}
