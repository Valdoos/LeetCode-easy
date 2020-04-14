func moveZeroes(nums []int)  {
    next := 0
    for _, n := range nums {
        if n != 0 {
            nums[next]=n
            next++
        }
    }
    for l:=len(nums); next<l; next++{
        nums[next] = 0
    }
}

func moveZeroes(nums []int)  {
    last := 0
    for i, n := range nums {
        if n != 0 {
            nums[last], nums[i] = n, nums[last]
            last++
        }
    }
}
