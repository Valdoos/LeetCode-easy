/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    ans := 1
    maxDepth(root,&ans)
    return ans-1
}

func maxDepth(root *TreeNode,ans *int) int {
    if root == nil {
        return 0
    }
    L := maxDepth(root.Left, ans)
    R := maxDepth(root.Right, ans)
    *ans = max(*ans,L+R+1)
    return max(L,R)+1
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
