package arrayhashing

/*
	===========================================================
	REMOVE ELEMENT
	===========================================================

	Problem:
	Given an integer array nums and an integer val, remove all
	occurrences of val IN-PLACE.

	Return the number of elements that are not equal to val.

	Example:

	Input:
		nums = [3, 2, 2, 3]
		val  = 3

	After removing 3:

		[2, 2]

	Return:

		2


	===========================================================
	IMPORTANT REQUIREMENT
	===========================================================

	The problem says:

		"Remove elements in-place."

	This means we should NOT create another array to store
	the answer.

	Instead, we modify the existing nums array itself.


	===========================================================
	APPROACH: TWO POINTERS
	===========================================================

	We use two indexes:

		1. readIndex
		2. writeIndex


	-----------------------------------------------------------
	readIndex
	-----------------------------------------------------------

	readIndex is responsible for:

		"Looking at every element."

	It moves from LEFT -> RIGHT through the entire array.

	Therefore:

		readIndex = 0
		readIndex = 1
		readIndex = 2
		...

	It never skips an element.


	-----------------------------------------------------------
	writeIndex
	-----------------------------------------------------------

	writeIndex tells us:

	"Where should the next valid element be written?"

	An element is valid if:

		nums[i] != val

	Therefore, writeIndex only moves when we find an element
	that we want to KEEP.


	===========================================================
	CORE IDEA
	===========================================================

	We read every element using readIndex.

	There are only two possibilities:


	CASE 1:
		nums[readIndex] == val

	We DON'T want this element.

	So:

		DO NOTHING

	We simply move readIndex forward.


	CASE 2:
		nums[readIndex] != val

	We WANT to keep this element.

	So we copy it to the position indicated by writeIndex:

		nums[writeIndex] = nums[readIndex]

	Then move writeIndex forward.

	At the end:

		writeIndex

	is exactly the number of elements that should remain.


	===========================================================
	EXAMPLE
	===========================================================

	nums = [3, 2, 2, 3]
	val  = 3


	Initially:

		writeIndex = 0
		readIndex  = 0


	-----------------------------------------------------------
	READ 3
	-----------------------------------------------------------

		nums[readIndex]
			=
		nums[0]
			=
		3

	3 == val

	So we DON'T want it.

	Do not write anything.

	Move:

		readIndex = 1

	Array is still:

		[3, 2, 2, 3]


	-----------------------------------------------------------
	READ 2
	-----------------------------------------------------------

		nums[1] = 2

	2 != 3

	This is a valid element.

	writeIndex = 0

	So:

		nums[0] = nums[1]

	Therefore:

		[2, 2, 2, 3]

	Then:

		writeIndex = 1
		readIndex  = 2


	-----------------------------------------------------------
	READ 2
	-----------------------------------------------------------

		nums[2] = 2

	2 != 3

	Keep it.

		writeIndex = 1

	So:

		nums[1] = nums[2]

	Array remains:

		[2, 2, 2, 3]

	Then:

		writeIndex = 2
		readIndex  = 3


	-----------------------------------------------------------
	READ 3
	-----------------------------------------------------------

		nums[3] = 3

	3 == val

	Don't keep it.

	Just:

		readIndex = 4


	Now:

		readIndex == len(nums)

	So we are done.


	Final array:

		[2, 2, 2, 3]

	The important part is only the first writeIndex elements:

		[2, 2]

	So:

		return 2


	===========================================================
	WHY DON'T WE NEED TO ACTUALLY DELETE ELEMENTS?
	===========================================================

	This is an important detail of this problem.

	We don't physically shrink the slice.

	For example, we might finish with:

		[2, 2, 2, 3]

	But return:

		2

	The caller only considers:

		nums[:2]

	which is:

		[2, 2]

	Everything after index 1 is irrelevant.


	===========================================================
	COMPLEXITY
	===========================================================

	Time:

		O(n)

	Why?

	Because readIndex goes through the array exactly once.


	Space:

		O(1)

	Why?

	Because we only use two integer variables:

		writeIndex
		readIndex

	We don't create another array.


	===========================================================
	INTERVIEW SUMMARY
	===========================================================

	"I use two pointers.

	readIndex scans every element.

	writeIndex tracks where the next element that should be
	kept needs to be placed.

	If nums[readIndex] equals val, I skip it.

	If it doesn't equal val, I copy it to nums[writeIndex]
	and increment writeIndex.

	At the end, writeIndex represents the number of elements
	that remain."

*/

func removeElement(nums []int, val int) int {

	// --------------------------------------------------------
	// writeIndex tells us WHERE to place the next element
	// that we want to keep.
	//
	// Initially, no valid elements have been found, so the
	// first valid element should be written at index 0.
	// --------------------------------------------------------
	writeIndex := 0

	// --------------------------------------------------------
	// readIndex scans every element in the original array.
	//
	// We start at index 0 because we need to examine every
	// element to determine whether it should be kept or removed.
	// --------------------------------------------------------
	readIndex := 0

	// --------------------------------------------------------
	// Continue until readIndex has examined every element.
	// --------------------------------------------------------
	for readIndex < len(nums) {

		// ----------------------------------------------------
		// Check whether the current element should be removed.
		//
		// If nums[readIndex] == val:
		//
		//     We DON'T want this element.
		//
		// Therefore, we don't write it anywhere.
		//
		// If nums[readIndex] != val:
		//
		//     We WANT to keep it.
		//
		// So we copy it to the next available position,
		// which is writeIndex.
		// ----------------------------------------------------
		if nums[readIndex] != val {

			// Copy the valid element to the next position
			// where a kept element should exist.
			//
			// Notice that writeIndex can be behind readIndex.
			//
			// This is what allows us to overwrite elements
			// that need to be removed.
			nums[writeIndex] = nums[readIndex]

			// We successfully placed one valid element.
			//
			// Therefore, the next valid element should go
			// into the next position.
			writeIndex++
		}

		// Move to the next element that we need to inspect.
		//
		// IMPORTANT:
		// readIndex ALWAYS moves, regardless of whether we
		// keep or remove the current element.
		readIndex++
	}

	// --------------------------------------------------------
	// writeIndex is now the number of elements that we kept.
	//
	// The valid portion of nums is:
	//
	//     nums[:writeIndex]
	//
	// So we return writeIndex.
	// --------------------------------------------------------
	return writeIndex
}