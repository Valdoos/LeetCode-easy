/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    middle := len(nums)/2
    head := TreeNode{Val: nums[middle] }
    head.Left = sortedArrayToBST(nums[:middle])
    head.Right = sortedArrayToBST(nums[middle+1:])
    return &head
}
