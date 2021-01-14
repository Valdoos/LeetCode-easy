/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
            it := l2 
            for it!=nil && it.Val >=10{
                modify(it)
                it=it.Next
            }
            return l2
    }
    if l2 == nil {
            it := l1 
            for it!=nil && it.Val >=10{
                modify(it)
                it=it.Next
            }
            return l1
    }
    l1.Val += l2.Val
    modify(l1)
    l1.Next = addTwoNumbers(l1.Next,l2.Next)
    return l1
}

func modify(l *ListNode) {
        if l.Val >= 10{
        l.Val %= 10
        if l.Next == nil {
             l.Next = &ListNode{
                        Val: 1,
                        Next: nil,
            } 
        } else {
            l.Next.Val++
        }
    }
}
