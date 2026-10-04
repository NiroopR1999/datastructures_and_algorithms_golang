package arrayhashing

/*
	===========================================================
	GROUP ANAGRAMS
	===========================================================

	Problem:
	Given an array of strings, group the anagrams together.

	Example:

	Input:
	["eat", "tea", "tan", "ate", "nat", "bat"]

	Output:
	[
		["eat", "tea", "ate"],
		["tan", "nat"],
		["bat"],
	]

	-----------------------------------------------------------
	WHAT IS AN ANAGRAM?
	-----------------------------------------------------------

	Two strings are anagrams if they contain:

		1. The same characters
		2. With the same frequency
		3. But possibly in a different order

	Example:

		"eat"
		"tea"
		"ate"

	All three contain:

		a -> 1
		e -> 1
		t -> 1

	So they are anagrams.

	But:

		"eat"
		"tan"

	have different character frequencies, so they are NOT
	anagrams.


	===========================================================
	APPROACH
	===========================================================

	Instead of sorting every string, we will create a
	"frequency signature" for every string.

	For every string, we maintain an array of size 26.

	Each position represents one lowercase English letter:

		index 0  -> 'a'
		index 1  -> 'b'
		index 2  -> 'c'
		...
		index 25 -> 'z'

	For example:

		"eat"

	contains:

		a -> 1
		e -> 1
		t -> 1

	So its frequency array will conceptually look like:

		[a=1, b=0, c=0, d=0, e=1, ..., t=1, ..., z=0]

	Now consider:

		"tea"

	It contains the exact same characters with the exact
	same frequencies.

	So it produces the EXACT SAME [26]int key.

	Therefore:

		"eat" -> same key
		"tea" -> same key
		"ate" -> same key

	We can use that key in a map.

	-----------------------------------------------------------
	MAP STRUCTURE
	-----------------------------------------------------------

	We use:

		map[[26]int][]string

	Why?

	Because:

		[26]int
			=
		the frequency signature

	And:

		[]string
			=
		all strings having that signature

	Conceptually:

		frequencyKey1 -> ["eat", "tea", "ate"]
		frequencyKey2 -> ["tan", "nat"]
		frequencyKey3 -> ["bat"]


	===========================================================
	WHY NOT SORT?
	===========================================================

	Another common solution is:

		"eat" -> sort -> "aet"
		"tea" -> sort -> "aet"
		"ate" -> sort -> "aet"

	Then use "aet" as the map key.

	But sorting a string of length K costs:

		O(K log K)

	Our frequency-counting approach only needs:

		O(K)

	because we simply visit every character once.

	So this solution is preferable when the problem specifically
	asks for a solution WITHOUT sorting.


	===========================================================
	TIME COMPLEXITY
	===========================================================

	Let:

		N = number of strings
		K = maximum length of a string

	For every string, we visit every character once:

		O(K)

	For N strings:

		O(N * K)

	The [26]int comparison used by the map is effectively O(1)
	because 26 is a fixed constant.


	===========================================================
	SPACE COMPLEXITY
	===========================================================

	The map stores all the input strings in groups.

	So the space required for storing the strings is:

		O(N * K)

	We also create a [26]int for each string, but because 26
	is constant, that additional key space is effectively O(1)
	per string during processing.

	Overall:

		O(N * K)


	===========================================================
	IMPORTANT GO CONCEPT
	===========================================================

	Why can [26]int be used as a map key?

	Go requires map keys to be comparable.

	Arrays are comparable if their elements are comparable.

	int is comparable.

	Therefore:

		[26]int

	is comparable and can safely be used as a map key.

	However, a slice like:

		[]int

	cannot be used directly as a map key.

	That's why we use:

		[26]int

	instead of:

		[]int
*/

