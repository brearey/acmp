package main

import (
	"fmt"
)

func acmp1() {
	var a int32
	var b int32

	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a + b)
}

func acmp2() {
	// 5 -> 15
	var N int
	fmt.Scan(&N)

	var sum int
	for i := 1; i <= N; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	acmp2()
}