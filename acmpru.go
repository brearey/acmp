package main

import "fmt"

func Acmp1() {
	var a int32
	var b int32

	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a + b)
}

func Acmp2() {
	// 5 -> 15
	var N int
	fmt.Scan(&N)
	var sum int

	if N > 0 {
		for i := 1; i <= N; i++ {
			sum += i
		}
		fmt.Println(sum)
	} else {
		for i := 1; i >= N; i-- {
			sum += i
		}
		fmt.Println(sum)
	}
}