func groupAnagrams(strs []string) [][]string {

	/*
		-------------------------------------------------------
		STEP 1: CREATE THE MAP
		-------------------------------------------------------

		We need a map where:

			key   = character frequency
			value = all strings having that frequency

		So:

			map[[26]int][]string

		means:

			[26]int  -> identifies an anagram group
			[]string -> stores the actual strings in that group

		Initially the map is empty.

		Example:

			groups = {}
	*/
	groups := make(map[[26]int][]string)

	/*
		-------------------------------------------------------
		STEP 2: PROCESS EVERY STRING
		-------------------------------------------------------

		We cannot determine whether strings are anagrams without
		looking at their characters.

		So we process every string independently.

		Example:

			strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

		First iteration:

			str = "eat"

		Second iteration:

			str = "tea"

		and so on.
	*/
	for _, str := range strs {

		/*
			---------------------------------------------------
			STEP 3: CREATE A FREQUENCY ARRAY
			---------------------------------------------------

			We need to count how many times each character occurs
			in the CURRENT string.

			There are only 26 lowercase English letters.

			So we create:

				[26]int

			Every integer starts at 0 automatically.

			Therefore:

				var key [26]int

			initially gives:

				[0, 0, 0, 0, 0, ..., 0]

			We will modify this array based on the characters
			in the current string.

			IMPORTANT:

			We create a NEW key for every string.

			Why?

			Because the frequency signature belongs to the
			current string only.

			We don't want the character counts from "eat" to
			carry over when processing "tea".
		*/
		var key [26]int

		/*
			---------------------------------------------------
			STEP 4: COUNT EVERY CHARACTER
			---------------------------------------------------

			Now examine every character in the current string.

			For example:

				str = "eat"

			Characters processed:

				'e'
				'a'
				't'
		*/
		for _, char := range str {

			/*
				------------------------------------------------
				WHY char - 'a'?
				------------------------------------------------

				We need to convert a character into an array
				index from 0 to 25.

				Characters have integer values internally.

				So:

					'a' - 'a' = 0
					'b' - 'a' = 1
					'c' - 'a' = 2
					'd' - 'a' = 3
					...
					'z' - 'a' = 25

				Therefore:

					char - 'a'

				gives us the correct position in our 26-element
				array.

				For example:

					char = 'e'

					'e' - 'a' = 4

				So:

					key[4]++

				means:

					increment the count of 'e'.

				Another example:

					char = 't'

					't' - 'a' = 19

				So:

					key[19]++

				means:

					increment the count of 't'.


				------------------------------------------------
				WHY ++?
				------------------------------------------------

				We are counting frequency.

				If we see a character once:

					count = 1

				If we see it again:

					count = 2

				Again:

					count = 3

				And so on.

				Therefore:

					key[char-'a']++

				is equivalent to:

					key[char-'a'] =
						key[char-'a'] + 1
			*/
			key[char-'a']++
		}

		/*
			---------------------------------------------------
			STEP 5: USE THE FREQUENCY ARRAY AS THE MAP KEY
			---------------------------------------------------

			At this point, "key" completely describes the
			characters inside the current string.

			For example:

				"eat"

			produces:

				a -> 1
				e -> 1
				t -> 1

			And:

				"tea"

			also produces:

				a -> 1
				e -> 1
				t -> 1

			Therefore:

				key("eat") == key("tea")

			This is exactly what we want.

			Now we use that frequency array as the map key.

			---------------------------------------------------
			WHAT DOES groups[key] MEAN?
			---------------------------------------------------

			Suppose this is our map:

				groups = {
					key1: ["eat", "tea"]
				}

			If the current string is "ate", it produces the
			same key1.

			So:

				groups[key]

			gives:

				["eat", "tea"]

			Then:

				append(groups[key], "ate")

			becomes:

				["eat", "tea", "ate"]

			Finally, we assign that slice back to the map:

				groups[key] = append(groups[key], str)

			---------------------------------------------------
			WHY DOES THIS GROUP ANAGRAMS?
			---------------------------------------------------

			Because anagrams ALWAYS have the same frequency
			array.

			Therefore:

				same frequency array
						↓
				same map key
						↓
				same map entry
						↓
				same anagram group
		*/
		groups[key] = append(groups[key], str)
	}

	/*
		-------------------------------------------------------
		STEP 6: CREATE THE FINAL RESULT
		-------------------------------------------------------

		Our map now contains all the groups.

		For example:

			groups =

				key1 -> ["eat", "tea", "ate"]
				key2 -> ["tan", "nat"]
				key3 -> ["bat"]

		But the required return type is:

			[][]string

		So we need to extract the values from the map.

		We initialize an empty 2D slice.

		Why len(groups)?

		Because every unique key represents exactly one group.

		If there are 3 unique keys, there will be 3 groups.

		So len(groups) is a good initial capacity.
	*/
	res := make([][]string, 0, len(groups))

	/*
		-------------------------------------------------------
		STEP 7: EXTRACT EVERY GROUP FROM THE MAP
		-------------------------------------------------------

		We only care about the values of the map here.

		The key was useful while grouping, but we don't need
		the frequency arrays in our final answer.

		For example:

			key1 -> ["eat", "tea", "ate"]

		We take:

			["eat", "tea", "ate"]

		and append it to res.
	*/
	for _, group := range groups {

		/*
			Add the current anagram group to the final result.

			Example:

				res = []

			after first group:

				res = [
					["eat", "tea", "ate"]
				]

			after second group:

				res = [
					["eat", "tea", "ate"],
					["tan", "nat"]
				]
		*/
		res = append(res, group)
	}

	/*
		-------------------------------------------------------
		STEP 8: RETURN THE RESULT
		-------------------------------------------------------

		At this point every string has been placed into the
		correct anagram group.

		We return the 2D slice.
	*/
	return res
}