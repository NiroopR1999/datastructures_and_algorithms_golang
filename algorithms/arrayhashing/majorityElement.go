package arrayhashing

/*
	===========================================================
	MAJORITY ELEMENT
	===========================================================

	Problem:
	Given an integer array nums, find the element that appears
	more than n/2 times.

	Example:

	Input:
		[2, 2, 1, 1, 1, 2, 2]

	Output:
		2

	Because:

		2 appears 4 times
		1 appears 3 times

	And:

		n = 7
		n/2 = 3.5

	So 2 appears more than n/2 times.


	===========================================================
	IMPORTANT OBSERVATION
	===========================================================

	The majority element appears MORE THAN HALF of the array.

	That means:

		majority frequency > all other elements combined

	For example:

		[2, 2, 2, 1, 1]

	2 appears 3 times.

	Everything else combined appears only 2 times.

	Therefore, even if every occurrence of the majority element
	is "cancelled" against an occurrence of a different element,
	the majority element will still have at least one occurrence
	left.


	===========================================================
	APPROACH: BOYER-MOORE VOTING ALGORITHM
	===========================================================

	We maintain two variables:

		candidate
		count


	-----------------------------------------------------------
	candidate
	-----------------------------------------------------------

	Represents the element that is currently considered the
	possible majority element.


	-----------------------------------------------------------
	count
	-----------------------------------------------------------

	Represents the current "vote balance" for candidate.

	If we see candidate:

		count++

	If we see something different:

		count--

	When count becomes 0, it means:

		"The elements seen so far have completely cancelled
		each other out."

	So we are free to choose the next element as a new candidate.


	===========================================================
	WHY DOES CANCELLING WORK?
	===========================================================

	Consider:

		[2, 2, 1, 1, 1, 2, 2]

	There are more 2s than all the other elements combined.

	We can conceptually cancel different elements against
	occurrences of the candidate:

		2 vs 1  -> CANCEL
		2 vs 1  -> CANCEL

	Remaining:

		1, 2, 2

	There are still more 2s.

	The majority element cannot be completely cancelled because
	it occurs more than all non-majority elements combined.


	===========================================================
	ALGORITHM
	===========================================================

	For every number:

		1. If count == 0:
				make current number the candidate.

		2. If current number == candidate:
				increment count.

		3. Otherwise:
				decrement count.

	At the end:

		candidate = majority element


	===========================================================
	WHY CAN WE RETURN candidate WITHOUT VERIFYING IT?
	===========================================================

	This is IMPORTANT.

	This solution assumes the problem guarantees that a majority
	element exists.

	For example, LeetCode 169 guarantees:

		"The majority element always exists."

	Because the majority element appears more than n/2 times,
	it cannot be completely cancelled.

	Therefore, the final candidate MUST be the majority element.

	If the problem does NOT guarantee that a majority exists,
	we would need to perform a second pass to verify:

		count(candidate) > n/2


	===========================================================
	DRY RUN
	===========================================================

	nums = [2, 2, 1, 1, 1, 2, 2]

	Initially:

		candidate = 0
		count = 0


	-----------------------------------------------------------
	i = 0
	nums[i] = 2
	-----------------------------------------------------------

	count == 0

	So choose:

		candidate = 2

	Now:

		nums[i] == candidate

	So:

		count++

	Therefore:

		candidate = 2
		count = 1


	-----------------------------------------------------------
	i = 1
	nums[i] = 2
	-----------------------------------------------------------

	2 == candidate

	So:

		count++

	Now:

		candidate = 2
		count = 2


	-----------------------------------------------------------
	i = 2
	nums[i] = 1
	-----------------------------------------------------------

	1 != candidate (2)

	So:

		count--

	Now:

		candidate = 2
		count = 1


	-----------------------------------------------------------
	i = 3
	nums[i] = 1
	-----------------------------------------------------------

	1 != 2

	So:

		count--

	Now:

		candidate = 2
	count = 0


	-----------------------------------------------------------
	i = 4
	nums[i] = 1
	-----------------------------------------------------------

	count == 0

	So we need a new candidate:

		candidate = 1

	Then:

		nums[i] == candidate

	So:

		count++

	Now:

		candidate = 1
		count = 1


	-----------------------------------------------------------
	i = 5
	nums[i] = 2
	-----------------------------------------------------------

	2 != 1

	So:

		count--

	Now:

		candidate = 1
		count = 0


	-----------------------------------------------------------
	i = 6
	nums[i] = 2
	-----------------------------------------------------------

	count == 0

	So:

		candidate = 2

	Then:

	2 == candidate

	So:

		count++

	Final state:

		candidate = 2
		count = 1


	Therefore:

		return 2


	===========================================================
	WHY DID CANDIDATE CHANGE FROM 2 -> 1 -> 2?
	===========================================================

	This is the part that often confuses people.

	"How can we keep changing the candidate and still know
	the final answer?"

	Because candidate is NOT saying:

		"This is definitely the majority."

	It is saying:

		"This is the current survivor after cancelling
		different elements."

	Whenever count reaches 0, the previous candidate has been
	completely cancelled by different elements.

	We then start a new cancellation group.

	Because the real majority occurs more than n/2 times,
	it survives ALL possible cancellations.

	So eventually it becomes the final candidate.


	===========================================================
	CORE INVARIANT
	===========================================================

	The algorithm maintains:

		count =

			number of candidate votes
			-
			number of opposing votes

	Whenever we encounter:

		same as candidate
			-> +1 vote

		different from candidate
			-> -1 vote

	When:

		count == 0

	the current group has no surviving candidate.

	So we select the next element.


	===========================================================
	WHY NOT USE A HASH MAP?
	===========================================================

	A straightforward solution is:

		map[number]frequency

	Then count how many times each number occurs.

	That would require:

		Time:  O(n)
		Space: O(n)

	But Boyer-Moore uses:

		Time:  O(n)
		Space: O(1)

	So we get constant extra space.


	===========================================================
	COMPLEXITY
	===========================================================

	Time:

		O(n)

	Why?

	We iterate through nums exactly once.


	Space:

		O(1)

	Why?

	We only store:

		candidate
		count

	No hash map.
	No additional array.


	===========================================================
	INTERVIEW EXPLANATION
	===========================================================

	"I use the Boyer-Moore Voting Algorithm.

	I maintain a candidate and a count.

	If the current number equals the candidate, I increment
	the count because it supports the candidate.

	If it differs, I decrement the count because it cancels
	one occurrence of the candidate.

	When the count reaches zero, the current candidate has been
	completely cancelled, so I choose the next number as the
	candidate.

	Because the majority element occurs more than n/2 times,
	it cannot be completely cancelled by all the other elements.
	Therefore, the final candidate is guaranteed to be the
	majority element.

	This gives O(n) time and O(1) space."
*/

