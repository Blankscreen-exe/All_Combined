package main

import (
	"fmt"
	"math"
)

func slope_formula_2(a, b float64) float64 {
	c := math.Pow(a+b, 2)
	return c

}

func Slope2() {
	var n1 float64
	var n2 float64
	fmt.Println("Enter Your First number: ")
	fmt.Scanln(&n1)
	fmt.Println("Enter Second number: ")
	fmt.Scanln(&n2)
	fmt.Printf("Result : %.2f", slope_formula_2(n1, n2))
}
