package arrayhashing

/*
	===========================================================
	SORT AN ARRAY - MERGE SORT
	===========================================================

	Problem:
	Given an integer array, sort it in ascending order.

	Example:

	Input:
		[5, 2, 3, 1]

	Output:
		[1, 2, 3, 5]


	===========================================================
	APPROACH: MERGE SORT
	===========================================================

	Merge Sort uses the:

		Divide -> Conquer -> Merge

	strategy.


	-----------------------------------------------------------
	STEP 1: DIVIDE
	-----------------------------------------------------------

	We repeatedly divide the array into two halves.

	Example:

		[5, 2, 3, 1]

			↓ divide

		[5, 2]       [3, 1]

			↓             ↓

		[5] [2]       [3] [1]

	We continue dividing until every subarray contains
	only one element.


	-----------------------------------------------------------
	WHY STOP AT ONE ELEMENT?
	-----------------------------------------------------------

	A single element is already sorted.

	For example:

		[5]

	There is nothing to compare because there is only one
	element.

	So our recursive base case is:

		len(nums) <= 1

	Once we reach this point, we simply return the array.


	-----------------------------------------------------------
	STEP 2: CONQUER
	-----------------------------------------------------------

	Once the array has been divided into single elements,
	the recursive calls return.

	Then we start combining the smaller sorted arrays.

	Example:

		[5] + [2]

	becomes:

		[2, 5]

	Then:

		[3] + [1]

	becomes:

		[1, 3]

	Finally:

		[2, 5] + [1, 3]

	becomes:

		[1, 2, 3, 5]


	-----------------------------------------------------------
	STEP 3: MERGE
	-----------------------------------------------------------

	The merge function receives TWO SORTED arrays.

	Example:

		left  = [2, 5]
		right = [1, 3]

	We compare the first unused element from each array:

		2 vs 1

	1 is smaller -> put 1 into result.

	Then:

		2 vs 3

	2 is smaller -> put 2 into result.

	Then:

		5 vs 3

	3 is smaller -> put 3 into result.

	Right side is exhausted.

	So append the remaining 5.

	Result:

		[1, 2, 3, 5]


	===========================================================
	IMPORTANT MERGE INSIGHT
	===========================================================

	Why can we simply compare left[i] and right[j]?

	Because BOTH arrays are already sorted.

	For example:

		left  = [2, 5, 8]
		right = [1, 3, 7]

	We know:

		left[0] = 2

	is the smallest remaining element on the LEFT.

	And:

		right[0] = 1

	is the smallest remaining element on the RIGHT.

	So the smallest element overall MUST be either:

		left[i]
		OR
		right[j]

	Therefore, comparing only these two elements is enough.


	===========================================================
	WHY DO WE NEED TWO POINTERS?
	===========================================================

	We maintain:

		i -> current position in left
		j -> current position in right

	Initially:

		i = 0
		j = 0

	Whenever we choose left[i]:

		i++

	Whenever we choose right[j]:

		j++

	This allows us to process both arrays exactly once.


	===========================================================
	DRY RUN
	===========================================================

	Input:

		[5, 2, 3, 1]


	-----------------------------------------------------------
	DIVIDE
	-----------------------------------------------------------

	[5, 2, 3, 1]

	            ↓

	[5, 2]          [3, 1]

	   ↓                ↓

	[5] [2]          [3] [1]


	-----------------------------------------------------------
	MERGE [5] AND [2]
	-----------------------------------------------------------

	Compare:

		5 vs 2

	2 is smaller.

	Result:

		[2]

	Left still has:

		[5]

	Right is exhausted.

	Append remaining 5:

		[2, 5]


	-----------------------------------------------------------
	MERGE [3] AND [1]
	-----------------------------------------------------------

	Compare:

		3 vs 1

	1 is smaller.

	Result:

		[1]

	Append remaining 3:

		[1, 3]


	-----------------------------------------------------------
	FINAL MERGE
	-----------------------------------------------------------

	left:

		[2, 5]

	right:

		[1, 3]


	i = 0
	j = 0


	Compare:

		2 vs 1

	Choose 1.

	result:

		[1]

	j++

	j = 1


	Compare:

		2 vs 3

	Choose 2.

	result:

		[1, 2]

	i++

	i = 1


	Compare:

		5 vs 3

	Choose 3.

	result:

		[1, 2, 3]

	j++

	j = 2


	Now:

		j == len(right)

	Right side is exhausted.

	Left still contains:

		[5]

	Append it.

	Result:

		[1, 2, 3, 5]


	===========================================================
	WHY DO WE APPEND THE REMAINING ELEMENTS?
	===========================================================

	Consider:

		left  = [2, 5]
		right = [1]

	After comparing:

		2 vs 1

	we choose 1.

	Now right is exhausted.

	But left still contains:

		[2, 5]

	Those elements are ALREADY SORTED.

	Therefore, we don't need to compare them again.

	We can directly append them.

	This is why we have:

		res = append(res, left[i:]...)
		res = append(res, right[j:]...)

	One of these will usually be empty, but handling both
	makes the merge function complete and safe.


	===========================================================
	IMPORTANT GO DETAIL: make() + append()
	===========================================================

	WRONG:

		res := make([]int, len(left)+len(right))

	Why?

	Because this creates a slice that ALREADY contains
	len(left)+len(right) zero values.

	Example:

		make([]int, 5)

	gives:

		[0, 0, 0, 0, 0]

	If we then do:

		append(res, 1)

	we get:

		[0, 0, 0, 0, 0, 1]

	That's NOT what we want.

	Instead use:

		res := make([]int, 0, len(left)+len(right))

	This means:

		length = 0
		capacity = len(left) + len(right)

	So initially:

		[]

	and we have enough allocated capacity to hold the final
	result.

	Then append works correctly.


	===========================================================
	TIME COMPLEXITY
	===========================================================

	At every level, we process all N elements during merging.

	Number of levels:

		O(log N)

	Work per level:

		O(N)

	Therefore:

		O(N log N)


	===========================================================
	SPACE COMPLEXITY
	===========================================================

	Merge creates a new result slice.

	Therefore the merge operations require:

		O(N)

	additional space.

	The recursive calls also use:

		O(log N)

	stack space.

	Overall:

		O(N)

	because O(N) dominates O(log N).


	===========================================================
	IMPORTANT PROPERTY
	===========================================================

	Merge Sort has guaranteed:

		O(N log N)

	time complexity.

	Unlike some implementations of Quick Sort, it doesn't
	degrade to O(N²) because of a bad pivot.

	Also, Merge Sort is naturally suited to sorting linked
	lists and external/large datasets because of its
	divide-and-merge structure.


	===========================================================
	INTERVIEW EXPLANATION
	===========================================================

	"I use Merge Sort.

	I recursively divide the array into two halves until each
	half contains at most one element, because a single element
	is already sorted.

	Then I merge the two sorted halves.

	During merging, I maintain two pointers, one for each half.
	I compare the current elements and append the smaller one
	to the result.

	Once one half is exhausted, I append the remaining elements
	from the other half because they are already sorted.

	This gives O(N log N) time and O(N) additional space."
*/


