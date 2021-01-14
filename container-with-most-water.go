func maxArea(height []int) int {
    var area int
    for i, j := 0, len(height)-1;i<j;{
        if height[i] < height[j] {
            if (j-i)* height[i] > area {
                area = (j-i)* height[i]
            }
            i++
        } else {
           if (j-i)* height[j] > area {
                area = (j-i)* height[j]
            }
            j--
        }
    }
    return area
}

