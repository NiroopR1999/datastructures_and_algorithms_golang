package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {

	/*
		APPROACH:

		We have K sorted linked lists.

		Instead of merging all lists at once, we merge
		them TWO at a time.

		Example:

		List 1: 1 → 4 → 5
		List 2: 1 → 3 → 4
		List 3: 2 → 6

		First:
		    result = List 1

		Then merge:
		    result + List 2

		Then:
		    result + List 3

		So:

		    List 1
		      ↓
		    merge with List 2
		      ↓
		    result
		      ↓
		    merge with List 3
		      ↓
		    final result

		mergeSort() is responsible for merging TWO sorted lists.

		Time:  O(N * K)
		Space: O(1)

		N = total number of nodes
		K = number of lists
	*/

	// If there are no lists, there is nothing to merge.
	if len(lists) == 0 {
		return nil
	}

	// Start with the first list as our initial result.
	result := lists[0]

	// Merge every remaining list into result one by one.
	for i := 1; i < len(lists); i++ {

		// mergeSort() merges two sorted lists
		// and returns one sorted list.
		result = mergeSort(result, lists[i])
	}

	// result now contains all K lists merged
	// into one sorted linked list.
	return result
}

func mergeSort(left, right *ListNode) *ListNode {

	/*
		WHY DO WE CHECK nil FIRST?

		If one of the lists is empty, we don't need
		to perform any merging.

		Example:

		    left  = nil
		    right = 1 → 2 → 3

		The answer is simply:

		    1 → 2 → 3

		These checks also prevent us from trying to use
		tail.Next when tail is nil.
	*/

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	/*
		head = first node of our final merged list.

		tail = last node of our final merged list.

		IMPORTANT:

		head NEVER moves.

		tail keeps moving forward whenever we add
		a new node.

		Think:

		    head
		     ↓
		    1 → 3 → 5
		          ↑
		         tail
	*/

	var head, tail *ListNode

	// Keep comparing nodes while BOTH lists still
	// have nodes remaining.
	for left != nil && right != nil {

		// If left node is smaller, take the left node.
		if left.Val < right.Val {

			// Store the current left node that
			// we want to add to our result.
			node := left

			// If this is the FIRST node,
			// both head and tail should point to it.
			if head == nil {
				head, tail = node, node
			} else {

				// Connect the new node after tail.
				tail.Next = node

				// Move tail forward because node
				// is now the last node.
				tail = tail.Next
			}

			// Move left forward because we have
			// already added the current left node.
			left = left.Next

		} else {

			// If right.Val is smaller OR equal,
			// take the right node.
			node := right

			// If this is the FIRST node,
			// initialize both head and tail.
			if head == nil {
				head, tail = node, node
			} else {

				// Connect the new node after tail.
				tail.Next = node

				// Move tail to the newly added node.
				tail = tail.Next
			}

			// Move right forward because we have
			// already added the current right node.
			right = right.Next
		}
	}

	/*
		At this point, one list is empty.

		Example:

		    left:  5 → 7 → 9
		    right: nil

		Because left is already sorted, we don't need
		to compare anything anymore.

		We can simply connect the remaining nodes
		directly after tail.
	*/

	if left != nil {
		// Right list is finished.
		// Attach the remaining left nodes.
		tail.Next = left
	} else {
		// Left list is finished.
		// Attach the remaining right nodes.
		tail.Next = right
	}

	// head points to the first node of our
	// completely merged sorted list.
	return head
}