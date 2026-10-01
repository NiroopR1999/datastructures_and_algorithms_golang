/*
# Persistent Stack

A Persistent Stack keeps all previous versions of the stack available.

When we perform an operation such as Push() or Pop(), the existing
stack is not destroyed.

Instead, a new version is created.

Example:

	Version 0: []

	Version 1: [10]

	Version 2: [10, 20]

	Version 3: [10, 20, 30]

All four versions continue to exist and can be accessed independently.

## Structural Sharing

The stack uses linked nodes.

When a new value is pushed, the new node points to the existing
top node.

For example:

	Version 1:

	10

	Version 2:

	20 -> 10

	Version 3:

	30 -> 20 -> 10

Version 3 reuses the nodes from Version 2.

This is called structural sharing.

We do not need to copy all existing elements when creating a new
version.

## Why structural sharing?

Suppose a stack contains 1,000,000 elements.

Creating another version should not require copying all 1,000,000
elements.

With structural sharing, Push() only creates one new node.

The previous version remains unchanged.

## Versions

Each PersistentStack value represents one version of the stack.

Example:

	stack1 := NewPersistentStack()

	stack2 := stack1.Push(10)
	stack3 := stack2.Push(20)

Now:

	stack1 = []
	stack2 = [10]
	stack3 = [10, 20]

We can continue using stack1 and stack2 even after creating stack3.

## Operations

	Push(value)
		Creates and returns a new stack version containing value.

	Pop()
		Returns the top value and a new stack version without that value.

	Peek()
		Returns the top value without creating a new version.

	IsEmpty()
		Checks whether the stack is empty.

	Size()
		Returns the number of elements.

	Contains(value)
		Searches for a value.

	Clear()
		Returns an empty stack version.

	ToSlice()
		Returns the elements from bottom to top.

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

	Push: O(1) additional space
	Pop:  O(1) additional space

The nodes are shared between versions, so previous versions do not
require duplicated storage for unchanged elements.
*/

package stack

import "errors"

var ErrPersistentStackEmpty = errors.New("persistent stack is empty")

// persistentNode represents one element in a persistent stack.
//
// Nodes are never modified after creation.
// This allows multiple stack versions to safely share nodes.
type persistentNode struct {
	value int
	next  *persistentNode
}

// PersistentStack represents one version of a persistent stack.
type PersistentStack struct {
	// top points to the top node of this version.
	top *persistentNode

	// size stores the number of elements in this version.
	size int
}

// NewPersistentStack creates an empty persistent stack.
func NewPersistentStack() *PersistentStack {
	return &PersistentStack{}
}

// Push creates and returns a new stack version.
//
// Why don't we modify the current stack?
// Previous versions must remain available.
//
// Why does the new node point to the old top?
// This allows the new version to reuse all existing nodes.
func (s *PersistentStack) Push(value int) *PersistentStack {
	newNode := &persistentNode{
		value: value,
		next:  s.top,
	}

	return &PersistentStack{
		top:  newNode,
		size: s.size + 1,
	}
}

// Pop returns the top value and creates a new stack version
// without that value.
//
// Why don't we delete the current top node?
// The current version may still need it.
func (s *PersistentStack) Pop() (int, *PersistentStack, error) {
	if s.top == nil {
		return 0, s, ErrPersistentStackEmpty
	}

	// The current top belongs to the current version.
	value := s.top.value

	// The next node becomes the top of the new version.
	newStack := &PersistentStack{
		top:  s.top.next,
		size: s.size - 1,
	}

	return value, newStack, nil
}

// Peek returns the top value without creating a new version.
func (s *PersistentStack) Peek() (int, error) {
	if s.top == nil {
		return 0, ErrPersistentStackEmpty
	}

	return s.top.value, nil
}

// IsEmpty reports whether this version contains no elements.
func (s *PersistentStack) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of elements in this version.
func (s *PersistentStack) Size() int {
	return s.size
}

// Contains checks whether value exists in this stack version.
func (s *PersistentStack) Contains(value int) bool {
	current := s.top

	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

// Clear returns an empty stack version.
//
// Why return a new version?
// The current version must remain unchanged.
func (s *PersistentStack) Clear() *PersistentStack {
	return NewPersistentStack()
}

// ToSlice returns the stack contents from bottom to top.
func (s *PersistentStack) ToSlice() []int {
	result := make([]int, 0, s.size)

	current := s.top

	// Traverse from top to bottom.
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	// Reverse so the result is bottom to top.
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result
}