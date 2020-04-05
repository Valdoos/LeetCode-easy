type MinStack struct {
    StackHead *Node
    MinHead *Node
}

type Node struct{
    Val int
    Previous *Node
    Num int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        StackHead : nil,
        MinHead : nil,
    }
}


func (this *MinStack) Push(x int)  {
    this.StackHead = &Node{
        Val: x,
        Previous: this.StackHead,
    }
    if this.MinHead == nil {
        this.MinHead = &Node {
            Val: x,
            Previous: nil,
            Num:1,
        }
    } else if this.MinHead.Val > x {
        this.MinHead = &Node {
            Val: x,
            Previous: this.MinHead,
            Num: 1,
        }
    } else if this.MinHead.Val == x {
        this.MinHead.Num++ 
    }
}


func (this *MinStack) Pop()  {
    if this.StackHead == nil {
        return
    }
    if this.StackHead.Val == this.MinHead.Val {
        this.MinHead.Num--
    }
    if this.MinHead.Num == 0 {
        this.MinHead = this.MinHead.Previous
    }
    this.StackHead = this.StackHead.Previous
}


func (this *MinStack) Top() int {
    if this.StackHead == nil {
        panic("error")
    }
    return this.StackHead.Val
}


func (this *MinStack) GetMin() int {
    if this.MinHead == nil {
        panic("error")
    }
   return this.MinHead.Val
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
