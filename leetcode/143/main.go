/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getMid(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
    }
    return prev
}

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    mid := getMid(head)
    secondHalf := mid.Next
    mid.Next = nil 

    rev := reverse(secondHalf)

    p1 := head
    p2 := rev
    for p2 != nil {
        tmp1 := p1.Next
        tmp2 := p2.Next
        
        p1.Next = p2
        p2.Next = tmp1
        
        p1 = tmp1
        p2 = tmp2
    }
}
