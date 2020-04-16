import "math"
func thirdMax(nums []int) int {
    max := math.MinInt64
    secondMax := math.MinInt64
    thirdMax := math.MinInt64
    count := 0
    for _, k := range nums {
        if k > max {
            thirdMax, secondMax, max = secondMax, max, k     
            count++
        } else if k > secondMax && k < max {
            thirdMax, secondMax = secondMax, k
            count++
        } else if k > thirdMax && k < secondMax {
            thirdMax = k
            count++
        }
    }
    if count > 2 {
        return thirdMax
    }
    return max
}
