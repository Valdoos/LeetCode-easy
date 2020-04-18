/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    return findSum(root,0,sum,map[int]int{0:1})
}

func findSum(root *TreeNode,current,target int, hash map[int]int) int {
    if root == nil {
        return 0
    }
    current+=root.Val
    sum := hash[current-target]
    hash[current]++
    sum += findSum(root.Left,current,target, hash)
    sum += findSum(root.Right,current,target,hash)
    hash[current]--
    return sum
}
