package arrayhashing

/*
	===========================================================
	SORT COLORS
	===========================================================

	Problem:

	Given an array containing only:

		0, 1, 2

	Sort the array in-place so that:

		0s come first
		1s come next
		2s come last

	Example:

	Input:
		[2, 0, 2, 1, 1, 0]

	Output:
		[0, 0, 1, 1, 2, 2]


	===========================================================
	APPROACH: DUTCH NATIONAL FLAG ALGORITHM
	===========================================================

	We have exactly THREE possible values:

		0
		1
		2

	Instead of using a sorting algorithm, we can partition
	the array into three regions.

	We maintain THREE pointers:

		left
		mid
		right


	-----------------------------------------------------------
	LEFT
	-----------------------------------------------------------

	Everything BEFORE left is guaranteed to be 0.

	Region:

		[0 ... left-1]

	contains only:

		0


	-----------------------------------------------------------
	MID
	-----------------------------------------------------------

	mid points to the element that we currently need to inspect.

	Everything between:

		left ... mid-1

	has already been processed and contains only:

		1

	But the element at nums[mid] has NOT necessarily been
	processed yet.


	-----------------------------------------------------------
	RIGHT
	-----------------------------------------------------------

	Everything AFTER right is guaranteed to be 2.

	Region:

		[right+1 ... n-1]

	contains only:

		2


	===========================================================
	THE ARRAY IS DIVIDED INTO FOUR REGIONS
	===========================================================

	At any point:

		[ 0s ][ 1s ][ unknown ][ 2s ]

		      left  mid       right


	More precisely:

		[0 ........ left-1]
				↓
			  all 0s

		[left .... mid-1]
				↓
			  all 1s

		[mid .... right]
				↓
			  UNKNOWN

		[right+1 .... n-1]
				↓
			  all 2s


	The algorithm's job is to eliminate the UNKNOWN region.

	When:

		mid > right

	the unknown region becomes empty.

	Therefore the array is sorted.


	===========================================================
	WHAT DO WE DO WITH EACH VALUE?
	===========================================================


	CASE 1:
	nums[mid] == 0


	We found a 0.

	0 belongs at the LEFT side.

	So we swap:

		nums[left]
		with
		nums[mid]

	After the swap, the 0 is correctly placed at left.

	Then:

		left++

	Why?

	Because the position we just filled is guaranteed to
	contain 0. We never need to examine it again.

	We also do:

		mid++

	Why?

	The element originally at mid was 0 and has now been
	placed into the correct 0 region.

	Therefore the unknown region moves forward.


	-----------------------------------------------------------
	CASE 2:
	nums[mid] == 1
	-----------------------------------------------------------

	1 already belongs in the middle.

	So we don't need to swap anything.

	We simply do:

		mid++

	Why?

	Because the current element has been completely processed.

	It belongs exactly where it is.


	-----------------------------------------------------------
	CASE 3:
	nums[mid] == 2
	-----------------------------------------------------------

	2 belongs at the RIGHT side.

	So we swap:

		nums[mid]
		with
		nums[right]

	Then:

		right--

	Why?

	We just placed a 2 at the right boundary, so that position
	is permanently sorted.

	BUT:

	IMPORTANT:

	DO NOT increment mid.


	===========================================================
	WHY DON'T WE INCREMENT MID AFTER A 2 SWAP?
	===========================================================

	This is the most important detail of the algorithm.

	Suppose:

		[1, 2, 0]

	Initially:

		left = 0
		mid  = 0
		right = 2

	We process:

		nums[mid] = 1

	So:

		mid++

	Now:

		left = 0
		mid = 1
		right = 2


	We have:

		[1, 2, 0]
		     ↑
		    mid


	nums[mid] == 2

	So we swap nums[mid] with nums[right]:

		[1, 0, 2]
		     ↑
		    mid


	Notice something VERY important:

	The 0 that came from right has NEVER been examined.

	If we now do:

		mid++

	we would move past the 0 without processing it.

	That would be wrong.

	So after swapping a 2:

		right--

	but:

		mid stays where it is.


	Now we inspect the new nums[mid]:

		0

	and correctly move it to the left.


	===========================================================
	DRY RUN
	===========================================================

	Input:

		[2, 0, 2, 1, 1, 0]


	Initially:

		left = 0
		mid  = 0
		right = 5


	-----------------------------------------------------------
	STEP 1
	-----------------------------------------------------------

	nums[mid] = 2

	Swap nums[mid] with nums[right]:

		[0, 0, 2, 1, 1, 2]
		 ↑           ↑
		mid         right

	Then:

		right--

	Now:

		left = 0
		mid = 0
		right = 4

	IMPORTANT:

	We DO NOT increment mid.


	-----------------------------------------------------------
	STEP 2
	-----------------------------------------------------------

	nums[mid] = 0

	Swap:

		nums[left] <-> nums[mid]

	They are actually the same position.

	Array:

		[0, 0, 2, 1, 1, 2]

	Then:

		left++
		mid++

	Now:

		left = 1
		mid = 1
		right = 4


	-----------------------------------------------------------
	STEP 3
	-----------------------------------------------------------

	nums[mid] = 0

	0 belongs on the left.

	Array remains:

		[0, 0, 2, 1, 1, 2]

	Then:

		left = 2
		mid = 2


	-----------------------------------------------------------
	STEP 4
	-----------------------------------------------------------

	nums[mid] = 2

	Swap nums[mid] with nums[right]:

		[0, 0, 1, 1, 2, 2]
		        ↑     ↑
		       mid  right

	Then:

		right--

	Now:

		left = 2
		mid = 2
		right = 3

	Again:

	DO NOT increment mid.


	-----------------------------------------------------------
	STEP 5
	-----------------------------------------------------------

	nums[mid] = 1

	1 is already in the correct middle region.

	So:

		mid++

	Now:

		left = 2
		mid = 3
		right = 3


	-----------------------------------------------------------
	STEP 6
	-----------------------------------------------------------

	nums[mid] = 1

	Again, 1 is already correct.

		mid++

	Now:

		left = 2
		mid = 4
		right = 3


	Now:

		mid > right

	So the unknown region is empty.

	Done.


	Final:

		[0, 0, 1, 1, 2, 2]


	===========================================================
	LOOP INVARIANT
	===========================================================

	At every iteration:

		[0 ... left-1]
			contains ONLY 0s

		[left ... mid-1]
			contains ONLY 1s

		[mid ... right]
			contains UNKNOWN elements

		[right+1 ... n-1]
			contains ONLY 2s

	Every iteration reduces the UNKNOWN region.

	When:

		mid > right

	there are no unknown elements remaining.


	===========================================================
	WHY THIS IS O(1) SPACE
	===========================================================

	We don't create another array.

	We only maintain three integer indexes:

		left
		mid
		right

	So:

		Space = O(1)


	===========================================================
	TIME COMPLEXITY
	===========================================================

	Every element is processed at most a constant number of
	times.

	Therefore:

		Time = O(n)


	===========================================================
	INTERVIEW EXPLANATION
	===========================================================

	"I use the Dutch National Flag algorithm with three
	pointers: left, mid, and right.

	The left pointer represents the boundary of the 0 region,
	the right pointer represents the boundary of the 2 region,
	and mid scans the unknown region.

	If nums[mid] is 0, I swap it with nums[left] and increment
	both left and mid.

	If it is 1, it is already in the correct region, so I only
	increment mid.

	If it is 2, I swap it with nums[right] and decrement right,
	but I don't increment mid because the element brought from
	the right has not been examined yet.

	When mid passes right, every element has been placed into
	the correct region.

	This gives O(n) time and O(1) space."
*/


