func removeDuplicates(nums []int) int {
    var n = len(nums)
    if n <= 1 {
        return n
    }
    var i = 0
    for j:=1; j<n; j++{
        if nums[i]!=nums[j] {
            i++
            nums[i]=nums[j]
        }
    }
    return i+1
}
