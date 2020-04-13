/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []string{fmt.Sprintf("%d",root.Val)}
    }
    left := binaryTreePaths(root.Left);
    right := binaryTreePaths(root.Right)
    for i,v := range left {
        left[i]=fmt.Sprintf("%d->%s",root.Val,v)
    }
    for i,v := range right {
        right[i]=fmt.Sprintf("%d->%s",root.Val,v)
    }
    return append(left,right...)
}
