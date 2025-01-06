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
	var H, W, Q int
	fmt.Scanf("%d %d %d", &H, &W, &Q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 100000000000)

	var orders [][]int
	for i := 0; i < Q; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		R, _ := strconv.ParseInt(inputs[0], 10, 64)
		C, _ := strconv.ParseInt(inputs[1], 10, 64)
		v := []int{int(R) - 1, int(C) - 1}
		orders = append(orders, v)
	}

	var matrix = make([][]int, H)
	for i := 0; i < H; i++ {
		matrix[i] = make([]int, W)
		for j := 0; j < W; j++ {
			matrix[i][j] = 1
		}
	}

	remains := H * W

	for _, v := range orders {
		i := v[0]
		j := v[1]
		if matrix[i][j] == 1 {
			matrix[i][j] = 0
			// fmt.Printf("%d %d found\n", i, j)
			remains--
		} else {
			if search_and_destroy(matrix, i, j, -1, 0) {
				remains--
			}
			if search_and_destroy(matrix, i, j, 1, 0) {
				remains--
			}
			if search_and_destroy(matrix, i, j, 0, -1) {
				remains--
			}
			if search_and_destroy(matrix, i, j, 0, 1) {
				remains--
			}
		}
	}

	fmt.Print(remains)
}

func search_and_destroy(matrix [][]int, i int, j int, di int, dj int) bool {
	H := len(matrix)
	W := len(matrix[0])
	if i < 0 || i > H-1 || j < 0 || j > W-1 {
		// not found
		return false
	} else if matrix[i][j] == 1 {
		matrix[i][j] = 0
		// fmt.Printf("%d %d found\n", i, j)
		return true
	} else {
		return search_and_destroy(matrix, i+di, j+dj, di, dj)
	}
}
