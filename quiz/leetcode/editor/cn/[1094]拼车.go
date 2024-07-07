package main

func carPooling_Dirfference(trips [][]int, capacity int) bool {
	d := NewDifference(make([]int, 1001))

	for _, trip := range trips {
		start, dest := trip[1], trip[2]
		numPassengers := trip[0]

		d.increment(start, dest-1, numPassengers)
	}

	for _, num := range d.result() {
		if num > capacity {
			return false
		}
	}
	return true
}

//leetcode submit region begin(Prohibit modification and deletion)
func carPooling(trips [][]int, capacity int) bool {
	nums, diff := make([]int, 1001), make([]int, 1001)

	for _, trip := range trips {
		start, dest := trip[1], trip[2]
		numPassengers := trip[0]

		diff[start] += numPassengers
		diff[dest] -= numPassengers
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[0] = diff[0]
		} else {
			nums[i] = nums[i-1] + diff[i]
		}

		if nums[i] > capacity {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
