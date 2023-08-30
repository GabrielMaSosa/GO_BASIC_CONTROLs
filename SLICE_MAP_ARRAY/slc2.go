package main

import (
	"fmt"
)

func main() {
	//a := make([]string, 6) creamos un slice con valores nulos
	var a []string
	a = append(a, "Lunes", "Martes", "Miercoles", "Jueves") //agregar al slice
	fmt.Println(a[0:3])                                     //imprimo con un rango
	fmt.Println(a)

	//si agrego otro elemento la capacidad se me duplica y el len aumenta al valor actual
	a = append(a, "Sabado")
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
