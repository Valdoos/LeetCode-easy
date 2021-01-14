/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || head.Next == nil{
        return nil
    }
    i:=0
    root := head
    slow := head
    for i < n {
        head = head.Next
        i++
    }
    if head != nil {
        for head.Next != nil {
            head = head.Next
            slow = slow.Next
        }
        slow.Next = slow.Next.Next
    } else {
        root = slow.Next
    }
    return root
}
