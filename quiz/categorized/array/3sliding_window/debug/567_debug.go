package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	targetCount, window := make(map[string]int), make(map[string]int)

	for _, c := range s1 {
		targetCount[string(c)]++
	}

	runes := []rune(s2)
	var satisfied int
	for right < len(s2) {
		toAdd := string(runes[right])
		if _, ok := targetCount[toAdd]; ok {
			window[toAdd]++
			if window[toAdd] == targetCount[toAdd] {
				satisfied++
			}
		}
		right++

		for right-left >= len(s1) {
			if satisfied == len(targetCount) {
				fmt.Printf("found answer %s, start=%d(%s), end=%d(%s)\n", s2[left:right], left, string(s2[left]), right-1, string(s2[right-1]))
				return true
			}

			toRemove := string(runes[left])
			if _, ok := targetCount[toRemove]; ok {
				if window[toRemove] == targetCount[toRemove] {
					satisfied--
				}
				window[toRemove]--
			}
			left++
		}
	}

	fmt.Println(s2[left-1 : right])

	return false
}

func main() {
	//s1 := "trinitrophenylmethylnitramine"
	//s2 := "dinitrophenylhydrazinetrinitrophenylmethylnitramine"

	//s2 := "dinitrophenylhydrazinetrinitrophenylmethylnitramine"
	//s1 := "itramine"

	s1 := "itramine"
	s2 := "itrophenylmethylnitramine"
	fmt.Println(checkInclusion(s1, s2))
}
