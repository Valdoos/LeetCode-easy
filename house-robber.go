func rob(nums []int) int {
    var a,b int
    for _, v := range nums {
        a, b = max(a,b),v+a
    }
    return max(a,b)
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
