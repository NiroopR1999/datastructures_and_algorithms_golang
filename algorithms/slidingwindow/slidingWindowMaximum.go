package slidingwindow

func maxSlidingWindow(nums []int, k int) []int {
	n, left, right := len(nums), 0, 0
	q, res := []int{}, []int{}

	for right < n {

		// q stores INDICES, not values.
		//
		// The indices in q are maintained in such a way that
		// their corresponding values are in decreasing order.
		//
		// Example:
		// nums = [1, 3, -1]
		// q = [1, 2]
		//
		// nums[q[0]] = 3
		// nums[q[1]] = -1
		//
		// So the FRONT of q always represents the maximum
		// element of the current window.


		// Remove indices that are no longer inside the window.
		//
		// Once left moves forward, any index smaller than left
		// belongs to an old window and cannot be used anymore.
		//
		// We only need to check q[0] because q contains indices
		// in increasing order.
		if len(q) > 0 && q[0] < left {
			q = q[1:]
		}


		// Remove smaller elements from the BACK of the deque.
		//
		// Why?
		//
		// Suppose the deque represents:
		//
		// [5, 3, 2]
		//
		// and the current number is 6.
		//
		// 6 is bigger than 2, 3, and 5.
		//
		// Therefore, none of those elements can ever become
		// the maximum of a future window while 6 is present.
		//
		// So we remove them from the BACK.
		//
		// We do this repeatedly because there may be multiple
		// smaller elements at the back.
		for len(q) > 0 && nums[q[len(q)-1]] < nums[right] {
			q = q[:len(q)-1]
		}


		// Add the current index to the BACK.
		//
		// At this point, every element before this index in q
		// has a value >= nums[right].
		//
		// Therefore q continues to maintain decreasing values.
		q = append(q, right)


		// Check whether we have a complete window of size k.
		//
		// Window size is:
		//
		// right - left + 1
		//
		// Once it becomes k, q[0] contains the index of the
		// largest element in this window.
		if right-left+1 == k {

			// q[0] is the index of the maximum element.
			// We store its value in the result.
			res = append(res, nums[q[0]])

			// Move the window one position to the right.
			//
			// Example:
			//
			// [1, 3, -1]
			//  ↑
			// left = 0
			//
			// After left++:
			//
			// [1, 3, -1]
			//     ↑
			//   left = 1
			//
			// The next iteration will add the next element
			// and create the next window.
			left += 1
		}

		// Expand the window by moving right forward.
		right += 1
	}

	return res
}