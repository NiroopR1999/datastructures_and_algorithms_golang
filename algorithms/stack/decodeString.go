package stack

func decodeString(s string) string {
	stringStack, numStack := []string{}, []int{}
	currentString, repeat := "", 0

	for _, c := range s {

		if c >= '0' && c <= '9' {

			// Build the repeat number.
			//
			// Example:
			// "123" → 1 → 12 → 123
			repeat = repeat*10 + int(c-'0')

		} else if c == '[' {

			// Save the repeat count because it belongs
			// to the string inside these brackets.
			numStack = append(numStack, repeat)

			// Save the string that existed BEFORE '['.
			//
			// Example:
			// "ab2[c]"
			//
			// Before entering [:
			// currentString = "ab"
			//
			// We need to remember "ab" so we can attach
			// the decoded "cc" to it later.
			stringStack = append(stringStack, currentString)

			// We are now starting to build the string
			// INSIDE the brackets.
			currentString = ""

			// The repeat number has been consumed.
			repeat = 0

		} else if c == ']' {

			// Get the repeat count for this bracket.
			prevNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			// Get the string that existed before '['.
			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			// Save the decoded string inside the brackets.
			//
			// Example:
			// 2[cd]
			//
			// currentString = "cd"
			temp := currentString

			// Repeat the INNER decoded string.
			//
			// 2[cd] → "cdcd"
			//
			// Then attach it to the previous string.
			//
			// If:
			// prevString = "ab"
			// temp       = "cd"
			//
			// result = "ab" + "cdcd"
			currentString = prevString

			for i := 0; i < prevNum; i++ {
				currentString += temp
			}

		} else {

			// Normal character.
			//
			// Add it to the current string.
			currentString += string(c)
		}
	}

	return currentString
}