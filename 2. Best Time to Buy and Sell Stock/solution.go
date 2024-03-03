package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	left, right, max := 0, 1, 0

	for right < len(prices) {
		profit := prices[right] - prices[left]
		if prices[left] < prices[right] {
			max = int(math.Max(float64(max), float64(profit)))
		} else {
			left = right
		}
		right += 1
	}
	return max
}