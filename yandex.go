package main

import "fmt"

func Triangle() {
	var a int
	var b int
	var c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	
	if (a > 0 && b > 0 && c > 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}