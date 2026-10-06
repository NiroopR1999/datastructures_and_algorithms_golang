package linkedlist
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates2(head *ListNode) *ListNode {

    /*
    APPROACH:

    This problem is different from LeetCode 83.

    Here, if a value appears more than once, we must remove
    ALL occurrences of that value.

    Example:

        1 → 2 → 3 → 3 → 4 → 4 → 5

    Result:

        1 → 2 → 5

    Since the list is sorted, duplicate values are always
    next to each other.

    We use a dummy node before head.

        dummy → 1 → 2 → 3 → 3 → 4 → 4 → 5

    `prev` points to the last node that we know is NOT duplicated.

    `current` checks the current group of values.

    If current.Val == current.Next.Val:
        We found a duplicate group.

        Keep moving current until the duplicate group ends.

        Then connect prev.Next directly to current.Next,
        effectively removing the entire duplicate group.

    Otherwise:
        The current node is unique, so move prev forward.

    WHY DUMMY?

    The first nodes themselves might be duplicates.

    Example:

        1 → 1 → 2 → 3

    The answer should start at 2.

    A dummy node gives us a node before the head,
    so we can remove duplicate nodes at the beginning
    using the exact same logic.

    Time:  O(n)
    Space: O(1)
    */

    // Dummy node makes it possible to remove duplicate nodes
    // even when the duplicates start at the head.
    dummy := &ListNode{Next: head}

    // prev points to the last node that we know is unique.
    prev := dummy

    // current is used to inspect the list.
    current := head

    for current != nil {

        // Check whether current is the beginning of
        // a duplicate group.
        if current.Next != nil && current.Val == current.Next.Val {

            // We found a duplicate.
            //
            // Keep moving current until we reach the end
            // of this duplicate group.
            duplicateValue := current.Val

            for current != nil && current.Val == duplicateValue {
                current = current.Next
            }

            // Remove the entire duplicate group.
            //
            // Example:
            //
            // prev → 3 → 3 → 4
            //        ↑
            //      current
            //
            // After:
            //
            // prev ─────────→ 4
            prev.Next = current

        } else {

            // current is unique.
            //
            // Move prev forward because current is now
            // confirmed to be part of the answer.
            prev = current

            // Move current forward to check the next node.
            current = current.Next
        }
    }

    // dummy.Next is the new head of the list.
    return dummy.Next
}