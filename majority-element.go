func majorityElement(nums []int) int {
    count := 0
    major := nums[0]
    for _,n :=range nums{
        if count == 0{
            major = n
        }
        if major == n {
            count++
        } else {
            count--
        }
    }
    return major
}
