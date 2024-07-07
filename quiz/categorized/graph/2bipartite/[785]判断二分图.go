package main

//leetcode submit region begin(Prohibit modification and deletion)

func isBipartite(graph [][]int) bool {
	res := true
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))

	for vertex := range graph {
		traverseGraph_785(vertex, graph, visited, color, &res)
	}
	return res
}

func traverseGraph_785(vertex int, graph [][]int, visited, color []bool, res *bool) {
	if !*res || visited[vertex] {
		return
	}

	visited[vertex] = true
	for _, neighbor := range graph[vertex] {
		if !visited[neighbor] {
			color[neighbor] = !color[vertex]
			traverseGraph_785(neighbor, graph, visited, color, res)
		} else if color[neighbor] == color[vertex] {
			*res = false
			return
		}
	}
}

func bfsIsBipartite(graph [][]int) bool {
	res := true
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))

	for cur := range graph {
		bfsGraph_785(cur, graph, visited, color, &res)
	}
	return res
}

func bfsGraph_785(node int, graph [][]int, visited, color []bool, res *bool) {
	if !*res || visited[node] {
		return
	}

	visited[node] = true
	var q []int
	q = append(q, node)

	for len(q) > 0 {
		cur := q[0]
		for _, neighbor := range graph[cur] {
			if !visited[neighbor] {
				visited[neighbor] = true
				color[neighbor] = !color[cur]
				q = append(q, neighbor)
			} else if color[neighbor] == color[cur] {
				*res = false
				return
			}
		}
		q = q[1:]
	}
}
//leetcode submit region end(Prohibit modification and deletion)
