package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s) // %dでint型を代入

	N := len(s)

	n := 0
	for i := 0; i < N/2; i++ {
		if s[i] != s[N-i-1] {
			n++
		}

	}
	fmt.Printf("%d", n)
}
