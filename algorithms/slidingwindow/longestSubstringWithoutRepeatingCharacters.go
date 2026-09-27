package slidingwindow

func lengthOfLongestSubstring(s string) int {
	// maxLen stores the longest substring length we have found so far.
	maxLen := 0

	// left and right represent the current sliding window.
	//
	// Example:
	// s = "abc"
	//      ^ ^
	//    left right
	//
	// Everything between left and right is our current substring.
	left, right := 0, 0

	// seen stores the latest index of each character.
	//
	// Example:
	// "abc"
	//
	// seen:
	// a -> 0
	// b -> 1
	// c -> 2
	//
	// We need the index because when we find a duplicate,
	// we need to know where the previous occurrence was.
	seen := make(map[byte]int)

	for right < len(s) {

		// Check whether the current character was seen before.
		//
		// index  = where we last saw this character
		// exists = whether the character exists in the map
		index, exists := seen[s[right]]

		if exists {

			// The previous occurrence is at 'index'.
			//
			// We cannot keep that old character inside our
			// current window because then the substring would
			// contain a duplicate.
			//
			// So left needs to move AFTER the old character.
			//
			// Example:
			//
			// s = "abca"
			//      0123
			//
			// When right = 3:
			// current character = 'a'
			// previous 'a'     = index 0
			//
			// Therefore left must move to:
			//
			// index + 1 = 1
			//
			// This gives us:
			//
			// "bca"
			//
			// Why max()?
			//
			// Because left must NEVER move backward.
			//
			// Example:
			// left = 2
			// index = 0
			//
			// The old character is already outside our window,
			// so we should keep left at 2.
			//
			// max(2, 0+1) = 2
			left = max(index+1, left)
		}

		// Store the current character's latest position.
		//
		// We use the latest position because if this character
		// appears again later, we need the most recent occurrence
		// to decide where left should move.
		seen[s[right]] = right

		// Calculate the size of the current window.
		//
		// Example:
		// left = 2
		// right = 5
		//
		// Window contains indices:
		// 2, 3, 4, 5
		//
		// That's 4 characters.
		//
		// So:
		// right - left + 1
		//
		// The +1 is needed because both left and right
		// are included in the window.
		maxLen = max(maxLen, right-left+1)

		// Move right forward to examine the next character.
		right++
	}

	// Return the longest substring length found.
	return maxLen
}