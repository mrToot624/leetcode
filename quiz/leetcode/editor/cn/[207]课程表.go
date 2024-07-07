package main

//leetcode submit region begin(Prohibit modification and deletion)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph_207(numCourses, prerequisites)
	state := make([]int, numCourses)
	var hasCycle bool
	for vertex := range graph {
		if state[vertex] == 0 {
			traverseGraph_207(graph, state, &hasCycle, vertex)
		}
	}
	return !hasCycle
}

func traverseGraph_207(graph [][]int, state []int, hasCycle *bool, vertex int) {
	if *hasCycle {
		return
	}

	state[vertex] = 1
	for _, neighbor := range graph[vertex] {
		if state[neighbor] == 0 {
			traverseGraph_207(graph, state, hasCycle, neighbor)
		} else if state[neighbor] == 1 {
			*hasCycle = true
			return
		}
	}
	state[vertex] = 2
}

func buildGraph_207(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	return graph
}

func bfsCanFinish(numCourses int, prerequisites [][]int) bool {
	graph, inDegrees := buildGraphAndInDegrees_207(numCourses, prerequisites)

	var q []int
	var count int
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		vertex := q[0]
		q = q[1:]
		count++

		for _, neighbor := range graph[vertex] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	return count == numCourses
}

func buildGraphAndInDegrees_207(numCourses int, prerequisites [][]int) ([][]int, []int) {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		inDegrees[prerequisite[0]]++
	}
	return graph, inDegrees
}
//leetcode submit region end(Prohibit modification and deletion)
