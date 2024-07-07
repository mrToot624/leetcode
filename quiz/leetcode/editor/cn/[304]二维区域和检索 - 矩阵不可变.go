package main

//leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	preSum[][]int
}


func Constructor_304(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])

	preSum := make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i][j+1] + preSum[i+1][j] - preSum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{preSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)
