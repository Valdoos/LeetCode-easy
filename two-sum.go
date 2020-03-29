func twoSum(nums []int, target int) []int {
    myMap := make(map[int]int)
    for i,n := range nums {
        if _,ok := myMap[n]; ok{
            return []int{myMap[n],i}
        }
        myMap[target-n] = i
    }
    return []int{}
}
