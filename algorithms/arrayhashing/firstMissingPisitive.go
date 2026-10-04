package arrayhashing

func firstMissingPositive(nums []int) int {

	/*
		========================================================
		APPROACH: IN-PLACE CYCLIC PLACEMENT
		========================================================

		We need to find the SMALLEST POSITIVE integer that does
		not exist in nums.

		Example:

			nums = [3, 4, -1, 1]

		Positive numbers are:

			1, 3, 4

		The smallest missing positive is:

			2


		========================================================
		IMPORTANT OBSERVATION
		========================================================

		If nums has length n, the answer MUST be between:

			1 and n+1

		Why?

		Consider:

			nums = [1, 2, 3]

		Every positive number from 1 to n exists.

		So the answer is:

			n + 1
			= 4


		Therefore, for an array of length n, we only care
		about numbers:

			1, 2, 3, ..., n


		Numbers such as:

			-5
			0
			100
			1000

		cannot be the answer if n is small enough, so we can
		ignore them.


		========================================================
		CORE IDEA
		========================================================

		We want every useful number to go to a specific index.

		For a number `x`:

			x should go to index x-1

		Why?

			1 → index 0
			2 → index 1
			3 → index 2
			4 → index 3
			...
			x → index x-1


		For example:

			nums = [3, 4, -1, 1]

		We want:

			index:    0   1   2   3
			          ↓   ↓   ↓   ↓

			value:    1   2   3   4

		After placing all useful numbers correctly, we can
		simply scan the array.

		The first index where:

			nums[i] != i+1

		gives us the answer:

			i+1


		========================================================
		WHY DO WE NEED THE INNER `for` LOOP?
		========================================================

		A number may not be in its correct position initially.

		Example:

			nums = [3, 4, -1, 1]

		At index 0 we have:

			nums[0] = 3

		3 belongs at:

			index = 3-1
			      = 2

		So we swap 3 into index 2.

		Now index 0 contains another number.

		We need to check that new number too.

		That's why we keep swapping until the current position
		contains a number that:

			1. is not useful
			OR
			2. is already in the correct position
			OR
			3. has a duplicate already occupying its target


		========================================================
		WHEN IS A NUMBER "USEFUL"?
		========================================================

		We only care about:

			1 <= nums[i] <= n

		Anything outside this range cannot affect the answer.

		So the condition:

			nums[i] > 0 && nums[i] <= l

		means:

			"Is this number one of the values that could
			 determine the answer?"


		========================================================
		WHY `nums[nums[i]-1] != nums[i]`?
		========================================================

		This is VERY important.

		Suppose:

			nums = [2, 2, 1]

		We are at index 0:

			nums[0] = 2

		2 belongs at:

			index 2-1 = 1

		So we would like to swap.

		But index 1 ALREADY contains 2:

			[2, 2, 1]
			    ↑
			    2 already exists at its target

		If we keep swapping:

			nums[1], nums[0] = nums[0], nums[1]

		we get:

			[2, 2, 1]

		Nothing changes.

		We would keep doing this forever.

		Therefore:

			nums[nums[i]-1] != nums[i]

		means:

			"Only swap if the destination does NOT already
			 contain the same value."

		This prevents an infinite loop caused by duplicates.


		========================================================
		PHASE 1: PLACE NUMBERS
		========================================================
	*/

	l := len(nums)

	for i := 0; i < l; i++ {

		/*
			Keep placing nums[i] into its correct position
			until it cannot/should not be moved anymore.

			For a value x:

				target index = x - 1
		*/
		for nums[i] > 0 &&
			nums[i] <= l &&
			nums[nums[i]-1] != nums[i] {

			/*
				Swap nums[i] with the element at its target
				position.

				Example:

					nums = [3, 4, -1, 1]

					i = 0
					nums[i] = 3

					target = 3 - 1
					       = 2

					Swap:

						index 0 ↔ index 2

					Result:

						[-1, 4, 3, 1]

					Now 3 is exactly where it belongs:

						index 2 → value 3
				*/

			nums[nums[i]-1], nums[i] =
				nums[i], nums[nums[i]-1]
		}
	}


	/*
		========================================================
		PHASE 2: FIND THE FIRST INCORRECT POSITION
		========================================================

		After Phase 1, every useful number that could be placed
		in its correct position has been placed.

		Therefore we expect:

			index 0 → 1
			index 1 → 2
			index 2 → 3
			index 3 → 4
			...

		In general:

			nums[i] should equal i+1


		Example:

			nums = [1, 2, -1, 4]

			index:    0   1   2   3
			expected: 1   2   3   4
			actual:   1   2  -1   4

		At index 2:

			nums[2] != 3

		So:

			3

		is the first missing positive.
	*/

	for i := range nums {

		if nums[i] != i+1 {

			// Since index i should contain i+1,
			// and it doesn't, i+1 is missing.
			return i + 1
		}
	}


	/*
		========================================================
		WHY RETURN `l + 1`?
		========================================================

		If we reach here, EVERY number from:

			1 ... l

		exists.

		Therefore the first missing positive must be:

			l + 1

		Example:

			nums = [1, 2, 3]

			l = 3

			1 exists
			2 exists
			3 exists

		So:

			4

		is the first missing positive.
	*/

	return l + 1
}