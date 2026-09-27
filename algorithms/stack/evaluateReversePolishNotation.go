package stack

import "strconv"

func evalRPN(tokens []string) int {

	// We use a stack because in Reverse Polish Notation,
	// an operator always works on the most recently seen
	// two numbers.
	stack := []int{}

	for _, token := range tokens {

		// If the token is not an operator, it must be a number.
		//
		// Numbers are pushed onto the stack because we may need
		// them later when we encounter an operator.
		if token != "+" &&
			token != "-" &&
			token != "*" &&
			token != "/" {

			val, _ := strconv.Atoi(token)

			stack = append(stack, val)

			// There is nothing else to do with a number.
			// We wait until an operator tells us to use it.
			continue
		}

		// The top of the stack is the RIGHT operand.
		//
		// Example:
		//     5 - 2
		//
		// Stack:
		//     [5, 2]
		//
		// 2 is popped first, so it is the right operand.
		right := stack[len(stack)-1]

		// The next element is the LEFT operand.
		//
		// 5 is popped second, so it is the left operand.
		left := stack[len(stack)-2]

		// Remove both operands because we are going to replace
		// them with the result of the operation.
		//
		// Example:
		//     [5, 2] → []
		stack = stack[:len(stack)-2]

		// Perform the operation and push the result back.
		//
		// Why push the result?
		// Because this result may itself be used by a later operator.
		switch token {
		case "+":
			stack = append(stack, left+right)

		case "-":
			// Order matters for subtraction.
			// We must do LEFT - RIGHT.
			stack = append(stack, left-right)

		case "*":
			stack = append(stack, left*right)

		case "/":
			// Order matters for division as well.
			// We must do LEFT / RIGHT.
			stack = append(stack, left/right)
		}
	}

	// After processing every token, the stack contains
	// exactly one value: the final result.
	return stack[0]
}