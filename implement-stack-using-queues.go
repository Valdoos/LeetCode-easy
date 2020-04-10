type MyStack struct {
    head *Node
}

type Node struct{
    Val int
    Next *Node 
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.head = &Node{
        Val: x,
        Next: this.head,
    }
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    current := this.head
    if current == nil {
        panic("There is no one node")
    }
    this.head = current.Next
    return current.Val
}


/** Get the top element. */
func (this *MyStack) Top() int {
    if this.head == nil {
        panic("There is no one node")
    }
    return this.head.Val
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.head == nil
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
