func missingNumber(nums []int) int {
    n := len(nums)+1
    ans := n*(n-1)/2
    for _, k := range nums {
        ans -=k
    }
    return ans
}

func missingNumber(nums []int) int {
    ans := len(nums)
    for i, k := range nums {
        ans^=i^k
    }
    return ans
}
