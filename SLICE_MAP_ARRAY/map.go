package main

import (
	"fmt"
)

func main() {
	//declarar map
	mymap := make(map[string]int)
	mymap1 := map[string]int{}
	mymap3 := map[string]int{"Gabriel1": 32, "Benjamin": 33, "Carlos": 44, "Pedro": 11, "Ricardo": 33}
	mymap1["Gabriel"] = 32
	mymap["Gabriel"] = 32
	fmt.Println(mymap1)
	fmt.Println(mymap["Gabriel"])
	//para editar
	mymap3["Pedro"] = 33
	//agregar
	mymap3["Raul"] = 55
	//eliminar
	delete(mymap3, "Raul")
	fmt.Println(mymap3)

	// para recorrer
	for k, v := range mymap3 {
		fmt.Println(k, v)
	}

}
