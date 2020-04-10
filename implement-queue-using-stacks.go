type MyQueue struct {
    head *Node
    back *Node
}

type Node struct{
    Val int
    Next *Node 
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    if this.head == nil {
        this.head = &Node{
            Val: x,
            Next: nil,
        }
        this.back = this.head
    } else {
        if this.back == nil {
            panic("There is no one node")
        }
        this.back.Next = &Node{
            Val: x,
            Next: nil,
        }
        this.back = this.back.Next
    }
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if this.head == nil {
        panic("There is no one node")
    }
    val := this.head.Val
    this.head = this.head.Next
    return val
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if this.head == nil {
        panic("There is no one node")
    }
    return this.head.Val
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.head == nil
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
