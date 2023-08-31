package main

import (
	"fmt"
)

func sumaslice(valores ...int) int {
	var acu int = 0

	for _, val := range valores {
		acu = acu + val

	}
	return acu

}

func main() {
	a, b, c, d, e := 3, 4, 2, 6, -1

	sum := sumaslice(a, b, c, d, e)

	fmt.Println("Suma muchos valores y vale", sum)

}
