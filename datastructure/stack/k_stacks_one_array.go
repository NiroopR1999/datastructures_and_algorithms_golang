/*
# K Stacks in One Array

K Stacks in One Array is a technique for storing multiple independent
stacks inside a single array.

Instead of allocating one separate array for every stack, we use:

	1. One shared array to store all values.
	2. One top array to track the top of every stack.
	3. One next array to connect elements and maintain free positions.

For example, if we want 3 stacks:

	Stack 0
	Stack 1
	Stack 2

All three stacks use the same underlying storage array.

## Why use this?

The main purpose is efficient memory utilization.

Suppose the total capacity is 10.

If we create three separate fixed-size arrays:

	Stack 0 -> capacity 3
	Stack 1 -> capacity 3
	Stack 2 -> capacity 4

The capacity of each stack is fixed.

With K Stacks in One Array, all stacks share the same pool of
10 positions.

If Stack 0 needs more space while Stack 1 has unused space,
Stack 0 can use that free space.

## Internal representation

We maintain three arrays:

	data[]
		Stores the actual stack values.

	top[]
		Stores the index of the top element of each stack.

	next[]
		Connects elements belonging to a stack and also keeps track
		of free positions.

We also maintain:

	free
		Stores the index of the first available free position.

Example:

	Stack 0 -> [10, 20]
	Stack 1 -> [30]
	Stack 2 -> [40, 50]

The stacks are logically independent even though their elements
are physically stored in the same data array.

## Stack numbering

This implementation uses zero-based stack numbers:

	0, 1, 2, ..., k-1

For example, if k = 3:

	Push(10, 0) -> pushes 10 into stack 0
	Push(20, 1) -> pushes 20 into stack 1
	Push(30, 2) -> pushes 30 into stack 2

## Overflow

Overflow occurs when there are no free positions left in the
shared array.

It does not matter which stack is being pushed into.

If the shared array is full, no stack can accept another element.

## Underflow

Underflow occurs when Pop() or Peek() is called on an empty stack.

One stack can be empty while the other stacks still contain values.

## Operations

	Push(value, stackID)
		Pushes a value into the specified stack.

	Pop(stackID)
		Removes and returns the top value from the specified stack.

	Peek(stackID)
		Returns the top value without removing it.

	IsEmpty(stackID)
		Checks whether a specific stack is empty.

	IsFull()
		Checks whether the shared array has no free positions.

	Size(stackID)
		Returns the number of elements in a specific stack.

	Clear(stackID)
		Removes all elements from a specific stack.

	Contains(stackID, value)
		Checks whether a value exists in a specific stack.

## Complexity

	Push:     O(1)
	Pop:      O(1)
	Peek:     O(1)
	IsEmpty:  O(1)
	IsFull:   O(1)
	Size:     O(1)
	Clear:    O(n)
	Contains: O(n)

Space:

	O(n + k)

where:

	n = total capacity
	k = number of stacks
*/

package stack

import "errors"

var (
	ErrKStacksFull        = errors.New("all shared stack space is full")
	ErrKStackEmpty        = errors.New("stack is empty")
	ErrInvalidStackNumber = errors.New("invalid stack number")
	ErrInvalidKStackSize  = errors.New("capacity and number of stacks must be greater than zero")
)

// KStacks represents K independent stacks sharing one array.
type KStacks struct {
	data []int

	// top[i] stores the index of the top element of stack i.
	top []int

	// next connects elements inside stacks and also connects free positions.
	next []int

	// free stores the first available position in data.
	free int

	// sizes stores the number of elements in each stack.
	sizes []int
}

// NewKStacks creates k stacks using one shared array.
//
// capacity = total number of elements that all stacks can store together.
// k        = number of independent stacks.
func NewKStacks(capacity, k int) (*KStacks, error) {
	if capacity <= 0 || k <= 0 {
		return nil, ErrInvalidKStackSize
	}

	data := make([]int, capacity)
	top := make([]int, k)
	next := make([]int, capacity)
	sizes := make([]int, k)

	// Initially every stack is empty.
	// -1 means that the stack has no top element.
	for i := range top {
		top[i] = -1
	}

	// Initially every position in data is free.
	// Each free position points to the next free position.
	for i := 0; i < capacity-1; i++ {
		next[i] = i + 1
	}

	// The last position marks the end of the free list.
	next[capacity-1] = -1

	return &KStacks{
		data:  data,
		top:   top,
		next:  next,
		free:  0,
		sizes: sizes,
	}, nil
}

