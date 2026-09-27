package slidingwindow

func minSubArrayLen(target int, nums []int) int {

	// total stores the sum of the elements inside
	// our current sliding window.
	//
	// left  -> beginning of the window
	// right -> end of the window
	//
	// We initialize minLength with len(nums)+1 because
	// len(nums) is the maximum possible length of any
	// subarray.
	//
	// Therefore len(nums)+1 can never be a valid answer.
	// It acts as a "not found yet" marker.
	total, minLength, left, right := 0, len(nums)+1, 0, 0

	for right < len(nums) {

		// Expand the window by adding the new element
		// pointed to by right.
		//
		// Example:
		// [2, 3, 1]
		//  ↑     ↑
		// left  right
		//
		// total contains the sum of this entire window.
		total += nums[right]

		// Once total >= target, the current window is valid.
		//
		// But we don't want to stop here because the current
		// window may contain unnecessary elements at the left.
		//
		// So we keep shrinking the window as long as it is
		// still valid.
		//
		// WHY "for" instead of "if"?
		// Because we may be able to remove multiple elements
		// from the left and still keep total >= target.
		for total >= target {

			// The current window is:
			//
			// nums[left : right+1]
			//
			// Its length is:
			// right - left + 1
			//
			// We compare it with the smallest valid window
			// found so far.
			minLength = min(right-left+1, minLength)

			// Remove the leftmost element from the window.
			//
			// WHY?
			// We already know the current window is valid.
			// Now we want to make it smaller and check whether
			// it can still satisfy the target.
			total -= nums[left]

			// Move left forward because nums[left] has just
			// been removed from the window.
			left++
		}

		// Move right forward to expand the window again.
		//
		// This allows us to include the next element and
		// search for another possible valid window.
		right++
	}

	// If minLength is still len(nums)+1, it means we never
	// found any subarray whose sum was >= target.
	//
	// len(nums)+1 was our "not found" marker.
	if minLength == len(nums)+1 {
		return 0
	}

	// Return the smallest valid subarray length found.
	return minLength
}