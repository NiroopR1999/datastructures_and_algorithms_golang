package linkedlist


// detectCycle returns the node where the cycle begins.
// If there is no cycle, it returns nil.
//
// Approach:
// We use Floyd's Cycle Detection algorithm.
//
// Phase 1:
// - slow moves 1 step.
// - fast moves 2 steps.
// - If they meet, a cycle exists.
//
// Phase 2:
// - Move slow back to head.
// - Keep fast at the meeting point.
// - Move both 1 step at a time.
// - They will meet exactly at the cycle's starting node.
//
// Why does Phase 2 work?
//
// Let:
//   L = distance from head to the cycle start.
//   K = distance from cycle start to the meeting point.
//
// When slow and fast meet:
//
//   slow traveled = L + K
//
// Since fast moves twice as fast, fast has traveled twice
// the distance of slow. The extra distance traveled by fast
// must therefore be a whole number of complete cycles.
//
// This means the distance from the meeting point back to
// the cycle start is equivalent to L within the cycle.
//
// Therefore:
//   - slow starts at head and needs L steps to reach cycle start.
//   - fast starts at meeting point and also needs L steps
//     to reach cycle start.
//
// So, when both move one step at a time, they meet at the
// cycle's starting node.
//
// Time Complexity:  O(n)
// Space Complexity: O(1)
func detectCycle(head *ListNode) *ListNode {

	// Both pointers start at head because we want to traverse
	// the list at different speeds.
	slow, fast := head, head

	// PHASE 1: Detect whether a cycle exists.
	//
	// fast moves two steps, so fast and fast.Next must exist
	// before we access fast.Next.Next.
	for fast != nil && fast.Next != nil {

		// Move slow one step.
		slow = slow.Next

		// Move fast two steps.
		// The speed difference allows fast to catch slow
		// if both are moving around a cycle.
		fast = fast.Next.Next

		// If they point to the same node, fast has caught slow.
		// Two pointers moving at different speeds can meet again
		// only if there is a cycle.
		if slow == fast {
			break
		}
	}

	// If fast reached the end, there is no cycle.
	//
	// Without a cycle, fast will eventually become nil
	// because it is moving twice as fast as slow.
	if fast == nil || fast.Next == nil {
		return nil
	}

	// PHASE 2: Find the beginning of the cycle.
	//
	// We know slow and fast are currently at the meeting point.
	//
	// Move slow back to head.
	//
	// Why?
	// The distance from head to the cycle start is exactly the
	// same effective distance that fast needs to travel from the
	// meeting point to reach the cycle start.
	slow = head

	// Now both pointers move at the same speed.
	//
	// slow:
	//   head → → → cycle start
	//
	// fast:
	//   meeting point → → → cycle start
	//
	// Because these distances are equal, they meet at
	// the beginning of the cycle.
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// The meeting node is the first node of the cycle.
	return slow
}