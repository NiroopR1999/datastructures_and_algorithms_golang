package arrayhashing

func subarraySum(nums []int, k int) int {

	/*
		========================================================
		APPROACH: PREFIX SUM + FREQUENCY MAP
		========================================================

		We need to count how many CONTIGUOUS subarrays have
		a sum equal to k.

		Instead of checking every possible subarray, we use:

			1. Prefix Sum
			2. Hash Map

		The key idea is:

			previousPrefix = currentPrefix - k


		========================================================
		WHY DOES `currentPrefix - k` WORK?
		========================================================

		Suppose the array looks like this:

			[ elements before subarray ][ desired subarray ]
			          sum = P                  sum = k

		The prefix sum at the end of the desired subarray is:

			currentPrefix = P + k

		Therefore:

			currentPrefix = previousPrefix + k

		Rearranging:

			previousPrefix = currentPrefix - k


		For example:

			nums = [2, 4, 1, 3]
			k = 4

		Prefix sums:

			2 → 6 → 7 → 10

		When we are at the final element:

			currentPrefix = 10

		We want a subarray with sum 4.

		So we ask:

			"What prefix sum should have existed BEFORE
			 the desired subarray started?"

			10 - 4 = 6

		And prefix sum 6 exists:

			[2, 4] [1, 3]
			    ↑      ↑
			  sum 6   sum 4

		So:

			10 - 6 = 4

		Therefore [1, 3] is a valid subarray.


		THIS IS WHY:

			freq[prefixSum-k]

		gives us the answer.

		If `prefixSum-k` has appeared before, then the elements
		between that previous prefix and the current position
		must have sum exactly k.


		========================================================
		WHY DO WE STORE A FREQUENCY?
		========================================================

		The same prefix sum can appear multiple times.

		For example:

			nums = [1, -1, 1, -1]
			k = 0

		Prefix sums:

			1 → 0 → 1 → 0

		When the current prefix is 0:

			currentPrefix - k
			= 0 - 0
			= 0

		There was a previous prefix sum of 0.

		That previous occurrence gives us one valid subarray.

		If the required prefix had appeared 3 times before,
		then there would be 3 different starting positions,
		and therefore 3 valid subarrays ending at the current
		position.

		Therefore we store:

			prefixSum → number of times we have seen it

		not simply:

			prefixSum → true/false


		========================================================
		WHY `{0: 1}`?
		========================================================

		We start with:

			freq := map[int]int{0: 1}

		This means:

			"Before processing any element, prefix sum 0 has
			 already occurred once."

		This is necessary for subarrays that start at index 0.

		Example:

			nums = [3]
			k = 3

		After processing 3:

			prefixSum = 3

		Required previous prefix:

			3 - 3 = 0

		We have:

			freq[0] = 1

		So we find one valid subarray:

			[3]


		Without `freq[0] = 1`, we would miss every valid
		subarray that starts from index 0.


		========================================================
		ALGORITHM
		========================================================

		For every number:

			1. Add it to prefixSum.

			2. Calculate:

				   prefixSum - k

			   This tells us what prefix sum must have existed
			   before the desired subarray started.

			3. Add the frequency of that prefix to count.

			4. Store the current prefix sum in the map.


		Why do we check the map BEFORE storing the current
		prefix?

		Because we need a prefix sum from an EARLIER position.

		The current prefix represents the END of our subarray.
		*/


	// freq stores:
	//
	//     prefix sum → how many times we've seen it
	//
	// Start with prefix sum 0 occurring once.
	//
	// This handles subarrays that begin at index 0.
	freq := map[int]int{0: 1}

	// Running prefix sum.
	//
	// We don't need to create a separate prefix-sum array.
	// We can calculate it while traversing nums.
	prefixSum := 0

	// Total number of subarrays whose sum is exactly k.
	count := 0


	for _, v := range nums {

		// Add the current value to the running prefix sum.
		//
		// Example:
		//
		// nums = [2, 4, 1]
		//
		// prefixSum:
		//
		// after 2 → 2
		// after 4 → 6
		// after 1 → 7
		prefixSum += v


		/*
			================================================
			THE MOST IMPORTANT LINE
			================================================

				count += freq[prefixSum-k]

			Why?

			We know:

				currentPrefix
					=
				previousPrefix + subarraySum


			We want:

				subarraySum = k


			So:

				currentPrefix
					=
				previousPrefix + k


			Therefore:

				previousPrefix
					=
				currentPrefix - k


			So we look for:

				freq[prefixSum-k]


			If it exists, every occurrence represents a
			different starting position for a subarray whose
			sum is exactly k.

			For example:

				currentPrefix = 10
				k = 4

			We need:

				previousPrefix = 10 - 4
				                = 6

			If prefix sum 6 occurred before, then:

				10 - 6 = 4

			Therefore the elements between those two prefix
			positions have sum exactly 4.

			If prefix sum 6 occurred 3 times, there are
			3 different valid subarrays ending here.

			That is why we add the FREQUENCY, not just 1.
		*/
		count += freq[prefixSum-k]


		/*
			Now record the current prefix sum.

			For example:

				prefixSum = 6

			Then:

				freq[6]++

			means:

				"We have now seen prefix sum 6 one more time."

			We do this AFTER checking `prefixSum-k` because
			the current prefix should only become available
			as a PREVIOUS prefix for future elements.
		*/
		freq[prefixSum]++
	}


	// Return the total number of valid subarrays.
	return count
}