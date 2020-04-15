func intersect(nums1 []int, nums2 []int) []int {
    p1:=&nums1
    p2:=&nums2
    if len(nums1) > len(nums2) {
        p1, p2 = p2, p1
    }
    hash := make(map[int]int,len(*p1))
    ans := make([]int,0,len(*p1))
    for _,n := range (*p1) {
        hash[n]++
    }
    for _,n := range (*p2) {
        if _,ok := hash[n];ok{
            ans = append(ans,n)
            hash[n]--
            if hash[n]==0{
                delete(hash,n)
            }
        }
    }
    return ans
}
