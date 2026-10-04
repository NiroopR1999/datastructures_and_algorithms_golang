package arrayhashing
/*
	===========================================================
	PRODUCT OF ARRAY EXCEPT SELF
	===========================================================

	Problem:

	Given an integer array nums, return an array result such
	that:

		result[i] = product of every element in nums
		            EXCEPT nums[i]

	Example:

	Input:

		[1, 2, 3, 4]

	Output:

		[24, 12, 8, 6]


	Why?

	For index 0:

		nums[0] = 1

	Everything except 1:

		2 * 3 * 4 = 24

	Therefore:

		result[0] = 24


	For index 1:

		nums[1] = 2

	Everything except 2:

		1 * 3 * 4 = 12

	Therefore:

		result[1] = 12


	For index 2:

		nums[2] = 3

	Everything except 3:

		1 * 2 * 4 = 8

	Therefore:

		result[2] = 8


	For index 3:

		nums[3] = 4

	Everything except 4:

		1 * 2 * 3 = 6

	Therefore:

		result[3] = 6


	===========================================================
	IMPORTANT CONSTRAINT
	===========================================================

	The usual problem says:

		DO NOT use division.

	Why does that matter?

	A tempting solution would be:

		totalProduct / nums[i]

	For example:

		nums = [1, 2, 3, 4]

		total = 24

	Then:

		result[0] = 24 / 1 = 24
		result[1] = 24 / 2 = 12
		result[2] = 24 / 3 = 8
		result[3] = 24 / 4 = 6

	But division creates a problem when nums contains zero.

	Example:

		[1, 2, 0, 4]

	The total product is:

		0

	And:

		0 / 0

	is undefined.

	So instead, we calculate the products on the left and
	right independently.


	===========================================================
	CORE OBSERVATION
	===========================================================

	For every index i:

		result[i] =
			product of elements LEFT of i
			*
			product of elements RIGHT of i


	For:

		[1, 2, 3, 4]

	At index 2:

		3

	Left side:

		1 * 2 = 2

	Right side:

		4

	So:

		result[2] = 2 * 4
		          = 8


	Therefore, the problem becomes:

		1. Calculate prefix products.
		2. Calculate suffix products.
		3. Multiply them together.


	===========================================================
	APPROACH
	===========================================================

	We use ONE result array.

	First pass:
		Store the product of everything to the LEFT of each
		index.

	Second pass:
		Calculate the product of everything to the RIGHT of
		each index and multiply it into result.


	-----------------------------------------------------------
	PASS 1: PREFIX PRODUCT
	-----------------------------------------------------------

	We maintain:

		prefix

	which represents:

		"The product of all elements BEFORE the current index."

	Initially:

		prefix = 1

	Why 1?

	Because 1 is the multiplicative identity:

		1 * x = x

	And before index 0 there are no elements.

	So the product of the empty set is treated as 1.


	Example:

		nums = [1, 2, 3, 4]


	At i = 0:

		Elements before 0:

			[]

		Product:

			1

	So:

			result[0] = 1

	Then we update prefix:

			prefix *= nums[0]

			prefix = 1 * 1
			       = 1


	At i = 1:

		Elements before 1:

			[1]

	Therefore:

			result[1] = 1

	Then:

			prefix = 1 * 2
			       = 2


	At i = 2:

		Elements before 2:

			[1, 2]

	Therefore:

			result[2] = 2

	Then:

			prefix = 2 * 3
			       = 6


	At i = 3:

		Elements before 3:

			[1, 2, 3]

	Therefore:

			result[3] = 6

	Then:

			prefix = 6 * 4
			       = 24


	After the first pass:

		result = [1, 1, 2, 6]

	IMPORTANT:

	These are NOT the final answers yet.

	They represent:

		result[i] =
			product of everything LEFT of i


	===========================================================
	PASS 2: SUFFIX PRODUCT
	===========================================================

	Now we need the product of everything to the RIGHT.

	We traverse from RIGHT -> LEFT.

	We maintain:

		suffix

	which represents:

		"The product of all elements AFTER the current index."

	Initially:

		suffix = 1

	Again, why 1?

	Because there are no elements after the last element.

	And:

		1 is the identity for multiplication.


	For:

		[1, 2, 3, 4]

	we start at index 3.


	At i = 3:

		nums[3] = 4

	Everything to the right:

		[]

	Product:

		1

	Current result:

		result[3] = 6

	Multiply by suffix:

		result[3] *= suffix

		result[3] = 6 * 1
		          = 6

	Then update suffix:

		suffix *= nums[3]

		suffix = 1 * 4
		       = 4


	At i = 2:

	Everything to the right:

		[4]

	So:

		suffix = 4

	Current result:

		result[2] = 2

	Multiply:

		result[2] = 2 * 4
		          = 8

	Then:

		suffix = 4 * 3
		       = 12


	At i = 1:

	Everything to the right:

		[3, 4]

	So:

		suffix = 12

	Current result:

		result[1] = 1

	Multiply:

		result[1] = 1 * 12
		          = 12

	Then:

		suffix = 12 * 2
		       = 24


	At i = 0:

	Everything to the right:

		[2, 3, 4]

	So:

		suffix = 24

	Current result:

		result[0] = 1

	Multiply:

		result[0] = 1 * 24
		          = 24


	Final:

		result = [24, 12, 8, 6]


	===========================================================
	WHY CAN WE MODIFY result IN THE SECOND PASS?
	===========================================================

	This is the clever part.

	After the first pass:

		result[i]

	already contains:

		product of everything LEFT of i

	During the second pass:

		suffix

	contains:

		product of everything RIGHT of i

	So:

		result[i] *= suffix

	becomes:

		result[i] =
			leftProduct
			*
			rightProduct

	which is exactly:

		product of everything except nums[i]


	Therefore we don't need separate:

		prefix[]
		suffix[]

	arrays.

	We only need:

		result[]
		prefix
		suffix


	===========================================================
	DRY RUN
	===========================================================

	nums:

		[1, 2, 3, 4]


	-----------------------------------------------------------
	FIRST PASS
	-----------------------------------------------------------

	Start:

		prefix = 1

	i = 0:

		result[0] = 1
		prefix = 1 * 1 = 1

	i = 1:

		result[1] = 1
		prefix = 1 * 2 = 2

	i = 2:

		result[2] = 2
		prefix = 2 * 3 = 6

	i = 3:

		result[3] = 6
		prefix = 6 * 4 = 24


	result:

		[1, 1, 2, 6]


	-----------------------------------------------------------
	SECOND PASS
	-----------------------------------------------------------

	Start:

		suffix = 1


	i = 3:

		result[3] = 6 * 1 = 6
		suffix = 1 * 4 = 4


	i = 2:

		result[2] = 2 * 4 = 8
		suffix = 4 * 3 = 12


	i = 1:

		result[1] = 1 * 12 = 12
		suffix = 12 * 2 = 24


	i = 0:

		result[0] = 1 * 24 = 24
		suffix = 24 * 1 = 24


	Final:

		[24, 12, 8, 6]


	===========================================================
	WHAT ABOUT ZERO?
	===========================================================

	This approach automatically handles zero.

	Example:

		nums = [1, 2, 0, 4]


	We NEVER divide.

	We simply calculate products on each side.


	Expected:

	[0, 0, 8, 0]


	Why?

	Index 0:

		2 * 0 * 4 = 0


	Index 1:

		1 * 0 * 4 = 0


	Index 2:

		1 * 2 * 4 = 8


	Index 3:

		1 * 2 * 0 = 0


	The prefix/suffix approach handles this naturally.


	===========================================================
	WHY NOT CREATE PREFIX AND SUFFIX ARRAYS?
	===========================================================

	We could do:

		prefix[i] = product before i
		suffix[i] = product after i

	Then:

		result[i] = prefix[i] * suffix[i]

	That would work.

	But it would require:

		O(n)

	additional space for prefix and another:

		O(n)

	for suffix.

	Our solution reuses result itself for the prefix values.

	So the only additional variables are:

		prefix
		suffix

	Therefore we achieve O(1) EXTRA SPACE
	(excluding the required output array).


	===========================================================
	COMPLEXITY
	===========================================================

	Time:

		O(n)

	Why?

	We make exactly two passes over the array.

	First pass:

		O(n)

	Second pass:

		O(n)

	Therefore:

		O(n) + O(n)
		= O(n)


	Space:

		O(1) EXTRA SPACE

	Why?

	We only use:

		prefix
		suffix

	The result array doesn't count as extra space because it
	is the required output.

	If an interviewer counts the output array, total memory is:

		O(n)

	but auxiliary/extra space is:

		O(1)


	===========================================================
	INTERVIEW EXPLANATION
	===========================================================

	"For each index, the answer is the product of everything
	to its left multiplied by the product of everything to its
	right.

	I first make a left-to-right pass and store the prefix
	product in result[i].

	Then I make a right-to-left pass while maintaining a suffix
	product.

	For every index, I multiply result[i] by the current suffix
	product.

	This lets me calculate the answer without division and
	without using separate prefix and suffix arrays.

	The time complexity is O(n), and the extra space complexity
	is O(1)."
*/


