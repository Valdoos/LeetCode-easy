func rotate(nums []int, k int)  {
    l := len(nums)
    if l <= 1 {
        return
    }
    if k >=l {
        k %= l
    }
    if k == 0 {
        return
    }

    copy(nums,append(nums[l-k:],nums[:l-k]...))
}
