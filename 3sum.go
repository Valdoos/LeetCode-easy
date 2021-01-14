import "sort"
func threeSum(nums []int) [][]int {
    var ans [][]int
    sort.Ints(nums)
    l := len(nums)
    var last int
    for i,k:= range nums {
        if i!= 0 && k == last {
            continue
        }
        last = k
        l1:=i+1
        l2:=l-1
        for l1 < l2{
            n1 := nums[l1]
            n2 := nums[l2]
            if n1+n2 == -k {
                ans = append(ans,[]int{k,n1,n2})
                l1++
                for l1 < l && n1 == nums[l1] {
                    l1++
                }
                l2--
                for l2 > l1 && n2 == nums[l2] {
                    l2--
                }
            } else if nums[l1]+nums[l2] < -k {
                l1++
            } else {
                l2--
            }
        }
    }
    return ans
}
