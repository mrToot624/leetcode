package main

// leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m)
	copy(tmp, nums1)

	p1, p2 := 0, 0
	for i := 0; i < m+n; i++ {
		if p1 == m {
			nums1[i] = nums2[p2]
			p2++
		} else if p2 == n {
			nums1[i] = tmp[p1]
			p1++
		} else if tmp[p1] <= nums2[p2] {
			nums1[i] = tmp[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
