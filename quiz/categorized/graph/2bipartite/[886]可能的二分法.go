package main


//leetcode submit region begin(Prohibit modification and deletion)
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := buildGraph(n, dislikes)
	res := true
	visited := make([]bool, n+1)
	color := make([]bool, n+1)

	for vertex := range graph {
		traverseGraph_886(vertex, graph, visited, color, &res)
	}
	return res
}

func traverseGraph_886(vertex int, graph [][]int, visited, color []bool, res *bool) {
	if !*res || visited[vertex] {
		return
	}

	visited[vertex] = true
	for _, neighbor := range graph[vertex] {
		if !visited[neighbor] {
			color[neighbor] = !color[vertex]
			traverseGraph_886(neighbor, graph, visited, color, res)
		} else if color[neighbor] == color[vertex] {
			*res = false
			return
		}
	}
}

func buildGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n+1)
	for _, pair := range dislikes {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}
	return graph
}
//leetcode submit region end(Prohibit modification and deletion)
