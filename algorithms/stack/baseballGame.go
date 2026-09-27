package stack

import "strconv"

func calPoints(operations []string) int {

	// We use a slice as a stack to store the valid scores.
	//
	// WHY a stack?
	// The operations "D", "C", and "+" depend on the most
	// recently added scores.
	//
	// Example:
	// [5, 2]
	//          ↑
	//       most recent
	//
	// A slice gives us easy access to the last element.
	res := []int{}

	// Process each operation from left to right.
	for _, ops := range operations {

		// Each operation tells us what to do with the scores
		// currently stored in the stack.
		switch ops {

		case "D":
			// "D" means: double the previous score.
			//
			// WHY len(res)-1?
			// The last element in the slice is the most recent
			// valid score.
			prev := res[len(res)-1]

			// Double the previous score.
			//
			// Example:
			// res = [5]
			// D → 10
			prev = prev * 2

			// Add the new score to the stack.
			//
			// WHY append?
			// The doubled score is a new valid round,
			// so it becomes the new most recent score.
			res = append(res, prev)

		case "C":
			// "C" means: invalidate/remove the previous score.
			//
			// Example:
			// res = [5, 10, 15]
			// C
			// res = [5, 10]
			//
			// WHY len(res)-1?
			// The last element is the score that needs
			// to be removed.
			res = res[:len(res)-1]

		case "+":
			// "+" means: create a new score equal to the
			// sum of the previous two valid scores.
			//
			// Example:
			// res = [5, 10]
			// + → 15
			//
			// The last score is at len(res)-1.
			prev1 := res[len(res)-1]

			// The second-last score is at len(res)-2.
			prev2 := res[len(res)-2]

			// Add the two previous scores and push the
			// result as a new valid score.
			res = append(res, prev1+prev2)

		default:
			// If the operation isn't "D", "C", or "+",
			// it must be a numeric score.
			//
			// Example:
			// "5"  → 5
			// "10" → 10
			// "-3" → -3
			//
			// WHY strconv.Atoi?
			// operations contains strings, but our stack
			// stores integers.
			val, _ := strconv.Atoi(ops)

			// Add the numeric score to the stack.
			res = append(res, val)
		}
	}

	// At this point, res contains all scores that are still valid.
	//
	// We now calculate their total.
	total := 0

	for i := range res {
		total += res[i]
	}

	// Return the final sum of all valid scores.
	return total
}