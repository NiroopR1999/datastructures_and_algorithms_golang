/*
# Two-Ended Stack

A Two-Ended Stack is a stack-like data structure that allows
elements to be inserted and removed from both ends.

The two ends are:

	Front
	Back

Operations can explicitly choose which end to use.

The structure follows LIFO behavior independently at each end.

Example:

	PushFront(10)
	PushFront(20)

	Front
	  ↓
	[20, 10]
	         ↑
	        Back

	PushBack(30)

	[20, 10, 30]

Now:

	PopFront() -> 20
	PopBack()  -> 30

## Why use a Two-Ended Stack?

Some problems need efficient access to both ends of a collection.

For example, we may want to:

	- Add elements from either side.
	- Remove elements from either side.
	- Inspect either end.
	- Maintain data from both directions.

## Internal representation

The implementation uses a slice.

	Front → [10, 20, 30] ← Back

The front is index 0.

The back is the last element.

Appending to the back is O(1) amortized.

Removing from the back is O(1).

For the front, elements are shifted when using a slice,
so PushFront and PopFront are O(n).

## Operations

	PushFront(value)
		Adds a value to the front.

	PushBack(value)
		Adds a value to the back.

	PopFront()
		Removes and returns the front value.

	PopBack()
		Removes and returns the back value.

	PeekFront()
		Returns the front value without removing it.

	PeekBack()
		Returns the back value without removing it.

	IsEmpty()
		Checks whether the structure is empty.

	Size()
		Returns the number of elements.

	Contains(value)
		Checks whether a value exists.

	Clear()
		Removes all elements.

	ToSlice()
		Returns the elements from front to back.

## Complexity

	PushFront: O(n)
	PushBack:  O(1) amortized
	PopFront:  O(n)
	PopBack:   O(1)
	PeekFront: O(1)
	PeekBack:  O(1)
	IsEmpty:   O(1)
	Size:       O(1)
	Contains:   O(n)
	Clear:      O(1)
	ToSlice:    O(n)

Space:

	O(n)

where n is the number of elements.
*/

package stack

import "errors"

var ErrTwoEndedStackEmpty = errors.New("two-ended stack is empty")

// TwoEndedStack represents a stack that supports operations
// from both the front and back.
type TwoEndedStack struct {
	data []int
}

// NewTwoEndedStack creates an empty two-ended stack.
func NewTwoEndedStack() *TwoEndedStack {
	return &TwoEndedStack{
		data: []int{},
	}
}

// PushFront adds value to the front.
//
// Why insert at index 0?
// The front of this structure is represented by the first element
// of the slice.
func (s *TwoEndedStack) PushFront(value int) {
	// Increase the slice size by one.
	s.data = append(s.data, 0)

	// Shift all existing elements one position to the right.
	copy(s.data[1:], s.data[:len(s.data)-1])

	// Place the new value at the front.
	s.data[0] = value
}

// PushBack adds value to the back.
//
// append adds the value after the current last element.
func (s *TwoEndedStack) PushBack(value int) {
	s.data = append(s.data, value)
}

// PopFront removes and returns the front value.
func (s *TwoEndedStack) PopFront() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrTwoEndedStackEmpty
	}

	// The first element represents the front.
	value := s.data[0]

	// Remove the first element.
	s.data = s.data[1:]

	return value, nil
}

// PopBack removes and returns the back value.
func (s *TwoEndedStack) PopBack() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrTwoEndedStackEmpty
	}

	lastIndex := len(s.data) - 1

	// Save the value before removing the last element.
	value := s.data[lastIndex]

	// Clear the removed value.
	s.data[lastIndex] = 0

	// Remove the last element.
	s.data = s.data[:lastIndex]

	return value, nil
}

// PeekFront returns the front value without removing it.
func (s *TwoEndedStack) PeekFront() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrTwoEndedStackEmpty
	}

	return s.data[0], nil
}

// PeekBack returns the back value without removing it.
func (s *TwoEndedStack) PeekBack() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrTwoEndedStackEmpty
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *TwoEndedStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements.
func (s *TwoEndedStack) Size() int {
	return len(s.data)
}

// Contains checks whether value exists in the structure.
func (s *TwoEndedStack) Contains(value int) bool {
	left := 0
	right := len(s.data) - 1

	// Check from both ends while moving toward the center.
	for left <= right {
		if s.data[left] == value || s.data[right] == value {
			return true
		}

		left++
		right--
	}

	return false
}

// Clear removes all elements.
func (s *TwoEndedStack) Clear() {
	s.data = []int{}
}

// ToSlice returns a copy of the elements from front to back.
func (s *TwoEndedStack) ToSlice() []int {
	result := make([]int, len(s.data))
	copy(result, s.data)

	return result
}