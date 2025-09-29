package main

// leetcode submit region begin(Prohibit modification and deletion)
func finalPrices(prices []int) []int {
	var mq []int
	for i := len(prices) - 1; i >= 0; i-- {
		for len(mq) > 0 && prices[i] < mq[len(mq)-1] {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, prices[i])
		if len(mq) > 1 {
			prices[i] -= mq[len(mq)-2]
		}
	}
	return prices
}

//leetcode submit region end(Prohibit modification and deletion)
