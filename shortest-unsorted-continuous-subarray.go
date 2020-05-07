func findUnsortedSubarray(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    first := 0
    last := 0
    b := true
    for i:=1;i<len(nums);i++ {
        if nums[i] < nums[i-1] {
            last = i
            b = false
            for first > 0 && nums[i] < nums[first-1] {
                first--
            }
            nums[i] = nums[i-1]
        } else if b{
            first++
        }
    }
    ans := last-first+1
    if ans < 0 {
        return 0
    }
    return ans
}
