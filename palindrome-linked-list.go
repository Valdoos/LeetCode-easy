/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(head *ListNode) *ListNode{
    var ans *ListNode
    for head != nil {
        next:=head.Next
        head.Next = ans
        ans = head
        head = next
    }
    return ans
}
func isPalindrome(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next!=nil {
        slow=slow.Next
        fast=fast.Next.Next
    }
    fast = head
    slow = reverse(slow)
    for slow != nil {
        if slow.Val != fast.Val {
            return false
        }
        slow = slow.Next
        fast = fast.Next
    }
    return true
}
