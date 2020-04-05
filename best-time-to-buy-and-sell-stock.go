func maxProfit(prices []int) int {
    if len(prices)==0{
        return 0
    }
    var min int = prices[0]
    var profit int
    for _,p:=range(prices) {
        if p < min{
            min=p
        }else if p-min > profit {
            profit = p-min
        }
    }
    return profit
}
