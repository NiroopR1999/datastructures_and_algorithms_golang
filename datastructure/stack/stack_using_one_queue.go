/*
# Stack Using One Queue

A Stack Using One Queue implements stack behavior using only one queue.

A queue follows:

	FIFO = First In, First Out

A stack requires:

	LIFO = Last In, First Out

Since the queue naturally removes elements from the front, we rearrange
the queue after every Push so that the newest element moves to the front.

Example:

	Push(10)

	Queue:
	[10]

	Push(20)

	Before rotation:
	[10, 20]

	After rotation:
	[20, 10]

	Push(30)

	Before rotation:
	[20, 10, 30]

	After rotation:
	[30, 20, 10]

Now the front of the queue always represents the top of the stack.

## Push

Push performs two steps:

	1. Add the new value to the back of the queue.
	2. Move all older elements behind the new value.

Example:

	Before:

	[10, 20, 30]

	Push(40)

	First:

	[10, 20, 30, 40]

	Rotate the older elements:

	[40, 10, 20, 30]

The new value is now at the front.

## Pop

Because the newest element is always at the front,
Pop simply removes the front queue element.

Example:

	[40, 30, 20, 10]

	Pop() -> 40

Result:

	[30, 20, 10]

## Peek

Peek returns the front element without removing it.

## Why only one queue?

The implementation uses exactly one queue data structure.

The rotation performed during Push is what allows the queue
to behave like a stack.

## Operations

	Push(value)
		Adds a value to the stack.

	Pop()
		Removes and returns the top value.

	Peek()
		Returns the top value without removing it.

	IsEmpty()
		Checks whether the stack is empty.

	Size()
		Returns the number of elements.

	Contains(value)
		Checks whether a value exists.

	Clear()
		Removes all elements.

	ToSlice()
		Returns the stack contents from top to bottom.

## Complexity

	Push:     O(n)
	Pop:      O(1)
	Peek:     O(1)
	IsEmpty:  O(1)
	Size:     O(1)
	Contains: O(n)
	Clear:    O(1)
	ToSlice:  O(n)

Space:

	O(n)

where n is the number of elements in the stack.
*/

package stack

import "errors"

var ErrOneQueueStackEmpty = errors.New("stack using one queue is empty")

// OneQueueStack represents a stack implemented using one queue.
type OneQueueStack struct {
	queue []int
}

// NewOneQueueStack creates an empty stack.
func NewOneQueueStack() *OneQueueStack {
	return &OneQueueStack{
		queue: []int{},
	}
}

// Push adds value to the top of the stack.
//
// Why rotate the queue?
// A queue adds elements at the back, but a stack needs the newest
// element at the front so that Pop() can remove it in O(1).
func (s *OneQueueStack) Push(value int) {
	// Add the new value to the back of the queue.
	s.queue = append(s.queue, value)

	// Remember how many older elements existed before the new value.
	oldSize := len(s.queue) - 1

	// Move every older element behind the newly inserted value.
	for i := 0; i < oldSize; i++ {
		front := s.queue[0]

		// Remove the current front element.
		s.queue = s.queue[1:]

		// Add it to the back.
		s.queue = append(s.queue, front)
	}
}

// Pop removes and returns the top value.
func (s *OneQueueStack) Pop() (int, error) {
	if len(s.queue) == 0 {
		return 0, ErrOneQueueStackEmpty
	}

	// The front of the queue is always the top of the stack.
	value := s.queue[0]

	// Remove the front element.
	s.queue = s.queue[1:]

	return value, nil
}

// Peek returns the top value without removing it.
func (s *OneQueueStack) Peek() (int, error) {
	if len(s.queue) == 0 {
		return 0, ErrOneQueueStackEmpty
	}

	return s.queue[0], nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *OneQueueStack) IsEmpty() bool {
	return len(s.queue) == 0
}

// Size returns the number of elements in the stack.
func (s *OneQueueStack) Size() int {
	return len(s.queue)
}

// Contains checks whether value exists in the stack.
func (s *OneQueueStack) Contains(value int) bool {
	left := 0
	right := len(s.queue) - 1

	// Check from both ends while moving toward the center.
	for left <= right {
		if s.queue[left] == value || s.queue[right] == value {
			return true
		}

		left++
		right--
	}

	return false
}

// Clear removes all elements from the stack.
func (s *OneQueueStack) Clear() {
	s.queue = []int{}
}

// ToSlice returns the stack contents from top to bottom.
//
// The queue is already arranged with the stack top at index 0,
// so no additional reversal is required.
func (s *OneQueueStack) ToSlice() []int {
	result := make([]int, len(s.queue))
	copy(result, s.queue)

	return result
}