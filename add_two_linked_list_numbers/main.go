package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Takes a reversed linked list and coverts it to int (e.g. 2 -> 4 -> 3 = 342)
func linkedListToInt(l *ListNode) (i int) {
	digit, base := 1, 10
	for l != nil {
		i = i + (l.Val * digit)
		digit = digit * base
		l = l.Next
	}
	return
}

// Splits an int into separate digits (e.g. 807 = [7, 0 8])
func splitDigits(i int) []int {
	if i == 0 {
		return []int{0}
	}

	digitS := make([]int, 0)
	for i != 0 {
		tmp := i % 10
		digitS = append(digitS, tmp)
		i = i / 10
	}

	return digitS
}

func newNode() *ListNode {
	return &ListNode{Val: 0}
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func getNext(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

// Takes two reversed list nodes, adds the corresponding ints, and returns a new
// linked list node that represents the Sum.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	ref := newNode()
	head := ref
	loop := true

	for loop {
		sum := getVal(l1) + getVal(l2)
		if carry {
			sum++
		}
		if sum >= 10 {
			carry = true
			ref.Val = sum % 10
		} else {
			carry = false
			ref.Val = sum
		}
		l1 = getNext(l1)
		l2 = getNext(l2)
		loop = l1 != nil || l2 != nil || carry
		if loop {
			ref.Next = newNode()
			ref = ref.Next
		}
	}

	return head
}

// takes a variadic number of ints and returns a linked list with the ints in reverse order
func newLinkedListInt(num int) (root *ListNode) {
	digits := splitDigits(num)

	var ref *ListNode

	for _, i := range digits {
		n := &ListNode{Val: i}

		if root == nil {
			root = n
			ref = n
			continue
		}

		ref.Next = n
		ref = n
	}

	return root
}

func main() {
	l1 := newLinkedListInt(342)
	l2 := newLinkedListInt(465)

	val := addTwoNumbers(l1, l2)
	fmt.Println(linkedListToInt(val))
}
