package main

import (
	"fmt"
	"math"
)

func main() {

	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入
	ans := math.Floor(math.Sqrt(float64(n)))

	fmt.Printf("%d", int(math.Pow(ans, 2)))
}
