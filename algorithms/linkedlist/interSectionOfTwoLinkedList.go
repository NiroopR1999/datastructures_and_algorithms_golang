package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  time optmised
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    seen:=make(map[*ListNode]bool) 
    first:=headA

    for first!=nil {
        seen[first]=true
        first=first.Next
    }
    second:=headB
    for second!=nil {
        if seen[second] {
            return second
        }
        second=second.Next
    }
    return nil
}


// time and space optimised

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
APPROACH:

Use two pointers, one for each linked list.

Why two pointers?

The two lists may have different lengths before they
merge.

Example:

A: 1 → 2 → 3 ─┐
               ↓
               7 → 8 → 9
               ↑
B:     4 → 5 ─┘

List A has 3 nodes before the intersection.
List B has 2 nodes before the intersection.

So if both pointers simply move forward, they will
not reach the intersection at the same time.

The trick:

1. Start one pointer at A and one at B.
2. Move both one node at a time.
3. When a pointer reaches the end of its list,
   move it to the HEAD of the OTHER list.

Why switch lists?

Suppose:

    A's unique part = a
    B's unique part = b
    Common part    = c

Pointer 1 travels:

    A + B
    = a + c + b + c

Pointer 2 travels:

    B + A
    = b + c + a + c

Therefore both pointers travel the SAME total distance.

The unequal parts (a and b) cancel each other out.

After switching lists, both pointers are effectively
aligned with respect to the common portion.

Therefore:

    - If the lists intersect, the pointers meet at
      the intersection node.
    - If they don't intersect, both eventually become nil.

Time:  O(m + n)
Space: O(1)
*/

func getIntersectionNode1(headA, headB *ListNode) *ListNode {

    // Start one pointer at each list.
    //
    // Why?
    // Each pointer initially walks through its own list.
    // Since the lists can have different lengths, they
    // may not reach the common part at the same time.
    first := headA
    second := headB

    for first != second {

        // Move first one node forward.
        //
        // If first reaches the end of List A, instead of
        // stopping, move it to the beginning of List B.
        //
        // Why?
        // This makes first walk BOTH lists:
        //
        //     A → B
        //
        // So first travels:
        //
        //     length(A) + length(B)
        //
        // This compensates for any extra nodes that A
        // had compared with B.
        if first == nil {
            first = headB
        } else {
            first = first.Next
        }

        // Move second one node forward.
        //
        // If second reaches the end of List B, move it
        // to the beginning of List A.
        //
        // Why?
        // This makes second walk BOTH lists:
        //
        //     B → A
        //
        // So second also travels:
        //
        //     length(B) + length(A)
        //
        // Now both pointers have travelled exactly the
        // same total distance.
        if second == nil {
            second = headA
        } else {
            second = second.Next
        }
    }

    // Why can we simply return first?
    //
    // The loop stops when:
    //
    //     first == second
    //
    // There are only two possibilities:
    //
    // 1. They meet at a common node.
    //    → That node is the intersection.
    //
    // 2. The lists do not intersect.
    //    → Both eventually become nil.
    //
    // Therefore first is either:
    //
    //     intersection node
    //
    // or
    //
    //     nil
    return first
}