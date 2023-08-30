package main

import (
	"fmt"
)

func main() {
	//a := make([]string, 6) creamos un slice con valores nulos
	var a []string
	a = append(a, "Lunes", "Martes", "Miercoles", "Jueves")
	fmt.Println(a[0:3])
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
