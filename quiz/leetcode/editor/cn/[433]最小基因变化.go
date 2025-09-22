package main

//leetcode submit region begin(Prohibit modification and deletion)
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	geneSet := make(map[string]bool)
	for _, gene := range bank {
		geneSet[gene] = true
	}

	var cnt int
	q := []string{startGene}
	traveled := map[string]bool{startGene: true}
	nucles := []byte{'A', 'C', 'G', 'T'}
	for len(q) > 0 {
		l := len(q)
		cnt++
		for i := 0; i < l; i++ {
			gene := q[0]
			q = q[1:]
			for j := range gene {
				for _, nucle := range nucles {
					geneByte := []byte(gene)
					if geneByte[j] == nucle {
						continue
					}
					geneByte[j] = nucle
					mutatedGene := string(geneByte)
					if geneSet[mutatedGene] && !traveled[mutatedGene] {
						if mutatedGene == endGene {
							return cnt
						}
						q = append(q, mutatedGene)
						traveled[mutatedGene] = true
					}
				}
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
