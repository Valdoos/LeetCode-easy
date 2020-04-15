import "math"
func isPowerOfFour(n int) bool {
    lg := math.Log2(float64(n))/2
    return lg-float64(int(lg))==0
}
