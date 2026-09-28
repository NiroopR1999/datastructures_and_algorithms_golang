package stack

func carFleet(target int, position []int, speed []int) int {

	// Each car needs two pieces of information:
	// position -> where the car currently is
	// speed    -> how fast the car is moving
	//
	// We store them together as:
	// [position, speed]
	cars := [][2]int{}

	n := len(position)

	// Combine position and speed for each car.
	for i := range n {
		cars = append(cars, [2]int{position[i], speed[i]})
	}

	// Sort cars from closest to the target -> farthest from the target.
	//
	// Why?
	// We process the cars from FRONT -> BACK.
	// This allows every car to be compared with the fleet immediately
	// in front of it.
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	// Store the time at which each fleet reaches the target.
	//
	// We don't need to store the actual cars in a fleet.
	// We only care about their arrival time.
	fleet := []float64{}

	for _, car := range cars {

		// Calculate the time this car would take to reach the target
		// if it continued at its current speed.
		//
		// time = distance / speed
		//
		// distance = target - current position
		time := float64(target-car[0]) / float64(car[1])

		// If this is the first car, it automatically creates a fleet.
		//
		// Otherwise, compare its arrival time with the fleet directly
		// in front of it.
		//
		// If this car takes MORE time:
		//
		//     current car = 5 sec
		//     front fleet = 3 sec
		//
		// The current car cannot catch the fleet ahead.
		// Therefore, it becomes a NEW fleet.
		if len(fleet) == 0 || time > fleet[len(fleet)-1] {
			fleet = append(fleet, time)
		}

		// If:
		//
		//     current car = 2 sec
		//     front fleet = 3 sec
		//
		// The current car would reach earlier, so it catches the
		// fleet ahead before reaching the target.
		//
		// Therefore, we do NOT create a new fleet.
	}

	// Each element in fleet represents one independent fleet.
	return len(fleet)
}

/*
APPROACH:

1. Combine position and speed
   --------------------------------
   Store every car as:
       [position, speed]

2. Sort by position in descending order
   -------------------------------------
   Process cars from closest to the target
   to farthest from the target.

   Why?
   Because the car in front determines whether
   the current car can catch it.

3. Calculate each car's arrival time
   ----------------------------------
   time = (target - position) / speed

   This tells us when the car would reach the
   target if it could travel independently.

4. Compare with the fleet ahead
   -----------------------------
   If:

       current time > previous fleet time

   The current car cannot catch the fleet ahead.

   Therefore:
       → NEW FLEET

   If:

       current time <= previous fleet time

   The current car catches the fleet ahead.

   Therefore:
       → SAME FLEET

5. Count the fleets
   -----------------
   Every time we create a new fleet, we append
   its arrival time to fleet.

   len(fleet) = number of car fleets.


Example:

target = 12

position = [10, 8, 5]
speed    = [2, 4, 1]

After sorting:

position   speed
   10        2
    8        4
    5        1

Arrival times:

Car at 10:
(12 - 10) / 2 = 1 sec

Car at 8:
(12 - 8) / 4 = 1 sec

Car at 5:
(12 - 5) / 1 = 7 sec


Process:

Car at 10:
fleet = [1]

Car at 8:
1 <= 1
→ catches the fleet ahead
→ fleet = [1]

Car at 5:
7 > 1
→ cannot catch the fleet ahead
→ NEW fleet
→ fleet = [1, 7]

Answer = 2 fleets.


TIME COMPLEXITY:
O(n log n)

Why?
Sorting takes O(n log n), and the single loop takes O(n).


SPACE COMPLEXITY:
O(n)

cars   -> O(n)
fleet  -> O(n)


KEY IDEA TO REMEMBER:

Sort from FRONT → BACK.

Then:

current time > front fleet time
    → cannot catch
    → NEW FLEET

current time <= front fleet time
    → catches
    → SAME FLEET
*/