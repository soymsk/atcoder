package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readNumbersInLine() ([]int, []int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 100000000000)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	var A []int
	for i := 0; i < len(inputs); i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		A = append(A, int(v))
	}
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")
	var B []int
	for i := 0; i < len(inputs); i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		B = append(B, int(v))
	}
	return A, B
}

func absDiff(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	A, B := readNumbersInLine()

	// fmt.Printf("%z %z\n", A, B)

	dp := make([][]bool, N)
	dp[0] = []bool{true, true}

	for i := 1; i < N; i++ {
		dp[i] = []bool{false, false}
		if dp[i-1][0] {
			// Can come from previous A
			if absDiff(A[i], A[i-1]) <= K {
				dp[i][0] = true
			}
			// fmt.Printf("%z %z %z", A, B, i)
			if absDiff(B[i], A[i-1]) <= K {
				dp[i][1] = true
			}
		}
		if dp[i-1][1] {
			// Can come from previous B
			if absDiff(A[i], B[i-1]) <= K {
				dp[i][0] = true
			}
			if absDiff(B[i], B[i-1]) <= K {
				dp[i][1] = true
			}
		}

		// fmt.Println("%z %z\n", dp, i)
		if !(dp[i][0] || dp[i][1]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
