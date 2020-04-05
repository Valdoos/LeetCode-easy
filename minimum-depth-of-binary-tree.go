/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := minDepth(root.Left)
    rightDepth := minDepth(root.Right)
    if leftDepth == 0 {
        return rightDepth+1
    } else if rightDepth == 0{
        return leftDepth+1
    }
    return min(leftDepth,rightDepth)+1
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}
