package arrayhashing

func topKFrequent(nums []int, k int) []int {

	/*
		========================================================
		APPROACH: FREQUENCY MAP + BUCKET SORT
		========================================================

		Goal:

			Return the K elements that appear most frequently.

		Example:

			nums = [1, 1, 1, 2, 2, 3]
			k = 2

		Frequencies:

			1 → 3
			2 → 2
			3 → 1

		Top 2:

			[1, 2]


		========================================================
		STEP 1: COUNT FREQUENCIES
		========================================================

		First create:

			number → frequency

		For:

			[1, 1, 1, 2, 2, 3]

		we get:

			1 → 3
			2 → 2
			3 → 1


		========================================================
		STEP 2: WHY BUCKET SORT?
		========================================================

		The maximum possible frequency of any number is:

			len(nums)

		For example, if:

			nums = [5, 5, 5, 5]

		then the frequency of 5 is 4.

		So we create buckets where:

			bucket[f] = all numbers that occur f times

		For:

			nums = [1, 1, 1, 2, 2, 3]

		we get:

			bucket[1] → [3]
			bucket[2] → [2]
			bucket[3] → [1]

		Now the important part:

		We want the MOST frequent elements.

		So instead of sorting the frequencies, we simply
		traverse the buckets from:

			highest frequency
				 ↓
			to lowest frequency

		and keep taking elements until we have k elements.

		This avoids sorting and gives us O(n) average time.


		========================================================
		TIME COMPLEXITY
		========================================================

		Building frequency map:
			O(n)

		Building buckets:
			O(n)

		Traversing buckets:
			O(n)

		Total:
			O(n)

		Space:
			O(n)
	*/


	// --------------------------------------------------------
	// STEP 1: COUNT HOW MANY TIMES EACH NUMBER APPEARS
	// --------------------------------------------------------

	freq := make(map[int]int)

	for _, num := range nums {

		// If num hasn't been seen before, freq[num] is 0.
		// Incrementing it gives us its frequency.
		freq[num]++
	}


	// --------------------------------------------------------
	// STEP 2: CREATE BUCKETS
	// --------------------------------------------------------

	/*
		bucket[i] contains all numbers that appear exactly
		i times.

		Why len(nums)+1?

		Because a number can appear at most len(nums) times.

		Example:

			nums = [5,5,5]

			frequency of 5 = 3

		So we need:

			bucket[3]
	*/
	bucket := make([][]int, len(nums)+1)

	for num, count := range freq {

		// Put the number into the bucket corresponding
		// to its frequency.
		//
		// Example:
		//
		// num = 1
		// count = 3
		//
		// bucket[3] = [1]
		bucket[count] = append(bucket[count], num)
	}


	// --------------------------------------------------------
	// STEP 3: COLLECT THE TOP K ELEMENTS
	// --------------------------------------------------------

	// We know there can be at most k results.
	//
	// len = 0
	// capacity = k
	//
	// This means the slice starts empty but has room
	// for k elements.
	res := make([]int, 0, k)


	/*
		Start from the HIGHEST frequency.

		Why?

		Because we want the most frequent elements first.

		Example:

			bucket[5] → [10]
			bucket[4] → []
			bucket[3] → [7, 8]
			bucket[2] → [4, 5]
			bucket[1] → [2]

		We visit:

			5 → 4 → 3 → 2 → 1

		So we encounter:

			10 → 7 → 8 → ...

		These are ordered from highest frequency to lowest.
	*/
	for count := len(bucket) - 1; count >= 0; count-- {

		// There may be multiple numbers with the same
		// frequency.
		//
		// Example:
		//
		// bucket[3] = [1, 7, 10]
		//
		// All three occur exactly 3 times.
		for _, num := range bucket[count] {

			// Add the number to our answer.
			res = append(res, num)

			// Once we have k elements, we're done.
			if len(res) == k {
				return res
			}
		}
	}


	// Normally we return from inside the loop once we have
	// k elements.
	//
	// This return is here as a safe fallback.
	return res
}