func sortColors(nums []int) {

	// --------------------------------------------------------
	// left:
	//
	// Everything before left is guaranteed to be 0.
	//
	// Initially there are no confirmed 0s, so left starts at 0.
	// --------------------------------------------------------
	left := 0

	// --------------------------------------------------------
	// mid:
	//
	// Points to the element we currently need to inspect.
	//
	// Initially we haven't processed anything, so we start
	// from index 0.
	// --------------------------------------------------------
	mid := 0

	// --------------------------------------------------------
	// right:
	//
	// Everything after right is guaranteed to be 2.
	//
	// Initially there are no confirmed 2s, so right starts at
	// the last index.
	// --------------------------------------------------------
	right := len(nums) - 1

	// --------------------------------------------------------
	// The region [mid ... right] contains elements that we
	// haven't completely classified yet.
	//
	// Once mid passes right, there are no unknown elements left.
	// --------------------------------------------------------
	for mid <= right {

		// ----------------------------------------------------
		// CASE 1: We found a 0.
		//
		// 0 belongs in the left region.
		// ----------------------------------------------------
		if nums[mid] == 0 {

			// Move the 0 into the next available position
			// in the 0 region.
			nums[left], nums[mid] = nums[mid], nums[left]

			// left now points to the next position where
			// another 0 can be placed.
			left++

			// The element originally at mid was 0, so after
			// the swap it has been correctly handled.
			//
			// Therefore move to the next unknown element.
			mid++

		// ----------------------------------------------------
		// CASE 2: We found a 1.
		//
		// 1 belongs in the middle region.
		//
		// No swap is necessary.
		// ----------------------------------------------------
		} else if nums[mid] == 1 {

			// This element is already in the correct region.
			//
			// So simply move to the next unknown element.
			mid++

		// ----------------------------------------------------
		// CASE 3: We found a 2.
		//
		// 2 belongs in the right region.
		// ----------------------------------------------------
		} else {

			// Move the 2 to the right side.
			nums[right], nums[mid] = nums[mid], nums[right]

			// The position at right now definitely contains 2.
			//
			// Therefore shrink the unknown/right boundary.
			right--

			// IMPORTANT:
			//
			// DO NOT increment mid here.
			//
			// Why?
			//
			// The element that came from nums[right] into
			// nums[mid] has not been examined yet.
			//
			// It could be:
			//
			//     0
			//     1
			//     2
			//
			// So we must inspect nums[mid] again.
		}
	}
}