// ===========================================================
// SORT ARRAY
// ===========================================================

func sortArray(nums []int) []int {

	// ---------------------------------------------------------
	// BASE CASE
	// ---------------------------------------------------------
	//
	// If the array has 0 or 1 element, it is already sorted.
	//
	// Examples:
	//
	//     []
	//     [5]
	//
	// There is nothing to sort.
	//
	// WHY <= 1 instead of == 1?
	//
	// Because an empty slice is also already sorted.
	//
	if len(nums) <= 1 {
		return nums
	}

	// ---------------------------------------------------------
	// FIND THE MIDDLE
	// ---------------------------------------------------------
	//
	// Example:
	//
	//     nums = [5, 2, 3, 1]
	//     len   = 4
	//
	//     len(nums) / 2 = 2
	//
	// So:
	//
	//     left  = nums[:2] = [5, 2]
	//     right = nums[2:] = [3, 1]
	//
	// We divide the array into two approximately equal halves.
	//
	mid := len(nums) / 2

	// ---------------------------------------------------------
	// SORT THE LEFT HALF
	// ---------------------------------------------------------
	//
	// We recursively call sortArray on the left half.
	//
	// The recursive call keeps dividing until it reaches the
	// base case.
	//
	// Eventually:
	//
	//     [5, 2]
	//
	// becomes:
	//
	//     [5]
	//     [2]
	//
	// and then merge produces:
	//
	//     [2, 5]
	//
	left := sortArray(nums[:mid])

	// ---------------------------------------------------------
	// SORT THE RIGHT HALF
	// ---------------------------------------------------------
	//
	// Same idea as the left side.
	//
	// Example:
	//
	//     [3, 1]
	//
	// becomes:
	//
	//     [3]
	//     [1]
	//
	// and merge produces:
	//
	//     [1, 3]
	//
	right := sortArray(nums[mid:])

	// ---------------------------------------------------------
	// MERGE THE TWO SORTED HALVES
	// ---------------------------------------------------------
	//
	// At this point we KNOW:
	//
	//     left  is sorted
	//     right is sorted
	//
	// So merge() can efficiently combine them into one sorted
	// array.
	//
	return merge(left, right)
}