func productExceptSelf(nums []int) []int {

	// --------------------------------------------------------
	// Create the output array.
	//
	// We need one result for every element in nums, so the
	// length must be the same as nums.
	//
	// Initially it contains zeros:
	//
	//     [0, 0, 0, 0]
	//
	// We will fill it during the prefix pass.
	// --------------------------------------------------------
	result := make([]int, len(nums))

	// --------------------------------------------------------
	// prefix represents:
	//
	//     product of all elements BEFORE the current index
	//
	// Initially there are no elements before index 0.
	//
	// The product of "nothing" is represented by 1 because:
	//
	//     1 * x = x
	//
	// This allows index 0 to work naturally.
	// --------------------------------------------------------
	prefix := 1

	// --------------------------------------------------------
	// FIRST PASS:
	//
	// Move from LEFT -> RIGHT.
	//
	// Before processing nums[i], prefix contains the product
	// of everything to the LEFT of i.
	// --------------------------------------------------------
	for i := range nums {

		// ----------------------------------------------------
		// Store the product of everything to the left of i.
		//
		// Example:
	//
	//     nums = [1, 2, 3, 4]
	//
	// At i = 2:
	//
	//     prefix = 1 * 2 = 2
	//
	// So:
	//
	//     result[2] = 2
	//
	// This is the left-side contribution to the answer.
	// ----------------------------------------------------
		result[i] = prefix

		// ----------------------------------------------------
		// Now include nums[i] in the prefix product so that
		// the NEXT index can use it.
		//
		// Example:
	//
	//     prefix = 2
	//     nums[i] = 3
	//
	//     prefix = 2 * 3 = 6
	//
	// At the next index, 6 represents:
	//
	//     1 * 2 * 3
	//
	// ----------------------------------------------------
		prefix *= nums[i]
	}

	// --------------------------------------------------------
	// At this point:
	//
	//     result[i]
	//
	// contains ONLY the product of elements to the LEFT.
	//
	// Example:
	//
	//     nums   = [1, 2, 3, 4]
	//     result = [1, 1, 2, 6]
	//
	// We still need the product of elements to the RIGHT.
	// --------------------------------------------------------

	// --------------------------------------------------------
	// suffix represents:
	//
	//     product of all elements AFTER the current index
	//
	// We start at 1 because there are no elements to the right
	// of the final element.
	// --------------------------------------------------------
	suffix := 1

	// --------------------------------------------------------
	// SECOND PASS:
	//
	// Move from RIGHT -> LEFT.
	//
	// Before processing nums[i], suffix contains the product
	// of everything to the RIGHT of i.
	// --------------------------------------------------------
	for i := len(nums) - 1; i >= 0; i-- {

		// ----------------------------------------------------
		// result[i] already contains:
		//
		//     product of everything LEFT of i
		//
		// suffix contains:
		//
		//     product of everything RIGHT of i
		//
		// Therefore multiplying them gives:
		//
		//     product of everything EXCEPT nums[i]
		// ----------------------------------------------------
		result[i] *= suffix

		// ----------------------------------------------------
		// Now include nums[i] in suffix.
		//
		// This is done AFTER using suffix for result[i].
		//
		// Why?
		//
		// nums[i] must NOT be included in its own answer.
		//
		// But it SHOULD be included when calculating the suffix
		// for the NEXT index to the left.
		//
		// Example:
		//
		// nums = [1, 2, 3, 4]
		//
		// At i = 2:
		//
		//     suffix = 4
		//
		// This is correct because only 4 is to the right of 3.
		//
		// AFTER calculating result[2], we update:
		//
		//     suffix = 4 * 3 = 12
		//
		// Now 12 correctly represents:
		//
		//     3 * 4
		//
		// which is everything to the right of index 1.
		// ----------------------------------------------------
		suffix *= nums[i]
	}

	// --------------------------------------------------------
	// Every result[i] now contains:
	//
	//     left product * right product
	//
	// which is exactly the product of every element except
	// nums[i].
	// --------------------------------------------------------
	return result
}