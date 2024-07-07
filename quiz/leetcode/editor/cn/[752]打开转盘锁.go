package main

//leetcode submit region begin(Prohibit modification and deletion)
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	reachDeadEnd := make(map[string]bool)
	for _, deadend := range deadends {
		reachDeadEnd[deadend] = true
	}

	var q []string
	q = append(q, "0000")
	if reachDeadEnd["0000"] {
		return -1
	} else {
		reachDeadEnd["0000"] = true
	}

	var step int
	for len(q) > 0 {
		l := len(q)
		step++
		for i := 0; i < l; i++ {
			cur := q[0]
			for j := 0; j < len(cur); j++ {
				up := scaleUp(cur, j)
				if up == target {
					return step
				}
				if !reachDeadEnd[up] {
					q = append(q, up)
					reachDeadEnd[up] = true
				}

				down := scaleDown(cur, j)
				if down == target {
					return step
				}
				if !reachDeadEnd[down] {
					q = append(q, down)
					reachDeadEnd[down] = true
				}
			}
			q = q[1:]
		}
	}

	return -1
}

func scaleUp(s string, index int) string {
	bytes := []byte(s)
	if bytes[index] == byte('9') {
		bytes[index] = byte('0')
	} else {
		bytes[index]++
	}
	return string(bytes)
}

func scaleDown(s string, index int) string {
	bytes := []byte(s)
	if bytes[index] == byte('0') {
		bytes[index] = byte('9')
	} else {
		bytes[index]--
	}
	return string(bytes)
}
//leetcode submit region end(Prohibit modification and deletion)
