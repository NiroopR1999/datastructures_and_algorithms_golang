package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {

    /*
    APPROACH:
    Since the linked list is sorted, duplicate values will always
    appear next to each other.

    We keep a pointer called current and compare:
    
        current.Val
        current.Next.Val

    If they are equal:
        current.Next is a duplicate, so skip it by doing:

        current.Next = current.Next.Next

    If they are different:
        Move current to the next node.

    IMPORTANT:
    When we remove a duplicate, we DO NOT move current.
    There may be more duplicates immediately after it.

    Example:

        1 → 1 → 1 → 2 → 3 → 3

    After removing duplicates:

        1 → 2 → 3

    Time:  O(n)
    Space: O(1)
    */

    // current points to the node we are currently checking.
    current := head

    // We need current.Next because we compare current
    // with the node immediately after it.
    for current != nil && current.Next != nil {

        if current.Val == current.Next.Val {

            // current.Next is a duplicate.
            //
            // Skip the duplicate and connect current
            // directly to the node after it.
            //
            // Example:
            //
            // Before:
            // 1 → 1 → 2
            // ↑   ↑
            // current
            //
            // After:
            // 1 ─────→ 2
            //
            current.Next = current.Next.Next

            // Don't move current.
            //
            // There could be another duplicate:
            //
            // 1 → 1 → 1 → 2
            // ↑
            // current
            //
            // After one deletion:
            //
            // 1 → 1 → 2
            // ↑
            // current
            //
            // So we check current.Next again.

        } else {

            // The next node has a different value,
            // so move current forward.
            current = current.Next
        }
    }

    return head
}