package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	var s []string
	for i := 1; i < len(paths); i++ {
		if paths[i] == "" || paths[i] == "." {
			continue
		}
		if paths[i] == ".." {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
			continue
		}
		s = append(s, paths[i])
	}
	return "/" + strings.Join(s, "/")
}

//leetcode submit region end(Prohibit modification and deletion)
