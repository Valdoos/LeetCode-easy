func intersection(nums1 []int, nums2 []int) []int {
    hash := make(map[int]struct{},len(nums1))
    ans := make([]int,0,len(nums1))
    for _,n := range nums1 {
        hash[n]=struct{}{}
    }
    for _,n := range nums2 {
        if _,ok := hash[n];ok{
           ans = append(ans,n)
            delete(hash,n)
        }
    }
    return ans
}
