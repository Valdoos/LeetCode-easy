type NumArray struct {
    sum []int
}


func Constructor(nums []int) NumArray {
    l := len(nums)
    if l == 0 {
        return NumArray{
            sum : []int{},
        }
    }
    sum := make([]int,l+1)
    for i:=2; i<l+1; i++ {
        sum[i-1]+=nums[i-2]
        sum[i]+=sum[i-1]
    }
    sum[l]+=nums[l-1]
    return NumArray{
        sum : sum,
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.sum[j+1]-this.sum[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
