/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
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
    }
    
    var pl int8
    var pr int8
    var ql int8
    var qr int8
    var sum int8
    
    if p.Left != nil {
        pl = 1
    }
    if q.Left != nil {
        ql = 1
    }
    if p.Right != nil {
        pr = 1
    }
    if q.Right != nil {
        qr = 1
    }
    sum = pl+qr+ql+qr
    if sum == 0 {
        return true
    } else if sum%2 == 1 {
        return false
    } else if sum == 2 {
        if pl == 1 {
            if ql == 1 && p.Left.Val == q.Left.Val{
                 return isSameTree(p.Left,q.Left)
            }
        }
         if pr == 1 {
             if qr == 1  && p.Right.Val == q.Right.Val {
                 return isSameTree(p.Right,q.Right)
            }
        }
    } else {  
        if p.Left.Val == q.Left.Val && p.Right.Val == q.Right.Val {
            return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
        }
    }
    return false
}
