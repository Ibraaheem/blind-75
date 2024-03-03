class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        profit = 0
        buy_price = prices[0]
        for price in prices[1:]:
            profit = max(profit, price - buy_price)
            buy_price = min(price, buy_price)
        return profit

        