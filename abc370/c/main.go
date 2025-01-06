package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_int() int {
	var N int
	fmt.Scanf("%d", &N)
	return N
}

func read_numbers() []int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 100000000000)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	var numbers []int
	for i := 0; i < len(inputs); i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 64)
		numbers = append(numbers, int(v))
	}
	return numbers
}

func read_lines(N int) [][]string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 100000000000)

	var lines [][]string
	for i := 0; i < N; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		lines = append(lines, inputs)
	}
	return lines
}

func sumDigit(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func findMinIndex(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, -1
	}

	minVal := nums[0]
	minIndex := 0

	for i, num := range nums {
		if num < minVal {
			minVal = num
			minIndex = i
		}
	}
	return minVal, minIndex
}

func main() {
	var S, T string
	var X []string
	fmt.Scanf("%s", &S)
	fmt.Scanf("%s", &T)

	var diffs []int
	for i := 0; i < len(S); i++ {
		// v := []int{int(S[i]) - int(T[i]), i}
		diffs = append(diffs, int(T[i])-int(S[i]))
	}

	// minVal := -100000
	// for minVal < 9999 {
	// 	minVal, i := findMinIndex(diffs)

	// 	if minVal != 0 {
	// 		S = S[0:i] + T[i:i+1] + S[i+1:]
	// 		X = append(X, S)
	// 		diffs[i] = 100000
	// 	}

	// }

	M := 0
	for i, num := range diffs {
		if num < 0 {
			S = S[0:i] + T[i:i+1] + S[i+1:]
			X = append(X, S)
			M++
		}
	}
	for i := len(diffs) - 1; i >= 0; i-- {
		num := diffs[i]
		if num > 0 {
			S = S[0:i] + T[i:i+1] + S[i+1:]
			X = append(X, S)
			M++
		}
	}
	fmt.Println(M)
	for _, v := range X {
		fmt.Printf("%s\n", v)
	}
}
