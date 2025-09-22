package main

// Union Find
type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, n}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot != qRoot {
		uf.parent[pRoot] = qRoot
	}
	uf.count--
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFind) Count() int {
	return uf.count
}
