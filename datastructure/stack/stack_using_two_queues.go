/*
# Stack Using Two Queues

A Stack Using Two Queues implements stack behavior using two queues.

A queue follows:

	FIFO = First In, First Out

A stack requires:

	LIFO = Last In, First Out

Since a queue removes elements from the front, we use two queues
to rearrange elements when performing Pop().

This implementation keeps newly pushed elements at the back of
the active queue.

Example:

	Push(10)
	Push(20)
	Push(30)

	Queue 1:

	[10, 20, 30]

	The stack logically represents:

	Top
	 ↓
	30
	20
	10

## Push

Push is simple:

	1. Add the new value to the active queue.
	2. No rearrangement is required.

Therefore:

	Push: O(1)

Example:

	[10, 20]

	Push(30)

	[10, 20, 30]

The newest value is at the back.

## Pop

The back of a queue cannot be removed directly.

To remove the newest element:

	1. Move all elements except the last one from the active
	   queue to the second queue.
	2. The remaining element is the stack's top.
	3. Remove that element.
	4. Swap the two queues.

Example:

	Queue 1:

	[10, 20, 30]

	Move 10:

	Queue 1: [30]
	Queue 2: [10]

	Move 20:

	Queue 1: [30]
	Queue 2: [10, 20]

	Remove 30.

	Queue 2 becomes the active queue:

	[10, 20]

## Peek

Peek uses the same process as Pop, except the last element
is not removed.

Therefore Peek is O(n).

## Queue roles

At any point:

	active queue
		Contains all current stack elements.

	helper queue
		Used temporarily during Pop() or Peek().

After an operation, the queues are swapped.

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

	Push:     O(1)
	Pop:      O(n)
	Peek:     O(n)
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

var ErrTwoQueueStackEmpty = errors.New("stack using two queues is empty")

// TwoQueueStack represents a stack implemented using two queues.
type TwoQueueStack struct {
	queue1 []int
	queue2 []int

	// active indicates which queue currently contains the stack.
	active int
}

// NewTwoQueueStack creates an empty stack.
func NewTwoQueueStack() *TwoQueueStack {
	return &TwoQueueStack{
		queue1: []int{},
		queue2: []int{},
		active: 1,
	}
}

// Push adds value to the top of the stack.
//
// The new element is added to the back of the active queue.
// No rearrangement is necessary.
func (s *TwoQueueStack) Push(value int) {
	if s.active == 1 {
		s.queue1 = append(s.queue1, value)
		return
	}

	s.queue2 = append(s.queue2, value)
}

// Pop removes and returns the top value.
//
// Why move n-1 elements?
// The newest element is at the back of the queue, but a queue
// only removes elements from the front. We move every older
// element to the helper queue so the newest element is left behind.
func (s *TwoQueueStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrTwoQueueStackEmpty
	}

	var activeQueue *[]int
	var helperQueue *[]int

	if s.active == 1 {
		activeQueue = &s.queue1
		helperQueue = &s.queue2
	} else {
		activeQueue = &s.queue2
		helperQueue = &s.queue1
	}

	// Move every element except the last one.
	for len(*activeQueue) > 1 {
		front := (*activeQueue)[0]

		*activeQueue = (*activeQueue)[1:]
		*helperQueue = append(*helperQueue, front)
	}

	// The remaining element is the newest element,
	// which represents the top of the stack.
	value := (*activeQueue)[0]

	// Remove the top element.
	*activeQueue = (*activeQueue)[:0]

	// The helper queue now contains the remaining stack elements.
	s.active = 3 - s.active

	return value, nil
}

// Peek returns the top value without removing it.
//
// The same movement is required as Pop(), but the last element
// is placed into the helper queue so that it remains in the stack.
func (s *TwoQueueStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrTwoQueueStackEmpty
	}

	var activeQueue *[]int
	var helperQueue *[]int

	if s.active == 1 {
		activeQueue = &s.queue1
		helperQueue = &s.queue2
	} else {
		activeQueue = &s.queue2
		helperQueue = &s.queue1
	}

	// Move every element except the last one.
	for len(*activeQueue) > 1 {
		front := (*activeQueue)[0]

		*activeQueue = (*activeQueue)[1:]
		*helperQueue = append(*helperQueue, front)
	}

	// Read the last element.
	value := (*activeQueue)[0]

	// Move it as well because Peek must not remove anything.
	*helperQueue = append(*helperQueue, value)

	// The helper queue now contains the complete stack.
	*activeQueue = (*activeQueue)[:0]

	s.active = 3 - s.active

	return value, nil
}

// IsEmpty reports whether the stack contains no elements.
func (s *TwoQueueStack) IsEmpty() bool {
	if s.active == 1 {
		return len(s.queue1) == 0
	}

	return len(s.queue2) == 0
}

// Size returns the number of elements in the stack.
func (s *TwoQueueStack) Size() int {
	if s.active == 1 {
		return len(s.queue1)
	}

	return len(s.queue2)
}

// Contains checks whether value exists in the stack.
func (s *TwoQueueStack) Contains(value int) bool {
	if s.active == 1 {
		left := 0
		right := len(s.queue1) - 1

		for left <= right {
			if s.queue1[left] == value || s.queue1[right] == value {
				return true
			}

			left++
			right--
		}

		return false
	}

	left := 0
	right := len(s.queue2) - 1

	for left <= right {
		if s.queue2[left] == value || s.queue2[right] == value {
			return true
		}

		left++
		right--
	}

	return false
}

// Clear removes all elements from the stack.
func (s *TwoQueueStack) Clear() {
	s.queue1 = []int{}
	s.queue2 = []int{}
	s.active = 1
}

// ToSlice returns the stack contents from top to bottom.
func (s *TwoQueueStack) ToSlice() []int {
	var queue []int

	if s.active == 1 {
		queue = s.queue1
	} else {
		queue = s.queue2
	}

	result := make([]int, len(queue))

	// The queue is bottom -> top, while a stack is top -> bottom.
	for i := range queue {
		result[len(queue)-1-i] = queue[i]
	}

	return result
}