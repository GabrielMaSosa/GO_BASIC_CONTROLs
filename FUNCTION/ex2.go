package main

import (
	"fmt"
)

func operacionAritme(a, b float64, tipo string) float64 {

	switch tipo {

	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b

		}

	}

	return 0.0
}

func main() {
	const (
		suma     = "+"
		resta    = "-"
		multip   = "*"
		division = "/"
	)
	var x, y float64
	x = 5.1
	y = 10.1
	fmt.Println(operacionAritme(x, y, division))

}
