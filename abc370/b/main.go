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

func main() {
	N := read_int()
	atoms := make([][]int, N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 100000000000)

	for i := 0; i < N; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		atoms[i] = make([]int, N)
		for j := 0; j < len(inputs); j++ {
			v, _ := strconv.ParseInt(inputs[j], 10, 64)
			atoms[i][j] = int(v)
		}
	}

	current := 1
	for k := 1; k <= N; k++ {
		if current < k {
			current = atoms[k-1][current-1]
		} else {
			current = atoms[current-1][k-1]
		}
	}

	fmt.Printf("%d", current)
}
