func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int{}
    }
    triangle := make([][]int,numRows)
    triangle[0]=[]int{1}
    for i:=1;i<numRows;i++{
        triangle[i]=make([]int,i+1)
        triangle[i][0] = 1
        triangle[i][i] = 1
        for j:=1;j<i;j++{
            triangle[i][j]=triangle[i-1][j]+triangle[i-1][j-1]
        }
    }
    return triangle
}
