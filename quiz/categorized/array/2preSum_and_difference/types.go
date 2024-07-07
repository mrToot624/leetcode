package main

// 差分数组
type difference struct {
	diff []int
}

func NewDifference(nums []int) *difference {
	diff := make([]int, len(nums))

	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &difference{diff}
}

func (d *difference) increment(start, end, delta int) {
	d.diff[start] += delta
	if end + 1 < len(d.diff) {
		d.diff[end+1] -= delta
	}
}

func (d *difference) result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}