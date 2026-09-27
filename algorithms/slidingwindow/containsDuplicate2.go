package slidingwindow

func containsNearbyDuplicate(nums []int, k int) bool {
	// Store the most recent index where each number appeared.
	//
	// Example:
	// nums = [1, 2, 3, 1]
	//
	// seen:
	// 1 -> 0
	// 2 -> 1
	// 3 -> 2
	seen := make(map[int]int)

	for i, num := range nums {

		// Check whether this number appeared previously.
		prevIndex, exists := seen[num]

		if exists {

			// prevIndex is always smaller than i because
			// it came from an earlier iteration.
			//
			// Therefore:
			// i - prevIndex
			//
			// is already the absolute distance between the
			// two occurrences. No abs() is necessary.
			if i-prevIndex <= k {
				return true
			}
		}

		// Store the CURRENT index as the latest occurrence.
		//
		// WHY update it?
		// If this number appears again later, its most recent
		// occurrence gives the smallest possible distance.
		seen[num] = i
	}

	// No duplicate pair was found within distance k.
	return false
}