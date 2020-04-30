package main

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(x)
}

// Given an array where elements are sorted in ascending order,
// convert it to a height balanced BST
// For this problem, a height-balanced binary tree is defined as a binary tree
// in which the depth of the two subtrees of every node never differ by more than 1.
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2

	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[0:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}
