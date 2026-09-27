package stack

type MinStack struct {
	// s is the actual stack.
	// It stores every value that we push.
	s []int

	// min stores the minimum value at every level of the stack.
	//
	// Example:
	// s:   [5, 3, 7, 2]
	// min: [5, 3, 3, 2]
	//
	// The value at min[i] represents the minimum value
	// from s[0] up to s[i].
	min []int
}

func Constructor() MinStack {
	return MinStack{
		// Start with two empty slices because initially
		// the stack contains no elements.
		s:   []int{},
		min: []int{},
	}
}

func (this *MinStack) Push(val int) {

	// Push the new value into the actual stack.
	this.s = append(this.s, val)

	if len(this.min) == 0 {

		// If this is the first element, it automatically
		// becomes the minimum because there is nothing else
		// in the stack to compare it with.
		this.min = append(this.min, val)

	} else {

		// The top of min contains the minimum value
		// of everything currently in the stack.
		//
		// We use it to determine whether the newly pushed
		// value becomes the new minimum.
		minVal := this.min[len(this.min)-1]

		if val < minVal {

			// If the new value is smaller than the previous
			// minimum, the new value becomes the minimum.
			minVal = val
		}

		// Store the minimum for this new level of the stack.
		//
		// We do this because when we later Pop(), we need to
		// know what the minimum was before this element was added.
		this.min = append(this.min, minVal)
	}
}

func (this *MinStack) Pop() {

	// Remove the top element from the actual stack.
	//
	// We don't need the value anymore because it has been popped.
	this.s = this.s[:len(this.s)-1]

	// Remove the corresponding minimum as well.
	//
	// Why?
	// Because min[i] belongs to the same stack level as s[i].
	//
	// If we don't remove it, min would contain information
	// about an element that no longer exists in the stack.
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {

	// The last element of a slice represents the top
	// of our stack.
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {

	// The top of min always contains the minimum value
	// of the entire actual stack.
	//
	// Therefore, we don't need to scan s.
	//
	// Scanning s would take O(n), but directly accessing
	// the top of min takes O(1).
	return this.min[len(this.min)-1]
}