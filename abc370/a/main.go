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
	var L, R int
	fmt.Scanf("%d %d", &L, &R)

	if L == 1 {
		if R == 1 {
			fmt.Print("Invalid")
		} else {
			fmt.Print("Yes")
		}
	} else {
		if R == 1 {
			fmt.Print("No")
		} else {
			fmt.Print("Invalid")
		}
	}
}
