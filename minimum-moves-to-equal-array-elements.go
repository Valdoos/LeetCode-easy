func minMoves(nums []int) int {
    min := nums[0]
    sum := nums[0]
    for i:=1;i<len(nums);i++ {
        sum+=nums[i]
        if min > nums[i] {
            min = nums[i]
        }
    }
    return sum - len(nums)*min
}
