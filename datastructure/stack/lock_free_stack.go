/*
# Lock-Free Stack

A Lock-Free Stack is a concurrent stack that allows multiple goroutines
to Push and Pop elements without using a mutex.

The stack follows the LIFO principle:

	LIFO = Last In, First Out

The stack is implemented using a linked list.

The important part is the atomic update of the top pointer.

When multiple goroutines try to modify the stack at the same time,
they use Compare-And-Swap (CAS) to safely update the top.

## How Push works

Suppose the stack is:

	top -> 30 -> 20 -> 10

To push 40:

	1. Create a new node containing 40.
	2. Point the new node to the current top.
	3. Atomically change top from 30 to 40.

Result:

	top -> 40 -> 30 -> 20 -> 10

If another goroutine changes top before the CAS succeeds,
the CAS fails.

The goroutine then reads the new top and tries again.

## How Pop works

Suppose:

	top -> 40 -> 30 -> 20

To pop:

	1. Read the current top.
	2. Read the node after the top.
	3. Atomically change top from the old node to the next node.

If another goroutine changes top before the CAS succeeds,
the operation retries.

## Compare-And-Swap

CAS means:

	"Change this value only if it is still the value I originally read."

Conceptually:

	if top == oldTop {
		top = newTop
	}

This entire check-and-update is performed atomically.

## Why CAS is needed

Without an atomic operation, two goroutines could read the same
top node and both try to modify it.

That could cause:

	- lost updates
	- corrupted links
	- incorrect Pop results
	- inconsistent stack state

CAS prevents an update from being applied when another goroutine
has already changed the stack.

## Lock-Free Progress

A lock-free algorithm guarantees system-wide progress.

If one goroutine is repeatedly interrupted or fails its CAS,
other goroutines can still make progress.

The implementation does not wait for a mutex to become available.

## Operations

	Push(value)
		Adds a value to the top of the stack.

	Pop()
		Removes and returns the top value.

	Peek()
		Reads the current top value.

	IsEmpty()
		Checks whether the stack is empty.

	Size()
		Returns the number of elements.

	Contains(value)
		Searches for a value.

	Clear()
		Removes all elements.

## Complexity

	Push:     O(1) expected
	Pop:      O(1) expected
	Peek:     O(1)
	IsEmpty:  O(1)
	Size:     O(1)
	Contains: O(n)
	Clear:    O(n)

Space:

	O(n)

Note:
Push and Pop may retry when there is contention because another
goroutine can modify the stack between the read and CAS operation.
*/

package stack

import (
	"errors"
	"sync/atomic"
)

var ErrLockFreeStackEmpty = errors.New("lock-free stack is empty")

// lockFreeNode represents one element in the stack.
type lockFreeNode struct {
	value int
	next  *lockFreeNode
}

// LockFreeStack is a stack that uses atomic operations instead of a mutex.
type LockFreeStack struct {
	// top stores the current top node.
	//
	// atomic.Pointer provides atomic Load and CompareAndSwap
	// operations on the pointer.
	top atomic.Pointer[lockFreeNode]

	// size tracks the number of elements.
	//
	// It is also accessed atomically because multiple goroutines
	// can modify it at the same time.
	size atomic.Int64
}

// NewLockFreeStack creates an empty lock-free stack.
func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

// Push adds value to the top of the stack.
func (s *LockFreeStack) Push(value int) {
	newNode := &lockFreeNode{
		value: value,
	}

	for {
		// Read the current top atomically.
		oldTop := s.top.Load()

		// The new node must point to the current top.
		newNode.next = oldTop

		// Change top only if it is still oldTop.
		//
		// If another goroutine changed top after our Load(),
		// CAS returns false and we retry with the new top.
		if s.top.CompareAndSwap(oldTop, newNode) {
			s.size.Add(1)
			return
		}
	}
}

// Pop removes and returns the top element.
func (s *LockFreeStack) Pop() (int, error) {
	for {
		// Read the current top atomically.
		oldTop := s.top.Load()

		// The stack is empty when top is nil.
		if oldTop == nil {
			return 0, ErrLockFreeStackEmpty
		}

		// The next node will become the new top.
		newTop := oldTop.next

		// Change top only if it is still oldTop.
		//
		// If another goroutine already changed top,
		// CAS fails and we retry.
		if s.top.CompareAndSwap(oldTop, newTop) {
			s.size.Add(-1)
			return oldTop.value, nil
		}
	}
}

// Peek returns the current top value without removing it.
func (s *LockFreeStack) Peek() (int, error) {
	top := s.top.Load()

	if top == nil {
		return 0, ErrLockFreeStackEmpty
	}

	return top.value, nil
}

// IsEmpty reports whether the stack is empty.
func (s *LockFreeStack) IsEmpty() bool {
	return s.top.Load() == nil
}

// Size returns the current number of elements.
func (s *LockFreeStack) Size() int {
	return int(s.size.Load())
}

// Contains checks whether value exists in the stack.
func (s *LockFreeStack) Contains(value int) bool {
	current := s.top.Load()

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
// Why use Compare-And-Swap?
// Another goroutine may be pushing or popping at the same time.
// CAS ensures we only clear the stack if the top has not changed
// since we read it.
func (s *LockFreeStack) Clear() {
	for {
		oldTop := s.top.Load()

		if oldTop == nil {
			return
		}

		if s.top.CompareAndSwap(oldTop, nil) {
			s.size.Store(0)
			return
		}
	}
}

// ToSlice returns the stack contents from top to bottom.
func (s *LockFreeStack) ToSlice() []int {
	result := make([]int, 0, s.Size())

	current := s.top.Load()

	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	return result
}