package main

import (
	"fmt"
	"math"
	"strconv"
)

func sumDigits(num int64) int64 {
	sum := int64(0)
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {

	// 1文字ずつデータ型を指定して受け取る
	var K int64         // int型の変数を宣言
	fmt.Scanf("%d", &K) // %dでint型を代入
	multiplier := int64(1)

	K_len := len(strconv.Itoa(int(K)))

	min := int64(math.MaxInt64)
	// for float64(K*multiplier) < math.Min(math.Pow(10, float64(min)), float64(K)*math.Pow(10, float64(K_len*2))) {
	for K*multiplier < K*int64(math.Pow(10, float64(K_len*2))) {
		v := sumDigits(K * multiplier)
		fmt.Printf("%d %d %d %d %d\n", K, multiplier, K*multiplier, v, min)
		if min > v {
			min = v
		}
		multiplier++
	}

	fmt.Printf("%d", min)
}
