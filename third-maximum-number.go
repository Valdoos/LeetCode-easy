import "math"
func thirdMax(nums []int) int {
    max := int64(math.MinInt64)
    secondMax := max
    thirdMax := max
    ok := 0
    for _, num := range nums {
        k := int64(num)
        if k > max {
            thirdMax = secondMax
            secondMax = max
            max = k        
            ok++
        } else if k > secondMax && max!=k{
            thirdMax = secondMax
            secondMax = k
            ok++
        } else if k > thirdMax && secondMax!=k && max!=k{
            thirdMax = k
            ok++
        }
    }
    if ok >= 3{
        return int(thirdMax)
    }
    return int(max)
}
