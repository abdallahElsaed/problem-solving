var maxProfit = function (prices) {
    let buy = 0;
    let sell = 1;
    let maxProfit = 0
    while (sell < prices.length) {
        if (prices[buy] < prices[sell]) {
            if (maxProfit < prices[sell] - prices[buy]) {
                maxProfit = prices[sell] - prices[buy]
            }

        } else {
            buy = sell // we update it here because all element after the buy is less than buy value. so if the sell value is less than buy, make buy value = sell
        }
        sell += 1 // in any case increase teh sell value by 1
    }
    return maxProfit
};

// 7,1,5,3,6,4 => 5
//7,6,4,3,1 => 0
//10,1,5,6,7,1 => 6
// 1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9 => 9
console.log(maxProfit([1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9]));