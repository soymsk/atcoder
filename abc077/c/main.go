package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	sc.Buffer(make([]byte, 64*1024), 100000000000)

	sc.Scan()

	var A []int64
	var B []int64
	var C []int64

	inputs := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		A = append(A, v)
	}
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		B = append(B, v)
	}
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		C = append(C, v)
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	// sort.Slice(B, func(i, j int) bool {
	// 	return B[i] < B[j]
	// })
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	var uniqs = 0

	for j := 0; j < n; j++ {
		b := B[j]
		// Search A
		i := sort.Search(n, func(j int) bool { return A[j] < b })

		// Search C
		k := sort.Search(n, func(k int) bool {
			return C[k] > b
		})
		uniqs += (n - i) * (n - k)

	}
	fmt.Printf("%d", uniqs)
}
