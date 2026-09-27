package stack

func dailyTemperatures(temperatures []int) []int {

	n := len(temperatures)

	// res[i] stores how many days we need to wait after
	// day i to get a warmer temperature.
	//
	// Example:
	// temperatures = [73, 74, 75, 71]
	//
	// res = [1, 1, 0, 0]
	//
	// We initialize it with n zeros because if no warmer
	// day exists, the answer is already 0.
	res := make([]int, n)

	// Stack stores the INDICES of temperatures.
	//
	// We store indices instead of the actual temperatures
	// because we need the index difference later:
	//
	//     warmerDayIndex - currentDayIndex
	//
	// The stack will maintain a decreasing sequence of
	// temperatures from bottom to top.
	stack := []int{}

	// We process from RIGHT to LEFT.
	//
	// Why?
	// For each day, we are looking for a warmer temperature
	// somewhere in the FUTURE.
	//
	// By starting from the right, the temperatures to the
	// right have already been processed and are available
	// in our stack.
	for i := n - 1; i >= 0; i-- {

		// Remove temperatures that cannot be the answer
		// for the current day.
		//
		// If today's temperature is greater than or equal to
		// the temperature at the top of the stack, that stack
		// temperature is useless.
		//
		// Example:
		//
		// Today's temperature = 75
		// Stack top temperature = 71
		//
		// 71 cannot be a warmer day for 75.
		//
		// Therefore, remove it.
		//
		// We also remove equal temperatures because the problem
		// asks for a STRICTLY warmer temperature.
		for len(stack) != 0 &&
			temperatures[i] >= temperatures[stack[len(stack)-1]] {

			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {

			// There is no warmer temperature to the right.
			//
			// Therefore, we leave res[i] as 0.
			res[i] = 0

		} else {

			// The stack is not empty.
			//
			// The index at the top of the stack represents
			// the nearest useful warmer temperature.
			//
			// Example:
			//
			// current index = 3
			// warmer index = 5
			//
			// Number of days to wait:
			//
			//     5 - 3 = 2
			//
			// This is why we store INDICES in the stack.
			res[i] = stack[len(stack)-1] - i
		}

		// Add the current index to the stack.
		//
		// Why?
		// An earlier day might need today's temperature as
		// its next warmer temperature.
		//
		// We add the index AFTER finding the answer for today
		// because today's temperature cannot be a future
		// temperature for itself.
		stack = append(stack, i)
	}

	return res
}