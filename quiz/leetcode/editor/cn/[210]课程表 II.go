package main

//leetcode submit region begin(Prohibit modification and deletion)
func dfsFindOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph_210(numCourses, prerequisites)
	state := make([]int, numCourses)
	hasCycle := false
	var res []int
	for vertex := range graph {
		if state[vertex] == 0 {
			traverseGraph_210(graph, state, &res, &hasCycle, vertex)
		}
	}

	if hasCycle {
		return nil
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func traverseGraph_210(graph [][]int, state []int, res *[]int, hasCycle *bool, vertex int) {
	if *hasCycle {
		return
	}

	state[vertex] = 1
	for _, neighbor := range graph[vertex] {
		if state[neighbor] == 0 {
			traverseGraph_210(graph, state, res, hasCycle, neighbor)
		} else if state[neighbor] == 1 {
			*hasCycle = true
			return
		}
	}

	state[vertex] = 2
	*res = append(*res, vertex)
}

func buildGraph_210(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph, inDegrees := buildGraphAndInDegrees_210(numCourses, prerequisites)

	var q, res []int
	for vertex, inDegree := range inDegrees {
		if inDegree == 0 {
			q = append(q, vertex)
		}
	}

	for len(q) > 0 {
		vertex := q[0]
		q = q[1:]
		res = append(res, vertex)

		for _, neighbor := range graph[vertex] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	if len(res) < numCourses {
		return nil
	}
	return res
}

func buildGraphAndInDegrees_210(numCourses int, prerequisites [][]int) ([][]int, []int) {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		inDegrees[prerequisite[0]]++
	}
	return graph, inDegrees
}
//leetcode submit region end(Prohibit modification and deletion)
