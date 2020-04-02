/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetricTree(root.Left,root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        if q == nil {
            return true
        } else {
            return false
        }
    } else if q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    } else {
        return isSymmetricTree(p.Left,q.Right) && isSymmetricTree(p.Right,q.Left)
    }
}
