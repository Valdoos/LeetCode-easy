func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    var res, sum = nums[0],0
    for i := range nums {
        sum +=nums[i]
        if res < sum {
            res = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return res
}
