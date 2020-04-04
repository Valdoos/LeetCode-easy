/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _,ok := check(root,0)
    return ok
}

func check(root *TreeNode, level int) (int, bool) {
    if root == nil {
        return level-1, true
    }
    left, okLeft := check(root.Left,level+1)
    right, okRight := check(root.Right,level+1)
    if !okLeft || !okRight {
        return level, false
    }
    if math.Abs(float64(left-right))>1 {
        return level, false
    } 
    return max(left,right), true
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
