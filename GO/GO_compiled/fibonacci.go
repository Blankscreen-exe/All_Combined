package main

import (
	"fmt"
)

func fibnocci(n int) int {
	// function body.....
	if n <= 1 {
		/* statement(s) will execute if the boolean expression is true */
		return n
	} else {
		/* statement(s) will execute if the boolean expression is false */
		return (fibnocci(n-1) + fibnocci(n-2))
	}
}
func Fibonacci() {
	var a int
	var b int
	fmt.Print("\nEnter the value>")
	fmt.Scan(&a)
	b = fibnocci(a)
	fmt.Println(b)
}
