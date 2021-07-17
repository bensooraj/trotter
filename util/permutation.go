package util

import (
	"gonum.org/v1/gonum/stat/combin"
)

func GetPermutations(n, k int) [][]int {
	list := combin.Permutations(n, k)
	return list
}
