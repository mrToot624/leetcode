package main

//leetcode submit region begin(Prohibit modification and deletion)
func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var track []int
	traverseGraph_797(graph, &res, track, 0)
	return res
}

func traverseGraph_797(graph [][]int, res *[][]int, track []int, num int) {
	track = append(track, num)

	if num == len(graph)-1 {
		path := make([]int, len(track))
		copy(path, track)
		*res = append(*res, path)
	}

	for _, next := range graph[num] {
		traverseGraph_797(graph, res, track, next)
	}

	track = track[:len(track)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
