package slidingwindow

func checkInclusion(s1 string, s2 string) bool {

	// WHY:
	// A permutation of s1 can only exist if s2 has at least
	// as many characters as s1.
	//
	// Example:
	// s1 = "abc"  -> needs 3 characters
	// s2 = "ab"   -> only has 2 characters
	//
	// Therefore, finding a permutation is impossible.
	if len(s1) > len(s2) {
		return false
	}

	// WHY [26]int?
	// We are assuming s1 and s2 contain only lowercase English
	// letters ('a' to 'z').
	//
	// Each index represents one character:
	//
	// index 0  -> 'a'
	// index 1  -> 'b'
	// index 2  -> 'c'
	// ...
	// index 25 -> 'z'
	//
	// count1 stores the frequency of every character in s1.
	// count2 stores the frequency of every character in our
	// current window inside s2.
	//
	// We use arrays instead of slices because:
	// 1. The size is fixed at 26.
	// 2. Arrays can be directly compared using == in Go.
	var count1, count2 [26]int

	// WHY do this loop?
	//
	// We need two things:
	//
	// 1. The character frequency of s1.
	// 2. The character frequency of the FIRST window of s2.
	//
	// The window must have the same size as s1 because a
	// permutation of s1 must contain exactly len(s1) characters.
	//
	// Example:
	//
	// s1 = "ab"
	// s2 = "eidbaooo"
	//
	// First window of s2 = "ei"
	//
	// We compare:
	//
	// s1     -> "ab" -> a:1, b:1
	// window -> "ei" -> e:1, i:1
	for i := range s1 {

		// WHY s1[i]-'a'?
		//
		// We convert the character into an array index.
		//
		// 'a' - 'a' = 0
		// 'b' - 'a' = 1
		// 'c' - 'a' = 2
		//
		// So if s1[i] is 'b', we increment count1[1].
		//
		// WHY increment?
		// We are counting how many times this character occurs
		// in s1.
		count1[s1[i]-'a'] = count1[s1[i]-'a'] + 1

		// WHY s2[i] here?
		//
		// During this first loop, we are also building the
		// FIRST window of s2.
		//
		// Since the loop runs len(s1) times, we take exactly
		// len(s1) characters from the beginning of s2.
		//
		// Example:
		// s1 = "ab" -> len = 2
		// s2 = "eidbaooo"
		//
		// i = 0 -> s2[0] = 'e'
		// i = 1 -> s2[1] = 'i'
		//
		// Therefore the first window is "ei".
		count2[s2[i]-'a'] = count2[s2[i]-'a'] + 1
	}

	// WHY check here?
	//
	// The first window has already been created.
	//
	// Example:
	//
	// s1 = "ab"
	// count1 = {a:1, b:1}
	//
	// first window = "ei"
	// count2 = {e:1, i:1}
	//
	// If both arrays are equal, the current window contains
	// exactly the same characters as s1.
	//
	// The order does NOT matter.
	//
	// "ab" -> a:1, b:1
	// "ba" -> a:1, b:1
	//
	// Therefore equal frequency arrays mean the window is
	// a permutation of s1.
	if count1 == count2 {
		return true
	}

	// WHY start i at len(s1)?
	//
	// The first window has already been checked.
	//
	// If len(s1) = 2:
	//
	// First window:
	// [0 1]
	//
	// The next character we haven't processed is index 2.
	//
	// So:
	//
	// i = 2
	//
	// We now slide the window one position at a time.
	for i := len(s1); i < len(s2); i++ {

		// WHY add s2[i]?
		//
		// s2[i] is the NEW character entering the window
		// from the right.
		//
		// Example:
		//
		// Before:
		// [e i] d b a
		//
		// i = 2 -> s2[2] = 'd'
		//
		// Temporarily:
		// [e i d]
		//
		// So we add 'd' to our frequency count.
		count2[s2[i]-'a'] = count2[s2[i]-'a'] + 1

		// WHY remove s2[i-len(s1)]?
		//
		// We just added a new character, but our window must
		// ALWAYS remain exactly len(s1) characters long.
		//
		// Therefore, we must remove the character that is
		// falling out from the LEFT side of the window.
		//
		// Example:
		//
		// s1 length = 2
		//
		// Current window:
		// [e i]
		//
		// Add 'd':
		// [e i d]
		//
		// 'e' is now outside the new 2-character window.
		//
		// Its index is:
		//
		// i - len(s1)
		// 2 - 2
		// = 0
		//
		// s2[0] = 'e'
		//
		// So we remove 'e' from count2.
		count2[s2[i-len(s1)]-'a'] = count2[s2[i-len(s1)]-'a'] - 1

		// WHY compare again?
		//
		// After adding the new character and removing the old
		// character, count2 now represents the NEW window.
		//
		// If count1 == count2, this window contains exactly
		// the same characters as s1.
		//
		// Therefore, this window is a permutation of s1.
		if count1 == count2 {
			return true
		}
	}

	// WHY return false?
	//
	// We checked:
	// 1. The first window.
	// 2. Every possible window after sliding.
	//
	// If none of the windows had the same character frequencies
	// as s1, then no permutation of s1 exists inside s2.
	return false
}