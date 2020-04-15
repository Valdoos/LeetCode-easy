import "math"
func isPowerOfFour(n int) bool {
    lg := math.Log2(float64(n))/2
    return lg-float64(int(lg))==0
}

func isPowerOfFour(n int) bool {
    return n>0 && 1073741824 % n == 0 && (n%10 == 4 || n%10 == 6 || n==1 )
}

