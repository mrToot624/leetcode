package main

func corpFlightBookings_Difference(bookings [][]int, n int) []int {
	d := NewDifference(make([]int, n))

	for _, booking := range bookings {
		start, end := booking[0], booking[1]
		order := booking[2]
		d.increment(start-1, end-1, order)
	}

	return d.result()
}

//leetcode submit region begin(Prohibit modification and deletion)

func corpFlightBookings(bookings [][]int, n int) []int {
	diff, res := make([]int, n), make([]int, n)

	for _, booking := range bookings {
		start, end := booking[0], booking[1]
		order := booking[2]
		diff[start-1] += order

		if end < n {
			diff[end] -= order
		}
	}

	res[0] = diff[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + diff[i]
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
