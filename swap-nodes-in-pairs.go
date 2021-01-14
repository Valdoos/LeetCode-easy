/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "fmt"

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    } else if head.Next == nil {
        return head
    }
    var k *ListNode = head
    head = head.Next
    if head.Next == nil {
        k.Next = nil
        head.Next = k
        return head
    } else if head.Next.Next == nil {
        k.Next = head.Next
        head.Next = k
        return head
    }
    k.Next = swapPairs(head.Next)
    head.Next = k
    return head
}
