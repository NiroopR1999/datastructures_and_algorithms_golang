package stack

type FreqStack struct {
	// freq[val] tells us how many times a value currently exists
	// in the stack.
	//
	// Example:
	// If we pushed 5 three times:
	// freq[5] = 3
	freq map[int]int

	// freqStack[f] stores all values whose CURRENT frequency is f.
	//
	// It behaves like a normal stack.
	//
	// Why do we need this?
	// When multiple values have the same highest frequency,
	// we need to remove the one that was added most recently.
	//
	// Example:
	// freqStack[2] = [5, 7, 5]
	//
	// Both 5 and 7 have frequency 2, but the latest one
	// among them is the last 5, so we can simply pop from
	// the end of this slice.
	freqStack map[int][]int

	// maxFreq stores the highest frequency of any value
	// currently present in the stack.
	//
	// This allows Pop() to immediately know which
	// frequency group to use instead of searching all
	// frequencies.
	maxFreq int
}

func Constructor1() FreqStack {
	return FreqStack{
		// Create an empty map to keep track of each
		// value's frequency.
		freq: make(map[int]int),

		// Create an empty map where each frequency will
		// have its own stack of values.
		freqStack: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {

	// Increase the frequency of this value.
	//
	// Example:
	// Before: freq[5] = 2
	// After:  freq[5] = 3
	this.freq[val]++

	// f is the NEW frequency of val.
	//
	// We need this because val now belongs to the
	// group of elements having frequency f.
	f := this.freq[val]

	// Put val into the stack for its new frequency.
	//
	// We append it because, when there is a tie in
	// frequency, the most recently added value should
	// be removed first.
	this.freqStack[f] = append(this.freqStack[f], val)

	// Update the highest frequency if necessary.
	if f > this.maxFreq {
		this.maxFreq = f
	}
}

func (this *FreqStack) Pop() int {

	// Get the stack containing all values with
	// the current highest frequency.
	popStack := this.freqStack[this.maxFreq]

	// Take the most recently added value from this
	// frequency group.
	popVal := popStack[len(popStack)-1]

	// Remove the last element from the frequency stack.
	this.freqStack[this.maxFreq] =
		this.freqStack[this.maxFreq][:len(this.freqStack[this.maxFreq])-1]

	// The value was removed, so decrease its frequency.
	this.freq[popVal]--

	// If no values remain at the current maximum frequency,
	// reduce maxFreq because that frequency no longer exists.
	if len(this.freqStack[this.maxFreq]) == 0 {
		this.maxFreq--
	}

	return popVal
}

/*
===========================================================
THE ONE IDEA TO REMEMBER
===========================================================

freq[value]
      ↓
"What is this value's frequency?"

freqStack[frequency]
      ↓
"Among values with this frequency,
 which one was inserted most recently?"

maxFreq
      ↓
"What is the highest frequency right now?"


So Pop() is basically:

1. Go to maxFreq
       ↓
2. Take the last element
       ↓
3. Decrease its frequency
       ↓
4. If that frequency group is empty,
   decrease maxFreq

This gives us:

Push → O(1) average
Pop  → O(1) average
Space → O(n)
===========================================================
*/

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor1()
 * obj.Push(val)
 * param2 := obj.Pop()
 */