/*
# Two Stacks in One Array

Two Stacks in One Array is a technique for storing two independent
stacks inside a single array.

The array is shared by both stacks.

Stack 1 grows from the left side toward the right.

Stack 2 grows from the right side toward the left.

Example with capacity 10:

	Stack 1 → → →       ← ← ← Stack 2

	[ ][ ][ ][ ][ ][ ][ ][ ][ ][ ]
	 ↑                           ↑
	top1                        top2

Stack 1 uses positions from the beginning.

Stack 2 uses positions from the end.

Both stacks continue growing toward the center.

## Why use one array?

The main advantage is efficient use of the available space.

Suppose the array has capacity 10.

If Stack 1 contains only 2 elements and Stack 2 needs 5 elements,
Stack 2 can use the remaining shared space.

Neither stack has a permanently reserved portion of the array.

## Push into Stack 1

Stack 1 grows from left to right.

Example:

	Push1(10)
	Push1(20)
	Push1(30)

	Array:

	[10][20][30][ ][ ][ ][ ][ ][ ][ ]

	top1 points to 30.

## Push into Stack 2

Stack 2 grows from right to left.

Example:

	Push2(100)
	Push2(200)

	Array:

	[10][20][30][ ][ ][ ][ ][200][100]

	top2 points to 200.

## Overflow

The array is full when the two stacks meet.

Condition:

	top1 + 1 == top2

At that point, there is no unused position between the stacks.

## Pop

Stack 1 removes elements from the right side of its region.

Stack 2 removes elements from the left side of its region.

Both Pop operations are O(1).

## Operations

	Push1(value)
		Pushes a value into Stack 1.

	Push2(value)
		Pushes a value into Stack 2.

	Pop1()
		Removes and returns the top of Stack 1.

	Pop2()
		Removes and returns the top of Stack 2.

	Peek1()
		Returns the top of Stack 1 without removing it.

	Peek2()
		Returns the top of Stack 2 without removing it.

	IsEmpty1()
		Checks whether Stack 1 is empty.

	IsEmpty2()
		Checks whether Stack 2 is empty.

	IsFull()
		Checks whether the shared array is full.

	Size1()
		Returns the size of Stack 1.

	Size2()
		Returns the size of Stack 2.

	Clear1()
		Removes all elements from Stack 1.

	Clear2()
		Removes all elements from Stack 2.

## Complexity

	Push1:    O(1)
	Push2:    O(1)
	Pop1:     O(1)
	Pop2:     O(1)
	Peek1:    O(1)
	Peek2:    O(1)
	IsEmpty1: O(1)
	IsEmpty2: O(1)
	IsFull:   O(1)
	Size1:    O(1)
	Size2:    O(1)
	Clear1:   O(n)
	Clear2:   O(n)

Space:

	O(n)

where n is the total capacity of the shared array.
*/

package stack

import "errors"

var (
	ErrTwoStacksArrayFull  = errors.New("two stacks shared array is full")
	ErrStack1Empty         = errors.New("stack 1 is empty")
	ErrStack2Empty         = errors.New("stack 2 is empty")
	ErrInvalidArrayCapacity = errors.New("array capacity must be greater than zero")
)

// TwoStacksOneArray stores two independent stacks in one array.
type TwoStacksOneArray struct {
	data []int

	// top1 is the index of the top element of Stack 1.
	top1 int

	// top2 is the index of the top element of Stack 2.
	top2 int
}

// NewTwoStacksOneArray creates two stacks sharing one array.
//
// Stack 1 grows from left to right.
// Stack 2 grows from right to left.
func NewTwoStacksOneArray(capacity int) (*TwoStacksOneArray, error) {
	if capacity <= 0 {
		return nil, ErrInvalidArrayCapacity
	}

	return &TwoStacksOneArray{
		data: make([]int, capacity),

		// -1 means Stack 1 is empty.
		top1: -1,

		// capacity means Stack 2 is empty.
		// Its first element will be stored at capacity - 1.
		top2: capacity,
	}, nil
}