func majorityElement(nums []int) int {

	// --------------------------------------------------------
	// candidate:
	//
	// The element currently competing to be the majority.
	//
	// We don't know the majority at the beginning, so we start
	// with a default value of 0.
	// --------------------------------------------------------
	candidate := 0

	// --------------------------------------------------------
	// count:
	//
	// Represents the current vote balance for candidate.
	//
	// Positive count:
	//     candidate currently has more support.
	//
	// Zero count:
	//     candidate has been completely cancelled.
	// --------------------------------------------------------
	count := 0

	// --------------------------------------------------------
	// Process every element exactly once.
	// --------------------------------------------------------
	for i := range nums {

		// ----------------------------------------------------
		// If count becomes zero, there is no active candidate.
		//
		// The previous candidate has been completely cancelled
		// by different elements.
		//
		// Therefore, the current element becomes the new
		// candidate.
		// ----------------------------------------------------
		if count == 0 {
			candidate = nums[i]
		}

		// ----------------------------------------------------
		// Now compare the current number with our candidate.
		//
		// If they are the same:
		//
		//     The current element supports the candidate.
		//
		// So we increase its vote count.
		//
		// If they are different:
		//
		//     The current element opposes the candidate.
		//
		// So one candidate vote gets cancelled.
		// ----------------------------------------------------
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	// --------------------------------------------------------
	// Because the problem guarantees that a majority element
	// exists, the final surviving candidate MUST be the
	// majority element.
	// --------------------------------------------------------
	return candidate
}