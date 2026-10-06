package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// hasCycle determines whether a linked list contains a cycle.
//
// Approach:
// We use Floyd's Cycle Detection algorithm (Tortoise and Hare).
//
// - slow moves one node at a time.
// - fast moves two nodes at a time.
// - If there is no cycle, fast will eventually reach nil.
// - If there is a cycle, fast will eventually catch slow.
//
// Why does fast catch slow?
// Once both pointers are inside the cycle, fast gains one node
// on slow during every iteration. Since the cycle is finite,
// fast must eventually land on the same node as slow.
//
// Time Complexity:  O(n)
// Space Complexity: O(1)
func hasCycle(head *ListNode) bool {

	// Both pointers start at the same node because we want
	// to traverse the list using two different speeds.
	slow, fast := head, head

	// We access fast.Next.Next below, so fast and fast.Next
	// must both exist before making the two-step movement.
	for fast != nil && fast.Next != nil {

		// Move slow by one step.
		// This gives fast a speed advantage of one node per iteration.
		slow = slow.Next

		// Move fast by two steps.
		// If a cycle exists, this extra speed will eventually
		// allow fast to catch up with slow.
		fast = fast.Next.Next

		// If both pointers reference the exact same node,
		// fast has caught slow, which can only happen inside a cycle.
		if slow == fast {
			return true
		}
	}

	// fast reached the end of the list, so there is no cycle.
	return false
}