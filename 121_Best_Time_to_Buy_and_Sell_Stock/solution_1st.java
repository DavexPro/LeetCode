public class Solution {
    public int maxProfit(int[] prices) {
        if(prices.length<=1)return 0;
        int min_price, max_profit, profit;
        min_price = prices[0];max_profit = 0;profit = 0;
        for(int i=0;i<prices.length;i++){
            if(min_price > prices[i]) min_price=prices[i];
            profit = prices[i] - min_price;
            if(max_profit < profit) max_profit=profit;
        }
        return max_profit;
    }
}
