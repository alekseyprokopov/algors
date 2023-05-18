package main

import "fmt"

func main() {
	test := []int{}
	test = append(test, nil...)
	fmt.Println(test)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	flag := true
	check(p, q, &flag)
	return flag
}

func check(p *TreeNode, q *TreeNode, flag *bool) {
	if p == nil && q == nil {
		return
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		*flag = false
		return
	}

	check(p.Left, q.Left, flag)
	check(p.Right, q.Right, flag)

}
