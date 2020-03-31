func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    return search(nums, target,0,len(nums)-1)
}
func search(nums []int, target int, b1 int, b2 int) int {
    index := (b2+b1)/2
    medium := nums[index]
    if medium == target{
        return index
    } else if medium > target {
        if b1==index{
            return b1
        }
        return search(nums, target, b1, index-1)
    } else {
        if b2==index {
            return index+1
        }
         return search(nums, target, index+1, b2)
    }
}
