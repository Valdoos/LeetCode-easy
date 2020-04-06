/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    A := headA
    B := headB
    for A != B{
        A = A.Next
        B = B.Next
        if A == B{
            return A
        }
        if A == nil{
            A = headB
        }
        if B == nil {
           B = headA 
        }
        
    }
    return A
}
