func containsDuplicate(nums []int) bool {
    hash := make(map[int]bool,len(nums))
    for _,n := range nums {
        if hash[n] {
            return true
        }
        hash[n] = true
    }
    return false
}
