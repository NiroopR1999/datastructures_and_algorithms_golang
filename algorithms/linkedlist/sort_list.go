package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// mergeSort merges two already-sorted linked lists.
//
// WHY:
// After recursively splitting the original list, we will have:
// left  = sorted first half
// right = sorted second half
//
// We now merge them into one sorted list.
func mergeSort1(left, right *ListNode) *ListNode {

    // head always points to the FIRST node of the merged list.
    // We return head at the end.
    var head *ListNode

    // tail always points to the LAST node of the merged list.
    // We use tail to append new nodes efficiently.
    var tail *ListNode

    // Continue while both lists still have nodes.
    //
    // WHY:
    // We need to compare the current node of both lists
    // and take the smaller one.
    for left != nil && right != nil {

        if left.Val < right.Val {

            // left is smaller, so add it to the merged list.
            node := left

            if head == nil {
                // This is the FIRST node we're adding.
                //
                // Both head and tail should point to it.
                head = node
                tail = node
            } else {
                // Add the node after the current last node.
                tail.Next = node

                // Move tail forward because node is now
                // the last node in our merged list.
                tail = tail.Next
            }

            // Move left forward because we already used
            // its current node.
            left = left.Next

        } else {

            // right is smaller (or equal), so add it.
            node := right

            if head == nil {
                // First node of the merged list.
                head = node
                tail = node

            } else {
                // Connect the new node after the current tail.
                tail.Next = node

                // Move tail to the newly added node.
                tail = tail.Next
            }

            // Move right forward because we already used
            // its current node.
            right = right.Next
        }
    }

    // At this point, one of the lists is empty.
    //
    // WHY:
    // The remaining nodes are already sorted, so we don't
    // need to compare them individually.
    if left != nil {
        tail.Next = left
    }

    if right != nil {
        tail.Next = right
    }

    // head points to the beginning of the complete sorted list.
    return head
}


func sortList(head *ListNode) *ListNode {

    // BASE CASE:
    // 0 nodes or 1 node is already sorted.
    //
    // WHY:
    // We need this to stop the recursion.
    // Without head.Next == nil, a one-node list would
    // keep recursively calling sortList on itself.
    if head == nil || head.Next == nil {
        return head
    }

    // Use slow/fast pointers to find the middle.
    //
    // slow moves 1 step.
    // fast moves 2 steps.
    //
    // Therefore, when fast reaches the end,
    // slow is around the middle.
    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Recursively sort the RIGHT half.
    //
    // slow.Next is the first node of the right half.
    right := sortList(slow.Next)

    // IMPORTANT:
    // Break the linked list into two separate lists.
    //
    // Before:
    //
    // left half → slow → right half
    //
    // After:
    //
    // left half → slow     right half
    //
    // Without this, the two halves would still be connected.
    slow.Next = nil

    // Recursively sort the LEFT half.
    left := sortList(head)

    // Now both halves are sorted.
    //
    // Merge them into one sorted list.
    return mergeSort(left, right)
}