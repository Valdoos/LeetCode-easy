/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    ans := [][]int{}
    helper(&ans,root,0)
    return reverse(ans)
}
func helper(arr *[][]int,node *TreeNode,level int ){
    if node!=nil{
        if len(*arr) <= level{
            *arr = append(*arr,[]int{})
        }
        (*arr)[level] = append((*arr)[level],node.Val)
        helper(arr,node.Left,level+1)
        helper(arr,node.Right,level+1)
    }
}
func reverse(input [][]int) [][]int{
    for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
        input[i], input[j] = input[j], input[i];
    }    
    return input
}
