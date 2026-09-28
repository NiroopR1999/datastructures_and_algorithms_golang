package stack

func asteroidCollision(asteroids []int) []int {

	// res acts as a stack.
	//
	// We store asteroids that are still alive.
	//
	// We only need to compare the current asteroid with
	// the most recently surviving asteroid because that is
	// the only asteroid it can collide with first.
	res := []int{}

	for i := range asteroids {

		// Take the current asteroid.
		asteroid := asteroids[i]

		// A collision can happen only when:
		//
		// 1. Current asteroid is moving LEFT  (< 0)
		// 2. Stack top is moving RIGHT         (> 0)
		//
		// Example:
		//
		// stack:   [5]
		// current: -3
		//
		// 5 →    ← -3
		//
		// They are moving toward each other.
		//
		// Other combinations don't collide:
		//
		// + +  → →
		// - -  ← ←
		// - +  ← →
		//
		// In those cases, the asteroids are moving away
		// from each other.
		for len(res) > 0 &&
			asteroid < 0 &&
			res[len(res)-1] > 0 {

			// Add the two asteroid values to determine
			// which asteroid is larger.
			//
			// Example:
			//
			// 5 and -3
			// 5 + (-3) = 2
			//
			// Positive result → right-moving asteroid wins.
			//
			// 3 and -5
			// 3 + (-5) = -2
			//
			// Negative result → left-moving asteroid wins.
			//
			// Equal sizes:
			// 5 + (-5) = 0
			//
			// Both are destroyed.
			collision := asteroid + res[len(res)-1]

			if collision == 0 {

				// Both asteroids have the same size,
				// so both are destroyed.
				//
				// Remove the right-moving asteroid
				// from the stack.
				res = res[:len(res)-1]

				// Mark the current asteroid as destroyed too.
				//
				// We use 0 as a marker so that we don't
				// append it to the result later.
				asteroid = 0

			} else if collision > 0 {

				// The right-moving asteroid is larger.
				//
				// Example:
				// 5 and -3
				//
				// -3 is destroyed.
				//
				// The stack asteroid survives, so we don't
				// remove anything from res.
				//
				// Mark current asteroid as destroyed.
				asteroid = 0

			} else {

				// The current left-moving asteroid is larger.
				//
				// Example:
				// 3 and -5
				//
				// 3 is destroyed.
				//
				// Remove the surviving stack asteroid and
				// continue the loop.
				//
				// Why continue?
				//
				// The current asteroid (-5) may now collide
				// with another asteroid further to the left.
				res = res[:len(res)-1]
			}
		}

		// If the current asteroid survived all collisions,
		// add it to the stack.
		//
		// asteroid == 0 means it was destroyed.
		if asteroid != 0 {
			res = append(res, asteroid)
		}
	}

	return res
}