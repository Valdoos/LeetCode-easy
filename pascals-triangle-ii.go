func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    row := make([]int,rowIndex+1)
    temp := make([]int,rowIndex+1)
    row[0]=1
    for i:=1;i<=rowIndex;i++{
         copy(temp,row)
         row[i]=1
         for j:=1;j<i;j++{
             row[j]=temp[j-1]+temp[j]
         }

    }
    return row
}
