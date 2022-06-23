package main

import (
	"fmt"
	"math"
)

func relativity(Mass float64, velocity float64) float64 {
	c := 299792458.0
	generalRelativity := float64(Mass) * math.Pow(c, 2)
	LF_sub := 1 - (float64(velocity) / math.Pow(c, 2))
	lorentzFactor := math.Pow(LF_sub, (1 / 2))
	r := generalRelativity / lorentzFactor
	return r
}

func Relativity() {
	var M, v float64

	fmt.Print("Value of Mass as INTEGER>")
	fmt.Scan(&M)
	fmt.Print("\nValue of Velocity as INTEGER>")
	fmt.Scan(&v)
	E := relativity(M, v)
	fmt.Println("\nCalculated result of E =")
	fmt.Println(E)
}
