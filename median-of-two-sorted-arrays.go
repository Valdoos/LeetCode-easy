func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    b := (len(nums1)+len(nums2)+1)%2
    k := (len(nums1) + len(nums2))/2
    ans := 0.0
    i,j:=0,0
    for i<len(nums1) && j<len(nums2) && k>b {
        if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
        k--
    }
    if i<len(nums1) && k>b {
        i+=k-b
        k=b
    }
    if j<len(nums2) && k>b {
        j+=k-b
        k=b
    }
    for k >= 0 {
        if (i < len(nums1) && j >= len(nums2)) || (i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j]) {
            ans+=float64(nums1[i])
            i++
        } else {
            ans+=float64(nums2[j])
            j++
        }
        k--
    }
    if b == 1 {
        ans/=2
    }
    return ans
}
