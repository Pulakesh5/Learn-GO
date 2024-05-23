package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const lim = 1e-6
	sqrtValue := 1.0
	prev := sqrtValue
	for {
		prev = sqrtValue
		sqrtValue -= (sqrtValue*sqrtValue - x) / (2 * sqrtValue)
		if sqrtValue-prev <= lim {
			break
		}
		fmt.Println(sqrtValue)
	}
	return sqrtValue
}

func main() {
	fmt.Println(Sqrt(2))
}
