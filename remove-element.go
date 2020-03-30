func removeElement(nums []int, val int) int {
    var n = len(nums)
    if n == 0 {
        return 0
    }
    var i = 0 
    for j := 0; j<n; j++ {
        if nums[j] != val {
            nums[i]=nums[j]
            i++
        }
    }
    return i
}