// ===========================================================
// MERGE TWO SORTED ARRAYS
// ===========================================================

func merge(left, right []int) []int {

	// ---------------------------------------------------------
	// CREATE AN EMPTY RESULT WITH ENOUGH CAPACITY
	// ---------------------------------------------------------
	//
	// IMPORTANT:
	//
	//     make([]int, 0, len(left)+len(right))
	//
	// means:
	//
	//     length   = 0
	//     capacity = len(left) + len(right)
	//
	// We want length 0 because we are going to build the
	// result using append().
	//
	// If we instead wrote:
	//
	//     make([]int, len(left)+len(right))
	//
	// we'd already have that many zero-valued elements.
	//
	res := make([]int, 0, len(left)+len(right))

	// ---------------------------------------------------------
	// i tracks the current element in left.
	//
	// j tracks the current element in right.
	//
	// Both start at 0 because we initially want to compare
	// the first element of each sorted array.
	// ---------------------------------------------------------
	i, j := 0, 0

	// ---------------------------------------------------------
	// CONTINUE WHILE BOTH ARRAYS HAVE ELEMENTS
	// ---------------------------------------------------------
	//
	// We can only compare:
	//
	//     left[i]
	//     right[j]
	//
	// if both indexes are still inside their arrays.
	//
	for i < len(left) && j < len(right) {

		// -----------------------------------------------------
		// Compare the smallest remaining element from each side.
		//
		// Because both arrays are sorted, left[i] is the
		// smallest remaining element on the left, and right[j]
		// is the smallest remaining element on the right.
		//
		// Therefore, whichever is smaller must be the next
		// smallest element in the final sorted array.
		// -----------------------------------------------------
		if left[i] < right[j] {

			// left[i] is smaller, so it belongs next in res.
			res = append(res, left[i])

			// We have consumed this element.
			//
			// Move i to the next unused element in left.
			i++

		} else {

			// right[j] is smaller OR equal.
			//
			// Put it into the result.
			res = append(res, right[j])

			// We have consumed this element.
			//
			// Move j to the next unused element in right.
			j++
		}
	}

	// ---------------------------------------------------------
	// APPEND REMAINING LEFT ELEMENTS
	// ---------------------------------------------------------
	//
	// The main loop stops as soon as ONE of the arrays is
	// exhausted.
	//
	// If left still has elements, they are already sorted.
	//
	// Example:
	//
	//     left  = [5]
	//     right = []
	//
	// There is no reason to compare anything anymore.
	//
	// Just append the remaining elements.
	//
	res = append(res, left[i:]...)

	// ---------------------------------------------------------
	// APPEND REMAINING RIGHT ELEMENTS
	// ---------------------------------------------------------
	//
	// Same reasoning.
	//
	// If right still has elements, they are already sorted,
	// so append them directly.
	//
	res = append(res, right[j:]...)

	// ---------------------------------------------------------
	// The result now contains every element from both arrays
	// in sorted order.
	// ---------------------------------------------------------
	return res
}