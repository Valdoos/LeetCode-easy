func merge(nums1 []int, n int, nums2 []int, m int)  {
    n--
    m--
    i := m+n+1
    for n >=0 && m>=0 {
        if nums1[n] > nums2[m]{
            nums1[i] = nums1[n]
            n--
            i--
        } else {
            nums1[i] = nums2[m]
            m--
            i--
        }
    }
    for m >= 0 {
        nums1[i] = nums2[m]
        i--
        m--
    }
    return 
}
