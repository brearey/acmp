package main

import "fmt"

func Triangle() {
	var a int
	var b int
	var c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	aIsGood := a < b+c
	bIsGood := b < a+c
	cIsGood := c < a+b

	if aIsGood && bIsGood && cIsGood {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
