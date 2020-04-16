/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return sumOfLeftLeavesHelper(root,false)
}

func sumOfLeftLeavesHelper(root *TreeNode,left bool) int {
    if root == nil {
        return 0
    }
    existLeft := root.Left != nil
    existRight := root.Right != nil
    var sumLeft, sumRight int
    if existLeft {
        sumLeft = sumOfLeftLeavesHelper(root.Left,true)
    }
    if existRight {
        sumRight = sumOfLeftLeavesHelper(root.Right,false)
    } else if !existLeft {
        if left {
            return root.Val
        } else {
            return 0
        }
    }
    return sumLeft+sumRight
}


func sumOfLeftLeaves(root *TreeNode) int {
    return sumOfLeftLeavesHelper(root,false)
}

func sumOfLeftLeavesHelper(root *TreeNode,left bool) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        if left {
            return root.Val
        }
        return 0
    }
    return sumOfLeftLeavesHelper(root.Left,true) + sumOfLeftLeavesHelper(root.Right,false)
}
