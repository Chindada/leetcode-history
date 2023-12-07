package main

// [[1,2],[2,3],[4,2]]
var testCase1 [][]int = [][]int{{1, 2}, {2, 3}, {4, 2}}

// [[1,2],[5,1],[1,3],[1,4]]
var testCase2 [][]int = [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
