package main

import "math"

func maxProfit(price []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for i := 0; i < len(price); i++ {
		if price[i] < minPrice {
			minPrice = price[i]
		} else if price[i]-minPrice > maxProfit {
			maxProfit = price[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	ouput := maxProfit(input)
	println(ouput)
}
