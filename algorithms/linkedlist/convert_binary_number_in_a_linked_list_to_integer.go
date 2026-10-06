package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
APPROACH:

The linked list contains binary digits.

Example:
    1 → 0 → 1 → 1

This represents:
    1011₂

We need to return its decimal value:
    11

For every new binary digit:

    result = result * 2 + current.Val

Why multiply by 2?

Because binary is base 2.

Multiplying by 2 shifts the existing binary number
one position to the left and creates space for the
new binary digit.

Example:

    10₂ → add 1 → 101₂

    10₂ × 2 = 100₂
    100₂ + 1 = 101₂

Therefore:

    result = result * 2 + current.Val


TIME COMPLEXITY:
    O(n)

We visit every node exactly once.

SPACE COMPLEXITY:
    O(1)

We only use a few variables.
*/

func getDecimalValue(head *ListNode) int {

	// Stores the decimal value of the binary
	// digits processed so far.
	result := 0

	current := head

	for current != nil {

		// Multiply by 2 because the number is binary.
		//
		// This shifts the existing binary digits
		// one position to the left.
		//
		// Then we add the current binary digit.
		result = result*2 + current.Val

		// Move to the next binary digit.
		current = current.Next
	}

	// Return the decimal value.
	return result
}
