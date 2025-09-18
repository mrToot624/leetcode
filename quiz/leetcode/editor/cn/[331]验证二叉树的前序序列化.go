package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func isValidSerialization(preorder string) bool {
	preList := strings.Split(preorder, ",")
	s := []int{1}
	for _, node := range preList {
		if node != "#" {
			if !decreaseParentEdge(&s){
				return false
			}
			s = append(s, 2)
		}
		if node == "#" {
			if !decreaseParentEdge(&s){
				return false
			}
		}
	}
	return len(s) == 0
}

func decreaseParentEdge(s *[]int) bool {
	if len(*s) > 0 {
		(*s)[len(*s)-1]--
		if (*s)[len(*s)-1] == 0 {
			*s = (*s)[:len(*s)-1]
		}
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