// Push1 adds value to Stack 1.
func (s *TwoStacksOneArray) Push1(value int) error {
	// If top1 and top2 are next to each other,
	// there is no free position left.
	if s.top1+1 == s.top2 {
		return ErrTwoStacksArrayFull
	}

	// Move Stack 1's top one position to the right.
	s.top1++

	// Store the value at the new top.
	s.data[s.top1] = value

	return nil
}

// Push2 adds value to Stack 2.
func (s *TwoStacksOneArray) Push2(value int) error {
	// Both stacks share the same free space.
	if s.top1+1 == s.top2 {
		return ErrTwoStacksArrayFull
	}

	// Move Stack 2's top one position to the left.
	s.top2--

	// Store the value at the new top.
	s.data[s.top2] = value

	return nil
}

// Pop1 removes and returns the top value from Stack 1.
func (s *TwoStacksOneArray) Pop1() (int, error) {
	if s.top1 == -1 {
		return 0, ErrStack1Empty
	}

	// Read the value before moving the top pointer.
	value := s.data[s.top1]

	// Clear the old position.
	s.data[s.top1] = 0

	// Move Stack 1's top backward.
	s.top1--

	return value, nil
}

// Pop2 removes and returns the top value from Stack 2.
func (s *TwoStacksOneArray) Pop2() (int, error) {
	if s.top2 == len(s.data) {
		return 0, ErrStack2Empty
	}

	// Read the value before moving the top pointer.
	value := s.data[s.top2]

	// Clear the old position.
	s.data[s.top2] = 0

	// Move Stack 2's top forward.
	s.top2++

	return value, nil
}

// Peek1 returns the top value of Stack 1 without removing it.
func (s *TwoStacksOneArray) Peek1() (int, error) {
	if s.top1 == -1 {
		return 0, ErrStack1Empty
	}

	return s.data[s.top1], nil
}

// Peek2 returns the top value of Stack 2 without removing it.
func (s *TwoStacksOneArray) Peek2() (int, error) {
	if s.top2 == len(s.data) {
		return 0, ErrStack2Empty
	}

	return s.data[s.top2], nil
}

// IsEmpty1 reports whether Stack 1 is empty.
func (s *TwoStacksOneArray) IsEmpty1() bool {
	return s.top1 == -1
}

// IsEmpty2 reports whether Stack 2 is empty.
func (s *TwoStacksOneArray) IsEmpty2() bool {
	return s.top2 == len(s.data)
}

// IsFull reports whether both stacks have consumed
// every position in the shared array.
func (s *TwoStacksOneArray) IsFull() bool {
	return s.top1+1 == s.top2
}

// Size1 returns the number of elements in Stack 1.
func (s *TwoStacksOneArray) Size1() int {
	return s.top1 + 1
}

// Size2 returns the number of elements in Stack 2.
func (s *TwoStacksOneArray) Size2() int {
	return len(s.data) - s.top2
}

// Capacity returns the total shared capacity.
func (s *TwoStacksOneArray) Capacity() int {
	return len(s.data)
}

// Clear1 removes every element from Stack 1.
func (s *TwoStacksOneArray) Clear1() {
	for i := 0; i <= s.top1; i++ {
		s.data[i] = 0
	}

	s.top1 = -1
}

// Clear2 removes every element from Stack 2.
func (s *TwoStacksOneArray) Clear2() {
	for i := s.top2; i < len(s.data); i++ {
		s.data[i] = 0
	}

	s.top2 = len(s.data)
}

// ToSlice1 returns Stack 1 from bottom to top.
func (s *TwoStacksOneArray) ToSlice1() []int {
	if s.IsEmpty1() {
		return []int{}
	}

	result := make([]int, s.Size1())
	copy(result, s.data[:s.top1+1])

	return result
}

// ToSlice2 returns Stack 2 from bottom to top.
func (s *TwoStacksOneArray) ToSlice2() []int {
	if s.IsEmpty2() {
		return []int{}
	}

	size := s.Size2()
	result := make([]int, size)

	// Stack 2 grows from right to left.
	// Its physical order is top -> bottom,
	// so reverse it to return bottom -> top.
	for i := 0; i < size; i++ {
		result[i] = s.data[len(s.data)-1-i]
	}

	return result
}