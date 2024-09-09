package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main_0() {

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
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	var uniqs = 0
	var low_bound_b, low_bound_c int = 0, 0
	for i := 0; i < n; i++ {
		for j := low_bound_b; j < n; j++ {
			if A[i] < B[j] {
				for k := low_bound_c; k < n; k++ {
					if B[j] < C[k] {
						uniqs += n - k
						break
					} else {
						low_bound_c = k
					}
				}
			} else if A[i] > B[j] {
				low_bound_b = j
			}
		}
		low_bound_c = 0
	}

	fmt.Printf("%d", uniqs)
}
