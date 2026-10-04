package arrayhashing

/*
	===========================================================
	LONGEST CONSECUTIVE SEQUENCE
	===========================================================

	Problem:

	Given an unsorted integer array nums, find the length of
	the longest sequence of consecutive integers.

	The numbers do NOT need to be next to each other in the
	original array.

	Example:

		Input:
			[100, 4, 200, 1, 3, 2]

		The longest consecutive sequence is:

			1, 2, 3, 4

		So the answer is:

			4


	===========================================================
	IMPORTANT OBSERVATION
	===========================================================

	The array is UNSORTED.

	Example:

		[100, 4, 200, 1, 3, 2]

	The consecutive sequence:

		1, 2, 3, 4

	is scattered throughout the array.

	We therefore cannot simply look at neighboring indexes.

	Instead, we care about whether a NUMBER exists.

	For example:

		Does 2 exist?
		Does 3 exist?
		Does 4 exist?
		Does 5 exist?

	A Hash Set gives us approximately O(1) lookup.


	===========================================================
	APPROACH
	===========================================================

	We use a Hash Set containing every number.

	Example:

		nums = [100, 4, 200, 1, 3, 2]

	Set:

		{
			100,
			4,
			200,
			1,
			3,
			2
		}


	Now for every number v, we ask:

		"Is v the START of a consecutive sequence?"


	-----------------------------------------------------------
	HOW DO WE KNOW IF v IS THE START?
	-----------------------------------------------------------

	A number v is the beginning of a sequence if:

		v - 1

	does NOT exist.

	Example:

		1, 2, 3, 4

	For 1:

		0 does not exist

	So 1 is the START.


	For 2:

		1 exists

	So 2 is NOT the start.


	For 3:

		2 exists

	So 3 is NOT the start.


	For 4:

		3 exists

	So 4 is NOT the start.


	Therefore, instead of starting a sequence from every
	number, we only start from the first number.


	===========================================================
	WHY IS THIS IMPORTANT?
	===========================================================

	Consider:

		[1, 2, 3, 4, 5]

	If we started counting from every number:

		Start at 1:
			1 -> 2 -> 3 -> 4 -> 5

		Start at 2:
			2 -> 3 -> 4 -> 5

		Start at 3:
			3 -> 4 -> 5

		Start at 4:
			4 -> 5

	We repeatedly traverse the same sequence.

	Instead, we only start from 1 because:

		0 does not exist

	For 2:

		1 exists

	so we skip it.

	For 3:

		2 exists

	so we skip it.

	And so on.

	This is the key optimization that gives us O(n)
	average time.


	===========================================================
	STEP-BY-STEP ALGORITHM
	===========================================================

	STEP 1:

	Create a Hash Set containing every number.


	STEP 2:

	Loop through nums again.


	STEP 3:

	For each number v, check:

		if v-1 does NOT exist

	Then v is the beginning of a consecutive sequence.


	STEP 4:

	Starting from v, repeatedly check:

		v + 1
		v + 2
		v + 3
		...

	until the next number doesn't exist.


	STEP 5:

	Keep track of the longest sequence found.


	===========================================================
	DRY RUN
	===========================================================

	Input:

		[100, 4, 200, 1, 3, 2]


	After building the set:

		{
			100,
			4,
			200,
			1,
			3,
			2
		}


	-----------------------------------------------------------
	v = 100
	-----------------------------------------------------------

	Check:

		100 - 1 = 99

	Does 99 exist?

		NO

	So 100 is a sequence start.

	Start counting:

		100 exists
		101 exists? NO

	Sequence:

		100

	count = 1

	longest = 1


	-----------------------------------------------------------
	v = 4
	-----------------------------------------------------------

	Check:

		4 - 1 = 3

	Does 3 exist?

		YES

	Therefore 4 is NOT a sequence start.

	Skip it.


	-----------------------------------------------------------
	v = 200
	-----------------------------------------------------------

	Check:

		199 exists?

		NO

	So 200 is a sequence start.

	Check:

		200 exists
		201 exists? NO

	count = 1

	longest remains 1.


	-----------------------------------------------------------
	v = 1
	-----------------------------------------------------------

	Check:

		0 exists?

		NO

	So 1 is a sequence start.

	Start counting:

		1 exists
		2 exists
		3 exists
		4 exists
		5 exists? NO

	So:

		sequence = [1, 2, 3, 4]

	count = 4

	Update:

		longest = 4


	-----------------------------------------------------------
	v = 3
	-----------------------------------------------------------

	Check:

		2 exists?

		YES

	So 3 is not the start.

	Skip.


	-----------------------------------------------------------
	v = 2
	-----------------------------------------------------------

	Check:

		1 exists?

		YES

	So 2 is not the start.

	Skip.


	Final answer:

		4


	===========================================================
	WHY DOES THE INNER LOOP NOT MAKE THE ALGORITHM O(n²)?
	===========================================================

	This is an extremely common interview question.

	At first glance we have:

		for every number
			for consecutive numbers

	which LOOKS like O(n²).

	But the important thing is:

	We only run the inner loop when v is the START of a
	sequence.

	For:

		[1, 2, 3, 4, 5]

	only 1 starts a sequence.

	So we do:

		1 -> 2 -> 3 -> 4 -> 5

	We DON'T repeat the traversal from 2, 3, 4, or 5.


	More generally, every number belongs to a sequence and
	is traversed by the inner loop only when we start from
	the beginning of that sequence.

	So across all sequences, the total number of successful
	inner-loop iterations is O(n).

	The Hash Set lookups are O(1) on average.

	Therefore:

		O(n) average time.


	===========================================================
	WHY USE A MAP/SET?
	===========================================================

	We need to answer questions like:

		"Does 73 exist?"

	Searching through a slice would take:

		O(n)

	Doing that repeatedly would become expensive.

	A hash set gives approximately:

		O(1)

	average lookup.

	In Go, we don't have a built-in Set type.

	So we commonly use:

		map[int]bool

	or:

		map[int]struct{}

	The second form is slightly more idiomatic when we only
	care about existence.


	===========================================================
	WHY map[int]struct{} IS OFTEN PREFERRED
	===========================================================

	Your code uses:

		map[int]bool

	This works perfectly.

	But if we don't care about a boolean value and only care
	whether the key exists, we can use:

		map[int]struct{}

	Example:

		set := make(map[int]struct{})

		set[10] = struct{}{}

	To check:

		_, exists := set[10]

	However, your:

		map[int]bool

	is easier to read for beginners:

		set[v] == true

	and:

		!set[v]

	are very clear.

	So your implementation is completely valid.


	===========================================================
	IMPORTANT GO DETAIL
	===========================================================

	Your code uses:

		if !set[v-1]

	This works because looking up a missing key in:

		map[int]bool

	returns the zero value of bool:

		false

	Therefore:

		if !set[v-1]

	means:

		"If v-1 does not exist in the set."


	Example:

		set = {1, 2, 3}

		set[2] -> true

		set[99] -> false

	There is no need to explicitly check whether the key
	exists.


	===========================================================
	COMPLEXITY
	===========================================================

	Building the set:

		O(n)


	Scanning nums:

		O(n)


	Inner consecutive traversal:

		O(n) TOTAL across all sequences

	NOT O(n²), because we only start traversing from sequence
	starts.


	Therefore:

		Time = O(n) average


	Space:

		The Hash Set stores every number:

		O(n)


	Final:

		Time:  O(n) average
		Space: O(n)


	===========================================================
	INTERVIEW EXPLANATION
	===========================================================

	"I put every number into a hash set so I can check whether
	a number exists in O(1) average time.

	Then for every number, I only start building a consecutive
	sequence if its previous number, v-1, doesn't exist.

	This identifies the beginning of a sequence.

	From that starting number, I keep checking v+1, v+2, and
	so on until the sequence ends, and I track the maximum
	length.

	The reason this is O(n) on average is that we only traverse
	a sequence from its beginning. We don't repeatedly traverse
	the same sequence starting from every element.

	Time complexity is O(n) average and space complexity is
	O(n)."
*/


