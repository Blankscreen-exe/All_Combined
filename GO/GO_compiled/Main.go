package main

import "fmt"

func main() {
	fmt.Println("Select from the options what u want to use:")
	fmt.Println("1.Trignometric Identities ")
	fmt.Println("2.Taylor Expansion")
	fmt.Println("3.Binomial Coefficient")
	fmt.Println("4.Loan Payment")
	fmt.Println("5.Gaussian")
	fmt.Println("6.Fourier Series")
	fmt.Println("7.Distance")
	fmt.Println("8.Slope Formula 01")
	fmt.Println("9.Slope Formula 02")
	fmt.Println("10.Slope Formula 03")
	fmt.Println("11.Quadratic Formula")
	fmt.Println("12.Laplace")
	fmt.Println("13.Prime Number")
	fmt.Println("14.Sum of Numbers")
	fmt.Println("15.Standard Deviation")
	fmt.Println("16.Relativity")
	fmt.Println("17.Euler RK1")
	fmt.Println("18.Fibonacci")

	fmt.Println("Select Any:")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		Trignometricidentity()
	case 2:
		TaylorExpansion()
	case 3:
		Binomial()
	case 4:
		LoanPayment()
	case 5:
		Gaussain()
	case 6:
		Fourierseries()
	case 7:
		Distance()
	case 8:
		Slope1()
	case 9:
		Slope2()
	case 10:
		Slope3()
	case 11:
		Quadraticformula()
	case 12:
		Laplacetransform()
	case 13:
		CheckPrime()
	case 14:
		Sum()
	case 15:
		Standarddeviation()
	case 16:
		Relativity()
	case 17:
		Euler()
	case 18:
		Fibonacci()
	default:
		fmt.Println("Invalid Input")
	}

}
