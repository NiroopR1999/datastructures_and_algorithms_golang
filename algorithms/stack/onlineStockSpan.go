package stack

type StockSpanner struct {
	// Each element stores:
	//
	// [0] = price
	// [1] = span
	//
	// Example:
	// [75, 4]
	//
	// means:
	// price = 75
	// span  = 4
	//
	// We use a stack because we only need to look at
	// the most recent previous prices.
	stack [][2]int
}

func Constructor2() StockSpanner {
	return StockSpanner{
		// Initially there are no previous stock prices.
		stack: [][2]int{},
	}
}

func (this *StockSpanner) Next(price int) int {

	// Today itself always counts as 1 day.
	//
	// Therefore, every span starts at 1.
	span := 1

	// Check the most recently stored price.
	//
	// We keep removing previous prices while they are
	// less than or equal to today's price.
	//
	// Why <= ?
	// Because the problem says that previous prices
	// equal to today's price are also included.
	for len(this.stack) != 0 &&
		this.stack[len(this.stack)-1][0] <= price {

		// The top element is:
		//
		// [previous price, previous span]
		//
		// We don't need to count those previous days
		// individually.
		//
		// Their span has already been calculated.
		//
		// Example:
		//
		// top = [75, 4]
		//
		// This means 75 represents 4 consecutive days.
		//
		// If today's price is 85, then 85 can include
		// all those 4 days.
		span += this.stack[len(this.stack)-1][1]

		// Remove the top element because its entire span
		// has now been included in today's span.
		//
		// Example:
		//
		// [100, 80, 75]
		//             ↑
		//            pop
		//
		// becomes:
		//
		// [100, 80]
		this.stack = this.stack[:len(this.stack)-1]
	}

	// Store today's price and its calculated span.
	//
	// Example:
	// price = 85
	// span  = 6
	//
	// Store:
	// [85, 6]
	this.stack = append(this.stack, [2]int{price, span})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */


/*
============================================================
APPROACH
============================================================

The stack stores:

    [price, span]

For every new price:

    1. Start with:
           span = 1

       Why?
       Because today itself always counts.

    2. Look at the top of the stack.

    3. If:

           top.price <= today's price

       then today's price can include that previous day
       and all the days represented by its stored span.

    4. Add the previous span:

           span += top.span

       Why?

       Because we already calculated that span earlier.
       We don't need to check those days again one by one.

    5. Pop that element from the stack.

    6. Continue checking the new top.

    7. Stop when:

           top.price > today's price

       because that price blocks us from going further
       into the past.

    8. Push today's:

           [price, span]

    9. Return span.


============================================================
THE CORE LOGIC
============================================================

    span = 1

    while stack is not empty
          AND top.price <= current price:

        span += top.span
        pop top

    push [current price, span]

    return span


============================================================
WHY THE STACK WORKS
============================================================

Suppose the stack contains:

    [100, 1]
    [80, 1]
    [75, 4]

Today's price = 85

First:

    85 >= 75

So we can absorb [75, 4]:

    span = 1 + 4
         = 5

Pop [75, 4].

Now:

    85 >= 80

So we can absorb [80, 1]:

    span = 5 + 1
         = 6

Pop [80, 1].

Now:

    85 < 100

So 100 blocks us.

Stop.

Push:

    [85, 6]


============================================================
IMPORTANT THING TO REMEMBER
============================================================

Each stack element is:

    [price, span]

Therefore:

    stack[len(stack)-1][0]
                         ↑
                       price

    stack[len(stack)-1][1]
                         ↑
                       span


The main trick is:

    "Don't count previous days one by one.

     Store their already-calculated span
     and jump over them."


============================================================
COMPLEXITY
============================================================

Time:  O(n) amortized

Why?

Every element is:

    pushed once
    popped at most once

So even though one Next() may pop many elements,
across all calls each element can only be popped once.

Space: O(n)

The stack can contain up to n elements.


============================================================
ONE-LINE MEMORY TRICK
============================================================

Current price beats previous price?

    YES → add previous span + pop

    NO  → stop

Then:

    push [current price, current span]
============================================================
*/