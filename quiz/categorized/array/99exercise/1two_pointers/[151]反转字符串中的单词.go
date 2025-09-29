package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	bytes, spaces := reverseAndTrim(s)

	start := 0
	for _, spaceIndex := range spaces {
		end := spaceIndex - 1
		for start < end {
			bytes[start], bytes[end] = bytes[end], bytes[start]
			start++
			end--
		}
		start = spaceIndex + 1
	}
	return string(bytes)
}

func reverseAndTrim(s string) ([]byte, []int) {
	var builder strings.Builder
	var spaces []int
	var meetWord bool
	for i := len(s) - 1; i >= 0; i-- {
		if !needTrim(s, i, meetWord) {
			if isSpace(s, i) {
				spaces = append(spaces, builder.Len())
			}
			builder.WriteByte(s[i])
			meetWord = true
		}
	}
	return []byte(builder.String()), append(spaces, builder.Len())
}

func needTrim(s string, i int, meetWord bool) bool {
	if !isSpace(s, i) {
		return false
	}

	if i == len(s)-1 || i == 0 {
		return true
	}

	return s[i-1] == ' ' || !meetWord
}

func isSpace(s string, i int) bool {
	return s[i] == ' '
}
//leetcode submit region end(Prohibit modification and deletion)
