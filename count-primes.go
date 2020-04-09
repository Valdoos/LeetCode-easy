func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    notPrime := make([]bool,n)
    sqrt := int(math.Sqrt(float64(n)))
    for i:=2;i<=sqrt;i+=1{
        if notPrime[i] {
            continue
        }
        for j:=i*i;j<n;j+=i {
            notPrime[j]=true
        }
    }
    counter:=0
    for i:=2; i<n; i++ {
        if !notPrime[i] {
            counter++
        }
    }
    return counter
}
