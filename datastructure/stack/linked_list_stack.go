/*
# Linked List Stack

A Linked List Stack is a stack implemented using linked nodes.

Each node contains:

	1. The value stored in the stack.
	2. A pointer to the next node.

The stack follows the LIFO principle:

	LIFO = Last In, First Out

The top of the stack is represented by the head node of the linked list.

Example:

	Push(10)
	Push(20)
	Push(30)

Stack:

	Top
	 ↓
	30 → 20 → 10 → nil

The most recently pushed element is always at the top.

## Push

A new node is created and placed at the beginning of the linked list.

Before:

	Top
	 ↓
	20 → 10 → nil

Push(30):

	Top
	 ↓
	30 → 20 → 10 → nil

Adding a node at the beginning is O(1).

## Pop

The top node is removed and the head pointer moves to the next node.

Before:

	Top
	 ↓
	30 → 20 → 10 → nil

Pop():

	Top
	 ↓
	20 → 10 → nil

The removed value is 30.

Removing the first node is O(1).

## Peek

Peek returns the value at the top without removing it.

## Empty Stack

The stack is empty when the top pointer is nil.

## Operations

	Push(value)
		Adds a value to the top.

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
		Returns the stack elements from bottom to top.

## Complexity

	Push:     O(1)
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

var ErrLinkedListStackEmpty = errors.New("linked list stack is empty")

// linkedListStackNode represents one node in the stack.
type linkedListStackNode struct {
	value int
	next  *linkedListStackNode
}

// LinkedListStack represents a stack implemented using linked nodes.
type LinkedListStack struct {
	// top points to the first node, which is the top of the stack.
	top *linkedListStackNode

	// size tracks the number of elements currently in the stack.
	size int
}

// NewLinkedListStack creates an empty linked-list stack.
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

// Push adds value to the top of the stack.
func (s *LinkedListStack) Push(value int) {
	newNode := &linkedListStackNode{
		value: value,

		// The new node points to the current top.
		// This keeps the existing elements connected.
		next: s.top,
	}

	// The new node becomes the new top.
	s.top = newNode

	// One element was added.
	s.size++
}

// Pop removes and returns the top value.
func (s *LinkedListStack) Pop() (int, error) {
	// We cannot pop when there is no top node.
	if s.top == nil {
		return 0, ErrLinkedListStackEmpty
	}

	// Save the value before moving the top pointer.
	value := s.top.value

	// Move top to the next node.
	// This removes the current top node from the stack.
	s.top = s.top.next

	// One element was removed.
	s.size--

	return value, nil
}

// Peek returns the top value without removing it.
func (s *LinkedListStack) Peek() (int, error) {
	if s.top == nil {
		return 0, ErrLinkedListStackEmpty
	}

	return s.top.value, nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *LinkedListStack) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of elements in the stack.
func (s *LinkedListStack) Size() int {
	return s.size
}

// Contains checks whether value exists in the stack.
func (s *LinkedListStack) Contains(value int) bool {
	current := s.top

	// Walk through every node until the value is found
	// or the end of the linked list is reached.
	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

// Clear removes all elements from the stack.
//
// Setting top to nil disconnects the entire linked structure
// from the stack. The garbage collector can reclaim the nodes
// when nothing else references them.
func (s *LinkedListStack) Clear() {
	s.top = nil
	s.size = 0
}

// ToSlice returns the stack elements from bottom to top.
func (s *LinkedListStack) ToSlice() []int {
	result := make([]int, 0, s.size)

	current := s.top

	// Traversal naturally gives top -> bottom.
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	// Reverse the result so it is returned bottom -> top.
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result
}