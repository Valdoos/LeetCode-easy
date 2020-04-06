func twoSum(numbers []int, target int) []int {
    i:=0
    j:=len(numbers)-1
    for i<j{
        sum :=numbers[i]+numbers[j]
        if sum==target{
            break
        } else if sum>target {
            j--
        } else {
            i++
        }
    }
    return []int{i+1,j+1}
}
