package arrayhashing

func majorityElement2(nums []int) []int {

	/*
		========================================================
		APPROACH
		========================================================

		We need to find ALL elements that appear more than n/3
		times.

		Example:

			nums = [3, 2, 3]

			n = 3
			n/3 = 1

			3 appears 2 times.

			2 > 1
			=> answer = [3]


		--------------------------------------------------------
		WHY CAN THERE BE AT MOST 2 ANSWERS?
		--------------------------------------------------------

		Suppose there were 3 different elements that each
		appeared more than n/3 times.

			A > n/3
			B > n/3
			C > n/3

		Then:

			A + B + C > n

		But the array contains only n elements.

		Therefore, having 3 elements each appearing more than
		n/3 times is impossible.

		So there can be at most TWO valid answers.

		That is why we maintain:

			candidate1 + count1
			candidate2 + count2


		--------------------------------------------------------
		IMPORTANT IDEA: VOTE CANCELLATION
		--------------------------------------------------------

		We process the array one element at a time.

		We maintain two possible candidates.

		If the current number matches one of our candidates,
		we increase that candidate's count.

		If one candidate has count 0, we use that empty slot
		for the current number.

		But if:

			1. current number != candidate1
			2. current number != candidate2
			3. both candidate counts are greater than 0

		then we do:

			count1--
			count2--

		Why?

		We can think of this as removing THREE DIFFERENT
		elements from consideration:

			candidate1
			candidate2
			current number

		For example:

			1 1 1 2 2 3

		When we have:

			candidate1 = 1
			candidate2 = 2
			current     = 3

		we can cancel one occurrence of each:

			1 + 2 + 3

		After cancellation:

			count1--
			count2--

		Why is this safe?

		Because we're looking for an element occurring MORE
		THAN n/3 times.

		An element that truly occurs more than n/3 times cannot
		be completely eliminated by repeatedly grouping it with
		two different elements and cancelling those groups.

		Therefore, after processing the entire array, the only
		elements that COULD still satisfy the requirement are
		the two candidates we have left.


		--------------------------------------------------------
		VERY IMPORTANT:
		--------------------------------------------------------

		The counts maintained during Phase 1 are NOT the actual
		frequencies in the array.

		They are only "vote counts" after cancellation.

		Therefore, after finding candidate1 and candidate2,
		we MUST scan the array again and calculate their REAL
		frequencies.

		This gives us two phases:

			PHASE 1:
				Find possible candidates using cancellation.

			PHASE 2:
				Count the actual occurrences of those candidates
				and verify whether they occur more than n/3 times.


		--------------------------------------------------------
		TIME AND SPACE
		--------------------------------------------------------

		Phase 1:
			One pass through nums
			=> O(n)

		Phase 2:
			One more pass through nums
			=> O(n)

		Total:
			O(n)

		Extra space:
			Only candidate1, candidate2, count1, count2
			and result.

			=> O(1) auxiliary space
			(ignoring the output slice)
	*/


	// Result will contain the elements that actually appear
	// more than n/3 times.
	res := []int{}


	// --------------------------------------------------------
	// PHASE 1: FIND THE TWO POSSIBLE CANDIDATES
	// --------------------------------------------------------

	// candidate1 and candidate2 store the TWO numbers that
	// currently have the strongest possibility of being
	// majority elements.
	candidate1, candidate2 := 0, 0

	// These are NOT actual frequencies.
	//
	// They represent the remaining "vote strength" of each
	// candidate after cancellation.
	count1, count2 := 0, 0


	// Process every number in the array.
	for _, v := range nums {

		// ----------------------------------------------------
		// CASE 1:
		// Current number belongs to candidate1.
		// ----------------------------------------------------

		if v == candidate1 && count1 > 0 {

			// Since we found another occurrence of candidate1,
			// increase its vote count.
			count1++


		// ----------------------------------------------------
		// CASE 2:
		// Current number belongs to candidate2.
		// ----------------------------------------------------

		} else if v == candidate2 && count2 > 0 {

			// Another occurrence of candidate2.
			count2++


		// ----------------------------------------------------
		// CASE 3:
		// candidate1 has no votes.
		// ----------------------------------------------------

		} else if count1 == 0 {

			// candidate1 has been completely cancelled out.
			//
			// Therefore, this slot is available for a new
			// candidate.
			candidate1 = v
			count1 = 1


		// ----------------------------------------------------
		// CASE 4:
		// candidate2 has no votes.
		// ----------------------------------------------------

		} else if count2 == 0 {

			// candidate2 has been completely cancelled out.
			//
			// Use this empty slot for the current number.
			candidate2 = v
			count2 = 1


		// ----------------------------------------------------
		// CASE 5:
		// Current number is different from BOTH candidates.
		//
		// AND both candidates currently have votes.
		// ----------------------------------------------------

		} else {

			/*
				We now have THREE DIFFERENT values:

					candidate1
					candidate2
					v

				Instead of keeping all three, we cancel one
				occurrence of each.

				Therefore:

					count1--
					count2--

				The current `v` is also effectively cancelled.

				Example:

					candidate1 = 1
					count1     = 3

					candidate2 = 2
					count2     = 2

					v = 3

				We remove:

					1
					2
					3

				Remaining vote counts:

					count1 = 2
					count2 = 1

				This cancellation is the core idea behind the
				extended Boyer-Moore Voting Algorithm.
			*/

			count1--
			count2--
		}
	}


	// --------------------------------------------------------
	// PHASE 2: VERIFY THE CANDIDATES
	// --------------------------------------------------------

	/*
		At this point:

			candidate1
			candidate2

		are only POSSIBLE answers.

		They are NOT guaranteed to actually appear more than
		n/3 times.

		Example:

			nums = [1, 2, 3, 4, 5, 6]

		There is no majority element.

		But the voting phase can still leave some numbers in
		candidate1 and candidate2.

		Therefore we MUST count their REAL frequencies.
	*/

	count1 = 0
	count2 = 0

	for _, v := range nums {

		// Count the actual number of times candidate1 occurs.
		if v == candidate1 {
			count1++
		}

		// Count the actual number of times candidate2 occurs.
		if v == candidate2 {
			count2++
		}
	}


	// --------------------------------------------------------
	// CHECK CANDIDATE 1
	// --------------------------------------------------------

	/*
		The requirement is:

			frequency > n/3

		NOT:

			frequency >= n/3

		Example:

			n = 6

			n/3 = 2

		A number appearing exactly 2 times does NOT qualify.

		It must appear at least 3 times.
	*/

	if count1 > len(nums)/3 {
		res = append(res, candidate1)
	}


	// --------------------------------------------------------
	// CHECK CANDIDATE 2
	// --------------------------------------------------------

	if count2 > len(nums)/3 {
		res = append(res, candidate2)
	}


	// Return only the candidates that passed the REAL
	// frequency check.
	return res
}