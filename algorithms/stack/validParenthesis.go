package stack

func isValid(s string) bool {
	// We use a slice as a stack.
	// WHY?
	// Opening brackets must be matched in the reverse order
	// in which they were encountered.
	//
	// Example:
	// "{[("
	//
	// The '(' was opened last, so ')' must close it first.
	// This is exactly LIFO (Last In, First Out), which is
	// what a stack provides.
	res := []rune{}

	// Map every closing bracket to the opening bracket
	// that it expects.
	//
	// WHY?
	// When we see a closing bracket, we need to know
	// which opening bracket should be at the top of the stack.
	//
	// ')' expects '('
	// ']' expects '['
	// '}' expects '{'
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Process every character from left to right.
	for _, ch := range s {

		// If the current character is a closing bracket,
		// we need to match it with an opening bracket.
		//
		// WHY?
		// A closing bracket cannot be pushed onto the stack.
		// Instead, it must close the most recently opened bracket.
		if ch == ')' || ch == '}' || ch == ']' {

			// If the stack is empty, there is no opening bracket
			// available to match this closing bracket.
			//
			// Example:
			// s = ")"
			//
			// There is nothing before ')' that it can close,
			// so the string is immediately invalid.
			//
			// WHY check this BEFORE res[len(res)-1]?
			// Accessing the last element of an empty slice
			// would cause a runtime panic.
			if len(res) == 0 {
				return false
			}

			// The closing bracket must match the opening bracket
			// at the top of the stack.
			//
			// WHY only check the top?
			// Because brackets must be closed in LIFO order.
			//
			// Example:
			// "{[}"
			//
			// Stack top is '['.
			// But '}' expects '{'.
			// So the order is wrong → invalid.
			if res[len(res)-1] != match[ch] {
				return false
			}

			// The opening bracket has now been successfully matched.
			// Remove it from the stack.
			//
			// WHY?
			// That opening bracket no longer needs to be matched.
			res = res[:len(res)-1]

			// IMPORTANT:
			// We don't push the closing bracket onto the stack.
			//
			// The closing bracket has already done its job:
			// it matched and removed its corresponding opening bracket.
			continue
		}

		// If the character is an opening bracket,
		// push it onto the stack.
		//
		// WHY?
		// We need to remember every unmatched opening bracket
		// until its corresponding closing bracket appears.
		res = append(res, ch)
	}

	// At the end, the stack must be empty.
	//
	// WHY?
	// If something is still in the stack, it means there is
	// an opening bracket that never received a closing bracket.
	//
	// Example:
	// s = "((("
	//
	// Stack still contains "((("
	// Therefore the string is invalid.
	return len(res) == 0
}