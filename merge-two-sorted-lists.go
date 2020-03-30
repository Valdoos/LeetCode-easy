/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l2 == nil {
        return l1
    }
    if l1 == nil {
        return l2
    }
    var answer *ListNode
    if l1.Val < l2.Val{
        answer=l1
        l1=l1.Next
    } else {
        answer=l2
        l2=l2.Next
    }
    var p *ListNode = answer
    for l1!=nil && l2!=nil {
        if l1.Val < l2.Val{
            answer.Next=l1
            l1=l1.Next
        } else {
            answer.Next=l2
            l2=l2.Next
        }
        answer=answer.Next
    }
    if l1!=nil{
        answer.Next=l1
    } else {
        answer.Next=l2
    }
    return p
}
