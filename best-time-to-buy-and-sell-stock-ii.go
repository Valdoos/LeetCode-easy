func maxProfit(prices []int) int {
    if len(prices)==0 {
        return 0
    }
    profit := 0
    last := prices[0]
    for _, p := range(prices){
        if p-last > 0 {
            profit += p-last
            last = p
        } else if p < last{
        last = p
        }
    }
    return profit 
}