// validateStackNumber makes sure the requested stack exists.
func (s *KStacks) validateStackNumber(stackNumber int) error {
	if stackNumber < 0 || stackNumber >= len(s.top) {
		return ErrInvalidStackNumber
	}

	return nil
}

// Push adds value to the specified stack.
//
// Why do we use free?
// All stacks share the same array, so we need to find an unused
// position before inserting the new value.
func (s *KStacks) Push(value int, stackNumber int) error {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return err
	}

	// No free position means the shared array is completely full.
	if s.free == -1 {
		return ErrKStacksFull
	}

	// Take the first available position.
	index := s.free

	// Move free to the next available position.
	s.free = s.next[index]

	// Store the value in the shared array.
	s.data[index] = value

	// Connect the new element to the previous top of this stack.
	s.next[index] = s.top[stackNumber]

	// Make the new element the top of this stack.
	s.top[stackNumber] = index

	// Track the number of elements in this stack.
	s.sizes[stackNumber]++

	return nil
}

// Pop removes and returns the top element from the specified stack.
func (s *KStacks) Pop(stackNumber int) (int, error) {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return 0, err
	}

	// -1 means this particular stack has no elements.
	if s.top[stackNumber] == -1 {
		return 0, ErrKStackEmpty
	}

	// Get the index of the current top element.
	index := s.top[stackNumber]

	// Read the value before releasing the position.
	value := s.data[index]

	// Move the stack top to the next element.
	s.top[stackNumber] = s.next[index]

	// Return this position to the shared free list.
	s.next[index] = s.free
	s.free = index

	// Clear the old value for easier debugging and inspection.
	s.data[index] = 0

	s.sizes[stackNumber]--

	return value, nil
}

// Peek returns the top element without removing it.
func (s *KStacks) Peek(stackNumber int) (int, error) {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return 0, err
	}

	if s.top[stackNumber] == -1 {
		return 0, ErrKStackEmpty
	}

	return s.data[s.top[stackNumber]], nil
}

// IsEmpty reports whether the specified stack is empty.
func (s *KStacks) IsEmpty(stackNumber int) bool {
	if stackNumber < 0 || stackNumber >= len(s.top) {
		return true
	}

	return s.top[stackNumber] == -1
}

// IsFull reports whether the shared array has no free positions.
func (s *KStacks) IsFull() bool {
	return s.free == -1
}

// Size returns the number of elements in the specified stack.
func (s *KStacks) Size(stackNumber int) int {
	if stackNumber < 0 || stackNumber >= len(s.sizes) {
		return 0
	}

	return s.sizes[stackNumber]
}

// Capacity returns the total number of positions available
// across all stacks.
func (s *KStacks) Capacity() int {
	return len(s.data)
}

// NumberOfStacks returns how many independent stacks are stored.
func (s *KStacks) NumberOfStacks() int {
	return len(s.top)
}

// Clear removes every element from the specified stack.
//
// Why rebuild the free-list links?
// Every removed position must become available to all stacks again.
func (s *KStacks) Clear(stackNumber int) error {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return err
	}

	for s.top[stackNumber] != -1 {
		index := s.top[stackNumber]

		// Move the top toward the bottom of the stack.
		s.top[stackNumber] = s.next[index]

		// Return this position to the shared free list.
		s.next[index] = s.free
		s.free = index

		s.data[index] = 0
	}

	s.sizes[stackNumber] = 0

	return nil
}

// Contains checks whether value exists in the specified stack.
func (s *KStacks) Contains(stackNumber int, value int) (bool, error) {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return false, err
	}

	// Start at the top of the requested stack.
	index := s.top[stackNumber]

	// Follow the stack's linked positions.
	for index != -1 {
		if s.data[index] == value {
			return true, nil
		}

		index = s.next[index]
	}

	return false, nil
}

// ToSlice returns the elements of one stack from bottom to top.
func (s *KStacks) ToSlice(stackNumber int) ([]int, error) {
	if err := s.validateStackNumber(stackNumber); err != nil {
		return nil, err
	}

	result := make([]int, 0, s.sizes[stackNumber])

	// Start from the top and follow the stack backwards.
	index := s.top[stackNumber]

	for index != -1 {
		result = append(result, s.data[index])
		index = s.next[index]
	}

	// The traversal above gives top -> bottom.
	// Reverse it so the returned slice is bottom -> top.
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result, nil
}