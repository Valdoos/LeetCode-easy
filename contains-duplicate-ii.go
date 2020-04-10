func containsNearbyDuplicate(nums []int, k int) bool {
    hash := make(map[int]int,len(nums))
    for i, n := range nums {
        if _, ok := hash[n];ok {
            j := hash[n]
            if i - j <= k {
                return true
            } 
        }
        hash[n] = i    
    }
    return false
}