func longestConsecutive(nums []int) int {

	// --------------------------------------------------------
	// longest stores the length of the longest consecutive
	// sequence found so far.
	//
	// Initially we haven't found any sequence, so:
	//
	//     longest = 0
	// --------------------------------------------------------
	longest := 0

	// --------------------------------------------------------
	// Create a Hash Set containing every number.
	//
	// map[int]bool:
	//
	//     key   -> number
	//     value -> whether the number exists
	//
	// We mainly care about the keys.
	//
	// Why use a set?
	//
	// Because we need to quickly answer:
	//
	//     "Does number X exist?"
	//
	// Hash map lookup is O(1) average.
	// --------------------------------------------------------
	set := make(map[int]bool)

	// --------------------------------------------------------
	// STEP 1:
	//
	// Put every number into the set.
	//
	// Example:
	//
	//     nums = [100, 4, 200, 1, 3, 2]
	//
	// becomes conceptually:
	//
	//     set = {
	//         100: true,
	//         4:   true,
	//         200: true,
	//         1:   true,
	//         3:   true,
	//         2:   true,
	//     }
	// --------------------------------------------------------
	for _, v := range nums {
		set[v] = true
	}

	// --------------------------------------------------------
	// STEP 2:
	//
	// Examine every number as a possible sequence START.
	// --------------------------------------------------------
	for _, v := range nums {

		// ----------------------------------------------------
		// A number is the START of a consecutive sequence only
		// if the previous number doesn't exist.
		//
		// Example:
		//
		//     v = 3
		//
		// If 2 exists:
		//
		//     2 -> 3
		//
		// then 3 clearly isn't the beginning.
		//
		// So we skip it.
		//
		// If 2 doesn't exist:
		//
		//     3
		//
		// then 3 could be the beginning.
		// ----------------------------------------------------
		if !set[v-1] {

			// ------------------------------------------------
			// v is the beginning of a sequence.
			//
			// v itself is already present, so the initial
			// sequence length is 1.
			// ------------------------------------------------
			count := 1

			// ------------------------------------------------
			// Now look for:
			//
			//     v + 1
			//     v + 2
			//     v + 3
			//     ...
			//
			// As long as each number exists, the consecutive
			// sequence continues.
			// ------------------------------------------------
			for set[v+count] {

				// We found the next consecutive number.
				//
				// Example:
				//
				//     v = 1
				//     count = 1
				//
				// Check:
				//
				//     set[1 + 1]
				//     set[2]
				//
				// If 2 exists, increase the sequence length.
				count++
			}

			// ------------------------------------------------
			// We have reached the end of the current sequence.
			//
			// Compare its length with the longest sequence
			// we've seen so far.
			// ------------------------------------------------
			if count > longest {
				longest = count
			}
		}
	}

	// --------------------------------------------------------
	// Return the maximum consecutive sequence length.
	// --------------------------------------------------------
	return longest
}