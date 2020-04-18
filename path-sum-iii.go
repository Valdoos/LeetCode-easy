/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    return findSum(root,sum,false)
}

func findSum(root *TreeNode, sum int, used bool) int {
    if root == nil {
        return 0
    }
    paths := findSum(root.Left, sum - root.Val, true) + findSum(root.Right, sum - root.Val, true)
    if !used {
        paths += findSum(root.Left, sum, false) + findSum(root.Right,sum, false)
    }
    if root.Val == sum {
        return paths+1
    }
    return paths
}
