package main

import "fmt"

func sumInts(s []int) int {
	sum := 0

	for _, i := range s {
		sum = sum + i
	}

	return sum
}

func removeZeroVals(s []int) []int {
	ret := make([]int, 0)

	for _, i := range s {
		if i == 0 {
			continue
		}
		ret = append(ret, i)
	}

	return ret
}

func removeIndex(s []int, index int) []int {
	if index >= len(s) {
		return s
	}

	cp := make([]int, len(s))
	copy(cp, s)

	return append(cp[:index], cp[index+1:]...)
}

/*
 * Given an array of integers nums and a positive integer k, find all unique subset partitions where
 * the sum of the set is equal to k.
 */
func findSubsetPartitionsWithSum(nums []int, k int) [][]int {
	subSets := [][]int{}
	var findRecursive func(set []int, sum int, idx int)

	findRecursive = func(subSet []int, sum int, idx int) {
		if sum == k {
			subSets = append(subSets, subSet)
			return
		}

		if sum < k || idx >= len(subSet) {
			return
		}

		// branch 1: continue with removing the item at the current index
		findRecursive(removeIndex(subSet, idx), sum-subSet[idx], idx)

		// branch 2: increment the index and continue
		findRecursive(subSet, sum, idx+1)
	}

	findRecursive(removeZeroVals(nums), sumInts(nums), 0)
	fmt.Println(subSets)
	return subSets
}

func main() {
	set := []int{4, 3, 2, 3, 5, 2, 1}
	k := 7
	findSubsetPartitionsWithSum(set, k)
}
