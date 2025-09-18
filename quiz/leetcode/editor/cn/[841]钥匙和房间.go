package main

//leetcode submit region begin(Prohibit modification and deletion)
func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))
	visited[0] = true
	var count int
	q := []int{0}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			room := q[0]
			q = q[1:]
			count++
			for _, key := range rooms[room] {
				if !visited[key] {
					q = append(q, key)
					visited[key] = true
				}
			}
		}
	}
	return count == len(rooms)
}
//leetcode submit region end(Prohibit modification and deletion)
