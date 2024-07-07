package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func slidingPuzzle(board [][]int) int {
	boardStr := boardToString(board)
	ans := "123450"
	if boardStr == ans {
		return 0
	}

	visited := make(map[string]bool)
	visited[boardStr] = true

	var step int
	var q []string
	q = append(q, boardStr)
	for len(q) > 0 {
		step++
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[0]
			for _, shifted := range shiftedBoards(cur) {
				if shifted == ans {
					return step
				}

				if !visited[shifted] {
					visited[shifted] = true
					q = append(q, shifted)
				}
			}
			q = q[1:]
		}
	}

	return -1
}

func shiftedBoards(boardStr string) []string {
	bytes := []byte(boardStr)
	var zeroIndex int
	for i, b := range bytes {
		if b == '0' {
			zeroIndex = i
			break
		}
	}

	var res []string
	for _, swap := range findSwaps(zeroIndex, 2, 3) {
		newBytes := make([]byte, len(bytes))
		copy(newBytes, bytes)

		newBytes[zeroIndex] = newBytes[swap]
		newBytes[swap] = '0'
		res = append(res, string(newBytes))
	}
	return res
}

func findSwaps(i, m, n int) []int {
	var res []int
	if i - n >= 0 {
		res = append(res, i-n)
	}
	if i+n <= m*n-1 {
		res = append(res, i+n)
	}
	if i%n != 0 {
		res = append(res, i-1)
	}
	if i%n != n-1 {
		res = append(res, i+1)
	}
	return res
}

func boardToString(board [][]int) string {
	var builder strings.Builder
	for _, row := range board {
		for _, num := range row {
			builder.WriteString(strconv.Itoa(num))
		}
	}
	return builder.String()
}
//leetcode submit region end(Prohibit modification and deletion)
