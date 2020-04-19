func numberOfBoomerangs(points [][]int) int {
    boomerangs := 0 
    for i, p1 :=range points {
        hash := make(map[int]int,len(points))
        for j, p2 := range points {
            if i == j {
                continue
            }
            distance := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])* (p1[1]-p2[1])
            if n, exist := hash[distance]; exist{
                boomerangs+=2*n
                hash[distance]++
            } else {
                hash[distance]=1
            }
            
        }
    }
    return boomerangs
}
